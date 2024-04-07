package main

import "fmt"

type Task struct {
	Title    string
	Estimate int
}

func main() {
	task1 := Task{
		Title:    "Learn Golang",
		Estimate: 3,
	}
	task1.Title = "Learning Go"
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)
}
