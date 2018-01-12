package common

type CResult struct {
	ID int64
	Job string
	Status uint8
}

type IResultListener interface {
	OnResult(*CResult)
}

type CResultListener struct {
	ResultChan chan *CResult
}

func (this CResultListener) OnResult(result *CResult) {
	this.ResultChan <- result
}

type IResultProvider interface {
	AddResultListener(IResultListener)
}

type CResultProvider struct {
	resultListeners []IResultListener
}

func (this CResultProvider) EmitResult(result *CResult) {
	for _, r := range this.resultListeners {
		r.OnResult(result)
	}
}

func (this CResultProvider) AddResultListener(listener IResultListener) {
	this.resultListeners = append(this.resultListeners,listener)
}
