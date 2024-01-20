package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task/tasks"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
