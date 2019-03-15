// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_Type int32

const (
	Message_PING                          Message_Type = 0
	Message_PONG                          Message_Type = 1
	Message_THREAD_ENVELOPE               Message_Type = 10
	Message_CAFE_CHALLENGE                Message_Type = 50
	Message_CAFE_NONCE                    Message_Type = 51
	Message_CAFE_REGISTRATION             Message_Type = 52
	Message_CAFE_DEREGISTRATION           Message_Type = 72
	Message_CAFE_DEREGISTRATION_ACK       Message_Type = 73
	Message_CAFE_SESSION                  Message_Type = 53
	Message_CAFE_REFRESH_SESSION          Message_Type = 54
	Message_CAFE_STORE                    Message_Type = 55
	Message_CAFE_STORE_ACK                Message_Type = 59
	Message_CAFE_UNSTORE                  Message_Type = 74
	Message_CAFE_UNSTORE_ACK              Message_Type = 75
	Message_CAFE_OBJECT                   Message_Type = 56
	Message_CAFE_OBJECT_LIST              Message_Type = 57
	Message_CAFE_STORE_THREAD             Message_Type = 58
	Message_CAFE_STORE_THREAD_ACK         Message_Type = 76
	Message_CAFE_UNSTORE_THREAD           Message_Type = 77
	Message_CAFE_UNSTORE_THREAD_ACK       Message_Type = 78
	Message_CAFE_DELIVER_MESSAGE          Message_Type = 60
	Message_CAFE_CHECK_MESSAGES           Message_Type = 61
	Message_CAFE_MESSAGES                 Message_Type = 62
	Message_CAFE_DELETE_MESSAGES          Message_Type = 63
	Message_CAFE_DELETE_MESSAGES_ACK      Message_Type = 64
	Message_CAFE_YOU_HAVE_MAIL            Message_Type = 65
	Message_CAFE_PUBLISH_PEER             Message_Type = 66
	Message_CAFE_PUBLISH_PEER_ACK         Message_Type = 67
	Message_CAFE_QUERY                    Message_Type = 70
	Message_CAFE_QUERY_RES                Message_Type = 71
	Message_CAFE_PUBSUB_QUERY             Message_Type = 102
	Message_CAFE_PUBSUB_QUERY_RES         Message_Type = 103
	Message_ERROR                         Message_Type = 500
	Message_CAFE_CONTACT_QUERY            Message_Type = 68  // Deprecated: Do not use.
	Message_CAFE_CONTACT_QUERY_RES        Message_Type = 69  // Deprecated: Do not use.
	Message_CAFE_PUBSUB_CONTACT_QUERY     Message_Type = 100 // Deprecated: Do not use.
	Message_CAFE_PUBSUB_CONTACT_QUERY_RES Message_Type = 101 // Deprecated: Do not use.
)

