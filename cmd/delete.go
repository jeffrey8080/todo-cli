package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	deleteAll   bool // 是否删除所有任务
	forceDelete bool // 强制删除（不确认）
)

var deleteCmd = &cobra.Command{
	Use:   "delete [任务ID...]",
	Short: "删除一个或多个待办事项",
	Long: `删除指定的待办事项。
可以删除单个任务、多个任务或所有任务。
支持使用逗号分隔的ID列表或ID范围。`,
	Example: `  todo delete 1                 # 删除任务1
  todo delete 1 3 5             # 删除任务1,3,5
  todo delete 1-3               # 删除任务1到3
  todo delete --all             # 删除所有任务
  todo delete 2,4,6-8           # 混合使用`,
	Args: cobra.MinimumNArgs(0), // 允许0个参数（使用--all时）
	RunE: func(cmd *cobra.Command, args []string) error {
		todos, err := readTodos()
		if err != nil {
			return fmt.Errorf("读取任务失败: %v", err)
		}

		if len(todos) == 0 {
			fmt.Println("📭 没有可删除的任务")
			return nil
		}

		// 情况1: 删除所有任务
		if deleteAll {
			if !forceDelete {
				fmt.Printf("⚠️  即将删除所有 %d 个任务，确定吗？ (y/N): ", len(todos))
				var confirm string
				fmt.Scanln(&confirm)
				if strings.ToLower(confirm) != "y" {
					fmt.Println("操作已取消")
					return nil
				}
			}

			// 清空文件
			if err := saveTodos([]TodoItem{}); err != nil {
				return fmt.Errorf("清空任务失败: %v", err)
			}

			fmt.Printf("🗑️  已删除所有 %d 个任务\n", len(todos))
			return nil
		}

		// 情况2: 没有指定任务ID
		if len(args) == 0 {
			return fmt.Errorf("请指定要删除的任务ID，或使用 --all 删除所有任务")
		}

		// 解析任务ID（支持多种格式）
		idsToDelete, err := parseTaskIDs(args, todos)
		if err != nil {
			return err
		}

		if len(idsToDelete) == 0 {
			fmt.Println("没有找到匹配的任务")
			return nil
		}

		// 显示将要删除的任务
		if !forceDelete {
			fmt.Println("将要删除以下任务:")
			for _, id := range idsToDelete {
				for _, todo := range todos {
					if todo.ID == id {
						status := "待办"
						if todo.Done {
							status = "已完成"
						}
						fmt.Printf("  #%d [%s] %s\n", todo.ID, status, todo.Task)
						break
					}
				}
			}

			fmt.Printf("\n确认删除 %d 个任务吗？ (y/N): ", len(idsToDelete))
			var confirm string
			fmt.Scanln(&confirm)
			if strings.ToLower(confirm) != "y" {
				fmt.Println("操作已取消")
				return nil
			}
		}

		// 删除任务
		var newTodos []TodoItem
		deletedCount := 0
		for _, todo := range todos {
			shouldDelete := false
			for _, id := range idsToDelete {
				if todo.ID == id {
					shouldDelete = true
					deletedCount++
					break
				}
			}
			if !shouldDelete {
				newTodos = append(newTodos, todo)
			}
		}

		// 重新编号（可选，保持ID连续）
		for i := range newTodos {
			newTodos[i].ID = i + 1
		}

		// 保存
		if err := saveTodos(newTodos); err != nil {
			return fmt.Errorf("保存任务失败: %v", err)
		}

		fmt.Printf("🗑️  已删除 %d 个任务，剩余 %d 个任务\n", deletedCount, len(newTodos))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// 添加命令别名
	deleteCmd.Aliases = []string{"del", "rm"}

	// 标志定义
	deleteCmd.Flags().BoolVarP(&deleteAll, "all", "a", false, "删除所有任务")
	deleteCmd.Flags().BoolVarP(&forceDelete, "force", "y", false, "强制删除，不确认")

	// 删除所有任务时，强制标志自动为true
	deleteCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if deleteAll && len(args) > 0 {
			return fmt.Errorf("不能同时使用 --all 和指定任务ID")
		}
		return nil
	}
}

// 解析任务ID（支持多种格式）
func parseTaskIDs(args []string, todos []TodoItem) ([]int, error) {
	var ids []int
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}

	for _, arg := range args {
		// 检查是否是范围格式 "1-3"
		if strings.Contains(arg, "-") {
			parts := strings.Split(arg, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("无效的范围格式: %s，请使用如 1-3 的格式", arg)
			}

			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("无效的范围: %s，必须是数字", arg)
			}

			if start < 1 || end > maxID || start > end {
				return nil, fmt.Errorf("无效的范围: %d-%d，有效ID为1-%d", start, end, maxID)
			}

			for i := start; i <= end; i++ {
				ids = append(ids, i)
			}
			continue
		}

		// 检查是否是用逗号分隔的列表 "1,3,5"
		if strings.Contains(arg, ",") {
			parts := strings.Split(arg, ",")
			for _, part := range parts {
				id, err := strconv.Atoi(strings.TrimSpace(part))
				if err != nil {
					return nil, fmt.Errorf("无效的ID: %s", part)
				}
				if id < 1 || id > maxID {
					return nil, fmt.Errorf("任务ID %d 不存在，有效ID为1-%d", id, maxID)
				}
				ids = append(ids, id)
			}
			continue
		}

		// 单个ID
		id, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("无效的ID: %s", arg)
		}
		if id < 1 || id > maxID {
			return nil, fmt.Errorf("任务ID %d 不存在，有效ID为1-%d", id, maxID)
		}
		ids = append(ids, id)
	}

	// 去重
	return removeDuplicates(ids), nil
}

// 去除重复的ID
func removeDuplicates(ids []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, id := range ids {
		if !seen[id] {
			seen[id] = true
			result = append(result, id)
		}
	}
	return result
}
