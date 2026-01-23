package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// 编辑模式选项
var (
	newDescription string // 新任务描述
	newPriority    string // 新优先级
	toggleStatus   bool   // 切换完成状态
	markUrgent     bool   // 标记为紧急
	markNormal     bool   // 取消紧急标记
)

var editCmd = &cobra.Command{
	Use:   "edit [任务ID]",
	Short: "编辑待办事项",
	Long: `编辑待办事项。
可以修改任务描述、优先级、状态和紧急标记。`,
	Example: `  todo edit 1 -d "新描述"        # 修改描述
  todo edit 2 -p high            # 修改优先级为高
  todo edit 3 -t                 # 切换完成状态
  todo edit 4 -u                 # 标记为紧急
  todo edit 5 -n                 # 取消紧急标记
  
  # 组合使用
  todo edit 1 -d "新任务" -p low -t`,
	Args: cobra.ExactArgs(1), // 必须有一个参数（任务ID）
	RunE: func(cmd *cobra.Command, args []string) error {
		idStr := args[0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return fmt.Errorf("任务ID必须是数字")
		}

		todos, err := readTodos()
		if err != nil {
			return fmt.Errorf("读取任务失败: %v", err)
		}

		// 查找要编辑的任务
		taskIndex := -1
		for i, todo := range todos {
			if todo.ID == id {
				taskIndex = i
				break
			}
		}

		if taskIndex == -1 {
			return fmt.Errorf("未找到任务 #%d", id)
		}

		originalTodo := todos[taskIndex]

		// 如果没有指定任何修改参数，显示错误
		if newDescription == "" && newPriority == "" && !toggleStatus && !markUrgent && !markNormal {
			fmt.Println("当前任务信息:")
			displayTodoDetails(originalTodo)
			fmt.Println("\n请至少指定一个修改选项：")
			fmt.Println("  -d, --desc     新的任务描述")
			fmt.Println("  -p, --priority 设置优先级 (low/normal/high)")
			fmt.Println("  -t, --toggle   切换完成状态")
			fmt.Println("  -u, --urgent   标记为紧急任务")
			fmt.Println("  -n, --normal   取消紧急标记")
			return fmt.Errorf("未指定修改内容")
		}

		modified := false

		// 1. 修改描述
		if newDescription != "" {
			todos[taskIndex].Task = newDescription
			modified = true
			fmt.Printf("📝 描述已更新: %s\n", newDescription)
		}

		// 2. 修改优先级
		if newPriority != "" {
			todos[taskIndex].Priority = newPriority
			modified = true
			fmt.Printf("📊 优先级已更新为: %s\n", newPriority)
		}

		// 3. 切换状态
		if toggleStatus {
			todos[taskIndex].Done = !todos[taskIndex].Done
			modified = true
			if todos[taskIndex].Done {
				fmt.Printf("✅ 任务标记为已完成\n")
			} else {
				fmt.Printf("🔄 任务重新标记为待办\n")
			}
		}

		// 4. 标记为紧急
		if markUrgent {
			if !strings.HasPrefix(todos[taskIndex].Task, "[紧急] ") {
				todos[taskIndex].Task = "[紧急] " + todos[taskIndex].Task
				modified = true
				fmt.Printf("🚨 标记为紧急任务\n")
			}
		}

		// 5. 取消紧急标记
		if markNormal {
			if strings.HasPrefix(todos[taskIndex].Task, "[紧急] ") {
				todos[taskIndex].Task = strings.TrimPrefix(todos[taskIndex].Task, "[紧急] ")
				modified = true
				fmt.Printf("📌 取消紧急标记\n")
			}
		}

		// 如果没有任何修改（所有指定的修改都是冗余的）
		if !modified {
			fmt.Println("⚠️  没有进行任何有效修改")
			return nil
		}

		// 保存修改
		if err := saveTodos(todos); err != nil {
			return fmt.Errorf("保存任务失败: %v", err)
		}

		// 显示修改后的结果
		fmt.Println("\n✅ 任务已更新:")
		displayTodoDetails(todos[taskIndex])

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// 命令别名
	editCmd.Aliases = []string{"modify", "update", "e"}

	// 标志定义
	editCmd.Flags().StringVarP(&newDescription, "desc", "d", "", "新的任务描述")
	editCmd.Flags().StringVarP(&newPriority, "priority", "p", "",
		"设置优先级 (low/normal/high)")
	editCmd.Flags().BoolVarP(&toggleStatus, "toggle", "t", false, "切换完成状态")
	editCmd.Flags().BoolVarP(&markUrgent, "urgent", "u", false, "标记为紧急任务")
	editCmd.Flags().BoolVarP(&markNormal, "normal", "n", false, "取消紧急标记")

	// 标志验证
	editCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		// 验证优先级值
		if newPriority != "" && newPriority != "low" &&
			newPriority != "normal" && newPriority != "high" {
			return fmt.Errorf("优先级必须是 low, normal 或 high")
		}

		// 紧急和普通标记不能同时使用
		if markUrgent && markNormal {
			return fmt.Errorf("不能同时使用 --urgent 和 --normal")
		}

		return nil
	}
}

// 显示任务详情
func displayTodoDetails(todo TodoItem) {
	fmt.Printf("ID:        #%d\n", todo.ID)

	// 状态和紧急标记
	statusMark := "⏳"
	if todo.Done {
		statusMark = "✅"
	}

	urgentMark := ""
	if strings.HasPrefix(todo.Task, "[紧急] ") {
		urgentMark = "🚨 "
	}

	fmt.Printf("状态:      %s ", statusMark)
	if todo.Done {
		fmt.Printf("已完成\n")
	} else {
		fmt.Printf("待办\n")
	}

	fmt.Printf("任务:      %s%s\n", urgentMark, todo.Task)

	// 优先级
	priority := todo.Priority
	if priority == "" {
		priority = "normal"
	}
	priorityIcon := "📊"
	switch priority {
	case "low":
		priorityIcon = "📉"
	case "high":
		priorityIcon = "📈"
	}
	fmt.Printf("优先级:    %s %s\n", priorityIcon, priority)

	fmt.Printf("创建时间:  %s\n", todo.CreatedAt.Format("2006-01-02 15:04"))
}
