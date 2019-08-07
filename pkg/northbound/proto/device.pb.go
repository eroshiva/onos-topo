// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/northbound/proto/device.proto

// Package admin defines the administrative and diagnostic gRPC interfaces.

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Device event type
type ListResponse_Type int32

const (
	// NONE indicates this response does not represent a state change
	ListResponse_NONE ListResponse_Type = 0
	// ADDED is an event which occurs when a device is added to the topology
	ListResponse_ADDED ListResponse_Type = 1
	// UPDATED is an event which occurs when a device is updated
	ListResponse_UPDATED ListResponse_Type = 2
	// REMOVED is an event which occurs when a device is removed from the topology
	ListResponse_REMOVED ListResponse_Type = 3
)

var ListResponse_Type_name = map[int32]string{
	0: "NONE",
	1: "ADDED",
	2: "UPDATED",
	3: "REMOVED",
}

var ListResponse_Type_value = map[string]int32{
	"NONE":    0,
	"ADDED":   1,
	"UPDATED": 2,
	"REMOVED": 3,
}

func (x ListResponse_Type) String() string {
	return proto.EnumName(ListResponse_Type_name, int32(x))
}

func (ListResponse_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{7, 0}
}

// AddDeviceRequest adds a device to the topology
type AddDeviceRequest struct {
	// device is the device to add
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDeviceRequest) Reset()         { *m = AddDeviceRequest{} }
func (m *AddDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*AddDeviceRequest) ProtoMessage()    {}
func (*AddDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{0}
}

