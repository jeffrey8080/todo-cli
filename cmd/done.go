package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var doneCmd = &cobra.Command{
	Use:   "done [任务ID]",
	Short: "标记任务为已完成",
	Args:  cobra.ExactArgs(1), // 必须且只能有一个参数
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

		found := false
		for i := range todos {
			if todos[i].ID == id {
				if todos[i].Done {
					fmt.Printf("⚠️  任务 #%d 已经是完成状态\n", id)
				} else {
					todos[i].Done = true
					fmt.Printf("🎉 已完成任务 #%d: %s\n", id, todos[i].Task)
				}
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("未找到任务 #%d", id)
		}

		// 保存更新
		if err := saveTodos(todos); err != nil {
			return fmt.Errorf("保存任务失败: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
