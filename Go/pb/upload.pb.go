// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: upload.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadRequest_File
	//	*UploadRequest_FileInfo
	Data isUploadRequest_Data `protobuf_oneof:"Data"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{0}
}

func (m *UploadRequest) GetData() isUploadRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadRequest) GetFile() []byte {
	if x, ok := x.GetData().(*UploadRequest_File); ok {
		return x.File
	}
	return nil
}

func (x *UploadRequest) GetFileInfo() *FileInfo {
	if x, ok := x.GetData().(*UploadRequest_FileInfo); ok {
		return x.FileInfo
	}
	return nil
}

type isUploadRequest_Data interface {
	isUploadRequest_Data()
}

type UploadRequest_File struct {
	File []byte `protobuf:"bytes,1,opt,name=file,proto3,oneof"`
}

type UploadRequest_FileInfo struct {
	FileInfo *FileInfo `protobuf:"bytes,2,opt,name=fileInfo,proto3,oneof"`
}

func (*UploadRequest_File) isUploadRequest_Data() {}

func (*UploadRequest_FileInfo) isUploadRequest_Data() {}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=filePath,proto3" json:"filePath,omitempty"`
	FileExt  string `protobuf:"bytes,2,opt,name=fileExt,proto3" json:"fileExt,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileInfo) GetFileExt() string {
	if x != nil {
		return x.FileExt
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SavePath string `protobuf:"bytes,1,opt,name=savePath,proto3" json:"savePath,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponse) GetSavePath() string {
	if x != nil {
		return x.SavePath
	}
	return ""
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5d, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x61, 0x76,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x61, 0x76,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x32, 0x47, 0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x15, 0x2e,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData = file_upload_proto_rawDesc
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_proto_rawDescData)
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_upload_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),  // 0: upload.UploadRequest
	(*FileInfo)(nil),       // 1: upload.FileInfo
	(*UploadResponse)(nil), // 2: upload.UploadResponse
}
var file_upload_proto_depIdxs = []int32{
	1, // 0: upload.UploadRequest.fileInfo:type_name -> upload.FileInfo
	0, // 1: upload.BaseUpload.Upload:input_type -> upload.UploadRequest
	2, // 2: upload.BaseUpload.Upload:output_type -> upload.UploadResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
	file_upload_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UploadRequest_File)(nil),
		(*UploadRequest_FileInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_rawDesc = nil
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BaseUploadClient is the client API for BaseUpload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BaseUploadClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (BaseUpload_UploadClient, error)
}

type baseUploadClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseUploadClient(cc grpc.ClientConnInterface) BaseUploadClient {
	return &baseUploadClient{cc}
}

func (c *baseUploadClient) Upload(ctx context.Context, opts ...grpc.CallOption) (BaseUpload_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BaseUpload_serviceDesc.Streams[0], "/upload.BaseUpload/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &baseUploadUploadClient{stream}
	return x, nil
}

type BaseUpload_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type baseUploadUploadClient struct {
	grpc.ClientStream
}

func (x *baseUploadUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *baseUploadUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BaseUploadServer is the server API for BaseUpload service.
type BaseUploadServer interface {
	Upload(BaseUpload_UploadServer) error
}

// UnimplementedBaseUploadServer can be embedded to have forward compatible implementations.
type UnimplementedBaseUploadServer struct {
}

func (*UnimplementedBaseUploadServer) Upload(BaseUpload_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterBaseUploadServer(s *grpc.Server, srv BaseUploadServer) {
	s.RegisterService(&_BaseUpload_serviceDesc, srv)
}

func _BaseUpload_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BaseUploadServer).Upload(&baseUploadUploadServer{stream})
}

type BaseUpload_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type baseUploadUploadServer struct {
	grpc.ServerStream
}

func (x *baseUploadUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *baseUploadUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BaseUpload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "upload.BaseUpload",
	HandlerType: (*BaseUploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _BaseUpload_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "upload.proto",
}