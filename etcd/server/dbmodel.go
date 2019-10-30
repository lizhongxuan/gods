package main

import "../protos/datapb"

type openplatformModel int

var OpenplatformModel openplatformModel

func (openplatformModel) GetEnterpriseStatusByRelationId(relationid string) (*datapb.EnterpriseStatus, error) {
	return &datapb.EnterpriseStatus{
		AppRelationId:    1,
		MaxOnlineUsers:   2,
		FileMaxSize:      3,
		MaxEditFileCount: 4,
	}, nil
}
