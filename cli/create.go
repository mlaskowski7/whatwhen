package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
	"whatwhen/bll"
)

var (
	title    string
	desc     string
	deadline string
	finished bool
)

var createCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a what",
	Run: func(cmd *cobra.Command, args []string) {
		if title == "" {
			fmt.Println("Title is required")
			return
		}

		parsedDeadline, err := time.Parse("2006-01-02T15:04", deadline)
		if err != nil {
			fmt.Println("Invalid deadline format, use YYYY-MM-DDTHH:MM")
			return
		}

		task, err := bll.CreateTask(title, desc, parsedDeadline, finished)
		if err != nil {
			fmt.Println("Failed to create the what:", err)
			return
		}

		fmt.Println("New what created :)")
		fmt.Println(task)
	},
}

func AddCreateCmdFlags() {
	createCmd.Flags().StringVar(&title, "title", "", "Task title (required)")
	createCmd.Flags().StringVar(&desc, "desc", "", "Task description")
	createCmd.Flags().StringVar(&deadline, "deadline", "", "Deadline in format YYYY-MM-DDTHH:MM")
	createCmd.Flags().BoolVar(&finished, "finished", false, "Whether the task is already finished")
}
