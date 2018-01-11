package module

import (
	"sync"
	common "github.com/kibamaple/cronx/common"
	model "github.com/kibamaple/cronx/model"
	tool "github.com/kibamaple/cronx/tool"
)

type CSpec struct {
	common.CErrorProvider
	common.CTaskProvider
	common.CEventListener
	abortChan chan struct{}
	startWG sync.WaitGroup
	stopWG sync.WaitGroup
}

func (this CSpec) runJob(timer *tool.CTimer,job *model.CJob){
	err := this.timer.Add(job.Name,job.Spec,func (id int64) {
		task := &common.CTask{id,job.Name,job.Task,job.Params,job.Status}
		this.EmitTask(task)
	})
	if err != nil {
		this.EmitError(err)
	}
}

func (this CSpec) process() {
	this.startWG.Wait()

	var events map[string]common.CEvent
	timer := tool.NewTimer()
	timer.start()
	defer timer.stop()
	for {
		select {
			case <-this.abortChan:
				this.stopWG.Done()
				return
			case event := <- this.EventChan:
				_event,ok := this.events[event.Key]
				if ok && _event.Time > event.Time {
					continue
				}
				if event.Name == Delete || event.Name == Put {
					timer.Remove(event.Key)
				}
				if event.Name != Delete {
					this.runJob(timer,&event.Job)
				}
				this.datas[event.Key] = event
		}
	}
}

func (this CSpec) Init() {
	go this.process()
}
func (this CSpec) Start() {
	this.startWG.Done()
}

func (this CSpec) Stop() {
	close(this.abortChan)
	this.stopWG.Done()
}