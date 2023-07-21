// Code generated by protoc-gen-permission. DO NOT EDIT.
package userv1

import commonpb "github.com/openfms/protos/gen/common/v1"

type MethodStreamInfo struct {
	Roles   []commonpb.UserRole `json:"roles"`
	RPCName string              `json:"rpc_name"`
}
type ServicePermission struct {
	ServiceName   string                       `json:"service_name"`
	MethodStreams map[string]*MethodStreamInfo `json:"method_streams"`
}

var UserServicePermission = &ServicePermission{
	MethodStreams: map[string]*MethodStreamInfo{
		"/user.v1.UserService/CreateUser": &MethodStreamInfo{
			RPCName: "CreateUser",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN},
		},
		"/user.v1.UserService/UpdateUser": &MethodStreamInfo{
			RPCName: "UpdateUser",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN, commonpb.UserRole_USER_ROLE_NORMAL},
		},
		"/user.v1.UserService/UploadAvatar": &MethodStreamInfo{
			RPCName: "UploadAvatar",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN, commonpb.UserRole_USER_ROLE_NORMAL},
		},
	},
	ServiceName: "user.v1.UserService",
}
