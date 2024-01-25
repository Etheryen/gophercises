package cmd

import (
	"fmt"
	"os"
	"task/tasks"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all undone tasks",
	Run: func(_ *cobra.Command, _ []string) {
		err := tasks.List()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
