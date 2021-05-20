// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitcoin/v1beta1/types.proto

package types

import (
	fmt "fmt"
	github_com_btcsuite_btcutil "github.com/btcsuite/btcutil"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type AddressRole int32

const (
	None          AddressRole = 0
	Deposit       AddressRole = 1
	Consolidation AddressRole = 2
)

var AddressRole_name = map[int32]string{
	0: "ADDRESS_ROLE_UNSPECIFIED",
	1: "ADDRESS_ROLE_DEPOSIT",
	2: "ADDRESS_ROLE_CONSOLIDATION",
}

var AddressRole_value = map[string]int32{
	"ADDRESS_ROLE_UNSPECIFIED":   0,
	"ADDRESS_ROLE_DEPOSIT":       1,
	"ADDRESS_ROLE_CONSOLIDATION": 2,
}

func (x AddressRole) String() string {
	return proto.EnumName(AddressRole_name, int32(x))
}

func (AddressRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca561ce6167cd5e4, []int{0}
}

type SignState int32

const (
	SIGN_STATE_UNSPECIFIED SignState = 0
	Signing                SignState = 1
	Signed                 SignState = 2
	Ready                  SignState = 3
)

var SignState_name = map[int32]string{
	0: "SIGN_STATE_UNSPECIFIED",
	1: "SIGN_STATE_SIGNING_PENDING_TRANSFERS",
	2: "SIGN_STATE_SIGNED_NOT_CONFIRMED",
	3: "SIGN_STATE_READY_TO_SIGN",
}

var SignState_value = map[string]int32{
	"SIGN_STATE_UNSPECIFIED":               0,
	"SIGN_STATE_SIGNING_PENDING_TRANSFERS": 1,
	"SIGN_STATE_SIGNED_NOT_CONFIRMED":      2,
	"SIGN_STATE_READY_TO_SIGN":             3,
}

func (x SignState) String() string {
	return proto.EnumName(SignState_name, int32(x))
}

func (SignState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca561ce6167cd5e4, []int{1}
}

// OutPointInfo describes all the necessary information to confirm the outPoint
// of a transaction
type OutPointInfo struct {
	OutPoint string                             `protobuf:"bytes,1,opt,name=out_point,json=outPoint,proto3" json:"out_point,omitempty"`
	Amount   github_com_btcsuite_btcutil.Amount `protobuf:"varint,2,opt,name=amount,proto3,casttype=github.com/btcsuite/btcutil.Amount" json:"amount,omitempty"`
	Address  string                             `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *OutPointInfo) Reset()      { *m = OutPointInfo{} }
func (*OutPointInfo) ProtoMessage() {}
func (*OutPointInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca561ce6167cd5e4, []int{0}
}
func (m *OutPointInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutPointInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutPointInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutPointInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutPointInfo.Merge(m, src)
}
func (m *OutPointInfo) XXX_Size() int {
	return m.Size()
}
func (m *OutPointInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OutPointInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OutPointInfo proto.InternalMessageInfo

// AddressInfo is a wrapper containing the Bitcoin P2WSH address, it's
// corresponding script and the underlying key
type AddressInfo struct {
	Address      string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Role         AddressRole `protobuf:"varint,2,opt,name=role,proto3,enum=bitcoin.v1beta1.AddressRole" json:"role,omitempty"`
	RedeemScript []byte      `protobuf:"bytes,3,opt,name=redeem_script,json=redeemScript,proto3" json:"redeem_script,omitempty"`
	KeyID        string      `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (m *AddressInfo) Reset()         { *m = AddressInfo{} }
func (m *AddressInfo) String() string { return proto.CompactTextString(m) }
func (*AddressInfo) ProtoMessage()    {}
func (*AddressInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca561ce6167cd5e4, []int{1}
}
func (m *AddressInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddressInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddressInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddressInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressInfo.Merge(m, src)
}
func (m *AddressInfo) XXX_Size() int {
	return m.Size()
}
func (m *AddressInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AddressInfo proto.InternalMessageInfo

type Network struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca561ce6167cd5e4, []int{2}
}
func (m *Network) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Network.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(m, src)
}
func (m *Network) XXX_Size() int {
	return m.Size()
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("bitcoin.v1beta1.AddressRole", AddressRole_name, AddressRole_value)
	proto.RegisterEnum("bitcoin.v1beta1.SignState", SignState_name, SignState_value)
	proto.RegisterType((*OutPointInfo)(nil), "bitcoin.v1beta1.OutPointInfo")
	proto.RegisterType((*AddressInfo)(nil), "bitcoin.v1beta1.AddressInfo")
	proto.RegisterType((*Network)(nil), "bitcoin.v1beta1.Network")
}

func init() { proto.RegisterFile("bitcoin/v1beta1/types.proto", fileDescriptor_ca561ce6167cd5e4) }

var fileDescriptor_ca561ce6167cd5e4 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xbd, 0x6d, 0x9a, 0x36, 0xdb, 0xf6, 0x7d, 0xc3, 0xaa, 0x42, 0x96, 0x0b, 0x8e, 0x15,
	0xa0, 0x54, 0x95, 0x88, 0x29, 0x7f, 0x2e, 0x1c, 0x90, 0xdc, 0xda, 0xad, 0x2c, 0x8a, 0x1d, 0xed,
	0x9a, 0x03, 0x5c, 0x2c, 0x27, 0x5e, 0xc2, 0xaa, 0x89, 0x37, 0xb2, 0x37, 0xd0, 0x7c, 0x03, 0x14,
	0x2e, 0x1c, 0xe1, 0x10, 0xa9, 0x12, 0x1c, 0xf8, 0x1c, 0x9c, 0x7a, 0xec, 0x91, 0x53, 0x05, 0xe9,
	0xb7, 0xe0, 0x84, 0xbc, 0x71, 0xd5, 0xb4, 0x27, 0xcf, 0xce, 0xfc, 0x9e, 0xd9, 0x67, 0xe4, 0x1d,
	0xb8, 0xde, 0x62, 0xa2, 0xcd, 0x59, 0x62, 0xbe, 0xdf, 0x6e, 0x51, 0x11, 0x6d, 0x9b, 0x62, 0xd8,
	0xa7, 0x59, 0xa3, 0x9f, 0x72, 0xc1, 0xd1, 0xff, 0x45, 0xb1, 0x51, 0x14, 0xb5, 0xb5, 0x0e, 0xef,
	0x70, 0x59, 0x33, 0xf3, 0x68, 0x8a, 0xd5, 0x3f, 0x01, 0xb8, 0xe2, 0x0f, 0x44, 0x93, 0xb3, 0x44,
	0xb8, 0xc9, 0x5b, 0x8e, 0xd6, 0x61, 0x85, 0x0f, 0x44, 0xd8, 0xcf, 0x13, 0x2a, 0x30, 0xc0, 0x66,
	0x05, 0x2f, 0xf1, 0x02, 0x40, 0xcf, 0x61, 0x39, 0xea, 0xf1, 0x41, 0x22, 0xd4, 0x39, 0x03, 0x6c,
	0xce, 0xef, 0x6c, 0xfc, 0x3d, 0xab, 0xd5, 0x3b, 0x4c, 0xbc, 0x1b, 0xb4, 0x1a, 0x6d, 0xde, 0x33,
	0x5b, 0xa2, 0x9d, 0x0d, 0x98, 0xa0, 0x79, 0x30, 0x10, 0xac, 0xdb, 0xb0, 0x24, 0x8d, 0x0b, 0x15,
	0x52, 0xe1, 0x62, 0x14, 0xc7, 0x29, 0xcd, 0x32, 0x75, 0x5e, 0xb6, 0xbe, 0x38, 0x3e, 0x2b, 0x7d,
	0x39, 0xae, 0x29, 0xf5, 0x63, 0x00, 0x97, 0xad, 0x69, 0x46, 0x9a, 0x99, 0xe1, 0xc1, 0x15, 0x1e,
	0x3d, 0x84, 0xa5, 0x94, 0x77, 0xa9, 0xf4, 0xf1, 0xdf, 0xa3, 0x5b, 0x8d, 0x6b, 0xd3, 0x36, 0x8a,
	0x2e, 0x98, 0x77, 0x29, 0x96, 0x24, 0xba, 0x03, 0x57, 0x53, 0x1a, 0x53, 0xda, 0x0b, 0xb3, 0x76,
	0xca, 0xfa, 0x42, 0x3a, 0x58, 0xc1, 0x2b, 0xd3, 0x24, 0x91, 0x39, 0x64, 0xc0, 0xf2, 0x21, 0x1d,
	0x86, 0x2c, 0x56, 0x4b, 0xf9, 0x7d, 0x3b, 0x95, 0xc9, 0x59, 0x6d, 0xe1, 0x05, 0x1d, 0xba, 0x36,
	0x5e, 0x38, 0xa4, 0x43, 0x37, 0xae, 0xdf, 0x86, 0x8b, 0x1e, 0x15, 0x1f, 0x78, 0x7a, 0x88, 0x10,
	0x2c, 0x25, 0x51, 0x8f, 0x16, 0xd6, 0x64, 0xbc, 0xf5, 0xf5, 0x72, 0x82, 0xfc, 0x6e, 0xb4, 0x01,
	0x55, 0xcb, 0xb6, 0xb1, 0x43, 0x48, 0x88, 0xfd, 0x03, 0x27, 0x7c, 0xe5, 0x91, 0xa6, 0xb3, 0xeb,
	0xee, 0xb9, 0x8e, 0x5d, 0x55, 0xb4, 0xa5, 0xd1, 0xd8, 0x28, 0x79, 0x3c, 0xa1, 0xe8, 0x1e, 0x5c,
	0xbb, 0xc2, 0xd9, 0x4e, 0xd3, 0x27, 0x6e, 0x50, 0x05, 0xda, 0xf2, 0x68, 0x6c, 0x2c, 0xda, 0xb4,
	0xcf, 0x33, 0x26, 0xd0, 0x36, 0xd4, 0xae, 0x60, 0xbb, 0xbe, 0x47, 0xfc, 0x03, 0xd7, 0xb6, 0x02,
	0xd7, 0xf7, 0xaa, 0x73, 0xda, 0x8d, 0xd1, 0xd8, 0x58, 0xdd, 0xe5, 0x49, 0xc6, 0xbb, 0x2c, 0x8e,
	0x04, 0xe3, 0x89, 0xb6, 0xf4, 0xf1, 0x9b, 0xae, 0xfc, 0xf8, 0xae, 0x83, 0xad, 0x9f, 0x00, 0x56,
	0x08, 0xeb, 0x24, 0x44, 0x44, 0x82, 0x22, 0x0d, 0xde, 0x24, 0xee, 0xbe, 0x17, 0x92, 0xc0, 0x0a,
	0xae, 0xf9, 0x42, 0x4f, 0xe1, 0xdd, 0x99, 0x5a, 0x1e, 0xba, 0xde, 0x7e, 0xd8, 0x74, 0x3c, 0x3b,
	0xff, 0x06, 0xd8, 0xf2, 0xc8, 0x9e, 0x83, 0xc9, 0x85, 0xbb, 0xbc, 0x29, 0x4b, 0x3a, 0xc8, 0x84,
	0xb5, 0x6b, 0x32, 0xc7, 0x0e, 0x3d, 0x3f, 0xc8, 0x6d, 0xee, 0xb9, 0xf8, 0xa5, 0x63, 0x57, 0xe7,
	0x34, 0x38, 0x1a, 0x1b, 0xe5, 0x5c, 0x41, 0x63, 0x74, 0x1f, 0xaa, 0x33, 0x02, 0xec, 0x58, 0xf6,
	0xeb, 0x30, 0xf0, 0xa5, 0xb2, 0x3a, 0xaf, 0x55, 0x46, 0x63, 0x63, 0x01, 0xd3, 0x28, 0x1e, 0x5e,
	0x0e, 0xb1, 0x83, 0x4f, 0xfe, 0xe8, 0xca, 0xc9, 0x44, 0x07, 0xa7, 0x13, 0x1d, 0xfc, 0x9e, 0xe8,
	0xe0, 0xf3, 0xb9, 0xae, 0x9c, 0x9e, 0xeb, 0xca, 0xaf, 0x73, 0x5d, 0x79, 0xf3, 0x64, 0xe6, 0x31,
	0x46, 0x47, 0xb4, 0x1b, 0xa5, 0xc9, 0xf4, 0x67, 0x15, 0xa7, 0x07, 0x6d, 0x9e, 0x52, 0xf3, 0xc8,
	0xbc, 0xd8, 0x1c, 0xb9, 0x31, 0xad, 0xb2, 0xdc, 0x85, 0xc7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0xd3, 0x4d, 0x1e, 0x51, 0x03, 0x00, 0x00,
}

func (m *OutPointInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutPointInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutPointInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.OutPoint) > 0 {
		i -= len(m.OutPoint)
		copy(dAtA[i:], m.OutPoint)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.OutPoint)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddressInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddressInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddressInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RedeemScript) > 0 {
		i -= len(m.RedeemScript)
		copy(dAtA[i:], m.RedeemScript)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RedeemScript)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Role != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Network) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Network) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Network) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutPointInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OutPoint)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTypes(uint64(m.Amount))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *AddressInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovTypes(uint64(m.Role))
	}
	l = len(m.RedeemScript)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Network) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutPointInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: OutPointInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutPointInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutPoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutPoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= github_com_btcsuite_btcutil.Amount(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *AddressInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: AddressInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddressInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= AddressRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemScript", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedeemScript = append(m.RedeemScript[:0], dAtA[iNdEx:postIndex]...)
			if m.RedeemScript == nil {
				m.RedeemScript = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Network) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Network: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Network: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
