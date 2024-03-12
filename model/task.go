package model

type Task struct {
	Id         int
	Opened     int
	Closed     int
	AuthorId   int
	AssignedId int
	Title      string
	Content    string
}

type UpdateTask struct {
	Title   string
	Content string
}
