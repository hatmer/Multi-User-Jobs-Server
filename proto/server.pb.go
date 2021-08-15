// Copyright 2015 gRPC authors.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/server.proto

package server

import (
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

// Request message for starting a job
type JobStartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job string `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *JobStartRequest) Reset() {
	*x = JobStartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStartRequest) ProtoMessage() {}

func (x *JobStartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStartRequest.ProtoReflect.Descriptor instead.
func (*JobStartRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *JobStartRequest) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

// Response message containing the JobID and current state of job
// returned after start, stop or query message
type JobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID  string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // current status of job
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *JobStatus) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *JobStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Message for stopping and getting information about submitted jobs
type JobControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID   string `protobuf:"bytes,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Request string `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"` // stop, status, stream
}

func (x *JobControlRequest) Reset() {
	*x = JobControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobControlRequest) ProtoMessage() {}

func (x *JobControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobControlRequest.ProtoReflect.Descriptor instead.
func (*JobControlRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{2}
}

func (x *JobControlRequest) GetJobID() string {
	if x != nil {
		return x.JobID
	}
	return ""
}

func (x *JobControlRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

// Line of job output for stream
type Line struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Line) Reset() {
	*x = Line{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Line) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Line) ProtoMessage() {}

func (x *Line) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Line.ProtoReflect.Descriptor instead.
func (*Line) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *Line) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_proto_server_proto protoreflect.FileDescriptor

var file_proto_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0f,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f,
	0x62, 0x22, 0x39, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x11,
	0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x1a, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x9f, 0x02,
	0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x35, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x04,
	0x53, 0x74, 0x6f, 0x70, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f,
	0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6e,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData = file_proto_server_proto_rawDesc
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_proto_rawDescData)
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_server_proto_goTypes = []interface{}{
	(*JobStartRequest)(nil),   // 0: server.JobStartRequest
	(*JobStatus)(nil),         // 1: server.JobStatus
	(*JobControlRequest)(nil), // 2: server.JobControlRequest
	(*Line)(nil),              // 3: server.Line
}
var file_proto_server_proto_depIdxs = []int32{
	0, // 0: server.Job.Start:input_type -> server.JobStartRequest
	2, // 1: server.Job.Stop:input_type -> server.JobControlRequest
	2, // 2: server.Job.Status:input_type -> server.JobControlRequest
	2, // 3: server.Job.Stream:input_type -> server.JobControlRequest
	2, // 4: server.Job.Output:input_type -> server.JobControlRequest
	1, // 5: server.Job.Start:output_type -> server.JobStatus
	1, // 6: server.Job.Stop:output_type -> server.JobStatus
	1, // 7: server.Job.Status:output_type -> server.JobStatus
	3, // 8: server.Job.Stream:output_type -> server.Line
	1, // 9: server.Job.Output:output_type -> server.JobStatus
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStartRequest); i {
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
		file_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatus); i {
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
		file_proto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobControlRequest); i {
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
		file_proto_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Line); i {
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
			RawDescriptor: file_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_rawDesc = nil
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}
