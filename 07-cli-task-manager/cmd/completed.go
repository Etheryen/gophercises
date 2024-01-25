package cmd

import (
	"fmt"
	"os"
	"task/tasks"

	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all tasks completed today",
	Run: func(_ *cobra.Command, _ []string) {
		err := tasks.Completed()
		if err != nil {
			fmt.Println("Error listing completed tasks:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
