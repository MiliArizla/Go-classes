package main

type Tasks struct {
	Id          string
	Nome        string
	ProjectId   string
	Description []string
	UserId      string
	DueDate     string
	TaskType    string
	Sprint      int
	Priority    Priority
	Status      Status
}

type Priority string

const (
	Low    string = "low"
	Medium        = "medium"
	High          = "high"
)

type Status string

const (
	inProgress string = "In Progress"
	open              = "Open"
	resolved          = "Resolved"
	closed            = "Closed"
	blocked           = "Blocked"
)
