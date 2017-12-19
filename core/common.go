package core

const EmptyString = ""

type CTask struct {
	ID int64
	Job string
	Name string
	Params map[string]string
	Single bool
}

const (
	Success uint8 = iota
	Failed
)

type CResult struct {
	ID int64
	Job string
	Status uint8
	Single bool
}

const (
	Put CEvent_Name = iota
	Delete
)

type CEvent struct {
	Name CEvent_Name
	Key string
	Value string
}

type CEvent_Name uint8

type IExecutor interface {
	Listen(<-chan *CTask,chan<- *CResult)
}

type IElection interface {
	Campaign(string) error
	Resign() error
	Destroy() error
}

type IStore interface {
	Put(string,string) error
	Get(string) (string,error)
	GetByPrefix(string) (map[string]string,error)
	Delete(string) error
	DeleteByPrefix(string) error
	Watch(string) ([]CEvent,error)
	WatchByPrefix(string) ([]CEvent,error)
}

type ITimer interface{
	Add(string,string,func(int64)) error
	Remove(string)
	Start()
	Stop()
}

type IScheduler interface {
	Dispatch(map[string]string,<-chan *CResult,chan<- *CTask)
}

type CDep struct {
	Job string
	Spec string
	Offset int64
	Status uint8
}

type CDepJob struct {
	Name string
	Spec string
	Offset int64
	Task string
	Params map[string]string
	Deps []CDep
}

type CSpecJob struct {
	Name string
	Spec string
	Task string
	Params map[string]string
}