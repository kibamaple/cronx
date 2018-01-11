package module

import (
	common "github.com/kibamaple/cronx/common"
)

type CCronx struct {
	common.CErrorProvider
	common.CTaskProvider
	common.CResultProvider
}

func (this CCronx) OnError(e error) {
	this.EmitError(err)
}

func (this CCronx) OnTask(task *common.CTask) {
	this.EmitTask(task)
}

func (this CCronx) OnResult(result *common.CResult) {
	this.EmitResult(result)
}

type CConcurrentCronx struct {
	CCronx
}

func (this CConcurrentCronx) OnError(e error) {
	go this.CCronx.OnError(err)
}

func (this CConcurrentCronx) OnTask(task *common.CTask) {
	go this.CCronx.OnTask(task)
}

func (this CConcurrentCronx) OnResult(result *common.CResult) {
	go this.CCronx.OnResult(result)
}

type CHandleCronx struct {
	CCronx
	common.CErrorListener
	common.CTaskListener
	common.CResultListener
	common.CConcurrent
}

func (this CHandleCronx) Process(abortChan <-chan struct{}){
	for {
		select	{
			case err := <-this.ErrorChan:
				this.CCronx.OnError(err)
			case task := <-this.TaskChan:
				this.CCronx.OnTask(task)
			case result := <-this.ResultChan:
				this.CCronx.OnResult(result)
			case <-abortChan:
				return
		}
	}
}