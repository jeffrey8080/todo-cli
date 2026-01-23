package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	todoFile = "todos.json" // 全局配置：存储文件
	version  = "1.0.1"      // 版本号
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "一个简单的待办事项管理工具",
	Long: `Todo CLI 是一个用于管理日常任务的命令行工具。
你可以添加、查看、标记完成你的待办事项。`,
	Version: version, // 设置版本号
	// 直接运行 `todo` 时显示帮助
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("使用 'todo --help' 查看所有可用命令")
	},
}

// 添加全局标志（所有子命令都可使用）
func init() {
	// 全局标志：指定不同的存储文件
	rootCmd.PersistentFlags().StringVarP(&todoFile, "file", "f", "todos.json",
		"指定待办事项存储文件")

	// 设置版本显示模板
	rootCmd.SetVersionTemplate(`{{printf "%s 版本 %s\n" .Name .Version}}`)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}
}
