// Code generated by protoc-gen-permission. DO NOT EDIT.
package usersv1

import commonpb "github.com/irisco88/protos/gen/common/v1"

type MethodStreamInfo struct {
	Roles   []commonpb.UserRole `json:"roles"`
	RPCName string              `json:"rpc_name"`
}
type ServicePermission struct {
	ServiceName   string                       `json:"service_name"`
	MethodStreams map[string]*MethodStreamInfo `json:"method_streams"`
}

var UsersServicePermission = &ServicePermission{
	MethodStreams: map[string]*MethodStreamInfo{"/users.v1.UsersService/GetUsers": &MethodStreamInfo{
		RPCName: "GetUsers",
		Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN},
	}},
	ServiceName: "users.v1.UsersService",
}
