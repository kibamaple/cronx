package core

import (
    "testing"
	"sync"
	"strconv" 
	"strings"
	"errors"
	"encoding/json"
	convey "github.com/smartystreets/goconvey/convey"
	gomock "github.com/golang/mock/gomock"
	gostub "github.com/prashantv/gostub"
)

func Test_CScheduler(t *testing.T) {

	convey.Convey("test dispatch dep", t, func() {

		convey.Convey("input single result",func(c convey.C) {
			id,job,status,single := int64(1),"job",Success,true
			dsn,data_value,dep_value := "dsn","data","dep"
			key := strings.Join([]string{data_value,job,strconv.FormatInt(id,CScheduler_Base)},CScheduler_Sep)
			context := map[string]string {CScheduler_DSN:dsn,CScheduler_Data:data_value,CScheduler_Dep:dep_value}
			input := make(chan *CResult)
			output := make(chan *CTask)
			defer func(){
				close(input)
				close(output)
			}()
			
			ctrl := gomock.NewController(t)
            defer ctrl.Finish()

			store := NewMockIStore(ctrl)
			var wg sync.WaitGroup
			wg.Add(1)
			store.EXPECT().Put(key,strconv.Itoa(int(status))).Times(1).Return(nil).Do(func(key string, value string) {
				wg.Done()
			})
			store.EXPECT().GetByPrefix(gomock.Any()).Return(nil,nil).Times(0)

			newStore_stubs := gostub.Stub(&newStore, func(_dsn string) (IStore) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				return store
			})
			defer newStore_stubs.Reset()

			scheduler := CScheduler{}
			scheduler.Dispatch(context,input,output)

			result := &CResult{id,job,status,single}
			input <- result
			wg.Wait()
		})

		convey.Convey("input result with non dep",func(c convey.C) {
			id,job,status,single := int64(1),"job",Success,false
			dsn,data_value,dep_value := "dsn","data","dep"
			key := strings.Join([]string{data_value,job,strconv.FormatInt(id,CScheduler_Base)},CScheduler_Sep)
			prefix := strings.Join([]string{dep_value,strconv.Itoa(int(status)),job},CScheduler_Sep)
			context := map[string]string {CScheduler_DSN:dsn,CScheduler_Data:data_value,CScheduler_Dep:dep_value}
			input := make(chan *CResult)
			output := make(chan *CTask)
			defer func(){
				close(input)
				close(output)
			}()
			
			ctrl := gomock.NewController(t)
            defer ctrl.Finish()

			var wg sync.WaitGroup
			wg.Add(1)
			store := NewMockIStore(ctrl)
			store.EXPECT().Put(key,strconv.Itoa(int(status))).Return(nil).Times(1)
			store.EXPECT().GetByPrefix(prefix).Return(nil,nil).Times(1).Do(func(key string) {
				wg.Done()
			})
			
			newStore_stubs := gostub.Stub(&newStore, func(_dsn string) (IStore) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				return store
			})
			defer newStore_stubs.Reset()

			scheduler := CScheduler{}
			scheduler.Dispatch(context,input,output)

			result := &CResult{id,job,status,single}
			input <- result
			wg.Wait()
		})

		convey.Convey("input result with dep",func(c convey.C) {
			id,job,status,single := int64(1),"job",Success,false
			dsn,data_value,dep_value := "dsn","data","dep"
			key := strings.Join([]string{data_value,job,strconv.FormatInt(id,CScheduler_Base)},CScheduler_Sep)
			prefix := strings.Join([]string{dep_value,strconv.Itoa(int(status)),job},CScheduler_Sep)
			context := map[string]string {CScheduler_DSN:dsn,CScheduler_Data:data_value,CScheduler_Dep:dep_value}
			job_name,job_spec,job_offset,job_task := "job_name","job_spec",int64(-2),"job_task"
			job_params := map[string]string {"job_name":"job_value","job_name2":"job_value2","job_name3":"job_value3"}
			dep_name,dep_spec,dep_status,dep_offset := "dep_name","dep_spec",Success,int64(2)
			deps := []CDep{CDep{dep_name,dep_spec,dep_offset,dep_status}}	
			data,err := json.Marshal(CDepJob{job_name,job_spec,job_offset,job_task,job_params,deps})
			if err!=nil {
				t.Error(err)
			}
			jobs := map[string]string{"key":string(data[:])}
			job_id,dep_id := int64(2),int64(3)
			dep_key := strings.Join([]string{data_value,dep_name,strconv.FormatInt(dep_id,CScheduler_Base)},CScheduler_Sep)
			input := make(chan *CResult)
			output := make(chan *CTask)
			defer func(){
				close(input)
				close(output)
			}()

			ctrl := gomock.NewController(t)
            defer ctrl.Finish()
			store := NewMockIStore(ctrl)
			store.EXPECT().Put(key,strconv.Itoa(int(status))).Return(nil).Times(1)
			store.EXPECT().GetByPrefix(prefix).Return(jobs,nil).Times(1)
			store.EXPECT().Get(dep_key).Return(strconv.Itoa(int(dep_status)),nil).Times(1)
			
			newStore_stubs := gostub.Stub(&newStore, func(_dsn string) (IStore) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				return store
			})
			defer newStore_stubs.Reset()
			compute_stubs := gostub.Stub(&compute, func(now int64,spec string,offset int64) (int64,error) {
				if now == id {
					c.So(job_spec,convey.ShouldEqual,spec)
					c.So(job_offset,convey.ShouldEqual,offset)
					return job_id,nil
				}else{
					c.So(job_id,convey.ShouldEqual,now)
					c.So(dep_spec,convey.ShouldEqual,spec)
					c.So(dep_offset,convey.ShouldEqual,offset)
					return dep_id,nil
				}
			})
			defer compute_stubs.Reset()

			scheduler := CScheduler{}
			scheduler.Dispatch(context,input,output)

			result := &CResult{id,job,status,single}
			input <- result
			task := <- output

			c.So(task.ID,convey.ShouldEqual,job_id)
			c.So(task.Job,convey.ShouldEqual,job_name)
			c.So(task.Name,convey.ShouldEqual,job_task)
			c.So(task.Single,convey.ShouldEqual,false)
			c.So(task.Params,convey.ShouldResemble,job_params)
		})

		convey.Convey("input result with failed dep",func(c convey.C) {
			id,job,status,single := int64(1),"job",Success,false
			dsn,data_value,dep_value := "dsn","data","dep"
			key := strings.Join([]string{data_value,job,strconv.FormatInt(id,CScheduler_Base)},CScheduler_Sep)
			prefix := strings.Join([]string{dep_value,strconv.Itoa(int(status)),job},CScheduler_Sep)
			context := map[string]string {CScheduler_DSN:dsn,CScheduler_Data:data_value,CScheduler_Dep:dep_value}
			job_name,job_spec,job_offset,job_task := "job_name","job_spec",int64(-2),"job_task"
			job_params := map[string]string {"job_name":"job_value","job_name2":"job_value2","job_name3":"job_value3"}
			dep_name,dep_spec,dep_status,dep_offset := "dep_name","dep_spec",Success,int64(2)
			deps := []CDep{CDep{dep_name,dep_spec,dep_offset,dep_status}}
			data,err := json.Marshal(CDepJob{job_name,job_spec,job_offset,job_task,job_params,deps})
			if err!=nil {
				t.Error(err)
			}
			jobs := map[string]string{"key":string(data[:])}
			job_id,dep_id := int64(2),int64(3)
			dep_key := strings.Join([]string{data_value,dep_name,strconv.FormatInt(dep_id,CScheduler_Base)},CScheduler_Sep)
			input := make(chan *CResult)
			output := make(chan *CTask)
			defer func(){
				close(input)
				close(output)
			}()

			var wg sync.WaitGroup
			wg.Add(1)
			ctrl := gomock.NewController(t)
            defer ctrl.Finish()
			store := NewMockIStore(ctrl)
			put := store.EXPECT().Put(key,strconv.Itoa(int(status))).Return(nil).Times(1)
			store.EXPECT().GetByPrefix(prefix).Return(jobs,nil).Times(1).After(put)
			store.EXPECT().Get(dep_key).Return(strconv.Itoa(int(Failed)),nil).Times(1).Do(func(key string) {
				wg.Done()
			})

			newStore_stubs := gostub.Stub(&newStore, func(_dsn string) (IStore) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				return store
			})
			defer newStore_stubs.Reset()
			compute_stubs := gostub.Stub(&compute, func(now int64,spec string,offset int64) (int64,error) {
				if now == id {
					c.So(job_spec,convey.ShouldEqual,spec)
					c.So(job_offset,convey.ShouldEqual,offset)
					return job_id,nil
				}else{
					c.So(job_id,convey.ShouldEqual,now)
					c.So(dep_spec,convey.ShouldEqual,spec)
					c.So(dep_offset,convey.ShouldEqual,offset)
					return dep_id,nil
				}
			})
			defer compute_stubs.Reset()

			scheduler := CScheduler{}
			scheduler.Dispatch(context,input,output)

			result := &CResult{id,job,status,single}
			input <- result
			wg.Wait()
		})

	})

	convey.Convey("test dispatch spec", t, func() {

		convey.Convey("dispatch by init",func(c convey.C) {
			task_id,task_name,task_spec,task_job := int64(1),"task_name","task_spec","task_job"
			task_params := map[string]string {"task_name":"task_value","task_name2":"task_value2","task_name3":"task_value3"}
			dsn,elect_value,spec_value := "dsn","elect","spec"
			key,ctime :="key",CSpecJob{task_job,task_spec,task_name,task_params}
			value,err := json.Marshal(ctime)
			if err!=nil {
				t.Error(err)
			}
			context := map[string]string {CScheduler_DSN:dsn,CScheduler_Election:elect_value,CScheduler_Spec:spec_value}
			events := map[string]string {key:string(value[:])}
			input := make(chan *CResult)
			output := make(chan *CTask)
			defer func(){
				close(input)
				close(output)
			}()

			ctrl := gomock.NewController(t)
            defer ctrl.Finish()
			timer := NewMockITimer(ctrl)
			election := NewMockIElection(ctrl)
			store := NewMockIStore(ctrl)
			campaign := election.EXPECT().Campaign(spec_value).Return(nil).Times(1)
			get := store.EXPECT().GetByPrefix(spec_value).Return(events,nil).Times(1).After(campaign)
			timer.EXPECT().Add(key,task_spec,gomock.Any()).Times(1).After(get).Do(func(_key string,_spec string,_cmd func(_time int64)) {
				_cmd(task_id)
			})
			timer.EXPECT().Start().Times(1)
			timer.EXPECT().Stop().Times(1)
			store.EXPECT().WatchByPrefix(spec_value).Return(nil,errors.New("close")).Times(1).After(get)
			election.EXPECT().Destroy().Times(1)
			election.EXPECT().Resign().Times(1)

			newElection_stubs := gostub.Stub(&newElection, func(_dsn string, _key string) (IElection) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				c.So(elect_value,convey.ShouldEqual,_key)
				return election
			})
			defer newElection_stubs.Reset()
			newStore_stubs := gostub.Stub(&newStore, func(_dsn string) (IStore) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				return store
			})
			defer newStore_stubs.Reset()
			newTimer_stubs := gostub.Stub(&newTimer, func() (ITimer) {
				return timer
			})
			defer newTimer_stubs.Reset()

			scheduler := CScheduler{}
			scheduler.Dispatch(context,input,output)

			task := <- output
			c.So(task.ID,convey.ShouldEqual,task_id)
			c.So(task.Job,convey.ShouldEqual,task_job)
			c.So(task.Name,convey.ShouldEqual,task_name)
			c.So(task.Single,convey.ShouldEqual,false)
			c.So(task.Params,convey.ShouldResemble,task_params)
		})

		convey.Convey("dispatch by update",func(c convey.C) {
			task_id,task_name,task_spec,task_job := int64(1),"task_name","task_spec","task_job"
			task_params := map[string]string {"task_name":"task_value","task_name2":"task_value2","task_name3":"task_value3"}
			dsn,elect_value,spec_value := "dsn","elect","spec"
			put_key,delete_key,ctime := "put_key","delete_key",CSpecJob{task_job,task_spec,task_name,task_params}
			value,err := json.Marshal(ctime)
			if err!=nil {
				t.Error(err)
			}
			put_event,delete_event := CEvent{Put,put_key,string(value[:])},CEvent{Delete,delete_key,string(value[:])}
			events := []CEvent{put_event,delete_event}
			context := map[string]string {CScheduler_DSN:dsn,CScheduler_Election:elect_value,CScheduler_Spec:spec_value}
			input := make(chan *CResult)
			output := make(chan *CTask)
			defer func(){
				close(input)
				close(output)
			}()

			ctrl := gomock.NewController(t)
            defer ctrl.Finish()
			timer := NewMockITimer(ctrl)
			election := NewMockIElection(ctrl)
			store := NewMockIStore(ctrl)
			campaign := election.EXPECT().Campaign(spec_value).Return(nil).Times(1)
			get := store.EXPECT().GetByPrefix(spec_value).Return(nil,nil).Times(1).After(campaign)
			watch := store.EXPECT().WatchByPrefix(spec_value).Return(events,nil).Times(1).After(get)
			clean := timer.EXPECT().Remove(put_key).Times(1).After(watch)
			timer.EXPECT().Add(put_key,task_spec,gomock.Any()).Times(1).After(clean).Do(func(_key string,_spec string,_cmd func(int64)) {
				_cmd(task_id)
			})
			timer.EXPECT().Remove(delete_key).Times(1).After(watch)
			store.EXPECT().WatchByPrefix(spec_value).Return(nil,errors.New("close")).Times(1).After(campaign)
			timer.EXPECT().Start().Times(1)
			timer.EXPECT().Stop().Times(1)
			election.EXPECT().Destroy().Times(1)
			election.EXPECT().Resign().Times(1)

			newElection_stubs := gostub.Stub(&newElection, func(_dsn string, _key string) (IElection) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				c.So(elect_value,convey.ShouldEqual,_key)
				return election
			})
			defer newElection_stubs.Reset()
			newStore_stubs := gostub.Stub(&newStore, func(_dsn string) (IStore) {
				c.So(dsn,convey.ShouldEqual,_dsn)
				return store
			})
			defer newStore_stubs.Reset()
			newTimer_stubs := gostub.Stub(&newTimer, func() (ITimer) {
				return timer
			})
			defer newTimer_stubs.Reset()

			scheduler := CScheduler{}
			scheduler.Dispatch(context,input,output)

			task := <- output
			c.So(task.ID,convey.ShouldEqual,task_id)
			c.So(task.Job,convey.ShouldEqual,task_job)
			c.So(task.Name,convey.ShouldEqual,task_name)
			c.So(task.Single,convey.ShouldEqual,false)
			c.So(task.Params,convey.ShouldResemble,task_params)
		})

	})

}