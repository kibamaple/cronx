package core

import (
    "testing"
	"strconv"
	convey "github.com/smartystreets/goconvey/convey"
	gostub "github.com/prashantv/gostub"
)

func Test_CExecutor(t *testing.T) {

	convey.Convey("test listen", t, func() {

		convey.Convey("on task",func(c convey.C) {
			path,id,job,name,single,status := "path",int64(1),"job","name",true,Success
			params := map[string]string {"name1":"value1","name2":"value2","name3":"value3"}
			task := CTask{id,job,name,params,single}
			executor := CExecutor{path}
			input := make(chan *CTask)
			output := make(chan *CResult)
			defer func(){
				close(input)
				close(output)
			}()

			stubs := gostub.Stub(&execute, func(t_path string,t_name string,t_params map[string]string,t_args ...string) uint8 {
				c.So(path,convey.ShouldEqual,t_path)
				c.So(name,convey.ShouldEqual,t_name)
				c.So(params,convey.ShouldResemble,t_params)
				c.So(len(t_args),convey.ShouldEqual,1)
				c.So(t_args,convey.ShouldContain,strconv.FormatInt(id,Execute_Base))
				return status
			})
			defer stubs.Reset()

			executor.Listen(input,output);
			input <- &task
			result := <- output

			c.So(result.ID,convey.ShouldEqual,id)
			c.So(result.Job,convey.ShouldEqual,job)
			c.So(result.Status,convey.ShouldEqual,status)
			c.So(result.Single,convey.ShouldEqual,single)
		})
		
		convey.Convey("on multi task",func(c convey.C) {
			path := "path"
			id1,job1,name1,single1,status1 := int64(1),"job1","name1",true,Success
			params1 := map[string]string {"name11":"value11","name12":"value12","name13":"value13"}
			task1 := CTask{id1,job1,name1,params1,single1}
			id2,job2,name2,single2,status2 := int64(2),"job2","name2",false,Failed
			params2 := map[string]string {"name21":"value21","name22":"value22","name23":"value23"}
			task2 := CTask{id2,job2,name2,params2,single2}
			executor := CExecutor{path}
			input := make(chan *CTask)
			output := make(chan *CResult)
			defer func(){
				close(input)
				close(output)
			}()

			stubs := gostub.Stub(&execute, func(t_path string,t_name string,t_params map[string]string,t_args ...string) uint8 {
				c.So(len(t_args),convey.ShouldEqual,1)
				if t_name == name1 {
					c.So(path,convey.ShouldEqual,t_path)
					c.So(params1,convey.ShouldResemble,t_params)
					c.So(t_args,convey.ShouldContain,strconv.FormatInt(id1,Execute_Base))
					return status1
				}else{
					c.So(path,convey.ShouldEqual,t_path)
					c.So(name2,convey.ShouldEqual,t_name)
					c.So(params2,convey.ShouldResemble,t_params)
					c.So(t_args,convey.ShouldContain,strconv.FormatInt(id2,Execute_Base))
					return status2
				}
			})
			defer stubs.Reset()

			executor.Listen(input,output);
			input <- &task1
			input <- &task2
			result1 := <- output
			result2 := <- output

			if result1.ID == id1 {
				c.So(result1.Job,convey.ShouldEqual,job1)
				c.So(result1.Status,convey.ShouldEqual,status1)
				c.So(result1.Single,convey.ShouldEqual,single1)
				c.So(result2.ID,convey.ShouldEqual,id2)
				c.So(result2.Job,convey.ShouldEqual,job2)
				c.So(result2.Status,convey.ShouldEqual,status2)
				c.So(result2.Single,convey.ShouldEqual,single2)
			}else{
				c.So(result1.ID,convey.ShouldEqual,id2)
				c.So(result1.Job,convey.ShouldEqual,job2)
				c.So(result1.Status,convey.ShouldEqual,status2)
				c.So(result1.Single,convey.ShouldEqual,single2)
				c.So(result2.ID,convey.ShouldEqual,id1)
				c.So(result2.Job,convey.ShouldEqual,job1)
				c.So(result2.Status,convey.ShouldEqual,status1)
				c.So(result2.Single,convey.ShouldEqual,single1)
			}
		})

	})

}

func Test_Execute(t *testing.T) {
	
	convey.Convey("test execute",t,func() {

		convey.Convey("success",func(c convey.C) {
			path,name := ".","./execute_test.sh"
			params := map[string]string {"STATUS":"success"}
			status := execute(path,name,params)
			c.So(status,convey.ShouldEqual,Success)
		})

		convey.Convey("failed",func(c convey.C) {
			path,name := ".","./execute_test.sh"
			params := map[string]string {"STATUS":"failed"}
			status := execute(path,name,params)
			c.So(status,convey.ShouldEqual,Failed)
		})

	})

}