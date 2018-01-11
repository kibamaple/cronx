package model

import (
	fmt "fmt"
	time "time"
	context "context"
	clientv3 "github.com/coreos/etcd/clientv3"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	common "github.com/kibamaple/cronx/common"
)

type IResultModel interface {
	Save(*common.CResult) error
	GetStatus(string,int64) (uint8,error)
}

type CResultModel struct {
	dsn string
	data string
	sep string
	timeout time.Duration
}

func (this CResultModel) Save(result *common.CResult) error {
	client,err := clientv3.NewFromURL(this.dsn)
	if err != nil {
		return err
	}
	defer client.Close()
	id,job,status := strconv.FormatInt(result.ID,Decimal),result.Job,result.Status
	key := strings.Join([]string{this.data,job,id},this.sep)
	val := strconv.Itoa(int(status))
	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	_, _err := client.Put(ctx, key, val)
	cancel()
	return _err
}

func (this CResultModel) GetStatus(job string,id int64) (uint8,error) {
	client,err := clientv3.NewFromURL(this.dsn)
	if err != nil {
		return err
	}
	defer client.Close()
	_id := strconv.FormatInt(id,Decimal)
	key := strings.Join([]string{this.data,job,_id},this.sep)

	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	res, _err := client.Get(ctx, key, clientv3.WithLimit(1))
	cancel()
	if _err != nil {
		return common.Failed,_err
	}
	if res.Count <= 0 {
		return common.Failed,fmt.Errorf("Error:%s is not found",key)
	}
	kv := res.Kvs[0]
	return uint8(strconv.Atoi(string(kv.Value))),_err
}