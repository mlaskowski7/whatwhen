package main

import (
	"fmt"
	"os"
	"whatwhen/cli"
	"whatwhen/dal"
)

func main() {
	err := dal.ConnectToPostgres()
	if err != nil {
		fmt.Println("Could not to connect to the database :O ->", err)
		os.Exit(1)
	}

	err = cli.Execute()
	if err != nil {
		fmt.Println("Command line execution failed :O ->", err)
		os.Exit(1)
	}
}
