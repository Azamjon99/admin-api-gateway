// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vehicle.proto

package logistics_vehicle

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Vehicle struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DriverId             string   `protobuf:"bytes,2,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Model                string   `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Make                 string   `protobuf:"bytes,4,opt,name=make,proto3" json:"make,omitempty"`
	PlateNumber          string   `protobuf:"bytes,5,opt,name=plate_number,proto3" json:"plate_number,omitempty"`
	ImageUrl             string   `protobuf:"bytes,6,opt,name=image_url,proto3" json:"image_url,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vehicle) Reset()         { *m = Vehicle{} }
func (m *Vehicle) String() string { return proto.CompactTextString(m) }
func (*Vehicle) ProtoMessage()    {}
func (*Vehicle) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{0}
}

func (m *Vehicle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vehicle.Unmarshal(m, b)
}
func (m *Vehicle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vehicle.Marshal(b, m, deterministic)
}
func (m *Vehicle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vehicle.Merge(m, src)
}
func (m *Vehicle) XXX_Size() int {
	return xxx_messageInfo_Vehicle.Size(m)
}
func (m *Vehicle) XXX_DiscardUnknown() {
	xxx_messageInfo_Vehicle.DiscardUnknown(m)
}

var xxx_messageInfo_Vehicle proto.InternalMessageInfo

func (m *Vehicle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vehicle) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

func (m *Vehicle) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Vehicle) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

func (m *Vehicle) GetPlateNumber() string {
	if m != nil {
		return m.PlateNumber
	}
	return ""
}

func (m *Vehicle) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *Vehicle) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Vehicle) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CreateVehicleRequest struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Make                 string   `protobuf:"bytes,3,opt,name=make,proto3" json:"make,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVehicleRequest) Reset()         { *m = CreateVehicleRequest{} }
func (m *CreateVehicleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateVehicleRequest) ProtoMessage()    {}
func (*CreateVehicleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{1}
}

func (m *CreateVehicleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVehicleRequest.Unmarshal(m, b)
}
func (m *CreateVehicleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVehicleRequest.Marshal(b, m, deterministic)
}
func (m *CreateVehicleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVehicleRequest.Merge(m, src)
}
func (m *CreateVehicleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateVehicleRequest.Size(m)
}
func (m *CreateVehicleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVehicleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVehicleRequest proto.InternalMessageInfo

func (m *CreateVehicleRequest) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

func (m *CreateVehicleRequest) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *CreateVehicleRequest) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

type CreateVehicleResponse struct {
	Vehicle              *Vehicle `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateVehicleResponse) Reset()         { *m = CreateVehicleResponse{} }
func (m *CreateVehicleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateVehicleResponse) ProtoMessage()    {}
func (*CreateVehicleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{2}
}

func (m *CreateVehicleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateVehicleResponse.Unmarshal(m, b)
}
func (m *CreateVehicleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateVehicleResponse.Marshal(b, m, deterministic)
}
func (m *CreateVehicleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateVehicleResponse.Merge(m, src)
}
func (m *CreateVehicleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateVehicleResponse.Size(m)
}
func (m *CreateVehicleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateVehicleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateVehicleResponse proto.InternalMessageInfo

func (m *CreateVehicleResponse) GetVehicle() *Vehicle {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

type UpdateVehicleRequest struct {
	Vehicle              *Vehicle `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVehicleRequest) Reset()         { *m = UpdateVehicleRequest{} }
func (m *UpdateVehicleRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateVehicleRequest) ProtoMessage()    {}
func (*UpdateVehicleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{3}
}

