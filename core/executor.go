package core

import (
	"os"
	"os/exec"
	"strconv"
)

const ENV_Equal = "="
const Execute_Base = 10

var execute = func(path string,name string,params map[string]string,args ...string) uint8 {
	cmd := exec.Command(name)
	env := os.Environ()
	for key,value := range params {
		env = append(env,key+ENV_Equal+value)
	}
	cmd.Dir = path
	cmd.Env = env
	cmd.Args = args
	if err := cmd.Run(); err == nil {
		return Success
	}
	return Failed
}

type CExecutor struct {
	Path string
}

func (this CExecutor) onTask(task *CTask,output chan<- *CResult) {
	status := execute(this.Path,task.Name,task.Params,strconv.FormatInt(task.ID,Execute_Base))
	result := &CResult{task.ID,task.Job,status,task.Single}
	output <- result
}

func (this CExecutor) listen(input <-chan *CTask,output chan<- *CResult) {
	for task:= range input {
		go this.onTask(task,output)
	}
}

func (this CExecutor) Listen(input <-chan *CTask,output chan<- *CResult) {
	go this.listen(input,output)
}