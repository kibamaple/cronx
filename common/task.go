package common

type CTask struct {
	ID int64
	Job string
	Name string
	Params *map[string]string
	Status uint8
}

type ITaskListener interface {
	OnTask(*CTask)
}

type CTaskListener struct {
	TaskChan chan *CTask
}

func (this CTaskListener) OnTask(task *CTask) {
	this.TaskChan <- task
}

type ITaskProvider interface {
	AddTaskListener(*ITaskListener)
}

type CTaskProvider struct {
	taskListeners []*ITaskListener
}

func (this CTaskProvider) EmitTask(task *CTask) {
	for _, r := range this.taskListeners {
		r.OnTask(task)
	}
}

func (this CTaskProvider) AddTaskListener(listener *ITaskListener) {
	this.taskListeners = append(this.taskListeners,listener)
}