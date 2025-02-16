// Code generated by protoc-gen-permission. DO NOT EDIT.
package ioelssv1

import commonpb "github.com/irisco88/protos/gen/common/v1"

type MethodStreamInfo struct {
	Roles   []commonpb.UserRole `json:"roles"`
	RPCName string              `json:"rpc_name"`
}
type ServicePermission struct {
	ServiceName   string                       `json:"service_name"`
	MethodStreams map[string]*MethodStreamInfo `json:"method_streams"`
}

var IoelssServicePermission = &ServicePermission{
	MethodStreams: map[string]*MethodStreamInfo{
		"/ioelss.v1.IoelssService/CreateIoelss": &MethodStreamInfo{
			RPCName: "CreateIoelss",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN},
		},
		"/ioelss.v1.IoelssService/DeleteIoelss": &MethodStreamInfo{
			RPCName: "DeleteIoelss",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN, commonpb.UserRole_USER_ROLE_NORMAL},
		},
		"/ioelss.v1.IoelssService/UpdateIoelss": &MethodStreamInfo{
			RPCName: "UpdateIoelss",
			Roles:   []commonpb.UserRole{commonpb.UserRole_USER_ROLE_ADMIN, commonpb.UserRole_USER_ROLE_NORMAL},
		},
	},
	ServiceName: "ioelss.v1.IoelssService",
}
