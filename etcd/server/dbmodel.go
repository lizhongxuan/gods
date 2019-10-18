package main

import "../protos/datapb"

type openplatformModel int

var OpenplatformModel openplatformModel

func (openplatformModel) GetEnterpriseStatusByRelationId(relationid string) (*datapb.EnterpriseStatus, error) {

	eps := &datapb.EnterpriseStatus{}
	//query := `SELECT app_relation_id,max_online_users,file_max_size,max_edit_file_count FROM enterprise_status WHERE app_relation_id=?`
	//if err := database.QueryRow(query, relationid).Scan(&eps.AppRelationId, &eps.MaxOnlineUsers, &eps.FileMaxSize, &eps.MaxEditFileCount); err != nil {
	//	klog.Errorf("query enterprise_status err:%v", err)
	//	return nil, err
	//}
	eps.AppRelationId = 1
	eps.MaxEditFileCount = 2
	eps.FileMaxSize = 3
	eps.MaxOnlineUsers = 4
	return eps, nil
}
