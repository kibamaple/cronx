package common

type CEvent_Name uint8

const (
	Put CEvent_Name = iota
	Delete
)

type CEvent struct {
	Name CEvent_Name
	Key string
	Job CJob
	time int64
}

type IEventListener interface {
	OnEvent(*CEvent)
}

type CEventListener struct {
	EventChan chan *CEvent
}

func (this CEventListener) OnEvent(event *CEvent) {
	this.EventChan <- event
}

type IEventProvider interface {
	AddEventListener(*IEventListener)
}

type CEventProvider struct {
	eventListeners []*IEventListener
}

func (this CEventProvider) EmitEvent(event *CEvent) {
	for _, r := range this.eventListeners {
		r.OnEvent(event)
	}
}

func (this CEventProvider) AddEventListener(listener *IEventListener) {
	this.eventListeners = append(this.eventListeners,listener)
}