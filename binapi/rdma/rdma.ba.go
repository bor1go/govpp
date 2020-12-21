// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0
//  VPP:              20.05.1-release
// source: /usr/share/vpp/api/plugins/rdma.api.json

// Package rdma contains generated bindings for API file rdma.api.
//
// Contents:
//   1 enum
//   4 messages
//
package rdma

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	interface_types "git.fd.io/govpp.git/binapi/interface_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "rdma"
	APIVersion = "1.0.0"
	VersionCrc = 0x8b33158c
)

// RdmaMode defines enum 'rdma_mode'.
type RdmaMode uint32

const (
	RDMA_API_MODE_AUTO RdmaMode = 0
	RDMA_API_MODE_IBV  RdmaMode = 1
	RDMA_API_MODE_DV   RdmaMode = 2
)

var (
	RdmaMode_name = map[uint32]string{
		0: "RDMA_API_MODE_AUTO",
		1: "RDMA_API_MODE_IBV",
		2: "RDMA_API_MODE_DV",
	}
	RdmaMode_value = map[string]uint32{
		"RDMA_API_MODE_AUTO": 0,
		"RDMA_API_MODE_IBV":  1,
		"RDMA_API_MODE_DV":   2,
	}
)

func (x RdmaMode) String() string {
	s, ok := RdmaMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RdmaMode(" + strconv.Itoa(int(x)) + ")"
}

// RdmaCreate defines message 'rdma_create'.
type RdmaCreate struct {
	HostIf  string   `binapi:"string[64],name=host_if" json:"host_if,omitempty"`
	Name    string   `binapi:"string[64],name=name" json:"name,omitempty"`
	RxqNum  uint16   `binapi:"u16,name=rxq_num,default=1" json:"rxq_num,omitempty"`
	RxqSize uint16   `binapi:"u16,name=rxq_size,default=1024" json:"rxq_size,omitempty"`
	TxqSize uint16   `binapi:"u16,name=txq_size,default=1024" json:"txq_size,omitempty"`
	Mode    RdmaMode `binapi:"rdma_mode,name=mode,default=0" json:"mode,omitempty"`
}

func (m *RdmaCreate) Reset()               { *m = RdmaCreate{} }
func (*RdmaCreate) GetMessageName() string { return "rdma_create" }
func (*RdmaCreate) GetCrcString() string   { return "076fe418" }
func (*RdmaCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *RdmaCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.HostIf
	size += 64 // m.Name
	size += 2  // m.RxqNum
	size += 2  // m.RxqSize
	size += 2  // m.TxqSize
	size += 4  // m.Mode
	return size
}
func (m *RdmaCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.HostIf, 64)
	buf.EncodeString(m.Name, 64)
	buf.EncodeUint16(m.RxqNum)
	buf.EncodeUint16(m.RxqSize)
	buf.EncodeUint16(m.TxqSize)
	buf.EncodeUint32(uint32(m.Mode))
	return buf.Bytes(), nil
}
func (m *RdmaCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HostIf = buf.DecodeString(64)
	m.Name = buf.DecodeString(64)
	m.RxqNum = buf.DecodeUint16()
	m.RxqSize = buf.DecodeUint16()
	m.TxqSize = buf.DecodeUint16()
	m.Mode = RdmaMode(buf.DecodeUint32())
	return nil
}

// RdmaCreateReply defines message 'rdma_create_reply'.
type RdmaCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *RdmaCreateReply) Reset()               { *m = RdmaCreateReply{} }
func (*RdmaCreateReply) GetMessageName() string { return "rdma_create_reply" }
func (*RdmaCreateReply) GetCrcString() string   { return "5383d31f" }
func (*RdmaCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *RdmaCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *RdmaCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *RdmaCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// RdmaDelete defines message 'rdma_delete'.
type RdmaDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *RdmaDelete) Reset()               { *m = RdmaDelete{} }
func (*RdmaDelete) GetMessageName() string { return "rdma_delete" }
func (*RdmaDelete) GetCrcString() string   { return "f9e6675e" }
func (*RdmaDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *RdmaDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *RdmaDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *RdmaDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// RdmaDeleteReply defines message 'rdma_delete_reply'.
type RdmaDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *RdmaDeleteReply) Reset()               { *m = RdmaDeleteReply{} }
func (*RdmaDeleteReply) GetMessageName() string { return "rdma_delete_reply" }
func (*RdmaDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*RdmaDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *RdmaDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *RdmaDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *RdmaDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_rdma_binapi_init() }
func file_rdma_binapi_init() {
	api.RegisterMessage((*RdmaCreate)(nil), "rdma_create_076fe418")
	api.RegisterMessage((*RdmaCreateReply)(nil), "rdma_create_reply_5383d31f")
	api.RegisterMessage((*RdmaDelete)(nil), "rdma_delete_f9e6675e")
	api.RegisterMessage((*RdmaDeleteReply)(nil), "rdma_delete_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*RdmaCreate)(nil),
		(*RdmaCreateReply)(nil),
		(*RdmaDelete)(nil),
		(*RdmaDeleteReply)(nil),
	}
}