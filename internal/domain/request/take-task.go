package request

type TakeTask struct {
	TaskTypeId  string `json:"taskTypeId"`
	AccessToken string
}
