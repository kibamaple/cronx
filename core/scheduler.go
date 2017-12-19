package core

import (
	"strings"
	"strconv"
	"encoding/json"
)

const CScheduler_DSN = "dsn"
const CScheduler_Data = "data"
const CScheduler_Dep = "dep"
const CScheduler_Election = "election"
const CScheduler_Spec = "spec"
const CScheduler_Unknown = ""
const CScheduler_Sep = ":"
const CScheduler_Base = 10

type CScheduler struct {
}

func (this CScheduler) compute(store IStore,id int64,data string,value string,output chan<- *CTask){
	var job CDepJob
	
	if err := json.Unmarshal([]byte(value), &job); err != nil {
		panic(err)
	}
	
	job_id,job_err := compute(id,job.Spec,job.Offset)
	if job_err != nil {
		panic(job_err)
	}
	count := 0
	for _,dep:= range job.Deps{
		dep_id,dep_err := compute(job_id,dep.Spec,dep.Offset)
		if dep_err !=nil {
			panic(dep_err)
		}
		dep_data,data_err := store.Get(strings.Join([]string{data,dep.Job,strconv.FormatInt(dep_id,CScheduler_Base)},CScheduler_Sep))
		if data_err != nil {
			break
		}
		if dep_status,dep_err := strconv.Atoi(dep_data); dep_err != nil || (uint(dep_status) != uint(dep.Status)) {
			break
		}
		count ++
	}
	if count >= len(job.Deps) {
		task := &CTask{job_id,job.Name,job.Task,job.Params,false}
		output <- task
	}
}

func (this CScheduler) onResult(store IStore,data string,dep string,result *CResult,output chan<- *CTask){
	id,job,status,single := result.ID,result.Job,result.Status,result.Single
	store.Put(strings.Join([]string{data,job,strconv.FormatInt(id,CScheduler_Base)},CScheduler_Sep),strconv.Itoa(int(status)))
	if !single {
		if jobs,err := store.GetByPrefix(strings.Join([]string{dep,strconv.Itoa(int(status)),job},CScheduler_Sep)); err == nil{
			for _,value:= range jobs{
				go this.compute(store,id,data,value,output)
			}
		}
	}
}

func (this CScheduler) dispatchDep(context map[string]string,input <-chan *CResult,output chan<- *CTask){
	var store IStore
	
	if dsn,dsn_ok := context[CScheduler_DSN]; dsn_ok{
		store = newStore(dsn)
	}

	data,data_ok := context[CScheduler_Data]
	dep,dep_ok := context[CScheduler_Dep]
	if store != nil && data_ok && dep_ok {
		for result:= range input{
			go this.onResult(store,data,dep,result,output)
		}
	}
}

func (this CScheduler) newSpecJob(timer ITimer,key string,value string,output chan<- *CTask){
	var job CSpecJob
	
	if err := json.Unmarshal([]byte(value), &job); err != nil {
		panic(err)
	}

	time_err := timer.Add(key,job.Spec,func (id int64) {
		task := &CTask{id,job.Name,job.Task,job.Params,false}
		output <- task
	})
	if time_err != nil {
		panic(time_err)
	}
}

func (this CScheduler) dispatchSpec(context map[string]string,output chan<- *CTask) {
	var store IStore
	var election IElection
	
	dsn,dsn_ok := context[CScheduler_DSN];
	if dsn_ok {
		store = newStore(dsn)
	}
	if elect,elect_ok := context[CScheduler_Election]; elect_ok{
		election = newElection(dsn,elect)
		defer election.Destroy()
	}
	if spec,spec_ok := context[CScheduler_Spec];store != nil && election != nil && spec_ok {
		
		if err := election.Campaign(spec); err != nil {
			return
		}
		defer election.Resign()
		
		kvs,kvs_err := store.GetByPrefix(spec)
		if kvs_err != nil {
			return
		}

		timer := newTimer()
		for key,value := range kvs {
			this.newSpecJob(timer,key,value,output)
		}
		timer.Start()
		defer timer.Stop()

		for {
			events,events_err := store.WatchByPrefix(spec)
			if events_err != nil {
				break;
			}
			for _,event := range events {
				if timer.Remove(event.Key); event.Name == Put {
					this.newSpecJob(timer,event.Key,event.Value,output)
				}
			}
		}

	}
}

func (this CScheduler) Dispatch(context map[string]string,input <-chan *CResult,output chan<- *CTask) {
	go this.dispatchDep(context,input,output)
	go this.dispatchSpec(context,output)
}