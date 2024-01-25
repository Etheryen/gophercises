package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task/tasks"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task with a given number",
	Run: func(_ *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(
				"Error removing task: you need to provide 1 task number",
			)
			os.Exit(1)
		}

		num, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf(
				"Error removing task: invalid number provided (%v)\n",
				err,
			)
			os.Exit(1)
		}

		err = tasks.Remove(num)
		if err != nil {
			fmt.Println("Error removing task:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
