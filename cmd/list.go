package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var (
	showAll  bool
	onlyDone bool
	sortBy   string
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "列出所有待办事项",
	Aliases: []string{"ls", "l"}, // 命令别名
	RunE: func(cmd *cobra.Command, args []string) error {
		todos, err := readTodos()
		if err != nil {
			return fmt.Errorf("读取任务失败: %v", err)
		}

		if len(todos) == 0 {
			fmt.Println("📭 没有待办事项")
			return nil
		}

		// 过滤任务
		var filtered []TodoItem
		for _, todo := range todos {
			if onlyDone && !todo.Done {
				continue
			}
			if !showAll && todo.Done && !onlyDone {
				continue
			}
			filtered = append(filtered, todo)
		}

		if len(filtered) == 0 {
			fmt.Println("没有符合条件的任务")
			return nil
		}

		// 使用 tabwriter 格式化输出
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\t状态\t优先级\t任务\t创建时间")
		fmt.Fprintln(w, "--\t----\t------\t----\t--------")

		for _, todo := range filtered {
			status := "待办"
			if todo.Done {
				status = "✅完成"
			}

			prio := todo.Priority
			if prio == "" {
				prio = "normal"
			}

			timeStr := todo.CreatedAt.Format("2006-01-02 15:04")
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
				todo.ID, status, prio, todo.Task, timeStr)
		}

		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// 互斥标志示例
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "显示所有任务（包括已完成）")
	listCmd.Flags().BoolVarP(&onlyDone, "done", "d", false, "只显示已完成的任务")
	listCmd.Flags().StringVarP(&sortBy, "sort", "s", "id", "排序方式 (id, time, priority)")

	// 标志互斥：--all 和 --done 不能同时使用
	listCmd.MarkFlagsMutuallyExclusive("all", "done")
}
