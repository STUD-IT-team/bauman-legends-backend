package response

type GetTaskTypes struct {
	TaskTypes []TaskType `json:"taskTypes"`
}

type TaskType struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	IsActive bool   `json:"is_active"`
}
