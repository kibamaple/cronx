package module

import (
	common "github.com/kibamaple/cronx/common"
	"os"
	"os/exec"
	"strconv"
)

const ENV_Equal = "="
const Execute_Base = 10

var runTask = func(path string,name string,params map[string]string,args ...string) err {
	cmd := exec.Command(name)
	env := os.Environ()
	for key,value := range params {
		env = append(env,key+ENV_Equal+value)
	}
	cmd.Dir = path
	cmd.Env = env
	cmd.Args = args
	err := cmd.Run()
	return err
}

type CExecutor struct {
	path string
	common.CErrorProvider
	common.CResultProvider
}

func (this CExecutor) OnTask(task *common.CTask) {
	status := task.Status
	id := strconv.FormatInt(task.ID,Decimal)
	
	if err := runTask(this.path,task.Name,task.Params,id); err != nil {
		status = Failed
		this.EmitError(err)
	}
	result := &common.CResult{task.ID,task.Job,status,task.Single}
	this.EmitResult(result)
}

type CConcurrentExecutor struct {
	CExecutor
}

func (this CExecutor) OnTask(task *common.CTask) {
	go this.CExecutor.OnTask(task)
}

type CHandleExecutor struct {
	CExecutor
	common.CTaskListener
	common.CConcurrent
}

func (this CHandleExecutor) Process(abortChan chan struct{}){
	for {
		select	{
			case task := <-this.TaskChan:
				this.CExecutor.OnTask(task)
			case <-abortChan:
				return
		}
	}
}