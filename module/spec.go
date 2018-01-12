package module

import (
	"sync"
	common "github.com/kibamaple/cronx/common"
	model "github.com/kibamaple/cronx/model"
	tool "github.com/kibamaple/cronx/tool"
)

type CSpec struct {
	dsn string
	key string
	val string
	common.CErrorProvider
	common.CTaskProvider
	eventModel *model.IEventModel
	signalChan chan struct{}
	startWG sync.WaitGroup
	stopWG sync.WaitGroup
}

func (this CSpec) runJob(timer *tool.ITimer,job *model.CJob){
	err := timer.Add(job.Name,job.Spec,func (id int64) {
		task := &common.CTask{id,job.Name,job.Task,job.Params,job.Status}
		this.EmitTask(task)
	})
	if err != nil {
		this.EmitError(err)
	}
}

func (this CSpec) run () {
	var events map[string]model.CEvent
	timer := tool.NewTimer()
	timer.Start()
	defer timer.Stop()
	elect := tool.NewElection(this.dsn,this.key,this.val)
	defer elect.Destroy()

	// election watcher
	for {
		if err := elect.Campaign(); err != nil {
			select {
				case _,ok := <-this.signalChan:
					if !ok {
						return true
					}
				default:
					this.EmitError(err)
					continue
			}
		}
	}
	defer elect.Resign()

	// watch spec job change
	eventChan,errorChan := this.eventModel.Watch(this.signalChan)
	for {
		select {
			case _,ok := <-this.signalChan:
				if !ok {
					return true
				}
			case err,ok := <-errorChan:
				if ok {
					this.EmitError(err)
				}else{
					break	
				}
			case event,ok := <-eventChan:
				if ok {
					this.update(event,events,timer)
				}else{
					break
				}
		}
	}
	return false
}

func (this CSpec) update(event *model.CEvent,events map[string]*model.CEvent,timer *tool.ITimer){
	_event,ok := events[event.Key]
	if ok && _event.Time > event.Time {
		continue
	}
	if event.Name == Delete || event.Name == Put {
		timer.Remove(event.Key)
	}
	if event.Name != Delete {
		this.runJob(timer,event.Job)
	}
	events[event.Key] = event
}

func (this CSpec) process() {
	this.startWG.Wait()

	for {
		if this.run() {
			this.stopWG.Done()
			break
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
	close(this.signalChan)
	this.stopWG.Done()
}