package model

import (
	fmt "fmt"
	time "time"
	context "context"
	clientv3 "github.com/coreos/etcd/clientv3"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	common "github.com/kibamaple/cronx/common"
)

type IRelationModel interface {
	GetAll(string,uint) ([]*common.CRelation,error)
}

type CRelationModel struct {
	dsn string
	rel string
	sep string
	timeout time.Duration
}

func (this CRelationModel) GetAll(job string,status uint8) ([]*common.CRelation,error) {
	client,err := clientv3.NewFromURL(this.dsn)
	if err != nil {
		return err
	}
	defer client.Close()
	key := strings.Join([]string{this.rel,strconv.Itoa(int(status)),job},this.sep)
	
	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	res, _err := client.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if _err != nil {
		return nil,_err
	}
	if res.Count <= 0 {
		return nil,fmt.Errorf("Error:%s is not found",key)
	}

	var rels []*common.CRelation
	for _, kv := range res.Kvs {
		var rel common.CRelation
		_err = json.Unmarshal([]byte(kv.Value), &rel)
		if _err != nil {
			continue
		}
		rels = append(rels,&rel)
	}
	return rels,_err
}