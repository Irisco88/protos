// Code generated by protoc-gen-permission. DO NOT EDIT.
package imeisv1

import commonpb "github.com/irisco88/protos/gen/common/v1"

type MethodStreamInfo struct {
	Roles   []commonpb.UserRole `json:"roles"`
	RPCName string              `json:"rpc_name"`
}
type ServicePermission struct {
	ServiceName   string                       `json:"service_name"`
	MethodStreams map[string]*MethodStreamInfo `json:"method_streams"`
}

var ImeisServicePermission = &ServicePermission{
	MethodStreams: map[string]*MethodStreamInfo{
		"/imeis.v1.ImeisService/CreateImeis": &MethodStreamInfo{
			RPCName: "CreateImeis",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN},
		},
		"/imeis.v1.ImeisService/DeleteImeis": &MethodStreamInfo{
			RPCName: "DeleteImeis",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN, commonpb.UserRole_USER_ROLE_NORMAL},
		},
		"/imeis.v1.ImeisService/UpdateImeis": &MethodStreamInfo{
			RPCName: "UpdateImeis",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN, commonpb.UserRole_USER_ROLE_NORMAL},
		},
	},
	ServiceName: "imeis.v1.ImeisService",
}
