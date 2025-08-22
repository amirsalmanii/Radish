package tasks

import (
	"os"
	"time"
)

func WriteFileTask1(fileName string , content []byte){
	// some io jobs
	time.Sleep(time.Second * 5)
	// then do
	os.WriteFile(fileName, []byte(content), 0644) 
}

func WriteFileTask2(){
	// some io jobs
	time.Sleep(time.Second * 2)
	// then do
	os.WriteFile("task2.txt", []byte("Hello From Task 2!"), 0644) 
}
