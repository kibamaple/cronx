package core

import (
	context "context"
	clientv3 "github.com/coreos/etcd/clientv3"
	concurrency "github.com/coreos/etcd/clientv3/concurrency"
)

var newElection = func(dsn string,key string) IElection{
	client,c_err := clientv3.NewFromURL(dsn)
	if c_err != nil {
		return nil
	}
	session,s_err := concurrency.NewSession(client)
	if s_err != nil {
		return nil
	}
	election := concurrency.NewElection(session,key)
	return CElection{session,election}
}

type CElection struct {
	session *concurrency.Session
	election *concurrency.Election
}

func (this CElection) Campaign(val string) error{
	return this.election.Campaign(context.Background(), val)
}

func (this CElection) Resign() error{
	return this.election.Resign(context.Background())
}

func (this CElection) Destroy() error{
	client := this.session.Client()
	err := this.session.Close()
	client.Close()
	return err
}