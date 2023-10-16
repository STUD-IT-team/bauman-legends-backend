package response

type TaskTypes struct {
	TaskTypes []TaskType `json:"task_types"`
}

type TaskType struct {
	Name     string `json:"name"`
	ID       int32  `json:"id"`
	IsActive bool   `json:"is_active"`
}
