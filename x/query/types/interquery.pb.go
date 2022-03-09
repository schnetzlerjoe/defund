// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query/interquery.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	crypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Interquery struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Storeid       string `protobuf:"bytes,2,opt,name=storeid,proto3" json:"storeid,omitempty"`
	Path          string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	TimeoutHeight uint64 `protobuf:"varint,4,opt,name=timeoutHeight,proto3" json:"timeoutHeight,omitempty"`
	ClientId      string `protobuf:"bytes,5,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (m *Interquery) Reset()         { *m = Interquery{} }
func (m *Interquery) String() string { return proto.CompactTextString(m) }
func (*Interquery) ProtoMessage()    {}
func (*Interquery) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfa9beb887ed5f4, []int{0}
}
func (m *Interquery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Interquery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Interquery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Interquery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interquery.Merge(m, src)
}
func (m *Interquery) XXX_Size() int {
	return m.Size()
}
func (m *Interquery) XXX_DiscardUnknown() {
	xxx_messageInfo_Interquery.DiscardUnknown(m)
}

var xxx_messageInfo_Interquery proto.InternalMessageInfo

func (m *Interquery) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Interquery) GetStoreid() string {
	if m != nil {
		return m.Storeid
	}
	return ""
}

func (m *Interquery) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Interquery) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *Interquery) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type InterqueryResult struct {
	Creator  string           `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Storeid  string           `protobuf:"bytes,2,opt,name=storeid,proto3" json:"storeid,omitempty"`
	Data     []byte           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Height   uint64           `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	ClientId string           `protobuf:"bytes,5,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Success  bool             `protobuf:"varint,6,opt,name=success,proto3" json:"success,omitempty"`
	Proof    *crypto.ProofOps `protobuf:"bytes,7,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *InterqueryResult) Reset()         { *m = InterqueryResult{} }
func (m *InterqueryResult) String() string { return proto.CompactTextString(m) }
func (*InterqueryResult) ProtoMessage()    {}
func (*InterqueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfa9beb887ed5f4, []int{1}
}
func (m *InterqueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterqueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterqueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterqueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterqueryResult.Merge(m, src)
}
func (m *InterqueryResult) XXX_Size() int {
	return m.Size()
}
func (m *InterqueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InterqueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_InterqueryResult proto.InternalMessageInfo

func (m *InterqueryResult) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *InterqueryResult) GetStoreid() string {
	if m != nil {
		return m.Storeid
	}
	return ""
}

func (m *InterqueryResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InterqueryResult) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *InterqueryResult) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *InterqueryResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *InterqueryResult) GetProof() *crypto.ProofOps {
	if m != nil {
		return m.Proof
	}
	return nil
}

type InterqueryTimeoutResult struct {
	Creator       string           `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Storeid       string           `protobuf:"bytes,2,opt,name=storeid,proto3" json:"storeid,omitempty"`
	TimeoutHeight uint64           `protobuf:"varint,3,opt,name=timeoutHeight,proto3" json:"timeoutHeight,omitempty"`
	ClientId      string           `protobuf:"bytes,4,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Proof         *crypto.ProofOps `protobuf:"bytes,5,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *InterqueryTimeoutResult) Reset()         { *m = InterqueryTimeoutResult{} }
func (m *InterqueryTimeoutResult) String() string { return proto.CompactTextString(m) }
func (*InterqueryTimeoutResult) ProtoMessage()    {}
func (*InterqueryTimeoutResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfa9beb887ed5f4, []int{2}
}
func (m *InterqueryTimeoutResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterqueryTimeoutResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterqueryTimeoutResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterqueryTimeoutResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterqueryTimeoutResult.Merge(m, src)
}
func (m *InterqueryTimeoutResult) XXX_Size() int {
	return m.Size()
}
func (m *InterqueryTimeoutResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InterqueryTimeoutResult.DiscardUnknown(m)
}

var xxx_messageInfo_InterqueryTimeoutResult proto.InternalMessageInfo

func (m *InterqueryTimeoutResult) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *InterqueryTimeoutResult) GetStoreid() string {
	if m != nil {
		return m.Storeid
	}
	return ""
}

func (m *InterqueryTimeoutResult) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *InterqueryTimeoutResult) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *InterqueryTimeoutResult) GetProof() *crypto.ProofOps {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*Interquery)(nil), "defundhub.defund.query.Interquery")
	proto.RegisterType((*InterqueryResult)(nil), "defundhub.defund.query.InterqueryResult")
	proto.RegisterType((*InterqueryTimeoutResult)(nil), "defundhub.defund.query.InterqueryTimeoutResult")
}

func init() { proto.RegisterFile("query/interquery.proto", fileDescriptor_edfa9beb887ed5f4) }

