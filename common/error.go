package common

type IErrorListener interface {
	OnError(error)
}

type CErrorListener struct {
	ErrorChan chan error
}

func (this *CErrorListener) OnError(e error) {
	this.ErrorChan <- e
}

type IErrorProvider interface {
	AddErrorListener(IErrorListener)
}

type CErrorProvider struct {
	errorListeners []IErrorListener
}

func (this *CErrorProvider) EmitError(err error) {
	for _, r := range this.errorListeners {
		r.OnError(err)
	}
}

func (this *CErrorProvider) AddErrorListener(listener IErrorListener) {
	this.errorListeners = append(this.errorListeners,listener)
}