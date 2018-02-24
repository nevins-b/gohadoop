// Code generated by protoc-gen-go.
// source: HAServiceProtocol.proto
// DO NOT EDIT!

package hadoop_common

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type HAServiceStateProto int32

const (
	HAServiceStateProto_INITIALIZING HAServiceStateProto = 0
	HAServiceStateProto_ACTIVE       HAServiceStateProto = 1
	HAServiceStateProto_STANDBY      HAServiceStateProto = 2
)

var HAServiceStateProto_name = map[int32]string{
	0: "INITIALIZING",
	1: "ACTIVE",
	2: "STANDBY",
}
var HAServiceStateProto_value = map[string]int32{
	"INITIALIZING": 0,
	"ACTIVE":       1,
	"STANDBY":      2,
}

func (x HAServiceStateProto) Enum() *HAServiceStateProto {
	p := new(HAServiceStateProto)
	*p = x
	return p
}
func (x HAServiceStateProto) String() string {
	return proto.EnumName(HAServiceStateProto_name, int32(x))
}
func (x HAServiceStateProto) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *HAServiceStateProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HAServiceStateProto_value, data, "HAServiceStateProto")
	if err != nil {
		return err
	}
	*x = HAServiceStateProto(value)
	return nil
}

type HARequestSource int32

const (
	HARequestSource_REQUEST_BY_USER        HARequestSource = 0
	HARequestSource_REQUEST_BY_USER_FORCED HARequestSource = 1
	HARequestSource_REQUEST_BY_ZKFC        HARequestSource = 2
)

var HARequestSource_name = map[int32]string{
	0: "REQUEST_BY_USER",
	1: "REQUEST_BY_USER_FORCED",
	2: "REQUEST_BY_ZKFC",
}
var HARequestSource_value = map[string]int32{
	"REQUEST_BY_USER":        0,
	"REQUEST_BY_USER_FORCED": 1,
	"REQUEST_BY_ZKFC":        2,
}

func (x HARequestSource) Enum() *HARequestSource {
	p := new(HARequestSource)
	*p = x
	return p
}
func (x HARequestSource) String() string {
	return proto.EnumName(HARequestSource_name, int32(x))
}
func (x HARequestSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *HARequestSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HARequestSource_value, data, "HARequestSource")
	if err != nil {
		return err
	}
	*x = HARequestSource(value)
	return nil
}

type HAStateChangeRequestInfoProto struct {
	ReqSource        *HARequestSource `protobuf:"varint,1,req,name=reqSource,enum=hadoop.common.HARequestSource" json:"reqSource,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HAStateChangeRequestInfoProto) Reset()         { *m = HAStateChangeRequestInfoProto{} }
func (m *HAStateChangeRequestInfoProto) String() string { return proto.CompactTextString(m) }
func (*HAStateChangeRequestInfoProto) ProtoMessage()    {}

func (m *HAStateChangeRequestInfoProto) GetReqSource() HARequestSource {
	if m != nil && m.ReqSource != nil {
		return *m.ReqSource
	}
	return 0
}

// *
// void request
type MonitorHealthRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MonitorHealthRequestProto) Reset()         { *m = MonitorHealthRequestProto{} }
func (m *MonitorHealthRequestProto) String() string { return proto.CompactTextString(m) }
func (*MonitorHealthRequestProto) ProtoMessage()    {}

// *
// void response
type MonitorHealthResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MonitorHealthResponseProto) Reset()         { *m = MonitorHealthResponseProto{} }
func (m *MonitorHealthResponseProto) String() string { return proto.CompactTextString(m) }
func (*MonitorHealthResponseProto) ProtoMessage()    {}

// *
// void request
type TransitionToActiveRequestProto struct {
	ReqInfo          *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *TransitionToActiveRequestProto) Reset()         { *m = TransitionToActiveRequestProto{} }
func (m *TransitionToActiveRequestProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToActiveRequestProto) ProtoMessage()    {}

func (m *TransitionToActiveRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

// *
// void response
type TransitionToActiveResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TransitionToActiveResponseProto) Reset()         { *m = TransitionToActiveResponseProto{} }
func (m *TransitionToActiveResponseProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToActiveResponseProto) ProtoMessage()    {}

// *
// void request
type TransitionToStandbyRequestProto struct {
	ReqInfo          *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *TransitionToStandbyRequestProto) Reset()         { *m = TransitionToStandbyRequestProto{} }
func (m *TransitionToStandbyRequestProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToStandbyRequestProto) ProtoMessage()    {}

func (m *TransitionToStandbyRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

// *
// void response
type TransitionToStandbyResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TransitionToStandbyResponseProto) Reset()         { *m = TransitionToStandbyResponseProto{} }
func (m *TransitionToStandbyResponseProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToStandbyResponseProto) ProtoMessage()    {}

// *
// void request
type GetServiceStatusRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetServiceStatusRequestProto) Reset()         { *m = GetServiceStatusRequestProto{} }
func (m *GetServiceStatusRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetServiceStatusRequestProto) ProtoMessage()    {}

// *
// Returns the state of the service
type GetServiceStatusResponseProto struct {
	State *HAServiceStateProto `protobuf:"varint,1,req,name=state,enum=hadoop.common.HAServiceStateProto" json:"state,omitempty"`
	// If state is STANDBY, indicate whether it is
	// ready to become active.
	ReadyToBecomeActive *bool `protobuf:"varint,2,opt,name=readyToBecomeActive" json:"readyToBecomeActive,omitempty"`
	// If not ready to become active, a textual explanation of why not
	NotReadyReason   *string `protobuf:"bytes,3,opt,name=notReadyReason" json:"notReadyReason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetServiceStatusResponseProto) Reset()         { *m = GetServiceStatusResponseProto{} }
func (m *GetServiceStatusResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetServiceStatusResponseProto) ProtoMessage()    {}

func (m *GetServiceStatusResponseProto) GetState() HAServiceStateProto {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *GetServiceStatusResponseProto) GetReadyToBecomeActive() bool {
	if m != nil && m.ReadyToBecomeActive != nil {
		return *m.ReadyToBecomeActive
	}
	return false
}

func (m *GetServiceStatusResponseProto) GetNotReadyReason() string {
	if m != nil && m.NotReadyReason != nil {
		return *m.NotReadyReason
	}
	return ""
}

func init() {
	proto.RegisterEnum("hadoop.common.HAServiceStateProto", HAServiceStateProto_name, HAServiceStateProto_value)
	proto.RegisterEnum("hadoop.common.HARequestSource", HARequestSource_name, HARequestSource_value)
}
