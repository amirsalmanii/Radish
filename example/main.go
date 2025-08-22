package main

import (
	"fmt"
	"net/http"
	"github.com/amirsalmanii/Radish/radish"
	"github.com/amirsalmanii/Radish/example/tasks"
)

func main() {
	// init Radish
	radishQ := radish.Radish{}

	// collect tasks
	radishQ.ListOfTasks = append(
		radishQ.ListOfTasks,
		func() {
			// if have parameters
			radish.Task(tasks.WriteFileTask1, "task1.txt", []byte("Hello From Task 1!"))
		},
		func() {
			// if doesn't have parameters
			radish.Task(tasks.WriteFileTask2)
		},
	)

	// start running task in background
	radishQ.RunTaskQueue()

	// running server or etc
	http.HandleFunc("/hello", Hello)
	fmt.Println("running server on port 8090")
	http.ListenAndServe(":8090", nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello \n")
}
