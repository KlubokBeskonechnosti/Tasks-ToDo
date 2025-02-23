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
		Read   = "read"
		Update = "update"
		Delete = "delete"
	)

	tasks := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the TO DO List CLI app!")
	
	for {
		fmt.Println()
		fmt.Println("Enter your command (create, read, update, delete):")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case Create:
			fmt.Println("Enter task name: ")
			newTask, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			newTask = newTask[:len(newTask)-1]
			tasks = append(tasks, newTask)

		case Read:
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		case Update:
			fmt.Println("Enter task name to update: ")
			input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			input = input[:len(input)-1]

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

			fmt.Println("Enter new task name: ")
			newTaskName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			newTaskName = newTaskName[:len(newTaskName)-1]

			if len(newTaskName) < 3 {
				fmt.Println("The new task namw is too short! Please, try again.")
				continue
			}
			tasks[indexToUpdate] = newTaskName
			fmt.Printf("Updated task #%d with name \"%s\" successfully!\n", indexToUpdate, newTaskName)

		case Delete:
			fmt.Println("Enter task name to remove: ")
			input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			input = input[:len(input)-1]

			indexToRemove := -1

			for i, task := range tasks {
				if task == input {
					indexToRemove = i
					break
				}
			}

			if indexToRemove == -1 {
				fmt.Println("Invalid name.Please try again.")
				continue
			}

			oldTaskName := tasks[indexToRemove]
			tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...)
			fmt.Printf("Removed tasks #%d with name \"%s\" successfully!\n", indexToRemove, oldTaskName)

		default:
			fmt.Println("Invalid command! Please, try again!")

		}

	}

}
