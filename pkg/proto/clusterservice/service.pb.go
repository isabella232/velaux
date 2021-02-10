// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/proto/clusterservice/service.proto

package clusterservice

import (
	proto "github.com/golang/protobuf/proto"
	model "github.com/oam-dev/velacp/pkg/datastore/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PutClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *model.Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *PutClusterRequest) Reset() {
	*x = PutClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutClusterRequest) ProtoMessage() {}

func (x *PutClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutClusterRequest.ProtoReflect.Descriptor instead.
func (*PutClusterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{0}
}

func (x *PutClusterRequest) GetCluster() *model.Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type PutClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutClusterResponse) Reset() {
	*x = PutClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutClusterResponse) ProtoMessage() {}

func (x *PutClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutClusterResponse.ProtoReflect.Descriptor instead.
func (*PutClusterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{1}
}

type GetClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterResponse) Reset() {
	*x = GetClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterResponse) ProtoMessage() {}

func (x *GetClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterResponse.ProtoReflect.Descriptor instead.
func (*GetClusterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{3}
}

type ListClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListClustersRequest) Reset() {
	*x = ListClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersRequest) ProtoMessage() {}

func (x *ListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersRequest.ProtoReflect.Descriptor instead.
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{4}
}

type ListClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*model.Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *ListClustersResponse) Reset() {
	*x = ListClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersResponse) ProtoMessage() {}

func (x *ListClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersResponse.ProtoReflect.Descriptor instead.
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListClustersResponse) GetClusters() []*model.Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type DelClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DelClusterRequest) Reset() {
	*x = DelClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelClusterRequest) ProtoMessage() {}

func (x *DelClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelClusterRequest.ProtoReflect.Descriptor instead.
func (*DelClusterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{6}
}

func (x *DelClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DelClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelClusterResponse) Reset() {
	*x = DelClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelClusterResponse) ProtoMessage() {}

func (x *DelClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_clusterservice_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelClusterResponse.ProtoReflect.Descriptor instead.
func (*DelClusterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_clusterservice_service_proto_rawDescGZIP(), []int{7}
}

var File_pkg_proto_clusterservice_service_proto protoreflect.FileDescriptor

var file_pkg_proto_clusterservice_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11,
	0x50, 0x75, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x75, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x4b, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x6c,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x27, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfa, 0x03, 0x0a,
	0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x77, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x32, 0x2e,
	0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x75, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x32, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x76, 0x65, 0x6c,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x7d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x34, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x77, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x32,
	0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x61, 0x6d, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x76, 0x65, 0x6c, 0x61, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_clusterservice_service_proto_rawDescOnce sync.Once
	file_pkg_proto_clusterservice_service_proto_rawDescData = file_pkg_proto_clusterservice_service_proto_rawDesc
)

func file_pkg_proto_clusterservice_service_proto_rawDescGZIP() []byte {
	file_pkg_proto_clusterservice_service_proto_rawDescOnce.Do(func() {
		file_pkg_proto_clusterservice_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_clusterservice_service_proto_rawDescData)
	})
	return file_pkg_proto_clusterservice_service_proto_rawDescData
}

var file_pkg_proto_clusterservice_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_clusterservice_service_proto_goTypes = []interface{}{
	(*PutClusterRequest)(nil),    // 0: vela.api.service.clusterservice.PutClusterRequest
	(*PutClusterResponse)(nil),   // 1: vela.api.service.clusterservice.PutClusterResponse
	(*GetClusterRequest)(nil),    // 2: vela.api.service.clusterservice.GetClusterRequest
	(*GetClusterResponse)(nil),   // 3: vela.api.service.clusterservice.GetClusterResponse
	(*ListClustersRequest)(nil),  // 4: vela.api.service.clusterservice.ListClustersRequest
	(*ListClustersResponse)(nil), // 5: vela.api.service.clusterservice.ListClustersResponse
	(*DelClusterRequest)(nil),    // 6: vela.api.service.clusterservice.DelClusterRequest
	(*DelClusterResponse)(nil),   // 7: vela.api.service.clusterservice.DelClusterResponse
	(*model.Cluster)(nil),        // 8: vela.api.model.Cluster
}
var file_pkg_proto_clusterservice_service_proto_depIdxs = []int32{
	8, // 0: vela.api.service.clusterservice.PutClusterRequest.cluster:type_name -> vela.api.model.Cluster
	8, // 1: vela.api.service.clusterservice.ListClustersResponse.clusters:type_name -> vela.api.model.Cluster
	0, // 2: vela.api.service.clusterservice.ClusterService.PutCluster:input_type -> vela.api.service.clusterservice.PutClusterRequest
	2, // 3: vela.api.service.clusterservice.ClusterService.GetCluster:input_type -> vela.api.service.clusterservice.GetClusterRequest
	4, // 4: vela.api.service.clusterservice.ClusterService.ListClusters:input_type -> vela.api.service.clusterservice.ListClustersRequest
	6, // 5: vela.api.service.clusterservice.ClusterService.DelCluster:input_type -> vela.api.service.clusterservice.DelClusterRequest
	1, // 6: vela.api.service.clusterservice.ClusterService.PutCluster:output_type -> vela.api.service.clusterservice.PutClusterResponse
	3, // 7: vela.api.service.clusterservice.ClusterService.GetCluster:output_type -> vela.api.service.clusterservice.GetClusterResponse
	5, // 8: vela.api.service.clusterservice.ClusterService.ListClusters:output_type -> vela.api.service.clusterservice.ListClustersResponse
	7, // 9: vela.api.service.clusterservice.ClusterService.DelCluster:output_type -> vela.api.service.clusterservice.DelClusterResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_clusterservice_service_proto_init() }
func file_pkg_proto_clusterservice_service_proto_init() {
	if File_pkg_proto_clusterservice_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_clusterservice_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_clusterservice_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_clusterservice_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_clusterservice_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_clusterservice_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_clusterservice_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_clusterservice_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_clusterservice_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelClusterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_clusterservice_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_clusterservice_service_proto_goTypes,
		DependencyIndexes: file_pkg_proto_clusterservice_service_proto_depIdxs,
		MessageInfos:      file_pkg_proto_clusterservice_service_proto_msgTypes,
	}.Build()
	File_pkg_proto_clusterservice_service_proto = out.File
	file_pkg_proto_clusterservice_service_proto_rawDesc = nil
	file_pkg_proto_clusterservice_service_proto_goTypes = nil
	file_pkg_proto_clusterservice_service_proto_depIdxs = nil
}
