package entity

//Task Entity
type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	DueDate string `json:"dueDate"`
	Status  int    `json:"status"`
}
