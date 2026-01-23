package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

type TodoItem struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	Priority  string    `json:"priority,omitempty"`
}

var (
	priority string // 局部标志变量
	urgent   bool   // 局部标志变量
)

var addCmd = &cobra.Command{
	Use:   "add [任务描述]",
	Short: "添加一个新的待办事项",
	Long:  `添加一个待办事项到列表中。可以指定优先级和紧急标志。`,
	Args:  cobra.MinimumNArgs(1), // 至少需要一个参数
	Example: `todo add "买牛奶"    #添加简单任务
todo add "写报告" -p high    #高优先级任务
todo add "紧急会议" -u    #紧急任务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 合并所有参数为任务描述
		taskDescription := args[0]
		if len(args) > 1 {
			for i := 1; i < len(args); i++ {
				taskDescription += " " + args[i]
			}
		}

		// 读取现有任务
		todos, err := readTodos()
		if err != nil {
			return fmt.Errorf("读取任务失败: %v", err)
		}

		// 创建新任务
		newTodo := TodoItem{
			ID:        len(todos) + 1,
			Task:      taskDescription,
			Done:      false,
			CreatedAt: time.Now(),
			Priority:  priority,
		}

		// 如果是紧急任务，添加标识
		if urgent {
			newTodo.Task = "[紧急] " + newTodo.Task
		}

		todos = append(todos, newTodo)

		// 保存到文件
		if err := saveTodos(todos); err != nil {
			return fmt.Errorf("保存任务失败: %v", err)
		}

		fmt.Printf("✅ 已添加任务 #%d: %s\n", newTodo.ID, newTodo.Task)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// 局部标志：仅对 add 命令有效
	addCmd.Flags().StringVarP(&priority, "priority", "p", "normal",
		"任务优先级 (low, normal, high)")
	addCmd.Flags().BoolVarP(&urgent, "urgent", "u", false, "标记为紧急任务")

	// 标志验证
	addCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if priority != "low" && priority != "normal" && priority != "high" {
			return fmt.Errorf("优先级必须是 low, normal 或 high")
		}
		return nil
	}
}

// 读取任务列表
func readTodos() ([]TodoItem, error) {
	var todos []TodoItem

	if _, err := os.Stat(todoFile); os.IsNotExist(err) {
		return todos, nil // 文件不存在时返回空列表
	}

	data, err := os.ReadFile(todoFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// 保存任务列表
func saveTodos(todos []TodoItem) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(todoFile, data, 0644)
}
