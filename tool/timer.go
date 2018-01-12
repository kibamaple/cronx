package tool

import (
	fmt "fmt"
	time "time"
	cron "github.com/kibamaple/cron"
)

type ITimer interface{
	Add(string,string,func(int64)) error
	Remove(string)
	Start()
	Stop()
}

var newTimer = func() ITimer{
	timer := cron.New();
	mappings := make(map[string]cron.EntryID)
	return CTimer{timer,mappings}
}

var compute = func(now int64,spec string,offset int64) (int64,error) {
	if offset == 0 {
		return now,nil;
	}

	sche,err := cron.Parse(spec)
	if err!= nil {
		return 0,err
	}

	_time := time.Unix(now,0)
	for ;offset>0;offset--{
		_time = sche.Next(_time)
	}

	for ;offset<0;offset++{
		_time = sche.Prev(_time)
	}
	return _time.Unix(),err
}

type CTimer struct {
	timer *cron.Cron
	mappings map[string]cron.EntryID
}

func (this CTimer) Add(key string,spec string,fn func(int64)) error {
	if _id,_ok:=this.mappings[key]; _ok{
		return fmt.Errorf("Error:%s is exists",_id)
	}

	id,err := this.timer.AddFunc(spec,func(t time.Time){
		fn(t.Unix())
	})

	if err!=nil {
		return err
	}

	this.mappings[key] = id
	return nil
}

func (this CTimer) Remove(key string) {
	if _id,_ok:=this.mappings[key]; _ok{
		this.timer.Remove(_id)
		delete(this.mappings,key)
	}
}

func (this CTimer) Start() {
	this.timer.Start()
}

func (this CTimer) Stop() {
	this.timer.Stop()
}