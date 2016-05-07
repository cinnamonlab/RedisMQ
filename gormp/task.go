package gormq

type QueueTask struct {
	Function QueueFunc
	Param string
}

func (task QueueTask) Execute()  {
	task.Function(task.Param)
}
