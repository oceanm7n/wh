package models

// json-bindable
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TaskList struct {
	Tasks []Task
}
