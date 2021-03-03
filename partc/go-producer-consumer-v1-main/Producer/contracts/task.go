package contracts

type Task struct {
	TaskName       string `json:"taskName"`
	TaskType       string `json:"taskType"`
	LastUpdateTime string `json:"lastUpdateTime"`
	ScheduledTime  string `json:"scheduledTime"`
	Periodicity    int    `json:"periodicity"`
	TaskStatus     string `json:"taskStatus"`
}
