package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)

func main() {
	const (
		Create = "create"
		Read = "read"
		Update = "update"
		Delete = "delete"
	)

	tasks := make([]string, 0)

	fmt.Println("Welcome to the TO DO List CLI app!")
	for {
	fmt.Println()
    fmt.Println("Enter your command (create, read, update, delete):")
		var command string
		fmt.Scan(&command)
		switch command {
		case Create:
			fmt.Println("Enter task name: ")
			newTask, _ := bufi.NewReader(os.Stdin).ReadString('\n')
			newTask = newTask[:len(newTask) - 1]
			tasks = append(tasks, newTasks)

		case Read:
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		case Update:
			fmt.Println("Enter task name to update: ")
			input, _ := bufoi.NewReader(os.Stdin).ReadString('\n')
			input = input[:len(input) - 1]

			indexToUpdate := -1
			for i, task := range tasks {
				if task == input {
					indexToUpdate = i
					break
				}
			}

			if indexToUpdate == -1 {
				fmt.Println("Invalid name.Please try again. ")
				continue
			}

			fmt.Println("enter new task name: ")
			new
		case Delete:

		}
	





	
}