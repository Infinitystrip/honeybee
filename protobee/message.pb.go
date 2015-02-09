// Code generated by protoc-gen-go.
// source: protobee/message.proto
// DO NOT EDIT!

/*
Package protobee is a generated protocol buffer package.

It is generated from these files:
	protobee/message.proto

It has these top-level messages:
	Message
	Server
	Connection
	SystemInfo
	Cpu
	FileSystem
	Memory
	Network
	Platform
*/
package protobee

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Message_Type int32

const (
	Message_SERVER      Message_Type = 0
	Message_SYSTEM_INFO Message_Type = 1
)

var Message_Type_name = map[int32]string{
	0: "SERVER",
	1: "SYSTEM_INFO",
}
var Message_Type_value = map[string]int32{
	"SERVER":      0,
	"SYSTEM_INFO": 1,
}

func (x Message_Type) Enum() *Message_Type {
	p := new(Message_Type)
	*p = x
	return p
}
func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}
func (x *Message_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_Type_value, data, "Message_Type")
	if err != nil {
		return err
	}
	*x = Message_Type(value)
	return nil
}

type Message struct {
	Type             *Message_Type `protobuf:"varint,1,req,name=type,enum=protobee.Message_Type" json:"type,omitempty"`
	Server           *Server       `protobuf:"bytes,2,opt,name=server" json:"server,omitempty"`
	SystemInfo       *SystemInfo   `protobuf:"bytes,3,opt,name=system_info" json:"system_info,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetType() Message_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Message_SERVER
}

func (m *Message) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *Message) GetSystemInfo() *SystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return nil
}

type Server struct {
	Hostname         *string       `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	Label            *string       `protobuf:"bytes,2,req,name=label" json:"label,omitempty"`
	Connections      []*Connection `protobuf:"bytes,3,rep,name=connections" json:"connections,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}

func (m *Server) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Server) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Server) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

type Connection struct {
	Transport        *string `protobuf:"bytes,1,req,name=transport" json:"transport,omitempty"`
	LocalAddress     *string `protobuf:"bytes,2,req,name=local_address" json:"local_address,omitempty"`
	LocalPort        *uint32 `protobuf:"varint,3,req,name=local_port" json:"local_port,omitempty"`
	RemoteAddress    *string `protobuf:"bytes,4,req,name=remote_address" json:"remote_address,omitempty"`
	RemotePort       *uint32 `protobuf:"varint,5,req,name=remote_port" json:"remote_port,omitempty"`
	Pid              *uint32 `protobuf:"varint,6,req,name=pid" json:"pid,omitempty"`
	Name             *string `protobuf:"bytes,7,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}

func (m *Connection) GetTransport() string {
	if m != nil && m.Transport != nil {
		return *m.Transport
	}
	return ""
}

func (m *Connection) GetLocalAddress() string {
	if m != nil && m.LocalAddress != nil {
		return *m.LocalAddress
	}
	return ""
}

func (m *Connection) GetLocalPort() uint32 {
	if m != nil && m.LocalPort != nil {
		return *m.LocalPort
	}
	return 0
}

func (m *Connection) GetRemoteAddress() string {
	if m != nil && m.RemoteAddress != nil {
		return *m.RemoteAddress
	}
	return ""
}

func (m *Connection) GetRemotePort() uint32 {
	if m != nil && m.RemotePort != nil {
		return *m.RemotePort
	}
	return 0
}

func (m *Connection) GetPid() uint32 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

func (m *Connection) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type SystemInfo struct {
	Cpu              *Cpu          `protobuf:"bytes,1,req,name=cpu" json:"cpu,omitempty"`
	FileSystems      []*FileSystem `protobuf:"bytes,2,rep,name=file_systems" json:"file_systems,omitempty"`
	Memory           *Memory       `protobuf:"bytes,3,req,name=memory" json:"memory,omitempty"`
	Networks         []*Network    `protobuf:"bytes,4,rep,name=networks" json:"networks,omitempty"`
	Platform         *Platform     `protobuf:"bytes,5,req,name=platform" json:"platform,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SystemInfo) Reset()         { *m = SystemInfo{} }
func (m *SystemInfo) String() string { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()    {}

func (m *SystemInfo) GetCpu() *Cpu {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *SystemInfo) GetFileSystems() []*FileSystem {
	if m != nil {
		return m.FileSystems
	}
	return nil
}

func (m *SystemInfo) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *SystemInfo) GetNetworks() []*Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *SystemInfo) GetPlatform() *Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

type Cpu struct {
	CpuCores         *uint32 `protobuf:"varint,1,req,name=cpu_cores" json:"cpu_cores,omitempty"`
	Family           *uint32 `protobuf:"varint,2,req,name=family" json:"family,omitempty"`
	Mhz              *uint32 `protobuf:"varint,3,req,name=mhz" json:"mhz,omitempty"`
	Model            *uint32 `protobuf:"varint,4,req,name=model" json:"model,omitempty"`
	ModelName        *string `protobuf:"bytes,5,req,name=model_name" json:"model_name,omitempty"`
	Stepping         *uint32 `protobuf:"varint,6,req,name=stepping" json:"stepping,omitempty"`
	VendorId         *string `protobuf:"bytes,7,req,name=vendor_id" json:"vendor_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Cpu) Reset()         { *m = Cpu{} }
func (m *Cpu) String() string { return proto.CompactTextString(m) }
func (*Cpu) ProtoMessage()    {}

func (m *Cpu) GetCpuCores() uint32 {
	if m != nil && m.CpuCores != nil {
		return *m.CpuCores
	}
	return 0
}

func (m *Cpu) GetFamily() uint32 {
	if m != nil && m.Family != nil {
		return *m.Family
	}
	return 0
}

func (m *Cpu) GetMhz() uint32 {
	if m != nil && m.Mhz != nil {
		return *m.Mhz
	}
	return 0
}

func (m *Cpu) GetModel() uint32 {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return 0
}

func (m *Cpu) GetModelName() string {
	if m != nil && m.ModelName != nil {
		return *m.ModelName
	}
	return ""
}

func (m *Cpu) GetStepping() uint32 {
	if m != nil && m.Stepping != nil {
		return *m.Stepping
	}
	return 0
}

func (m *Cpu) GetVendorId() string {
	if m != nil && m.VendorId != nil {
		return *m.VendorId
	}
	return ""
}

type FileSystem struct {
	KbSize           *uint32 `protobuf:"varint,1,req,name=kb_size" json:"kb_size,omitempty"`
	MountedOn        *string `protobuf:"bytes,2,req,name=mounted_on" json:"mounted_on,omitempty"`
	Name             *string `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FileSystem) Reset()         { *m = FileSystem{} }
func (m *FileSystem) String() string { return proto.CompactTextString(m) }
func (*FileSystem) ProtoMessage()    {}

func (m *FileSystem) GetKbSize() uint32 {
	if m != nil && m.KbSize != nil {
		return *m.KbSize
	}
	return 0
}

func (m *FileSystem) GetMountedOn() string {
	if m != nil && m.MountedOn != nil {
		return *m.MountedOn
	}
	return ""
}

func (m *FileSystem) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type Memory struct {
	SwapTotal        *string `protobuf:"bytes,1,req,name=swap_total" json:"swap_total,omitempty"`
	Total            *uint32 `protobuf:"varint,2,req,name=total" json:"total,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}

func (m *Memory) GetSwapTotal() string {
	if m != nil && m.SwapTotal != nil {
		return *m.SwapTotal
	}
	return ""
}

func (m *Memory) GetTotal() uint32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

type Network struct {
	MacAddress       *string `protobuf:"bytes,1,req,name=mac_address" json:"mac_address,omitempty"`
	IpAddress        *string `protobuf:"bytes,2,req,name=ip_address" json:"ip_address,omitempty"`
	IpAddressV6      *string `protobuf:"bytes,3,req,name=ip_address_v6" json:"ip_address_v6,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}

func (m *Network) GetMacAddress() string {
	if m != nil && m.MacAddress != nil {
		return *m.MacAddress
	}
	return ""
}

func (m *Network) GetIpAddress() string {
	if m != nil && m.IpAddress != nil {
		return *m.IpAddress
	}
	return ""
}

func (m *Network) GetIpAddressV6() string {
	if m != nil && m.IpAddressV6 != nil {
		return *m.IpAddressV6
	}
	return ""
}

type Platform struct {
	Hostname         *string `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	KernelName       *string `protobuf:"bytes,2,req,name=kernel_name" json:"kernel_name,omitempty"`
	KernelRelease    *string `protobuf:"bytes,3,req,name=kernel_release" json:"kernel_release,omitempty"`
	Machine          *string `protobuf:"bytes,4,req,name=machine" json:"machine,omitempty"`
	Processor        *string `protobuf:"bytes,5,req,name=processor" json:"processor,omitempty"`
	Os               *string `protobuf:"bytes,6,req,name=os" json:"os,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}

func (m *Platform) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Platform) GetKernelName() string {
	if m != nil && m.KernelName != nil {
		return *m.KernelName
	}
	return ""
}

func (m *Platform) GetKernelRelease() string {
	if m != nil && m.KernelRelease != nil {
		return *m.KernelRelease
	}
	return ""
}

func (m *Platform) GetMachine() string {
	if m != nil && m.Machine != nil {
		return *m.Machine
	}
	return ""
}

func (m *Platform) GetProcessor() string {
	if m != nil && m.Processor != nil {
		return *m.Processor
	}
	return ""
}

func (m *Platform) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

func init() {
	proto.RegisterEnum("protobee.Message_Type", Message_Type_name, Message_Type_value)
}