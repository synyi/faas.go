// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

package proto

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventType int32

const (
	HTTP EventType = 0
)

var EventType_name = map[int32]string{
	0: "HTTP",
}

var EventType_value = map[string]int32{
	"HTTP": 0,
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

type BodyType int32

const (
	RAW  BodyType = 0
	LINK BodyType = 1
)

var BodyType_name = map[int32]string{
	0: "RAW",
	1: "LINK",
}

var BodyType_value = map[string]int32{
	"RAW":  0,
	"LINK": 1,
}

func (BodyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

type Event struct {
	EventId     string    `protobuf:"bytes,1,opt,name=EventId,proto3" json:"EventId,omitempty"`
	SenderId    string    `protobuf:"bytes,2,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	Type        EventType `protobuf:"varint,3,opt,name=Type,proto3,enum=proto.EventType" json:"Type,omitempty"`
	Method      string    `protobuf:"bytes,4,opt,name=Method,proto3" json:"Method,omitempty"`
	Url         string    `protobuf:"bytes,5,opt,name=Url,proto3" json:"Url,omitempty"`
	Body        []byte    `protobuf:"bytes,6,opt,name=Body,proto3" json:"Body,omitempty"`
	BodyType    BodyType  `protobuf:"varint,7,opt,name=BodyType,proto3,enum=proto.BodyType" json:"BodyType,omitempty"`
	ContentType string    `protobuf:"bytes,8,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	Headers     []string  `protobuf:"bytes,9,rep,name=Headers,proto3" json:"Headers,omitempty"`
	Source      string    `protobuf:"bytes,10,opt,name=source,proto3" json:"source,omitempty"`
}

func (m *Event) Reset()      { *m = Event{} }
func (*Event) ProtoMessage() {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetSenderId() string {
	if m != nil {
		return m.SenderId
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return HTTP
}

func (m *Event) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Event) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Event) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Event) GetBodyType() BodyType {
	if m != nil {
		return m.BodyType
	}
	return RAW
}

func (m *Event) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Event) GetHeaders() []string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Event) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type Response struct {
	EventId     string   `protobuf:"bytes,1,opt,name=EventId,proto3" json:"EventId,omitempty"`
	Body        []byte   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	BodyType    BodyType `protobuf:"varint,3,opt,name=BodyType,proto3,enum=proto.BodyType" json:"BodyType,omitempty"`
	ContentType string   `protobuf:"bytes,4,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	Headers     []string `protobuf:"bytes,5,rep,name=Headers,proto3" json:"Headers,omitempty"`
	Status      int32    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	UsedTime    uint32   `protobuf:"varint,7,opt,name=UsedTime,proto3" json:"UsedTime,omitempty"`
}

func (m *Response) Reset()      { *m = Response{} }
func (*Response) ProtoMessage() {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetBodyType() BodyType {
	if m != nil {
		return m.BodyType
	}
	return RAW
}

func (m *Response) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Response) GetHeaders() []string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetUsedTime() uint32 {
	if m != nil {
		return m.UsedTime
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("proto.BodyType", BodyType_name, BodyType_value)
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0x77, 0x9b, 0xa4, 0x4d, 0xa7, 0xbf, 0x9f, 0x86, 0x05, 0x65, 0x11, 0x5c, 0x42, 0xf1,
	0x50, 0x2a, 0xf4, 0xa0, 0xbe, 0x80, 0x15, 0xa1, 0xc5, 0x3f, 0xc8, 0x36, 0xc5, 0x73, 0x35, 0x0b,
	0x0a, 0x9a, 0x2d, 0x49, 0x2a, 0xf4, 0xe6, 0x23, 0xf8, 0x0c, 0x9e, 0x7c, 0x14, 0x8f, 0xbd, 0x08,
	0x3d, 0xda, 0xed, 0xc5, 0x63, 0x1f, 0x41, 0x76, 0x93, 0x46, 0x41, 0x2c, 0x78, 0xda, 0xf9, 0x7e,
	0x67, 0x77, 0xe7, 0x33, 0x33, 0x50, 0x13, 0x0f, 0x22, 0x4a, 0x5b, 0xc3, 0x58, 0xa6, 0x92, 0x38,
	0xe6, 0xa8, 0x3f, 0x97, 0xc0, 0x39, 0xd6, 0x36, 0xa1, 0x50, 0x31, 0x41, 0x37, 0xa4, 0xd8, 0xc7,
	0x8d, 0x2a, 0x5f, 0x4a, 0xb2, 0x05, 0x6e, 0x4f, 0x44, 0xa1, 0x88, 0xbb, 0x21, 0x2d, 0x99, 0x54,
	0xa1, 0xc9, 0x0e, 0xd8, 0xc1, 0x78, 0x28, 0xa8, 0xe5, 0xe3, 0xc6, 0xda, 0x9e, 0x97, 0x7d, 0xde,
	0x32, 0x2f, 0xb5, 0xcf, 0x4d, 0x96, 0x6c, 0x42, 0xf9, 0x4c, 0xa4, 0x37, 0x32, 0xa4, 0xb6, 0x79,
	0x9f, 0x2b, 0xe2, 0x81, 0xd5, 0x8f, 0xef, 0xa8, 0x63, 0x4c, 0x1d, 0x12, 0x02, 0x76, 0x5b, 0x86,
	0x63, 0x5a, 0xf6, 0x71, 0xe3, 0x1f, 0x37, 0x31, 0xd9, 0x05, 0x57, 0x9f, 0xa6, 0x4e, 0xc5, 0xd4,
	0x59, 0xcf, 0xeb, 0x2c, 0x6d, 0x5e, 0x5c, 0x20, 0x3e, 0xd4, 0x8e, 0x64, 0x94, 0xe6, 0xf5, 0xa9,
	0x6b, 0xbe, 0xfe, 0x6e, 0xe9, 0x46, 0x3b, 0x62, 0x10, 0x8a, 0x38, 0xa1, 0x55, 0xdf, 0xd2, 0x8d,
	0xe6, 0x52, 0x63, 0x26, 0x72, 0x14, 0x5f, 0x0b, 0x0a, 0x19, 0x66, 0xa6, 0xea, 0x6f, 0x18, 0x5c,
	0x2e, 0x92, 0xa1, 0x8c, 0x12, 0xb1, 0x62, 0x4e, 0x4b, 0xf6, 0xd2, 0x2f, 0xec, 0xd6, 0x1f, 0xd9,
	0xed, 0x95, 0xec, 0xce, 0x0f, 0xf6, 0x5e, 0x3a, 0x48, 0x47, 0x89, 0x19, 0x9d, 0xc3, 0x73, 0xa5,
	0x97, 0xd7, 0x4f, 0x44, 0x18, 0xdc, 0xde, 0x67, 0xc3, 0xfb, 0xcf, 0x0b, 0xdd, 0xdc, 0x80, 0x6a,
	0xb1, 0x29, 0xe2, 0x82, 0xdd, 0x09, 0x82, 0x0b, 0x0f, 0x35, 0xb7, 0xbf, 0x98, 0x49, 0x05, 0x2c,
	0x7e, 0x78, 0xe9, 0x21, 0x9d, 0x3e, 0xed, 0x9e, 0x9f, 0x78, 0xb8, 0x7d, 0x30, 0x99, 0x31, 0x34,
	0x9d, 0x31, 0xb4, 0x98, 0x31, 0xfc, 0xa8, 0x18, 0x7e, 0x51, 0x0c, 0xbf, 0x2a, 0x86, 0x27, 0x8a,
	0xe1, 0x77, 0xc5, 0xf0, 0x87, 0x62, 0x68, 0xa1, 0x18, 0x7e, 0x9a, 0x33, 0x34, 0x99, 0x33, 0x34,
	0x9d, 0x33, 0x74, 0x55, 0x36, 0x5d, 0xef, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x53, 0x66, 0x3a,
	0x3e, 0x85, 0x02, 0x00, 0x00,
}

func (x EventType) String() string {
	s, ok := EventType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x BodyType) String() string {
	s, ok := BodyType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.EventId != that1.EventId {
		return false
	}
	if this.SenderId != that1.SenderId {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	if this.BodyType != that1.BodyType {
		return false
	}
	if this.ContentType != that1.ContentType {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if this.Source != that1.Source {
		return false
	}
	return true
}
func (this *Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Response)
	if !ok {
		that2, ok := that.(Response)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.EventId != that1.EventId {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	if this.BodyType != that1.BodyType {
		return false
	}
	if this.ContentType != that1.ContentType {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if this.Status != that1.Status {
		return false
	}
	if this.UsedTime != that1.UsedTime {
		return false
	}
	return true
}
func (this *Event) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 14)
	s = append(s, "&proto.Event{")
	s = append(s, "EventId: "+fmt.Sprintf("%#v", this.EventId)+",\n")
	s = append(s, "SenderId: "+fmt.Sprintf("%#v", this.SenderId)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Method: "+fmt.Sprintf("%#v", this.Method)+",\n")
	s = append(s, "Url: "+fmt.Sprintf("%#v", this.Url)+",\n")
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "BodyType: "+fmt.Sprintf("%#v", this.BodyType)+",\n")
	s = append(s, "ContentType: "+fmt.Sprintf("%#v", this.ContentType)+",\n")
	s = append(s, "Headers: "+fmt.Sprintf("%#v", this.Headers)+",\n")
	s = append(s, "Source: "+fmt.Sprintf("%#v", this.Source)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Response) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&proto.Response{")
	s = append(s, "EventId: "+fmt.Sprintf("%#v", this.EventId)+",\n")
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "BodyType: "+fmt.Sprintf("%#v", this.BodyType)+",\n")
	s = append(s, "ContentType: "+fmt.Sprintf("%#v", this.ContentType)+",\n")
	s = append(s, "Headers: "+fmt.Sprintf("%#v", this.Headers)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "UsedTime: "+fmt.Sprintf("%#v", this.UsedTime)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEvent(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Headers[iNdEx])
			copy(dAtA[i:], m.Headers[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.Headers[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ContentType) > 0 {
		i -= len(m.ContentType)
		copy(dAtA[i:], m.ContentType)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ContentType)))
		i--
		dAtA[i] = 0x42
	}
	if m.BodyType != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.BodyType))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x22
	}
	if m.Type != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SenderId) > 0 {
		i -= len(m.SenderId)
		copy(dAtA[i:], m.SenderId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.SenderId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EventId) > 0 {
		i -= len(m.EventId)
		copy(dAtA[i:], m.EventId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EventId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UsedTime != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.UsedTime))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Headers[iNdEx])
			copy(dAtA[i:], m.Headers[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.Headers[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ContentType) > 0 {
		i -= len(m.ContentType)
		copy(dAtA[i:], m.ContentType)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ContentType)))
		i--
		dAtA[i] = 0x22
	}
	if m.BodyType != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.BodyType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EventId) > 0 {
		i -= len(m.EventId)
		copy(dAtA[i:], m.EventId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EventId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EventId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.SenderId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovEvent(uint64(m.Type))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.BodyType != 0 {
		n += 1 + sovEvent(uint64(m.BodyType))
	}
	l = len(m.ContentType)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, s := range m.Headers {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EventId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.BodyType != 0 {
		n += 1 + sovEvent(uint64(m.BodyType))
	}
	l = len(m.ContentType)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, s := range m.Headers {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovEvent(uint64(m.Status))
	}
	if m.UsedTime != 0 {
		n += 1 + sovEvent(uint64(m.UsedTime))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`EventId:` + fmt.Sprintf("%v", this.EventId) + `,`,
		`SenderId:` + fmt.Sprintf("%v", this.SenderId) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`Url:` + fmt.Sprintf("%v", this.Url) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`BodyType:` + fmt.Sprintf("%v", this.BodyType) + `,`,
		`ContentType:` + fmt.Sprintf("%v", this.ContentType) + `,`,
		`Headers:` + fmt.Sprintf("%v", this.Headers) + `,`,
		`Source:` + fmt.Sprintf("%v", this.Source) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Response) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Response{`,
		`EventId:` + fmt.Sprintf("%v", this.EventId) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`BodyType:` + fmt.Sprintf("%v", this.BodyType) + `,`,
		`ContentType:` + fmt.Sprintf("%v", this.ContentType) + `,`,
		`Headers:` + fmt.Sprintf("%v", this.Headers) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`UsedTime:` + fmt.Sprintf("%v", this.UsedTime) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEvent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= EventType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyType", wireType)
			}
			m.BodyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BodyType |= BodyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BodyType", wireType)
			}
			m.BodyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BodyType |= BodyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedTime", wireType)
			}
			m.UsedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsedTime |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