var Message_Type_name = map[int32]string{
	0:   "PING",
	1:   "PONG",
	10:  "THREAD_ENVELOPE",
	50:  "CAFE_CHALLENGE",
	51:  "CAFE_NONCE",
	52:  "CAFE_REGISTRATION",
	72:  "CAFE_DEREGISTRATION",
	73:  "CAFE_DEREGISTRATION_ACK",
	53:  "CAFE_SESSION",
	54:  "CAFE_REFRESH_SESSION",
	55:  "CAFE_STORE",
	59:  "CAFE_STORE_ACK",
	74:  "CAFE_UNSTORE",
	75:  "CAFE_UNSTORE_ACK",
	56:  "CAFE_OBJECT",
	57:  "CAFE_OBJECT_LIST",
	58:  "CAFE_STORE_THREAD",
	76:  "CAFE_STORE_THREAD_ACK",
	77:  "CAFE_UNSTORE_THREAD",
	78:  "CAFE_UNSTORE_THREAD_ACK",
	60:  "CAFE_DELIVER_MESSAGE",
	61:  "CAFE_CHECK_MESSAGES",
	62:  "CAFE_MESSAGES",
	63:  "CAFE_DELETE_MESSAGES",
	64:  "CAFE_DELETE_MESSAGES_ACK",
	65:  "CAFE_YOU_HAVE_MAIL",
	66:  "CAFE_PUBLISH_PEER",
	67:  "CAFE_PUBLISH_PEER_ACK",
	70:  "CAFE_QUERY",
	71:  "CAFE_QUERY_RES",
	102: "CAFE_PUBSUB_QUERY",
	103: "CAFE_PUBSUB_QUERY_RES",
	500: "ERROR",
	68:  "CAFE_CONTACT_QUERY",
	69:  "CAFE_CONTACT_QUERY_RES",
	100: "CAFE_PUBSUB_CONTACT_QUERY",
	101: "CAFE_PUBSUB_CONTACT_QUERY_RES",
}
var Message_Type_value = map[string]int32{
	"PING":                          0,
	"PONG":                          1,
	"THREAD_ENVELOPE":               10,
	"CAFE_CHALLENGE":                50,
	"CAFE_NONCE":                    51,
	"CAFE_REGISTRATION":             52,
	"CAFE_DEREGISTRATION":           72,
	"CAFE_DEREGISTRATION_ACK":       73,
	"CAFE_SESSION":                  53,
	"CAFE_REFRESH_SESSION":          54,
	"CAFE_STORE":                    55,
	"CAFE_STORE_ACK":                59,
	"CAFE_UNSTORE":                  74,
	"CAFE_UNSTORE_ACK":              75,
	"CAFE_OBJECT":                   56,
	"CAFE_OBJECT_LIST":              57,
	"CAFE_STORE_THREAD":             58,
	"CAFE_STORE_THREAD_ACK":         76,
	"CAFE_UNSTORE_THREAD":           77,
	"CAFE_UNSTORE_THREAD_ACK":       78,
	"CAFE_DELIVER_MESSAGE":          60,
	"CAFE_CHECK_MESSAGES":           61,
	"CAFE_MESSAGES":                 62,
	"CAFE_DELETE_MESSAGES":          63,
	"CAFE_DELETE_MESSAGES_ACK":      64,
	"CAFE_YOU_HAVE_MAIL":            65,
	"CAFE_PUBLISH_PEER":             66,
	"CAFE_PUBLISH_PEER_ACK":         67,
	"CAFE_QUERY":                    70,
	"CAFE_QUERY_RES":                71,
	"CAFE_PUBSUB_QUERY":             102,
	"CAFE_PUBSUB_QUERY_RES":         103,
	"ERROR":                         500,
	"CAFE_CONTACT_QUERY":            68,
	"CAFE_CONTACT_QUERY_RES":        69,
	"CAFE_PUBSUB_CONTACT_QUERY":     100,
	"CAFE_PUBSUB_CONTACT_QUERY_RES": 101,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}
func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_7b943d204ab56303, []int{0, 0}
}

type Message struct {
	Type                 Message_Type `protobuf:"varint,1,opt,name=type,proto3,enum=Message_Type" json:"type,omitempty"`
	Payload              *any.Any     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	RequestId            int32        `protobuf:"varint,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	IsResponse           bool         `protobuf:"varint,4,opt,name=isResponse,proto3" json:"isResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7b943d204ab56303, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_PING
}

func (m *Message) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Message) GetIsResponse() bool {
	if m != nil {
		return m.IsResponse
	}
	return false
}

