package core

import (
	fmt "fmt"
	time "time"
	context "context"
	clientv3 "github.com/coreos/etcd/clientv3"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
)

const CStore_Timeout = "1m"

var newStore = func(dsn string) IStore{
	if timeout,err :=time.ParseDuration(CStore_Timeout);err == nil {
		return CStore{dsn,timeout}
	}
	return nil
}

type CStore struct {
	url string
	timeout time.Duration
}

func (this CStore) Put(key string,val string) error{
	client,err := clientv3.NewFromURL(this.url)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	_, _err := client.Put(ctx, key, val)
	cancel()
	return _err
}

func (this CStore) Get(key string) (string,error){
	client,err := clientv3.NewFromURL(this.url)
	if err != nil {
		return EmptyString,err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	res, _err := client.Get(ctx, key, clientv3.WithLimit(1))
	cancel()
	if _err != nil {
		return EmptyString,_err
	}
	if res.Count <= 0 {
		return EmptyString,fmt.Errorf("Error:%s is not found",key)
	}
	kv := res.Kvs[0]
	return string(kv.Value),_err
}

func (this CStore) GetByPrefix(prefix string) (map[string]string,error){
	client,err := clientv3.NewFromURL(this.url)
	if err != nil {
		return nil,err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	res, _err := client.Get(ctx, prefix,clientv3.WithPrefix())
	cancel()
	if _err != nil {
		return nil,_err
	}
	if res.Count <= 0 {
		return nil,fmt.Errorf("Error:%s is not found",prefix)
	}
	mappings := make(map[string]string)
	for _, kv := range res.Kvs {
		key := string(kv.Key)
		val := string(kv.Value)
		mappings[key] = val
	}
	return mappings,_err
}

func (this CStore) Delete(key string) error{
	client,err := clientv3.NewFromURL(this.url)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	_, _err := client.Delete(ctx, key, clientv3.WithLimit(1))
	cancel()
	return _err
}

func (this CStore) DeleteByPrefix(prefix string) error{
	client,err := clientv3.NewFromURL(this.url)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	_, _err := client.Delete(ctx, prefix,clientv3.WithPrefix())
	cancel()
	return _err
}

func (this CStore) Watch(key string) ([]CEvent,error){
	client,err := clientv3.NewFromURL(this.url)
	if err != nil {
		return nil,err
	}
	defer client.Close()

	var events []CEvent
	rchan := client.Watch(context.Background(), key)
	for res := range rchan {
		for _, e := range res.Events {
			var ctype CEvent_Name
			kv := e.Kv
			if e.Type == mvccpb.PUT {
				ctype = Put
			}else{
				ctype = Delete
			}
			event := CEvent{ctype,string(kv.Key),string(kv.Value)}
			events = append(events,event)
		}
	}
	return events,nil
}

func (this CStore) WatchByPrefix(prefix string) ([]CEvent,error){
	client,err := clientv3.NewFromURL(this.url)
	if err != nil {
		return nil,err
	}
	defer client.Close()

	var events []CEvent
	rchan := client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for res := range rchan {
		for _, e := range res.Events {
			var ctype CEvent_Name
			kv := e.Kv
			if e.Type == mvccpb.PUT {
				ctype = Put
			}else{
				ctype = Delete
			}
			event := CEvent{ctype,string(kv.Key),string(kv.Value)}
			events = append(events,event)
		}
	}
	return events,nil
}