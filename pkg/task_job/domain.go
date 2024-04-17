package taskjob

type TestTaskPayload struct {
	TaskType string
	DivviupId string
}

type CollectorRunPayload struct {
	TaskId uint
	TaskType string
	TaskName string
	DivviUpId string
}
