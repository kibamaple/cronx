package model

import (
	"encoding/json"
	time "time"
	context "context"
	clientv3 "github.com/coreos/etcd/clientv3"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	common "github.com/kibamaple/cronx/common"
)

type CEvent_Name uint8

type CEvent struct {
	Name CEvent_Name
	Key string
	Job *common.CJob
	time int64
}

const (
	Put CEvent_Name = iota
	Delete
)

type IEventModel interface {
	Watch(<-chan struct{}) (<-chan *CEvent,<-chan error)
}

type CEventModel struct {
	dsn string
	key string
	sep string
	timeout time.Duration
}

func (this CEventModel) buildEvent(name CEvent_Name,kv *mvccpb.KeyValue)(*CEvent,error){
	var job common.CJob
	err := json.Unmarshal([]byte(kv.Value), &job)
	if err != nil {
		return nil,err
	}
	return &CEvent{name,string(kv.Key),&job,kv.ModRevision},nil
}

func (this CEventModel) init(eventChan chan<- *CEvent,errorChan chan<- error) {
	client,err := clientv3.NewFromURL(this.dsn)
	if err != nil {
		errorChan <- err
		return
	}
	defer client.Close()
	
	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	res, _err := client.Get(ctx, this.key, clientv3.WithPrefix())
	cancel()
	if _err != nil {
		errorChan <- _err
		return
	}
	if res.Count <= 0 {
		return
	}

	for _, kv := range res.Kvs {
		event,event_err := this.buildEvent(Put,kv)
		if event_err != nil {
			errorChan <- event_err
		}else{
			eventChan <- event
		}
	}
}

func (this CEventModel) Watch(signalChan <-chan struct{}) (<-chan *CEvent,<-chan error){
	eventChan := make(chan *CEvent)
	errorChan := make(chan error)
	go this.run(signalChan,eventChan,errorChan)
	go this.init(eventChan,errorChan)
	return eventChan,errorChan
}

func (this CEventModel) run(signalChan <-chan struct{},eventChan chan<- *CEvent,errorChan chan<- error){
	for {
		if this.watch(signalChan,eventChan,errorChan) {
			break
		}
	}
	close(errorChan)
	close(eventChan)
}

func (this CEventModel) watch(signalChan <-chan struct{},eventChan chan<- *CEvent,errorChan chan<- error)(bool){
	
	client,err := clientv3.NewFromURL(this.dsn)
	if err != nil {
		errorChan <- err
		return true
	}
	defer client.Close()
	
	rchan := client.Watch(context.Background(), this.key, clientv3.WithPrefix())
	for {
		select {
			case _,ok:= <- signalChan:
				if !ok {
					return true
				}
			case res,ok:= <-rchan:
				if ok {
					for _, e := range res.Events {
						var name CEvent_Name
						kv := e.Kv
						if e.Type == mvccpb.PUT {
							name = Put
						}else{
							name = Delete
						}
						event,event_err := this.buildEvent(name,kv)
						if event_err != nil {
							errorChan <- event_err
						}else{
							eventChan <- event
						}
					}
				}else{
					break	
				}
		}
	}
	return false
}