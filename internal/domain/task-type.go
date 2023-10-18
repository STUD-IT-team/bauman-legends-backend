package domain

type TaskType struct {
	ID         int    `db:"id"`
	Title      string `db:"title"`
	IsActive   bool
	Count      int
	TeamAmount int
}

type TaskTypes []TaskType
