// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package servicedirectory aliases all exported identifiers in package
// "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb".
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package servicedirectory

import (
	src "cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use vars in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
var (
	File_google_cloud_servicedirectory_v1beta1_endpoint_proto             = src.File_google_cloud_servicedirectory_v1beta1_endpoint_proto
	File_google_cloud_servicedirectory_v1beta1_lookup_service_proto       = src.File_google_cloud_servicedirectory_v1beta1_lookup_service_proto
	File_google_cloud_servicedirectory_v1beta1_namespace_proto            = src.File_google_cloud_servicedirectory_v1beta1_namespace_proto
	File_google_cloud_servicedirectory_v1beta1_registration_service_proto = src.File_google_cloud_servicedirectory_v1beta1_registration_service_proto
	File_google_cloud_servicedirectory_v1beta1_service_proto              = src.File_google_cloud_servicedirectory_v1beta1_service_proto
)

// The request message for
// [RegistrationService.CreateEndpoint][google.cloud.servicedirectory.v1beta1.RegistrationService.CreateEndpoint].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type CreateEndpointRequest = src.CreateEndpointRequest

// The request message for
// [RegistrationService.CreateNamespace][google.cloud.servicedirectory.v1beta1.RegistrationService.CreateNamespace].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type CreateNamespaceRequest = src.CreateNamespaceRequest

// The request message for
// [RegistrationService.CreateService][google.cloud.servicedirectory.v1beta1.RegistrationService.CreateService].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type CreateServiceRequest = src.CreateServiceRequest

// The request message for
// [RegistrationService.DeleteEndpoint][google.cloud.servicedirectory.v1beta1.RegistrationService.DeleteEndpoint].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type DeleteEndpointRequest = src.DeleteEndpointRequest

// The request message for
// [RegistrationService.DeleteNamespace][google.cloud.servicedirectory.v1beta1.RegistrationService.DeleteNamespace].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type DeleteNamespaceRequest = src.DeleteNamespaceRequest

// The request message for
// [RegistrationService.DeleteService][google.cloud.servicedirectory.v1beta1.RegistrationService.DeleteService].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type DeleteServiceRequest = src.DeleteServiceRequest

// An individual endpoint that provides a
// [service][google.cloud.servicedirectory.v1beta1.Service]. The service must
// already exist to create an endpoint.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type Endpoint = src.Endpoint

// The request message for
// [RegistrationService.GetEndpoint][google.cloud.servicedirectory.v1beta1.RegistrationService.GetEndpoint].
// This should not be used to lookup endpoints at runtime. Instead, use the
// `resolve` method.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type GetEndpointRequest = src.GetEndpointRequest

// The request message for
// [RegistrationService.GetNamespace][google.cloud.servicedirectory.v1beta1.RegistrationService.GetNamespace].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type GetNamespaceRequest = src.GetNamespaceRequest

// The request message for
// [RegistrationService.GetService][google.cloud.servicedirectory.v1beta1.RegistrationService.GetService].
// This should not be used for looking up a service. Insead, use the `resolve`
// method as it contains all endpoints and associated metadata.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type GetServiceRequest = src.GetServiceRequest

// The request message for
// [RegistrationService.ListEndpoints][google.cloud.servicedirectory.v1beta1.RegistrationService.ListEndpoints].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ListEndpointsRequest = src.ListEndpointsRequest

// The response message for
// [RegistrationService.ListEndpoints][google.cloud.servicedirectory.v1beta1.RegistrationService.ListEndpoints].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ListEndpointsResponse = src.ListEndpointsResponse

// The request message for
// [RegistrationService.ListNamespaces][google.cloud.servicedirectory.v1beta1.RegistrationService.ListNamespaces].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ListNamespacesRequest = src.ListNamespacesRequest

// The response message for
// [RegistrationService.ListNamespaces][google.cloud.servicedirectory.v1beta1.RegistrationService.ListNamespaces].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ListNamespacesResponse = src.ListNamespacesResponse

// The request message for
// [RegistrationService.ListServices][google.cloud.servicedirectory.v1beta1.RegistrationService.ListServices].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ListServicesRequest = src.ListServicesRequest

// The response message for
// [RegistrationService.ListServices][google.cloud.servicedirectory.v1beta1.RegistrationService.ListServices].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ListServicesResponse = src.ListServicesResponse

// LookupServiceClient is the client API for LookupService service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type LookupServiceClient = src.LookupServiceClient

// LookupServiceServer is the server API for LookupService service.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type LookupServiceServer = src.LookupServiceServer

// A container for [services][google.cloud.servicedirectory.v1beta1.Service].
// Namespaces allow administrators to group services together and define
// permissions for a collection of services.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type Namespace = src.Namespace

// RegistrationServiceClient is the client API for RegistrationService
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type RegistrationServiceClient = src.RegistrationServiceClient

// RegistrationServiceServer is the server API for RegistrationService
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type RegistrationServiceServer = src.RegistrationServiceServer

// The request message for
// [LookupService.ResolveService][google.cloud.servicedirectory.v1beta1.LookupService.ResolveService].
// Looks up a service by its name, returns the service and its endpoints.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ResolveServiceRequest = src.ResolveServiceRequest

// The response message for
// [LookupService.ResolveService][google.cloud.servicedirectory.v1beta1.LookupService.ResolveService].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type ResolveServiceResponse = src.ResolveServiceResponse

// An individual service. A service contains a name and optional metadata. A
// service must exist before
// [endpoints][google.cloud.servicedirectory.v1beta1.Endpoint] can be added to
// it.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type Service = src.Service

// UnimplementedLookupServiceServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type UnimplementedLookupServiceServer = src.UnimplementedLookupServiceServer

// UnimplementedRegistrationServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type UnimplementedRegistrationServiceServer = src.UnimplementedRegistrationServiceServer

// The request message for
// [RegistrationService.UpdateEndpoint][google.cloud.servicedirectory.v1beta1.RegistrationService.UpdateEndpoint].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type UpdateEndpointRequest = src.UpdateEndpointRequest

// The request message for
// [RegistrationService.UpdateNamespace][google.cloud.servicedirectory.v1beta1.RegistrationService.UpdateNamespace].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type UpdateNamespaceRequest = src.UpdateNamespaceRequest

// The request message for
// [RegistrationService.UpdateService][google.cloud.servicedirectory.v1beta1.RegistrationService.UpdateService].
//
// Deprecated: Please use types in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
type UpdateServiceRequest = src.UpdateServiceRequest

// Deprecated: Please use funcs in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
func NewLookupServiceClient(cc grpc.ClientConnInterface) LookupServiceClient {
	return src.NewLookupServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
func NewRegistrationServiceClient(cc grpc.ClientConnInterface) RegistrationServiceClient {
	return src.NewRegistrationServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
func RegisterLookupServiceServer(s *grpc.Server, srv LookupServiceServer) {
	src.RegisterLookupServiceServer(s, srv)
}

// Deprecated: Please use funcs in: cloud.google.com/go/servicedirectory/apiv1beta1/servicedirectorypb
func RegisterRegistrationServiceServer(s *grpc.Server, srv RegistrationServiceServer) {
	src.RegisterRegistrationServiceServer(s, srv)
}
