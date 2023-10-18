package domain

type TaskType struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	IsActive bool
	Count    int
}

type TaskTypes []TaskType