func (m *AddDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDeviceRequest.Unmarshal(m, b)
}
func (m *AddDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDeviceRequest.Marshal(b, m, deterministic)
}
func (m *AddDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDeviceRequest.Merge(m, src)
}
func (m *AddDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_AddDeviceRequest.Size(m)
}
func (m *AddDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDeviceRequest proto.InternalMessageInfo

func (m *AddDeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

// AddDeviceResponse is sent in response to an AddDeviceRequest
type AddDeviceResponse struct {
	// metadata is the added device metadata
	Metadata             *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddDeviceResponse) Reset()         { *m = AddDeviceResponse{} }
func (m *AddDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*AddDeviceResponse) ProtoMessage()    {}
func (*AddDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{1}
}

func (m *AddDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDeviceResponse.Unmarshal(m, b)
}
func (m *AddDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDeviceResponse.Marshal(b, m, deterministic)
}
func (m *AddDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDeviceResponse.Merge(m, src)
}
func (m *AddDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_AddDeviceResponse.Size(m)
}
func (m *AddDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddDeviceResponse proto.InternalMessageInfo

func (m *AddDeviceResponse) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// UpdateDeviceRequest updates a device
type UpdateDeviceRequest struct {
	// device is the updated device
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDeviceRequest) Reset()         { *m = UpdateDeviceRequest{} }
func (m *UpdateDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDeviceRequest) ProtoMessage()    {}
func (*UpdateDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{2}
}

func (m *UpdateDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeviceRequest.Unmarshal(m, b)
}
func (m *UpdateDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeviceRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeviceRequest.Merge(m, src)
}
func (m *UpdateDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDeviceRequest.Size(m)
}
func (m *UpdateDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeviceRequest proto.InternalMessageInfo

func (m *UpdateDeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

// UpdateDeviceResponse is sent in response to an UpdateDeviceRequest
type UpdateDeviceResponse struct {
	// metadata is the updated device metadata
	Metadata             *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateDeviceResponse) Reset()         { *m = UpdateDeviceResponse{} }
func (m *UpdateDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateDeviceResponse) ProtoMessage()    {}
func (*UpdateDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{3}
}

func (m *UpdateDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeviceResponse.Unmarshal(m, b)
}
func (m *UpdateDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeviceResponse.Marshal(b, m, deterministic)
}
func (m *UpdateDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeviceResponse.Merge(m, src)
}
func (m *UpdateDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateDeviceResponse.Size(m)
}
func (m *UpdateDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeviceResponse proto.InternalMessageInfo

func (m *UpdateDeviceResponse) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// GetDeviceRequest gets a device by ID
type GetDeviceRequest struct {
	// device_id is the unique device ID with which to lookup the device
	DeviceId             string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceRequest) Reset()         { *m = GetDeviceRequest{} }
func (m *GetDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceRequest) ProtoMessage()    {}
func (*GetDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{4}
}

func (m *GetDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceRequest.Unmarshal(m, b)
}
func (m *GetDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceRequest.Merge(m, src)
}
func (m *GetDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceRequest.Size(m)
}
func (m *GetDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceRequest proto.InternalMessageInfo

func (m *GetDeviceRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

// GetDeviceResponse carries a device
type GetDeviceResponse struct {
	// device is the device object
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceResponse) Reset()         { *m = GetDeviceResponse{} }
func (m *GetDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceResponse) ProtoMessage()    {}
func (*GetDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{5}
}

func (m *GetDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceResponse.Unmarshal(m, b)
}
func (m *GetDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceResponse.Merge(m, src)
}
func (m *GetDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceResponse.Size(m)
}
func (m *GetDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceResponse proto.InternalMessageInfo

func (m *GetDeviceResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

// ListRequest requests a stream of devices and changes
// By default, the request requests a stream of all devices that are present in the topology when
// the request is received by the service. However, if `subscribe` is `true`, the stream will remain
// open after all devices have been sent and events that occur following the last device will be
// streamed to the client until the stream is closed.
type ListRequest struct {
	// subscribe indicates whether to subscribe to events (e.g. ADD, UPDATE, and REMOVE) that occur
	// after all devices have been streamed to the client
	Subscribe            bool     `protobuf:"varint,1,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{6}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetSubscribe() bool {
	if m != nil {
		return m.Subscribe
	}
	return false
}

// ListResponse carries a single device event
type ListResponse struct {
	// type is the type of the event
	Type ListResponse_Type `protobuf:"varint,1,opt,name=type,proto3,enum=proto.ListResponse_Type" json:"type,omitempty"`
	// device is the device on which the event occurred
	Device               *Device  `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{7}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetType() ListResponse_Type {
	if m != nil {
		return m.Type
	}
	return ListResponse_NONE
}

func (m *ListResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

// RemoveDeviceRequest removes a device by ID
type RemoveDeviceRequest struct {
	// device is the device to remove
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveDeviceRequest) Reset()         { *m = RemoveDeviceRequest{} }
func (m *RemoveDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveDeviceRequest) ProtoMessage()    {}
func (*RemoveDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{8}
}

func (m *RemoveDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDeviceRequest.Unmarshal(m, b)
}
func (m *RemoveDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDeviceRequest.Marshal(b, m, deterministic)
}
func (m *RemoveDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDeviceRequest.Merge(m, src)
}
func (m *RemoveDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveDeviceRequest.Size(m)
}
func (m *RemoveDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDeviceRequest proto.InternalMessageInfo

func (m *RemoveDeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

// RemoveDeviceResponse is sent in response to a RemoveDeviceRequest
type RemoveDeviceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveDeviceResponse) Reset()         { *m = RemoveDeviceResponse{} }
func (m *RemoveDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveDeviceResponse) ProtoMessage()    {}
func (*RemoveDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{9}
}

func (m *RemoveDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDeviceResponse.Unmarshal(m, b)
}
func (m *RemoveDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDeviceResponse.Marshal(b, m, deterministic)
}
func (m *RemoveDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDeviceResponse.Merge(m, src)
}
func (m *RemoveDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveDeviceResponse.Size(m)
}
func (m *RemoveDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDeviceResponse proto.InternalMessageInfo

// Device contains information about a device
type Device struct {
	// metadata is the store metadata used for concurrency control
	Metadata *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// id is a globally unique device identifier
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// address is the host:port of the device
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// target is the device target
	Target string `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	// software_version is the device software version
	SoftwareVersion string `protobuf:"bytes,5,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty"`
	// timeout indicates the device request timeout
	Timeout *duration.Duration `protobuf:"bytes,6,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// credentials contains the credentials for connecting to the device
	Credentials *Credentials `protobuf:"bytes,7,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// tls is the device TLS configuration
	Tls                  *TlsConfig `protobuf:"bytes,8,opt,name=tls,proto3" json:"tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{10}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Device) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Device) GetSoftwareVersion() string {
	if m != nil {
		return m.SoftwareVersion
	}
	return ""
}

func (m *Device) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *Device) GetCredentials() *Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *Device) GetTls() *TlsConfig {
	if m != nil {
		return m.Tls
	}
	return nil
}

// Credentials is the device credentials
type Credentials struct {
	// user is the user with which to connect to the device
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// password is the password for connecting to the device
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{11}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// Device TLS configuration
type TlsConfig struct {
	// caCert is the name of the device's CA certificate
	CaCert string `protobuf:"bytes,3,opt,name=caCert,proto3" json:"caCert,omitempty"`
	// cert is the name of the device's certificate
	Cert string `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	// key is the name of the device's TLS key
	Key string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	// plain indicates whether to connect to the device over plaintext
	Plain bool `protobuf:"varint,6,opt,name=plain,proto3" json:"plain,omitempty"`
	// insecure indicates whether to connect to the device with insecure communication
	Insecure             bool     `protobuf:"varint,7,opt,name=insecure,proto3" json:"insecure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsConfig) Reset()         { *m = TlsConfig{} }
func (m *TlsConfig) String() string { return proto.CompactTextString(m) }
func (*TlsConfig) ProtoMessage()    {}
func (*TlsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{12}
}

func (m *TlsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsConfig.Unmarshal(m, b)
}
func (m *TlsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsConfig.Marshal(b, m, deterministic)
}
func (m *TlsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsConfig.Merge(m, src)
}
func (m *TlsConfig) XXX_Size() int {
	return xxx_messageInfo_TlsConfig.Size(m)
}
func (m *TlsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TlsConfig proto.InternalMessageInfo

func (m *TlsConfig) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *TlsConfig) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *TlsConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TlsConfig) GetPlain() bool {
	if m != nil {
		return m.Plain
	}
	return false
}

func (m *TlsConfig) GetInsecure() bool {
	if m != nil {
		return m.Insecure
	}
	return false
}

// ObjectMetadata is the metadata required by the store for concurrency control
type ObjectMetadata struct {
	// id is the unique identifier for the object
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// version is the store version of the object
	Version              uint64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectMetadata) Reset()         { *m = ObjectMetadata{} }
func (m *ObjectMetadata) String() string { return proto.CompactTextString(m) }
func (*ObjectMetadata) ProtoMessage()    {}
func (*ObjectMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd506e14b3f7d725, []int{13}
}

func (m *ObjectMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMetadata.Unmarshal(m, b)
}
func (m *ObjectMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMetadata.Marshal(b, m, deterministic)
}
func (m *ObjectMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMetadata.Merge(m, src)
}
func (m *ObjectMetadata) XXX_Size() int {
	return xxx_messageInfo_ObjectMetadata.Size(m)
}
func (m *ObjectMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMetadata proto.InternalMessageInfo

func (m *ObjectMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectMetadata) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.ListResponse_Type", ListResponse_Type_name, ListResponse_Type_value)
	proto.RegisterType((*AddDeviceRequest)(nil), "proto.AddDeviceRequest")
	proto.RegisterType((*AddDeviceResponse)(nil), "proto.AddDeviceResponse")
	proto.RegisterType((*UpdateDeviceRequest)(nil), "proto.UpdateDeviceRequest")
	proto.RegisterType((*UpdateDeviceResponse)(nil), "proto.UpdateDeviceResponse")
	proto.RegisterType((*GetDeviceRequest)(nil), "proto.GetDeviceRequest")
	proto.RegisterType((*GetDeviceResponse)(nil), "proto.GetDeviceResponse")
	proto.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto.RegisterType((*RemoveDeviceRequest)(nil), "proto.RemoveDeviceRequest")
	proto.RegisterType((*RemoveDeviceResponse)(nil), "proto.RemoveDeviceResponse")
	proto.RegisterType((*Device)(nil), "proto.Device")
	proto.RegisterType((*Credentials)(nil), "proto.Credentials")
	proto.RegisterType((*TlsConfig)(nil), "proto.TlsConfig")
	proto.RegisterType((*ObjectMetadata)(nil), "proto.ObjectMetadata")
}

func init() { proto.RegisterFile("pkg/northbound/proto/device.proto", fileDescriptor_fd506e14b3f7d725) }

var fileDescriptor_fd506e14b3f7d725 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x6f, 0x6f, 0xd3, 0x3e,
	0x10, 0xee, 0x9f, 0xb4, 0x4d, 0xaf, 0xbf, 0xed, 0x97, 0x79, 0x63, 0x84, 0x0e, 0x21, 0x88, 0x84,
	0x04, 0x02, 0xb5, 0xb0, 0xc1, 0x0b, 0x2a, 0x78, 0x51, 0x35, 0x65, 0x9a, 0xc4, 0x36, 0x64, 0xb6,
	0xbd, 0x9d, 0xd2, 0xd8, 0x2b, 0x61, 0x5d, 0x1c, 0x6c, 0x67, 0xd3, 0xc4, 0x37, 0xe1, 0x8b, 0xf0,
	0x0d, 0xf8, 0x5c, 0x28, 0xb6, 0x93, 0xa5, 0x5b, 0x85, 0x50, 0x5f, 0x25, 0x77, 0xf7, 0x3c, 0xbe,
	0xc7, 0x77, 0xbe, 0x83, 0x27, 0xc9, 0xf9, 0xb4, 0x1f, 0x33, 0x2e, 0xbf, 0x4e, 0x58, 0x1a, 0x93,
	0x7e, 0xc2, 0x99, 0x64, 0x7d, 0x42, 0x2f, 0xa3, 0x90, 0xf6, 0x94, 0x81, 0x1a, 0xea, 0xd3, 0x7d,
	0x34, 0x65, 0x6c, 0x3a, 0xa3, 0x1a, 0x31, 0x49, 0xcf, 0xfa, 0x24, 0xe5, 0x81, 0x8c, 0x58, 0xac,
	0x61, 0xde, 0x3b, 0x70, 0x86, 0x84, 0xf8, 0x8a, 0x89, 0xe9, 0xf7, 0x94, 0x0a, 0x89, 0x9e, 0x42,
	0x53, 0x1f, 0xe5, 0x56, 0x1f, 0x57, 0x9f, 0x75, 0xb6, 0x57, 0x34, 0xb6, 0x67, 0x50, 0x26, 0xe8,
	0x7d, 0x84, 0xb5, 0x12, 0x55, 0x24, 0x2c, 0x16, 0x14, 0xbd, 0x06, 0xfb, 0x82, 0xca, 0x80, 0x04,
	0x32, 0x30, 0xec, 0x7b, 0x86, 0x7d, 0x38, 0xf9, 0x46, 0x43, 0xb9, 0x6f, 0x82, 0xb8, 0x80, 0x79,
	0xef, 0x61, 0xfd, 0x38, 0x21, 0x81, 0xa4, 0x4b, 0xa9, 0xd8, 0x83, 0x8d, 0x79, 0xf6, 0xf2, 0x42,
	0xfa, 0xe0, 0xec, 0x52, 0x39, 0xaf, 0x62, 0x0b, 0xda, 0x3a, 0xd1, 0x69, 0x44, 0xd4, 0x39, 0x6d,
	0x6c, 0x6b, 0xc7, 0x1e, 0xf1, 0x06, 0xb0, 0x56, 0x22, 0x98, 0xc4, 0xff, 0xa8, 0xfb, 0x05, 0x74,
	0x3e, 0x45, 0x42, 0xe6, 0x79, 0x1e, 0x42, 0x5b, 0xa4, 0x13, 0x11, 0xf2, 0x68, 0xa2, 0x89, 0x36,
	0xbe, 0x71, 0x78, 0x3f, 0xab, 0xf0, 0x9f, 0x46, 0x9b, 0x24, 0x2f, 0xc1, 0x92, 0xd7, 0x89, 0x46,
	0xae, 0x6e, 0xbb, 0x26, 0x45, 0x19, 0xd2, 0x3b, 0xba, 0x4e, 0x28, 0x56, 0xa8, 0x92, 0xa4, 0xda,
	0xdf, 0x24, 0xbd, 0x05, 0x2b, 0x23, 0x21, 0x1b, 0xac, 0x83, 0xc3, 0x83, 0xb1, 0x53, 0x41, 0x6d,
	0x68, 0x0c, 0x7d, 0x7f, 0xec, 0x3b, 0x55, 0xd4, 0x81, 0xd6, 0xf1, 0x67, 0x7f, 0x78, 0x34, 0xf6,
	0x9d, 0x5a, 0x66, 0xe0, 0xf1, 0xfe, 0xe1, 0xc9, 0xd8, 0x77, 0xea, 0x59, 0xff, 0x30, 0xbd, 0x60,
	0x97, 0xcb, 0xf5, 0x6f, 0x13, 0x36, 0xe6, 0xd9, 0x5a, 0xbe, 0xf7, 0xab, 0x06, 0x4d, 0xed, 0x5a,
	0xa2, 0x95, 0x68, 0x15, 0x6a, 0x11, 0x51, 0xb7, 0x6d, 0xe3, 0x5a, 0x44, 0x90, 0x0b, 0xad, 0x80,
	0x10, 0x4e, 0x85, 0x70, 0xeb, 0xca, 0x99, 0x9b, 0x68, 0x13, 0x9a, 0x32, 0xe0, 0x53, 0x2a, 0x5d,
	0x4b, 0x05, 0x8c, 0x85, 0x9e, 0x83, 0x23, 0xd8, 0x99, 0xbc, 0x0a, 0x38, 0x3d, 0xbd, 0xa4, 0x5c,
	0x44, 0x2c, 0x76, 0x1b, 0x0a, 0xf1, 0x7f, 0xee, 0x3f, 0xd1, 0x6e, 0xb4, 0x03, 0x2d, 0x19, 0x5d,
	0x50, 0x96, 0x4a, 0xb7, 0xa9, 0xe4, 0x3d, 0xe8, 0xe9, 0xa9, 0xeb, 0xe5, 0x53, 0xd7, 0xf3, 0xcd,
	0xd4, 0xe1, 0x1c, 0x89, 0xde, 0x40, 0x27, 0xe4, 0x94, 0xd0, 0x58, 0x46, 0xc1, 0x4c, 0xb8, 0x2d,
	0x45, 0x44, 0xe6, 0x5e, 0xa3, 0x9b, 0x08, 0x2e, 0xc3, 0x90, 0x07, 0x75, 0x39, 0x13, 0xae, 0xad,
	0xd0, 0x8e, 0x41, 0x1f, 0xcd, 0xc4, 0x88, 0xc5, 0x67, 0xd1, 0x14, 0x67, 0x41, 0xef, 0x03, 0x74,
	0x4a, 0x7c, 0x84, 0xc0, 0x4a, 0x05, 0xe5, 0xe6, 0xf1, 0xaa, 0x7f, 0xd4, 0x05, 0x3b, 0x09, 0x84,
	0xb8, 0x62, 0x3c, 0x2f, 0x52, 0x61, 0x7b, 0x3f, 0xa0, 0x5d, 0x1c, 0x98, 0x55, 0x27, 0x0c, 0x46,
	0x94, 0x4b, 0x53, 0x36, 0x63, 0x65, 0x87, 0x86, 0x99, 0x57, 0xd7, 0x4c, 0xfd, 0x23, 0x07, 0xea,
	0xe7, 0xf4, 0xda, 0x14, 0x29, 0xfb, 0x45, 0x1b, 0xd0, 0x48, 0x66, 0x41, 0x14, 0xab, 0xb2, 0xd8,
	0x58, 0x1b, 0x59, 0xf2, 0x28, 0x16, 0x34, 0x4c, 0x39, 0x55, 0xd7, 0xb6, 0x71, 0x61, 0x7b, 0x03,
	0x58, 0x9d, 0xef, 0xa9, 0xe9, 0x64, 0xb5, 0xdc, 0xc9, 0xbc, 0x1d, 0x99, 0x72, 0x0b, 0xe7, 0xe6,
	0xf6, 0xef, 0x1a, 0xac, 0xe8, 0x17, 0xf3, 0x85, 0x72, 0xf5, 0x70, 0x06, 0x50, 0x1f, 0x12, 0x82,
	0xee, 0x9b, 0x3a, 0xdd, 0x5e, 0x74, 0x5d, 0xf7, 0x6e, 0xc0, 0xbc, 0xbe, 0x0a, 0x1a, 0x41, 0x53,
	0xef, 0x15, 0xd4, 0x35, 0xa8, 0x05, 0x4b, 0xaa, 0xbb, 0xb5, 0x30, 0x56, 0x1c, 0x32, 0x80, 0xfa,
	0x2e, 0x95, 0x85, 0x80, 0xdb, 0xdb, 0xa5, 0x10, 0x70, 0x67, 0x8b, 0x78, 0x15, 0xb4, 0x03, 0x56,
	0x36, 0xcf, 0x08, 0xcd, 0x0d, 0xb7, 0xe6, 0xad, 0x2f, 0x18, 0x78, 0xaf, 0xf2, 0xaa, 0x9a, 0xa9,
	0xd6, 0xd3, 0x54, 0xa8, 0x5e, 0x30, 0x9a, 0x85, 0xea, 0x85, 0x83, 0x57, 0x99, 0x34, 0x55, 0x74,
	0xe7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xf8, 0xa3, 0x6f, 0x66, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceServiceClient interface {
	// Add adds a device to the topology
	Add(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error)
	// Update updates a device
	Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error)
	// Get gets a device by ID
	Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
	// List gets a stream of device add/update/remove events
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (DeviceService_ListClient, error)
	// Remove removes a device from the topology
	Remove(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error)
}

type deviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceServiceClient(cc *grpc.ClientConn) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) Add(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error) {
	out := new(AddDeviceResponse)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*UpdateDeviceResponse, error) {
	out := new(UpdateDeviceResponse)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (DeviceService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DeviceService_serviceDesc.Streams[0], "/proto.DeviceService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &deviceServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DeviceService_ListClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type deviceServiceListClient struct {
	grpc.ClientStream
}

func (x *deviceServiceListClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deviceServiceClient) Remove(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error) {
	out := new(RemoveDeviceResponse)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
type DeviceServiceServer interface {
	// Add adds a device to the topology
	Add(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error)
	// Update updates a device
	Update(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error)
	// Get gets a device by ID
	Get(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	// List gets a stream of device add/update/remove events
	List(*ListRequest, DeviceService_ListServer) error
	// Remove removes a device from the topology
	Remove(context.Context, *RemoveDeviceRequest) (*RemoveDeviceResponse, error)
}

// UnimplementedDeviceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (*UnimplementedDeviceServiceServer) Add(ctx context.Context, req *AddDeviceRequest) (*AddDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedDeviceServiceServer) Update(ctx context.Context, req *UpdateDeviceRequest) (*UpdateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDeviceServiceServer) Get(ctx context.Context, req *GetDeviceRequest) (*GetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDeviceServiceServer) List(req *ListRequest, srv DeviceService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedDeviceServiceServer) Remove(ctx context.Context, req *RemoveDeviceRequest) (*RemoveDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Add(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Update(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Get(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeviceServiceServer).List(m, &deviceServiceListServer{stream})
}

type DeviceService_ListServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type deviceServiceListServer struct {
	grpc.ServerStream
}

func (x *deviceServiceListServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DeviceService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Remove(ctx, req.(*RemoveDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _DeviceService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeviceService_Get_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _DeviceService_Remove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _DeviceService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/northbound/proto/device.proto",
}
