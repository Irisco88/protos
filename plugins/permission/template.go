package main

import commonpb "github.com/irisco88/protos/gen/common/v1"

type ServicePermission struct {
	ServiceName   string                       `json:"service_name"`
	MethodStreams map[string]*MethodStreamInfo `json:"method_streams"`
}

type MethodStreamInfo struct {
	Roles   []commonpb.UserRole `json:"roles"`
	RPCName string              `json:"rpc_name"`
}

var UserServicePermissions = &ServicePermission{
	ServiceName: "user.v1.UserService",
	MethodStreams: map[string]*MethodStreamInfo{
		"/user.v1.UserService/signup": {
			RPCName: "signup",
			Roles: []commonpb.UserRole{
				commonpb.UserRole_USER_ROLE_NORMAL,
				commonpb.UserRole_USER_ROLE_READER,
			},
		},
		"/user.v1.UserService/signin": {
			RPCName: "signin",
			Roles: []commonpb.UserRole{
				commonpb.UserRole_USER_ROLE_NORMAL,
			},
		},
	},
}

var _ = UserServicePermissions
