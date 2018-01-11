package module

import (
	"sync"
	common "github.com/kibamaple/cronx/common"
)

type CEmitter struct {
	dsn string
	key string
	sep string
	timeout time.Duration
	common.CErrorProvider
	common.CEventProvider
	abortChan chan struct{}
	startWG sync.WaitGroup
	stopWG sync.WaitGroup
}

func (this CEmitter) init(){
	jobs,jobs_err := this.jobModel.GetSpec()
	if jobs_err != nil {
		this.EmitError(jobs_err)
		return
	}
	for key,job := range jobs {
		this.EmitEvent(&common.CEvent{Put,key,job})
	}
}

func (this CEmitter) process() {
	this.startWG.Wait()
	
	for {
		if this.watch() {
			this.stopWG.Done()
			break
		}
	}


}

func (this CEmitter) watch()(bool){
	// election watcher
	for {
		if err := this.elect.Campaign(); err != nil {
			select {
				case <-this.abortChan:
					return true
				default:
					this.EmitError(err)
					continue
			}
		}
	}
	defer this.elect.Resign()

	// watch spec job change
	eventChan,errorChan := this.jobModel.WatchSpec(this.abortChan)

	//init sepc job
	this.init()
	for {
		select {
			case <-this.abortChan:
				return true
			case err := <-errorChan:
				this.EmitError(err)
				break
			case event := <-eventChan:
				this.EmitEvent(event)
		}
	}
	return false
}

func (this CEmitter) Init() {
	go this.process()
}

func (this CEmitter) Start() {
	this.startWG.Done()
}

func (this CEmitter) Stop() {
	close(this.abortChan)
	this.stopWG.Done()
}