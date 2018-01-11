package module

import (
	common "github.com/kibamaple/cronx/common"
	model "github.com/kibamaple/cronx/model"
)

var computeTime = func(now int64,spec string,offset int64) (int64,error) {
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

type CTrigger struct {
	resultModel *model.IResultModel
	relationModel *model.IRelationModel
	common.CErrorProvider
	common.CTaskProvider
}

func (this CTrigger) OnResult(result *common.CResult) {
	rels,rels_err := this.relationModel.GetAll(result.Job,result.Status)
	if err != nil {
		this.EmitError(rels_err)
		return
	}

	for _,rel := range rels{
		job_id,job_err := computeTime(result.ID,rel.job.Spec,rel.Offset)
		if job_err != nil {
			this.EmitError(job_err)
			continue
		}

		count := 0
		for _,dep:= range rel.Deps{
			dep_id,dep_err := computeTime(job_id,dep.Spec,dep.Offset)
			if dep_err !=nil {
				this.EmitError(dep_err)
				break
			}

			dep_status,dep_status_err := this.resultModel.GetStatus(dep.Job,dep_id)
			if dep_status_err != nil {
				this.EmitError(dep_status_err)
				break
			}
			if dep_status != dep.Status {
				break
			}
			count ++
		}

		if count >= len(rel.Deps) {
			task := &common.CTask{job_id,rel.job.Name,rel.job.Task,rel.job.Params,rel.job.Status}
			this.EmitTask(task)
		}
	}
}

type CConcurrentTrigger struct {
	CTrigger
}

func (this CTrigger) OnResult(result *common.CResult) {
	go this.CTrigger.OnResult(result)
}

type CHandleTrigger struct {
	CTrigger
	common.CResultListener
	common.CConcurrent
}

func (this CHandleTrigger) Process(abortChan chan struct{}){
	for {
		select	{
			case result := <-this.ResultChan:
				this.CTrigger.OnResult(result)
			case <-abortChan:
				return
		}
	}
}