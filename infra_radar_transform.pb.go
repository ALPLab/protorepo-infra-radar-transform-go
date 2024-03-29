// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra_radar_transform.proto

package infra_radar_transform

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
	osi3 "github.com/ALPLab/protorepo-osi-go/v3"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransformRequest struct {
	// Car positioning information in the GPS Exchange Format (gpx) (https://www.topografix.com/gpx/1/1/gpx.xsd)
	Car []byte `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	// Asfinag odo radar interface (version 1.20) unzipped json data
	Radar                []byte   `protobuf:"bytes,2,opt,name=radar,proto3" json:"radar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransformRequest) Reset()         { *m = TransformRequest{} }
func (m *TransformRequest) String() string { return proto.CompactTextString(m) }
func (*TransformRequest) ProtoMessage()    {}
func (*TransformRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf90ebcdf5750bd2, []int{0}
}

func (m *TransformRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformRequest.Unmarshal(m, b)
}
func (m *TransformRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformRequest.Marshal(b, m, deterministic)
}
func (m *TransformRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformRequest.Merge(m, src)
}
func (m *TransformRequest) XXX_Size() int {
	return xxx_messageInfo_TransformRequest.Size(m)
}
func (m *TransformRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransformRequest proto.InternalMessageInfo

func (m *TransformRequest) GetCar() []byte {
	if m != nil {
		return m.Car
	}
	return nil
}

func (m *TransformRequest) GetRadar() []byte {
	if m != nil {
		return m.Radar
	}
	return nil
}

type TransformResponse struct {
	SensorData           []*osi3.SensorData `protobuf:"bytes,1,rep,name=sensor_data,json=sensorData,proto3" json:"sensor_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TransformResponse) Reset()         { *m = TransformResponse{} }
func (m *TransformResponse) String() string { return proto.CompactTextString(m) }
func (*TransformResponse) ProtoMessage()    {}
func (*TransformResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf90ebcdf5750bd2, []int{1}
}

func (m *TransformResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformResponse.Unmarshal(m, b)
}
func (m *TransformResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformResponse.Marshal(b, m, deterministic)
}
func (m *TransformResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformResponse.Merge(m, src)
}
func (m *TransformResponse) XXX_Size() int {
	return xxx_messageInfo_TransformResponse.Size(m)
}
func (m *TransformResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransformResponse proto.InternalMessageInfo

func (m *TransformResponse) GetSensorData() []*osi3.SensorData {
	if m != nil {
		return m.SensorData
	}
	return nil
}

func init() {
	proto.RegisterType((*TransformRequest)(nil), "infra_radar_transform.TransformRequest")
	proto.RegisterType((*TransformResponse)(nil), "infra_radar_transform.TransformResponse")
}

func init() { proto.RegisterFile("infra_radar_transform.proto", fileDescriptor_cf90ebcdf5750bd2) }

var fileDescriptor_cf90ebcdf5750bd2 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0x3f, 0x4b, 0xc5, 0x30,
	0x10, 0xb7, 0x3e, 0x14, 0xbc, 0xe7, 0x50, 0x83, 0x42, 0x7d, 0x6f, 0x79, 0x74, 0xb1, 0x53, 0xc4,
	0xd7, 0xcd, 0x59, 0x04, 0x37, 0x89, 0xee, 0xf1, 0xb4, 0x29, 0x64, 0x30, 0x57, 0xef, 0xe2, 0xec,
	0x57, 0x97, 0x24, 0x58, 0x44, 0x0a, 0x6e, 0x97, 0xdc, 0xef, 0xef, 0xc1, 0xd6, 0x87, 0x91, 0xd1,
	0x32, 0x0e, 0xc8, 0x36, 0x32, 0x06, 0x19, 0x89, 0xdf, 0xf5, 0xc4, 0x14, 0x49, 0x5d, 0x2c, 0x2e,
	0x37, 0x97, 0x24, 0xbe, 0xbf, 0x26, 0xf1, 0x56, 0x5c, 0x10, 0xe2, 0x01, 0x23, 0x16, 0x46, 0x7b,
	0x0b, 0xf5, 0xf3, 0x0f, 0xce, 0xb8, 0x8f, 0x4f, 0x27, 0x51, 0xd5, 0xb0, 0x7a, 0x43, 0x6e, 0xaa,
	0x5d, 0xd5, 0x9d, 0x9a, 0x34, 0xaa, 0x73, 0x38, 0xca, 0x9a, 0xcd, 0x61, 0xfe, 0x2b, 0x8f, 0xf6,
	0x1e, 0xce, 0x7e, 0x71, 0x65, 0xa2, 0x20, 0x4e, 0xdd, 0xc0, 0xba, 0x98, 0xd8, 0xe4, 0xd2, 0x54,
	0xbb, 0x55, 0xb7, 0xde, 0xd7, 0x3a, 0x25, 0xd0, 0x4f, 0x79, 0x71, 0x87, 0x11, 0x0d, 0xc8, 0x3c,
	0xef, 0xbf, 0x60, 0xfb, 0x90, 0x72, 0x9b, 0xa4, 0xfa, 0x48, 0xe2, 0xa3, 0xa7, 0x30, 0x2b, 0xab,
	0x17, 0x38, 0x99, 0xab, 0xa8, 0x2b, 0xbd, 0xdc, 0xff, 0x6f, 0x89, 0x4d, 0xf7, 0x3f, 0xb0, 0x24,
	0x6e, 0x0f, 0x5e, 0x8f, 0xf3, 0x2d, 0xfa, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xb6, 0xfd,
	0x8c, 0x5c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InfraRadarPositionTransformClient is the client API for InfraRadarPositionTransform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InfraRadarPositionTransformClient interface {
	Transform(ctx context.Context, in *TransformRequest, opts ...grpc.CallOption) (*TransformResponse, error)
}

type infraRadarPositionTransformClient struct {
	cc *grpc.ClientConn
}

func NewInfraRadarPositionTransformClient(cc *grpc.ClientConn) InfraRadarPositionTransformClient {
	return &infraRadarPositionTransformClient{cc}
}

func (c *infraRadarPositionTransformClient) Transform(ctx context.Context, in *TransformRequest, opts ...grpc.CallOption) (*TransformResponse, error) {
	out := new(TransformResponse)
	err := c.cc.Invoke(ctx, "/infra_radar_transform.InfraRadarPositionTransform/transform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfraRadarPositionTransformServer is the server API for InfraRadarPositionTransform service.
type InfraRadarPositionTransformServer interface {
	Transform(context.Context, *TransformRequest) (*TransformResponse, error)
}

func RegisterInfraRadarPositionTransformServer(s *grpc.Server, srv InfraRadarPositionTransformServer) {
	s.RegisterService(&_InfraRadarPositionTransform_serviceDesc, srv)
}

func _InfraRadarPositionTransform_Transform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfraRadarPositionTransformServer).Transform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra_radar_transform.InfraRadarPositionTransform/Transform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfraRadarPositionTransformServer).Transform(ctx, req.(*TransformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InfraRadarPositionTransform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infra_radar_transform.InfraRadarPositionTransform",
	HandlerType: (*InfraRadarPositionTransformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "transform",
			Handler:    _InfraRadarPositionTransform_Transform_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra_radar_transform.proto",
}
