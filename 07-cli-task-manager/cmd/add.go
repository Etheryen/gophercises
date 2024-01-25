package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/tasks"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task with a given title",
	Run: func(_ *cobra.Command, args []string) {
		title := strings.Join(args, " ")

		err := tasks.Add(title)
		if err != nil {
			fmt.Println("Error adding task:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