var fileDescriptor_edfa9beb887ed5f4 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xbd, 0x4e, 0xc3, 0x30,
	0x18, 0xac, 0x69, 0xfa, 0x83, 0x01, 0x09, 0x79, 0x28, 0x56, 0x11, 0x51, 0x54, 0x31, 0x84, 0x25,
	0x11, 0xf0, 0x06, 0xb0, 0xd0, 0x09, 0x14, 0x31, 0xb1, 0xa5, 0xc9, 0xd7, 0xc6, 0x52, 0x1b, 0x07,
	0xe7, 0x8b, 0x44, 0xdf, 0x82, 0x89, 0x07, 0x62, 0x62, 0xec, 0x88, 0xc4, 0x82, 0xda, 0x17, 0x41,
	0xb5, 0xdb, 0x84, 0x1f, 0x09, 0x10, 0x6c, 0x77, 0xbe, 0x38, 0xbe, 0x3b, 0x1d, 0xed, 0xdc, 0x16,
	0xa0, 0xa6, 0xbe, 0x48, 0x11, 0x94, 0x86, 0x5e, 0xa6, 0x24, 0x4a, 0xd6, 0x89, 0x61, 0x58, 0xa4,
	0x71, 0x52, 0x0c, 0x3c, 0x83, 0x3c, 0xad, 0x76, 0x0f, 0x10, 0xd2, 0x18, 0xd4, 0x44, 0xa4, 0xe8,
	0x47, 0x6a, 0x9a, 0xa1, 0xf4, 0x33, 0x25, 0xe5, 0xd0, 0x5c, 0xeb, 0x3d, 0x10, 0x4a, 0xfb, 0xe5,
	0xbf, 0x18, 0xa7, 0xad, 0x48, 0x41, 0x88, 0x52, 0x71, 0xe2, 0x10, 0x77, 0x33, 0x58, 0xd3, 0xa5,
	0x92, 0xa3, 0x54, 0x20, 0x62, 0xbe, 0x61, 0x94, 0x15, 0x65, 0x8c, 0x5a, 0x59, 0x88, 0x09, 0xaf,
	0xeb, 0x63, 0x8d, 0xd9, 0x21, 0xdd, 0x41, 0x31, 0x01, 0x59, 0xe0, 0x05, 0x88, 0x51, 0x82, 0xdc,
	0x72, 0x88, 0x6b, 0x05, 0x1f, 0x0f, 0x59, 0x97, 0xb6, 0xa3, 0xb1, 0x80, 0x14, 0xfb, 0x31, 0x6f,
	0xe8, 0xdb, 0x25, 0xef, 0xbd, 0x10, 0xba, 0x5b, 0x19, 0x0b, 0x20, 0x2f, 0xc6, 0xf8, 0x57, 0x7b,
	0x71, 0x88, 0xa1, 0xb6, 0xb7, 0x1d, 0x68, 0xcc, 0x3a, 0xb4, 0x99, 0xbc, 0xf7, 0xb5, 0x62, 0xdf,
	0x19, 0xd2, 0x2f, 0x14, 0x51, 0x04, 0x79, 0xce, 0x9b, 0x0e, 0x71, 0xdb, 0xc1, 0x9a, 0xb2, 0x63,
	0xda, 0xd0, 0x95, 0xf2, 0x96, 0x43, 0xdc, 0xad, 0x93, 0x7d, 0xaf, 0xaa, 0xdc, 0x33, 0x95, 0x7b,
	0x57, 0x4b, 0xfd, 0x32, 0xcb, 0x03, 0xf3, 0x65, 0xef, 0x91, 0xd0, 0xbd, 0x2a, 0xdd, 0xb5, 0x69,
	0xe5, 0x1f, 0x21, 0xbf, 0xf4, 0x5d, 0xff, 0xa9, 0x6f, 0xeb, 0x53, 0xbc, 0x32, 0x44, 0xe3, 0xb7,
	0x21, 0xce, 0xce, 0x9f, 0xe6, 0x36, 0x99, 0xcd, 0x6d, 0xf2, 0x3a, 0xb7, 0xc9, 0xfd, 0xc2, 0xae,
	0xcd, 0x16, 0x76, 0xed, 0x79, 0x61, 0xd7, 0x6e, 0x8e, 0x46, 0x02, 0x97, 0x4b, 0x8c, 0xe4, 0xc4,
	0x2f, 0x77, 0xb9, 0x42, 0xfe, 0x9d, 0x6f, 0x26, 0x8c, 0xd3, 0x0c, 0xf2, 0x41, 0x53, 0xef, 0xf0,
	0xf4, 0x2d, 0x00, 0x00, 0xff, 0xff, 0xea, 0x0c, 0xa4, 0x0b, 0xd8, 0x02, 0x00, 0x00,
}

func (m *Interquery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Interquery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Interquery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintInterquery(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Storeid) > 0 {
		i -= len(m.Storeid)
		copy(dAtA[i:], m.Storeid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Storeid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterqueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterqueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterqueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterquery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Height != 0 {
		i = encodeVarintInterquery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Storeid) > 0 {
		i -= len(m.Storeid)
		copy(dAtA[i:], m.Storeid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Storeid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterqueryTimeoutResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterqueryTimeoutResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterqueryTimeoutResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterquery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x22
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintInterquery(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Storeid) > 0 {
		i -= len(m.Storeid)
		copy(dAtA[i:], m.Storeid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Storeid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterquery(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterquery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Interquery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Storeid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovInterquery(uint64(m.TimeoutHeight))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	return n
}

func (m *InterqueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Storeid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovInterquery(uint64(m.Height))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.Success {
		n += 2
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovInterquery(uint64(l))
	}
	return n
}

func (m *InterqueryTimeoutResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Storeid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovInterquery(uint64(m.TimeoutHeight))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovInterquery(uint64(l))
	}
	return n
}

func sovInterquery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterquery(x uint64) (n int) {
	return sovInterquery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Interquery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterquery
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
			return fmt.Errorf("proto: Interquery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Interquery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterquery
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
func (m *InterqueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterquery
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
			return fmt.Errorf("proto: InterqueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterqueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &crypto.ProofOps{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterquery
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
func (m *InterqueryTimeoutResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterquery
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
			return fmt.Errorf("proto: InterqueryTimeoutResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterqueryTimeoutResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
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
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &crypto.ProofOps{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterquery
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
func skipInterquery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterquery
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
					return 0, ErrIntOverflowInterquery
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
					return 0, ErrIntOverflowInterquery
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
				return 0, ErrInvalidLengthInterquery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterquery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterquery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterquery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterquery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterquery = fmt.Errorf("proto: unexpected end of group")
)