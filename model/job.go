package model

import (
	fmt "fmt"
	time "time"
	context "context"
	clientv3 "github.com/coreos/etcd/clientv3"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	common "github.com/kibamaple/cronx/common"
)

type IJobModel interface {
	GetSpec() ([]common.CJob,error)
	WatchSpec(<-chan struct{}) (<-chan *common.CEvent,<-chan error)
}

type CJobModel struct {
	dsn string
	key string
	sep string
	timeout time.Duration
}

func (this CJobModel) GetSpec() ([]common.CJob,error) {
	client,err := clientv3.NewFromURL(this.dsn)
	if err != nil {
		return nil,err
	}
	defer client.Close()
	
	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	res, _err := client.Get(ctx, this.key, clientv3.WithPrefix())
	cancel()
	if _err != nil {
		return nil,_err
	}
	if res.Count <= 0 {
		return nil,nil
	}

	var jobs []common.CJob
	for _, kv := range res.Kvs {
		var job common.CJob
		_err = json.Unmarshal([]byte(kv.Value), &job)
		if _err != nil {
			continue
		}
		jobs = append(jobs,job)
	}
	return jobs,_err
}

func (this CJobModel) WatchSpec(abortChan <-chan struct{}) (<-chan *common.CEvent,<-chan error){
	eventChan := make(chan *common.CEvent)
	errorChan := make(chan error)
	go this.watchSpec(abortChan,eventChan,errorChan)
	return eventChan,errorChan
}

func (this CJobModel) watchSpec(abortChan <-chan struct{},eventChan chan<- *common.CEvent,errorChan chan<- error){
	
	client,err := clientv3.NewFromURL(this.dsn)
	if err != nil {
		errorChan <- err
	}
	defer client.Close()
	
	rchan := client.Watch(context.Background(), this.key, clientv3.WithPrefix())
	for {
		select {
			
		}
	}
	for res := range rchan {
		for _, e := range res.Events {
			var name common.CEvent_Name
			kv := e.Kv
			if e.Type == mvccpb.PUT {
				name = Put
			}else{
				name = Delete
			}
			event := &common.CEvent{name,string(kv.Key),string(kv.Value)}
			eventChan <- event
		}
	}
}

