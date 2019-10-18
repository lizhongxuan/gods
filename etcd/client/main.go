package main

import (
	"../protos/datapb"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	es, err := RequirEnterpriseStatus("1")
	if err != nil {
		glog.Info("RequirEnterpriseStatus err:", RequirEnterpriseStatus)
		return
	}
	glog.Info("RequirEnterpriseStatus es:", es)
}

func RequirEnterpriseStatus(relationid string) (*datapb.EnterpriseStatus, error) {
	ctx := context.Background()
	pbin := &datapb.GetEnterpriseStatusReq{
		AppRelationId: relationid,
	}
	eps, err := DB.EnterpriseStatusGet(ctx, pbin)
	if err != nil {
		glog.Errorf("error : DB EnterpriseStatusGet , %s", err.Error())
		return nil, err
	}
	return eps, nil
}
