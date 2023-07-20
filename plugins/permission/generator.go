package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	commonpb "github.com/openfms/protos/gen/common/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_permission.pb.go"
	if len(file.Services) == 0 {
		return nil
	}
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	f := jen.NewFilePath(string(file.GoPackageName))
	f.PackageComment("// Code generated by protoc-gen-permission. DO NOT EDIT.")

	// Import required packages
	f.ImportAlias("github.com/openfms/protos/gen/common/v1", "commonpb")

	// Generate MethodStreamInfo struct
	f.Type().Id("MethodStreamInfo").Struct(
		jen.Id("Roles").Index().Qual("github.com/openfms/protos/gen/common/v1", "UserRole").Tag(map[string]string{"json": "roles"}),
		jen.Id("RPCName").String().Tag(map[string]string{"json": "rpc_name"}),
	)

	// Generate ServicePermission struct
	f.Type().Id("ServicePermission").Struct(
		jen.Id("ServiceName").String().Tag(map[string]string{"json": "service_name"}),
		jen.Id("MethodStreams").Map(jen.String()).Op("*").Id("MethodStreamInfo").Tag(map[string]string{"json": "method_streams"}),
	)
	for _, service := range file.Services {
		srvName := service.GoName + "Permission"
		methodsDicts := make([]jen.Code, 0)
		for _, method := range service.Methods {
			perms, ok := GetMethodPermissionOpts(method.Desc)
			if !ok || perms == nil {
				continue
			}
			fullMethod := BuildFullMethod(*file.Proto.Package, service.GoName, method.GoName)
			roleValues := make([]jen.Code, 0)
			for _, role := range perms.Roles {
				roleValues = append(roleValues, jen.Qual("github.com/openfms/protos/gen/common/v1", "UserRole_"+role.String()))
			}
			methodDetails := jen.Dict{
				jen.Lit(fullMethod): jen.Op("&").Id("MethodStreamInfo").Values(jen.Dict{
					jen.Id("RPCName"): jen.Lit(method.GoName),
					jen.Id("Roles"):   jen.Index().Qual("github.com/openfms/protos/gen/common/v1", "UserRole").Values(roleValues...),
				}),
			}
			methodsDicts = append(methodsDicts, methodDetails)
		}
		f.Var().Id(srvName).Op("=").Op("&").Id("ServicePermission").Values(jen.Dict{
			jen.Id("ServiceName"):   jen.Lit("/user.v1.UserService"),
			jen.Id("MethodStreams"): jen.Map(jen.String()).Op("*").Id("MethodStreamInfo").Values(methodsDicts...),
		})
	}
	g.P(fmt.Sprintf("%#v", f))
	return g
}

func GetMethodPermissionOpts(methodDesc protoreflect.MethodDescriptor) (*commonpb.RoleAccess, bool) {
	methodOpts, ok := methodDesc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return &commonpb.RoleAccess{}, false
	}
	permOpts, ok := proto.GetExtension(methodOpts, commonpb.E_Permission).(*commonpb.RoleAccess)
	if !ok {
		return &commonpb.RoleAccess{}, false
	}
	return permOpts, true
}

func BuildFullMethod(packageName, serviceName, methodStream string) string {
	return "/" + packageName + "." + serviceName + "/" + methodStream
}
