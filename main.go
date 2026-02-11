package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// терминальный ту ду лист
// сюда надо добавить аля круды наверное самих задач
//

// var count int
var tasks []string

func main() {

	data, err := os.ReadFile("data.json")
	if err == nil {

		err = json.Unmarshal(data, &tasks)

	} else {
		tasks = []string{} // Новый файл — пустой список
	}

	for {
		// size++;
		slice := tasks[:]

		fmt.Println("Your tasks:")
		for i, v := range slice {
			i++
			fmt.Println(i, v)
		}
		fmt.Printf("Select command task(0 - create new task): \n")

		command := 0
		fmt.Scan(&command)

		if command == 0 {
			fmt.Printf("You select command - new task\n")
			fmt.Printf("Enter you task:\n")
			newTask := ""
			fmt.Scan(&newTask)
			createTask(newTask)
		} else if command > 0 {
			deleteTask(command)
		} else {
			fmt.Println("you piece of shit")
			os.Exit(1)
		}
	}
}

func createTask(task string) []string {
	tasks = append(tasks, task)
	saveTasks()
	fmt.Printf("Create!")
	return tasks
}

func saveTasks() {

	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tasks)
	if err != nil {
		panic(err)
	}
}

func deleteTask(selectedTask int) {
	// Индекс в слайсе (нумерация с 1 для пользователя, 0 для слайса)
	index := selectedTask - 1

	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task number")
		return
	}

	fmt.Printf("You choose task %d: '%s'\n", selectedTask, tasks[index])
	fmt.Printf("Delete? Send 'YES' if you want to delete\n")

	var chooseDelete string
	fmt.Scan(&chooseDelete)

	if chooseDelete == "YES" {
		// Удаляем элемент из слайса
		tasks = append(tasks[:index], tasks[index+1:]...)
		saveTasks()
		fmt.Printf("Task %d deleted\n", selectedTask)
	} else {
		fmt.Println("Cancelled")
	}
}
