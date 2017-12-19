package main

import (
	os "os"
	signal "os/signal"
	core "github.com/kibamaple/cronx/core"
)

func main() {
	tasks := make(chan *core.CTask)
	results := make(chan *core.CResult)

	path := os.Getenv("JOBKEEPER_EXEC")
	executor := core.CExecutor{path}
	executor.Listen(tasks,results);

	ctx := make(map[string]string)
	ctx[core.CScheduler_DSN] = os.Getenv("JOBKEEPER_DSN")
	ctx[core.CScheduler_Election] = os.Getenv("JOBKEEPER_ELEC")
	ctx[core.CScheduler_Data] = os.Getenv("JOBKEEPER_DATA")
	ctx[core.CScheduler_Dep] = os.Getenv("JOBKEEPER_DEP")
	ctx[core.CScheduler_Spec] = os.Getenv("JOBKEEPER_SPEC")

	scheduler := core.CScheduler{}
	scheduler.Dispatch(ctx,results,tasks)

	signals := make(chan os.Signal)
	signal.Notify(signals,os.Kill)
	<- signals
}