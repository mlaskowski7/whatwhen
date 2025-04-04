package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"whatwhen/bll"
)

var viewCmd = &cobra.Command{
	Use:   "whats",
	Short: "View all whats",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := bll.GetAllTasks()
		if err != nil {
			fmt.Println("An error occurred while retrieving tasks -> ", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Tasks list is empty")
			return
		}

		fmt.Println("Your tasks:")
		for _, task := range tasks {
			fmt.Println(task)
		}
	},
}
