// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: address.proto

package address

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

type GenerateDepositAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId string `protobuf:"bytes,1,opt,name=businessId,proto3" json:"businessId,omitempty"`
	Currency   string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *GenerateDepositAddressRequest) Reset() {
	*x = GenerateDepositAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDepositAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDepositAddressRequest) ProtoMessage() {}

func (x *GenerateDepositAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDepositAddressRequest.ProtoReflect.Descriptor instead.
func (*GenerateDepositAddressRequest) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateDepositAddressRequest) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *GenerateDepositAddressRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type GenerateDepositAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GenerateDepositAddressResponse) Reset() {
	*x = GenerateDepositAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDepositAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDepositAddressResponse) ProtoMessage() {}

func (x *GenerateDepositAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDepositAddressResponse.ProtoReflect.Descriptor instead.
func (*GenerateDepositAddressResponse) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateDepositAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5b, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x3a, 0x0a, 0x1e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x32, 0x7d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x2e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_address_proto_goTypes = []interface{}{
	(*GenerateDepositAddressRequest)(nil),  // 0: address.GenerateDepositAddressRequest
	(*GenerateDepositAddressResponse)(nil), // 1: address.GenerateDepositAddressResponse
}
var file_address_proto_depIdxs = []int32{
	0, // 0: address.AddressService.GenerateDepositAddress:input_type -> address.GenerateDepositAddressRequest
	1, // 1: address.AddressService.GenerateDepositAddress:output_type -> address.GenerateDepositAddressResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDepositAddressRequest); i {
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
		file_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDepositAddressResponse); i {
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
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddressServiceClient interface {
	GenerateDepositAddress(ctx context.Context, in *GenerateDepositAddressRequest, opts ...grpc.CallOption) (*GenerateDepositAddressResponse, error)
}

type addressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressServiceClient(cc grpc.ClientConnInterface) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) GenerateDepositAddress(ctx context.Context, in *GenerateDepositAddressRequest, opts ...grpc.CallOption) (*GenerateDepositAddressResponse, error) {
	out := new(GenerateDepositAddressResponse)
	err := c.cc.Invoke(ctx, "/address.AddressService/GenerateDepositAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
type AddressServiceServer interface {
	GenerateDepositAddress(context.Context, *GenerateDepositAddressRequest) (*GenerateDepositAddressResponse, error)
}

// UnimplementedAddressServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (*UnimplementedAddressServiceServer) GenerateDepositAddress(context.Context, *GenerateDepositAddressRequest) (*GenerateDepositAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDepositAddress not implemented")
}

func RegisterAddressServiceServer(s *grpc.Server, srv AddressServiceServer) {
	s.RegisterService(&_AddressService_serviceDesc, srv)
}

func _AddressService_GenerateDepositAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDepositAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GenerateDepositAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/address.AddressService/GenerateDepositAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GenerateDepositAddress(ctx, req.(*GenerateDepositAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "address.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateDepositAddress",
			Handler:    _AddressService_GenerateDepositAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}