type Envelope struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Sig                  []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7b943d204ab56303, []int{1}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (dst *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(dst, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type Error struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_7b943d204ab56303, []int{2}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterEnum("Message_Type", Message_Type_name, Message_Type_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_7b943d204ab56303) }

var fileDescriptor_message_7b943d204ab56303 = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x6f, 0x4f, 0xe2, 0x4a,
	0x14, 0xc6, 0x2f, 0x8a, 0x82, 0x47, 0xd1, 0xf1, 0xf8, 0xaf, 0x7a, 0xbd, 0x06, 0x49, 0x6e, 0xc2,
	0xab, 0x9a, 0xe0, 0xf5, 0xee, 0xff, 0x5d, 0x4b, 0x39, 0xd2, 0x6a, 0x69, 0xd9, 0x69, 0x31, 0x71,
	0xdf, 0x34, 0xb0, 0x54, 0x42, 0xe2, 0xd2, 0x2e, 0xc5, 0xcd, 0xf2, 0x49, 0xf7, 0x8b, 0xf8, 0x01,
	0x36, 0x4c, 0xcb, 0x58, 0xa3, 0xfb, 0x6e, 0xe6, 0x79, 0x9e, 0xf3, 0x9b, 0x39, 0x33, 0x99, 0x81,
	0xd2, 0xb7, 0x20, 0x8e, 0xbb, 0x83, 0x40, 0x8d, 0xc6, 0xe1, 0x24, 0x3c, 0xd8, 0x1f, 0x84, 0xe1,
	0xe0, 0x2e, 0x38, 0x11, 0xb3, 0xde, 0xfd, 0xed, 0x49, 0x77, 0x34, 0x4d, 0xac, 0xca, 0x43, 0x01,
	0x0a, 0xad, 0x24, 0x8c, 0xc7, 0x90, 0x9f, 0x4c, 0xa3, 0x40, 0xc9, 0x95, 0x73, 0xd5, 0xf5, 0x5a,
	0x49, 0x4d, 0x75, 0xd5, 0x9b, 0x46, 0x01, 0x17, 0x16, 0xaa, 0x50, 0x88, 0xba, 0xd3, 0xbb, 0xb0,
	0xdb, 0x57, 0x16, 0xca, 0xb9, 0xea, 0x6a, 0x6d, 0x5b, 0x4d, 0xd8, 0xea, 0x9c, 0xad, 0x6a, 0xa3,
	0x29, 0x9f, 0x87, 0xf0, 0x10, 0x56, 0xc6, 0xc1, 0xf7, 0xfb, 0x20, 0x9e, 0x98, 0x7d, 0x65, 0xb1,
	0x9c, 0xab, 0x2e, 0xf1, 0x47, 0x01, 0x8f, 0x00, 0x86, 0x31, 0x0f, 0xe2, 0x28, 0x1c, 0xc5, 0x81,
	0x92, 0x2f, 0xe7, 0xaa, 0x45, 0x9e, 0x51, 0x2a, 0xbf, 0x96, 0x21, 0x3f, 0x5b, 0x1c, 0x8b, 0x90,
	0x6f, 0x9b, 0x76, 0x93, 0xfd, 0x25, 0x46, 0x8e, 0xdd, 0x64, 0x39, 0xdc, 0x82, 0x0d, 0xcf, 0xe0,
	0xa4, 0x35, 0x7c, 0xb2, 0xaf, 0xc9, 0x72, 0xda, 0xc4, 0x00, 0x11, 0xd6, 0x75, 0xed, 0x82, 0x7c,
	0xdd, 0xd0, 0x2c, 0x8b, 0xec, 0x26, 0xb1, 0x1a, 0xae, 0x03, 0x08, 0xcd, 0x76, 0x6c, 0x9d, 0xd8,
	0x29, 0xee, 0xc0, 0xa6, 0x98, 0x73, 0x6a, 0x9a, 0xae, 0xc7, 0x35, 0xcf, 0x74, 0x6c, 0xf6, 0x1f,
	0xee, 0xc1, 0x96, 0x90, 0x1b, 0xf4, 0xc4, 0x30, 0xf0, 0x6f, 0xd8, 0x7b, 0xc1, 0xf0, 0x35, 0xfd,
	0x8a, 0x99, 0xc8, 0x60, 0x4d, 0x98, 0x2e, 0xb9, 0xee, 0x2c, 0x7e, 0x86, 0x0a, 0x6c, 0xa7, 0xf8,
	0x0b, 0x4e, 0xae, 0x21, 0x9d, 0xff, 0xe5, 0x46, 0x5c, 0xcf, 0xe1, 0xc4, 0x5e, 0xc9, 0xcd, 0x8a,
	0xb9, 0xe0, 0xbd, 0x93, 0xbc, 0x8e, 0x9d, 0xa4, 0x2e, 0x71, 0x1b, 0x58, 0x56, 0x11, 0xb9, 0x2b,
	0xdc, 0x80, 0x55, 0xa1, 0x3a, 0xf5, 0x4b, 0xd2, 0x3d, 0xf6, 0x5a, 0xc6, 0x12, 0xc1, 0xb7, 0x4c,
	0xd7, 0x63, 0x6f, 0x64, 0xaf, 0x49, 0x69, 0x72, 0x5e, 0xec, 0x2d, 0xee, 0xc3, 0xce, 0x33, 0x59,
	0x80, 0x2d, 0x79, 0x0c, 0xf3, 0xe5, 0xd2, 0x9a, 0x96, 0x3c, 0x86, 0xa7, 0x86, 0xa8, 0xb2, 0x65,
	0xd3, 0x0d, 0xb2, 0xcc, 0x6b, 0xe2, 0x7e, 0x8b, 0x5c, 0x57, 0x6b, 0x12, 0x7b, 0x2f, 0x79, 0xba,
	0x41, 0xfa, 0xd5, 0x5c, 0x77, 0xd9, 0x07, 0xdc, 0x84, 0x92, 0x30, 0xa4, 0xf4, 0x31, 0x4b, 0x21,
	0x2f, 0xe3, 0x7c, 0xc2, 0x43, 0x50, 0x5e, 0x72, 0xc4, 0xea, 0xe7, 0xb8, 0x0b, 0x28, 0xdc, 0x1b,
	0xa7, 0xe3, 0x1b, 0xda, 0x35, 0xf9, 0x2d, 0xcd, 0xb4, 0x98, 0x26, 0xbb, 0x6f, 0x77, 0xea, 0x96,
	0xe9, 0x1a, 0x7e, 0x9b, 0x88, 0xb3, 0xba, 0xec, 0x3e, 0x2b, 0x0b, 0x92, 0x2e, 0xaf, 0xe8, 0x73,
	0x87, 0xf8, 0x0d, 0xbb, 0x90, 0x57, 0x24, 0xe6, 0x3e, 0x27, 0x97, 0x35, 0xb3, 0x54, 0xb7, 0x53,
	0x4f, 0xa3, 0xb7, 0x59, 0xaa, 0x94, 0x45, 0xc5, 0x00, 0x01, 0x96, 0x88, 0x73, 0x87, 0xb3, 0x87,
	0x45, 0x3c, 0x48, 0xf7, 0xaa, 0x3b, 0xb6, 0xa7, 0xe9, 0x5e, 0x5a, 0xde, 0x38, 0x58, 0x28, 0xe6,
	0xf0, 0x08, 0x76, 0x9f, 0x7b, 0x82, 0x41, 0xc2, 0x3f, 0x86, 0xfd, 0xec, 0x12, 0x4f, 0x11, 0x7d,
	0x11, 0xf9, 0x17, 0xfe, 0xf9, 0x63, 0x44, 0x90, 0x82, 0x59, 0xac, 0x72, 0x0e, 0x45, 0x1a, 0xfd,
	0x08, 0xee, 0xc2, 0x28, 0xc0, 0x0a, 0x14, 0xd2, 0xef, 0x42, 0xbc, 0xfc, 0xd5, 0x5a, 0x71, 0xfe,
	0xf2, 0xf9, 0xdc, 0x40, 0x06, 0x8b, 0xf1, 0x70, 0x20, 0xde, 0xfc, 0x1a, 0x9f, 0x0d, 0x2b, 0x67,
	0xb0, 0x44, 0xe3, 0x71, 0x38, 0x46, 0x84, 0xfc, 0xd7, 0xb0, 0x9f, 0xd4, 0x96, 0xb8, 0x18, 0xa3,
	0xf2, 0x88, 0x9c, 0x95, 0xac, 0x48, 0x50, 0x7d, 0x0b, 0x4a, 0xc3, 0x50, 0x9d, 0x04, 0x3f, 0x27,
	0xc3, 0xd9, 0xa7, 0xd1, 0xfb, 0xb2, 0x10, 0xf5, 0x7a, 0xcb, 0xe2, 0xf3, 0x38, 0xfd, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0xb6, 0xae, 0x03, 0x3a, 0xb7, 0x04, 0x00, 0x00,
}
