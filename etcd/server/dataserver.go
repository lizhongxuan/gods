package main

import (
	"../protos/datapb"
	"context"
)

type DataConf struct {
	//DbAddr  string
	//RdsAddr string
}

type DatabaseServer struct {
	dataConf DataConf
}

func (this *DatabaseServer) EnterpriseStatusGet(ctx context.Context, in *datapb.GetEnterpriseStatusReq) (*datapb.EnterpriseStatus, error) {
	eps, err := OpenplatformModel.GetEnterpriseStatusByRelationId(in.AppRelationId)
	if err != nil {
		return nil, err
	}
	return eps, nil
}
