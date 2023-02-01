package models

type Student struct {
	Id      int
	Email   string
	Name    string
	Age     int
	PhoneNo int
}

type DateAndTime struct {
	Date string
	Time string
}

type Todo struct {
	Id       int
	TaskName string
	DateTime DateAndTime
}

//todo list -> create a task, update a task, delete a task, read all the task
