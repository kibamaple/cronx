package module

import (
	common "github.com/kibamaple/cronx/common"
	model "github.com/kibamaple/cronx/model"
)

type CRecorder struct {
	resultModel *model.IResultModel
	common.CErrorProvider
}

func (this CRecorder) OnResult(result *common.CResult) {
	err := this.resultModel.Save(result)
	this.EmitError(err)
}

type CConcurrentRecorder struct {
	CRecorder
}

func (this CRecorder) OnResult(result *common.CResult) {
	go this.CRecorder.OnResult(result)
}

type CHandleRecorder struct {
	CRecorder
	common.CResultListener
	common.CConcurrent
}

func (this CHandleRecorder) Process(abortChan chan struct{}){
	for {
		select	{
			case result := <-this.ResultChan:
				this.CRecorder.OnResult(result)
			case <-abortChan:
				return
		}
	}
}