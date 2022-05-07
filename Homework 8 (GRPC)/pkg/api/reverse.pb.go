package api

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reverse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_reverse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_reverse_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_reverse_proto protoreflect.FileDescriptor

var file_reverse_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x36, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x72, 0x12,
	0x2d, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reverse_proto_rawDescOnce sync.Once
	file_reverse_proto_rawDescData = file_reverse_proto_rawDesc
)

func file_reverse_proto_rawDescGZIP() []byte {
	file_reverse_proto_rawDescOnce.Do(func() {
		file_reverse_proto_rawDescData = protoimpl.X.CompressGZIP(file_reverse_proto_rawDescData)
	})
	return file_reverse_proto_rawDescData
}

var file_reverse_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reverse_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: protos.Message
}
var file_reverse_proto_depIdxs = []int32{
	0, // 0: protos.Adder.Reverse:input_type -> protos.Message
	0, // 1: protos.Adder.Reverse:output_type -> protos.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reverse_proto_init() }
func file_reverse_proto_init() {
	if File_reverse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reverse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_reverse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reverse_proto_goTypes,
		DependencyIndexes: file_reverse_proto_depIdxs,
		MessageInfos:      file_reverse_proto_msgTypes,
	}.Build()
	File_reverse_proto = out.File
	file_reverse_proto_rawDesc = nil
	file_reverse_proto_goTypes = nil
	file_reverse_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdderClient is the client API for Adder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdderClient interface {
	Reverse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type adderClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderClient(cc grpc.ClientConnInterface) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) Reverse(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/protos.Adder/Reverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServer is the server API for Adder service.
type AdderServer interface {
	Reverse(context.Context, *Message) (*Message, error)
}

// UnimplementedAdderServer can be embedded to have forward compatible implementations.
type UnimplementedAdderServer struct {
}

func (*UnimplementedAdderServer) Reverse(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reverse not implemented")
}

func RegisterAdderServer(s *grpc.Server, srv AdderServer) {
	s.RegisterService(&_Adder_serviceDesc, srv)
}

func _Adder_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Adder/Reverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Reverse(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reverse",
			Handler:    _Adder_Reverse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reverse.proto",
}