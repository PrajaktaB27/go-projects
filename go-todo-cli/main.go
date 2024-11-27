package main

import (
	"fmt"
	"net/http"
)

func main() {
	// var taskOne = "watch go crash course"
	// taskTwo := "update gh with go project"

	// taskItems := []string{taskOne, taskTwo}
	// printTasks(taskItems)

	// taskItems = addTask(taskItems, "find a new video")

	// printTasks(taskItems)

	fmt.Println("############### Welcome to our todo list app! ###############")
	http.HandleFunc("/go-todo", helloUser)
	http.HandleFunc("/go-todo", showTasks)

	http.ListenAndServe(":3000", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "############### Welcome to our todo list app! ###############"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	var greeting = "############### Welcome to our todo list app! ###############"
	fmt.Fprintln(writer, greeting)
}

func printTasks(taskItems []string) {
	for index, task := range taskItems {
		// fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	updatedList := append(taskItems, newTask)
	fmt.Println("updated tasks")

	// for index, task := range updatedList {
	// 	fmt.Printf("%d. %s\n", index+1, task)
	// }
	return updatedList
}