func (m *UpdateVehicleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVehicleRequest.Unmarshal(m, b)
}
func (m *UpdateVehicleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVehicleRequest.Marshal(b, m, deterministic)
}
func (m *UpdateVehicleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVehicleRequest.Merge(m, src)
}
func (m *UpdateVehicleRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateVehicleRequest.Size(m)
}
func (m *UpdateVehicleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVehicleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVehicleRequest proto.InternalMessageInfo

func (m *UpdateVehicleRequest) GetVehicle() *Vehicle {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

type UpdateVehicleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateVehicleResponse) Reset()         { *m = UpdateVehicleResponse{} }
func (m *UpdateVehicleResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateVehicleResponse) ProtoMessage()    {}
func (*UpdateVehicleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{4}
}

func (m *UpdateVehicleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateVehicleResponse.Unmarshal(m, b)
}
func (m *UpdateVehicleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateVehicleResponse.Marshal(b, m, deterministic)
}
func (m *UpdateVehicleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateVehicleResponse.Merge(m, src)
}
func (m *UpdateVehicleResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateVehicleResponse.Size(m)
}
func (m *UpdateVehicleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateVehicleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateVehicleResponse proto.InternalMessageInfo

type DeleteVehicleRequest struct {
	VehicleId            string   `protobuf:"bytes,1,opt,name=vehicle_id,proto3" json:"vehicle_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteVehicleRequest) Reset()         { *m = DeleteVehicleRequest{} }
func (m *DeleteVehicleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteVehicleRequest) ProtoMessage()    {}
func (*DeleteVehicleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{5}
}

func (m *DeleteVehicleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteVehicleRequest.Unmarshal(m, b)
}
func (m *DeleteVehicleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteVehicleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteVehicleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteVehicleRequest.Merge(m, src)
}
func (m *DeleteVehicleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteVehicleRequest.Size(m)
}
func (m *DeleteVehicleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteVehicleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteVehicleRequest proto.InternalMessageInfo

func (m *DeleteVehicleRequest) GetVehicleId() string {
	if m != nil {
		return m.VehicleId
	}
	return ""
}

type DeleteVehicleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteVehicleResponse) Reset()         { *m = DeleteVehicleResponse{} }
func (m *DeleteVehicleResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteVehicleResponse) ProtoMessage()    {}
func (*DeleteVehicleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{6}
}

func (m *DeleteVehicleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteVehicleResponse.Unmarshal(m, b)
}
func (m *DeleteVehicleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteVehicleResponse.Marshal(b, m, deterministic)
}
func (m *DeleteVehicleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteVehicleResponse.Merge(m, src)
}
func (m *DeleteVehicleResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteVehicleResponse.Size(m)
}
func (m *DeleteVehicleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteVehicleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteVehicleResponse proto.InternalMessageInfo

type GetVehicleRequest struct {
	VehicleId            string   `protobuf:"bytes,1,opt,name=vehicle_id,proto3" json:"vehicle_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVehicleRequest) Reset()         { *m = GetVehicleRequest{} }
func (m *GetVehicleRequest) String() string { return proto.CompactTextString(m) }
func (*GetVehicleRequest) ProtoMessage()    {}
func (*GetVehicleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{7}
}

func (m *GetVehicleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVehicleRequest.Unmarshal(m, b)
}
func (m *GetVehicleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVehicleRequest.Marshal(b, m, deterministic)
}
func (m *GetVehicleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVehicleRequest.Merge(m, src)
}
func (m *GetVehicleRequest) XXX_Size() int {
	return xxx_messageInfo_GetVehicleRequest.Size(m)
}
func (m *GetVehicleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVehicleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVehicleRequest proto.InternalMessageInfo

func (m *GetVehicleRequest) GetVehicleId() string {
	if m != nil {
		return m.VehicleId
	}
	return ""
}

type GetVehicleResponse struct {
	Vehicle              *Vehicle `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVehicleResponse) Reset()         { *m = GetVehicleResponse{} }
func (m *GetVehicleResponse) String() string { return proto.CompactTextString(m) }
func (*GetVehicleResponse) ProtoMessage()    {}
func (*GetVehicleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{8}
}

func (m *GetVehicleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVehicleResponse.Unmarshal(m, b)
}
func (m *GetVehicleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVehicleResponse.Marshal(b, m, deterministic)
}
func (m *GetVehicleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVehicleResponse.Merge(m, src)
}
func (m *GetVehicleResponse) XXX_Size() int {
	return xxx_messageInfo_GetVehicleResponse.Size(m)
}
func (m *GetVehicleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVehicleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVehicleResponse proto.InternalMessageInfo

func (m *GetVehicleResponse) GetVehicle() *Vehicle {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

type ListVehicleRequest struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVehicleRequest) Reset()         { *m = ListVehicleRequest{} }
func (m *ListVehicleRequest) String() string { return proto.CompactTextString(m) }
func (*ListVehicleRequest) ProtoMessage()    {}
func (*ListVehicleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{9}
}

func (m *ListVehicleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVehicleRequest.Unmarshal(m, b)
}
func (m *ListVehicleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVehicleRequest.Marshal(b, m, deterministic)
}
func (m *ListVehicleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVehicleRequest.Merge(m, src)
}
func (m *ListVehicleRequest) XXX_Size() int {
	return xxx_messageInfo_ListVehicleRequest.Size(m)
}
func (m *ListVehicleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVehicleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVehicleRequest proto.InternalMessageInfo

func (m *ListVehicleRequest) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

type ListVehicleResponse struct {
	Vehicles             []*Vehicle `protobuf:"bytes,1,rep,name=vehicles,proto3" json:"vehicles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListVehicleResponse) Reset()         { *m = ListVehicleResponse{} }
func (m *ListVehicleResponse) String() string { return proto.CompactTextString(m) }
func (*ListVehicleResponse) ProtoMessage()    {}
func (*ListVehicleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_416ab71f8212867c, []int{10}
}

func (m *ListVehicleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVehicleResponse.Unmarshal(m, b)
}
func (m *ListVehicleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVehicleResponse.Marshal(b, m, deterministic)
}
func (m *ListVehicleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVehicleResponse.Merge(m, src)
}
func (m *ListVehicleResponse) XXX_Size() int {
	return xxx_messageInfo_ListVehicleResponse.Size(m)
}
func (m *ListVehicleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVehicleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVehicleResponse proto.InternalMessageInfo

func (m *ListVehicleResponse) GetVehicles() []*Vehicle {
	if m != nil {
		return m.Vehicles
	}
	return nil
}

func init() {
	proto.RegisterType((*Vehicle)(nil), "Vehicle")
	proto.RegisterType((*CreateVehicleRequest)(nil), "CreateVehicleRequest")
	proto.RegisterType((*CreateVehicleResponse)(nil), "CreateVehicleResponse")
	proto.RegisterType((*UpdateVehicleRequest)(nil), "UpdateVehicleRequest")
	proto.RegisterType((*UpdateVehicleResponse)(nil), "UpdateVehicleResponse")
	proto.RegisterType((*DeleteVehicleRequest)(nil), "DeleteVehicleRequest")
	proto.RegisterType((*DeleteVehicleResponse)(nil), "DeleteVehicleResponse")
	proto.RegisterType((*GetVehicleRequest)(nil), "GetVehicleRequest")
	proto.RegisterType((*GetVehicleResponse)(nil), "GetVehicleResponse")
	proto.RegisterType((*ListVehicleRequest)(nil), "ListVehicleRequest")
	proto.RegisterType((*ListVehicleResponse)(nil), "ListVehicleResponse")
}

func init() {
	proto.RegisterFile("vehicle.proto", fileDescriptor_416ab71f8212867c)
}

var fileDescriptor_416ab71f8212867c = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xfa, 0xff, 0x7e, 0x9f, 0x82, 0xd3, 0x14, 0x67, 0x21, 0xa5, 0x0c, 0x2e, 0x5c,
	0x55, 0x68, 0x41, 0xc4, 0xee, 0x54, 0x70, 0xe3, 0xaa, 0xa0, 0x0b, 0x17, 0x86, 0xb4, 0x73, 0xa9,
	0x83, 0x93, 0x26, 0xce, 0x4c, 0xfa, 0xc2, 0xbe, 0x88, 0x64, 0x32, 0x6d, 0x9a, 0xa8, 0x60, 0x77,
	0xc9, 0xef, 0xcc, 0x39, 0x77, 0xce, 0x0d, 0x81, 0xa3, 0x0d, 0xbe, 0x89, 0xa5, 0xc4, 0x71, 0xaa,
	0x12, 0x93, 0xb0, 0x4f, 0x0f, 0x3a, 0xcf, 0x05, 0x21, 0xc7, 0xe0, 0x0b, 0x4e, 0xbd, 0x91, 0x77,
	0xd1, 0x9b, 0xfb, 0x82, 0x93, 0x33, 0xe8, 0x71, 0x25, 0x36, 0xa8, 0x42, 0xc1, 0xa9, 0x6f, 0x71,
	0x09, 0x48, 0x00, 0xad, 0x38, 0xe1, 0x28, 0x69, 0xc3, 0x2a, 0xc5, 0x0b, 0x21, 0xd0, 0x8c, 0xa3,
	0x77, 0xa4, 0x4d, 0x0b, 0xed, 0x33, 0x61, 0xf0, 0x3f, 0x95, 0x91, 0xc1, 0x70, 0x9d, 0xc5, 0x0b,
	0x54, 0xb4, 0x65, 0xb5, 0x0a, 0xcb, 0x67, 0x89, 0x38, 0x5a, 0x61, 0x98, 0x29, 0x49, 0xdb, 0xc5,
	0xac, 0x1d, 0x20, 0x43, 0x80, 0xa5, 0xc2, 0xc8, 0x20, 0x0f, 0x23, 0x43, 0x3b, 0x56, 0xde, 0x23,
	0xb9, 0x9e, 0xa5, 0x7c, 0xab, 0x77, 0x0b, 0xbd, 0x24, 0xec, 0x15, 0x82, 0x3b, 0x7b, 0xda, 0x55,
	0x9d, 0xe3, 0x47, 0x86, 0xda, 0x54, 0x1b, 0x7a, 0xbf, 0x36, 0xf4, 0x7f, 0x6a, 0xd8, 0x28, 0x1b,
	0xb2, 0x19, 0x0c, 0x6a, 0xf9, 0x3a, 0x4d, 0xd6, 0x3a, 0xaf, 0xde, 0x71, 0xfb, 0xb6, 0xf1, 0xff,
	0x26, 0xdd, 0xf1, 0xf6, 0xc8, 0x56, 0x60, 0x37, 0x10, 0x3c, 0xd9, 0xab, 0xd6, 0x2e, 0xf7, 0x17,
	0xef, 0x29, 0x0c, 0x6a, 0xde, 0x62, 0x30, 0xbb, 0x82, 0xe0, 0x1e, 0x25, 0x7e, 0x0b, 0x1d, 0x02,
	0x38, 0x6f, 0x59, 0x79, 0x8f, 0xe4, 0x81, 0x35, 0x9f, 0x0b, 0x9c, 0xc2, 0xc9, 0x03, 0x9a, 0x03,
	0xd3, 0xae, 0x81, 0xec, 0x9b, 0x0e, 0x58, 0xca, 0x04, 0xc8, 0xa3, 0xd0, 0xe6, 0x90, 0xef, 0xc5,
	0x66, 0xd0, 0xaf, 0x78, 0xdc, 0xb8, 0x73, 0xe8, 0xba, 0x54, 0x4d, 0xbd, 0x51, 0xa3, 0x32, 0x6f,
	0xa7, 0xdc, 0x0e, 0x5e, 0xfa, 0xe3, 0x4b, 0x99, 0xac, 0x84, 0x36, 0x62, 0xa9, 0x43, 0xc7, 0x17,
	0x6d, 0xfb, 0x9b, 0x4c, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x90, 0xec, 0x66, 0x37, 0x03,
	0x00, 0x00,
}
