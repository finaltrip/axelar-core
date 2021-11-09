// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: snapshot/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
	exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorIllegibility int32

const (
	// these enum values are used for bitwise operations, therefore they need to
	// be powers of 2
	None                  ValidatorIllegibility = 0
	Tombstoned            ValidatorIllegibility = 1
	Jailed                ValidatorIllegibility = 2
	MissedTooManyBlocks   ValidatorIllegibility = 4
	NoProxyRegistered     ValidatorIllegibility = 8
	TssSuspended          ValidatorIllegibility = 16
	ProxyInsuficientFunds ValidatorIllegibility = 32
)

var ValidatorIllegibility_name = map[int32]string{
	0:  "VALIDATOR_ILLEGIBILITY_UNSPECIFIED",
	1:  "VALIDATOR_ILLEGIBILITY_TOMBSTONED",
	2:  "VALIDATOR_ILLEGIBILITY_JAILED",
	4:  "VALIDATOR_ILLEGIBILITY_MISSED_TOO_MANY_BLOCKS",
	8:  "VALIDATOR_ILLEGIBILITY_NO_PROXY_REGISTERED",
	16: "VALIDATOR_ILLEGIBILITY_TSS_SUSPENDED",
	32: "VALIDATOR_ILLEGIBILITY_PROXY_INSUFICIENT_FUNDS",
}

var ValidatorIllegibility_value = map[string]int32{
	"VALIDATOR_ILLEGIBILITY_UNSPECIFIED":             0,
	"VALIDATOR_ILLEGIBILITY_TOMBSTONED":              1,
	"VALIDATOR_ILLEGIBILITY_JAILED":                  2,
	"VALIDATOR_ILLEGIBILITY_MISSED_TOO_MANY_BLOCKS":  4,
	"VALIDATOR_ILLEGIBILITY_NO_PROXY_REGISTERED":     8,
	"VALIDATOR_ILLEGIBILITY_TSS_SUSPENDED":           16,
	"VALIDATOR_ILLEGIBILITY_PROXY_INSUFICIENT_FUNDS": 32,
}

func (ValidatorIllegibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7425dedf1b9aad3, []int{0}
}

type Validator struct {
	SDKValidator *types.Any `protobuf:"bytes,1,opt,name=sdk_validator,json=sdkValidator,proto3" json:"sdk_validator,omitempty"`
	ShareCount   int64      `protobuf:"varint,2,opt,name=share_count,json=shareCount,proto3" json:"share_count,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7425dedf1b9aad3, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

type Snapshot struct {
	Validators                 []Validator                            `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators"`
	Timestamp                  time.Time                              `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Height                     int64                                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	TotalShareCount            github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=total_share_count,json=totalShareCount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_share_count"`
	Counter                    int64                                  `protobuf:"varint,5,opt,name=counter,proto3" json:"counter,omitempty"`
	KeyShareDistributionPolicy exported.KeyShareDistributionPolicy    `protobuf:"varint,6,opt,name=key_share_distribution_policy,json=keyShareDistributionPolicy,proto3,enum=tss.exported.v1beta1.KeyShareDistributionPolicy" json:"key_share_distribution_policy,omitempty"`
	CorruptionThreshold        int64                                  `protobuf:"varint,7,opt,name=corruption_threshold,json=corruptionThreshold,proto3" json:"corruption_threshold,omitempty"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7425dedf1b9aad3, []int{1}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("snapshot.exported.v1beta1.ValidatorIllegibility", ValidatorIllegibility_name, ValidatorIllegibility_value)
	proto.RegisterType((*Validator)(nil), "snapshot.exported.v1beta1.Validator")
	proto.RegisterType((*Snapshot)(nil), "snapshot.exported.v1beta1.Snapshot")
}

func init() {
	proto.RegisterFile("snapshot/exported/v1beta1/types.proto", fileDescriptor_d7425dedf1b9aad3)
}

var fileDescriptor_d7425dedf1b9aad3 = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcf, 0x6e, 0xa3, 0x46,
	0x18, 0x37, 0xeb, 0x34, 0xeb, 0x9d, 0xa4, 0x5b, 0x96, 0x4d, 0x5a, 0x07, 0x69, 0x31, 0x8d, 0xd2,
	0xca, 0x5d, 0xc9, 0xb0, 0x49, 0x55, 0xa9, 0xea, 0xcd, 0x04, 0xb2, 0x22, 0xb1, 0xb1, 0x05, 0x24,
	0x6a, 0xf6, 0x82, 0x30, 0xcc, 0xe2, 0x91, 0xf1, 0x8c, 0xc5, 0x8c, 0xb7, 0xe1, 0x0d, 0x2a, 0x9f,
	0x56, 0xbd, 0xfb, 0xd4, 0x1e, 0xfa, 0x00, 0x7d, 0x88, 0xa8, 0xa7, 0xdc, 0x5a, 0xf5, 0x90, 0xb6,
	0xc9, 0x8b, 0x54, 0xc1, 0x98, 0x58, 0x8d, 0xd8, 0x93, 0xfd, 0xcd, 0xf7, 0xfb, 0xf7, 0xcd, 0x07,
	0x80, 0x2f, 0x28, 0xf6, 0x27, 0x74, 0x48, 0x98, 0x0a, 0x2f, 0x26, 0x24, 0x61, 0x30, 0x54, 0xdf,
	0xed, 0x0f, 0x20, 0xf3, 0xf7, 0x55, 0x96, 0x4e, 0x20, 0x55, 0x26, 0x09, 0x61, 0x44, 0xd8, 0x59,
	0xc2, 0x94, 0x25, 0x4c, 0xc9, 0x61, 0xe2, 0x56, 0x44, 0x22, 0x92, 0xa1, 0xd4, 0xbb, 0x7f, 0x0b,
	0x82, 0xd8, 0x88, 0x08, 0x89, 0x62, 0xa8, 0x66, 0xd5, 0x60, 0xfa, 0x56, 0x65, 0x68, 0x0c, 0x29,
	0xf3, 0xc7, 0x93, 0x1c, 0xb0, 0xf3, 0x7f, 0x80, 0x8f, 0xd3, 0xbc, 0x25, 0x05, 0x84, 0x8e, 0x09,
	0x55, 0x07, 0x3e, 0x85, 0x45, 0x9a, 0x80, 0x20, 0x9c, 0xf7, 0x65, 0x46, 0xe9, 0x07, 0xe3, 0x8a,
	0x3b, 0x0b, 0x05, 0x6f, 0x11, 0x6b, 0x51, 0xe4, 0xad, 0xbd, 0x5c, 0x9c, 0x32, 0x7f, 0x84, 0x70,
	0x54, 0xd0, 0xf3, 0x7a, 0x81, 0xda, 0xfd, 0x89, 0x03, 0x4f, 0xce, 0xfc, 0x18, 0x85, 0x3e, 0x23,
	0x89, 0x10, 0x82, 0x8f, 0x69, 0x38, 0xf2, 0xde, 0x2d, 0x0f, 0xea, 0x9c, 0xcc, 0x35, 0x37, 0x0e,
	0xb6, 0x94, 0xc5, 0x0c, 0xca, 0x72, 0x06, 0xa5, 0x8d, 0x53, 0xed, 0xab, 0x9b, 0xeb, 0xc6, 0xa6,
	0xa3, 0x9f, 0x14, 0xf4, 0xdf, 0x7f, 0x6b, 0x6d, 0x17, 0xb7, 0xb6, 0xda, 0xb0, 0x37, 0x69, 0x38,
	0xba, 0x77, 0x69, 0x80, 0x0d, 0x3a, 0xf4, 0x13, 0xe8, 0x05, 0x64, 0x8a, 0x59, 0xfd, 0x91, 0xcc,
	0x35, 0xab, 0x36, 0xc8, 0x8e, 0x0e, 0xef, 0x4e, 0x76, 0xaf, 0xaa, 0xa0, 0xe6, 0xe4, 0x7b, 0x10,
	0x8e, 0x01, 0x28, 0xf2, 0xd0, 0x3a, 0x27, 0x57, 0x9b, 0x1b, 0x07, 0x7b, 0x4a, 0xe9, 0x9a, 0x94,
	0xc2, 0x47, 0x5b, 0xbb, 0xbc, 0x6e, 0x54, 0xec, 0x15, 0xb6, 0xa0, 0x81, 0x27, 0xc5, 0x7a, 0x32,
	0xdf, 0x8d, 0x03, 0xf1, 0xc1, 0x6c, 0xee, 0x12, 0xa1, 0xd5, 0xee, 0x04, 0xde, 0xff, 0xdd, 0xe0,
	0xec, 0x7b, 0x9a, 0xf0, 0x29, 0x58, 0x1f, 0x42, 0x14, 0x0d, 0x59, 0xbd, 0x9a, 0x05, 0xcf, 0x2b,
	0xe1, 0x0d, 0x78, 0xc6, 0x08, 0xf3, 0x63, 0x6f, 0x75, 0xb6, 0x35, 0x99, 0x6b, 0x6e, 0x6a, 0xca,
	0x9d, 0xce, 0x5f, 0xd7, 0x8d, 0x2f, 0x23, 0xc4, 0x86, 0xd3, 0x81, 0x12, 0x90, 0x71, 0xbe, 0xab,
	0xfc, 0xa7, 0x45, 0xc3, 0x51, 0xbe, 0x57, 0x13, 0x33, 0xfb, 0x93, 0x4c, 0xc8, 0x29, 0x2e, 0x44,
	0xa8, 0x83, 0xc7, 0x99, 0x1e, 0x4c, 0xea, 0x1f, 0x65, 0xa6, 0xcb, 0x52, 0xa0, 0xe0, 0xc5, 0x08,
	0xa6, 0xb9, 0x67, 0x88, 0x28, 0x4b, 0xd0, 0x60, 0xca, 0x10, 0xc1, 0xde, 0x84, 0xc4, 0x28, 0x48,
	0xeb, 0xeb, 0x32, 0xd7, 0x7c, 0x7a, 0xf0, 0x4a, 0x61, 0x94, 0x3e, 0xbc, 0xab, 0x13, 0x98, 0x66,
	0x2e, 0xfa, 0x0a, 0xb1, 0x9f, 0xf1, 0x6c, 0x71, 0x54, 0xda, 0x13, 0xf6, 0xc1, 0x56, 0x40, 0x92,
	0x64, 0x3a, 0xc9, 0x8c, 0xd8, 0x30, 0x81, 0x74, 0x48, 0xe2, 0xb0, 0xfe, 0x38, 0xcb, 0xf6, 0xfc,
	0xbe, 0xe7, 0x2e, 0x5b, 0x2f, 0xff, 0xa8, 0x82, 0xed, 0x62, 0x33, 0x66, 0x1c, 0xc3, 0x08, 0x0d,
	0x50, 0x8c, 0x58, 0x2a, 0xbc, 0x02, 0xbb, 0x67, 0xed, 0x8e, 0xa9, 0xb7, 0xdd, 0x9e, 0xed, 0x99,
	0x9d, 0x8e, 0xf1, 0xda, 0xd4, 0xcc, 0x8e, 0xe9, 0x9e, 0x7b, 0xa7, 0x96, 0xd3, 0x37, 0x0e, 0xcd,
	0x23, 0xd3, 0xd0, 0xf9, 0x8a, 0x58, 0x9b, 0xcd, 0xe5, 0x35, 0x8b, 0x60, 0x28, 0x7c, 0x03, 0x3e,
	0x2f, 0x61, 0xb8, 0xbd, 0xae, 0xe6, 0xb8, 0x3d, 0xcb, 0xd0, 0x79, 0x4e, 0x7c, 0x3a, 0x9b, 0xcb,
	0xc0, 0x25, 0xe3, 0x01, 0x65, 0x04, 0xc3, 0x50, 0x68, 0x81, 0x17, 0x25, 0xb4, 0xe3, 0xb6, 0xd9,
	0x31, 0x74, 0xfe, 0x91, 0x08, 0x66, 0x73, 0x79, 0xfd, 0xd8, 0x47, 0x31, 0x0c, 0x85, 0x63, 0xd0,
	0x2a, 0x81, 0x77, 0x4d, 0xc7, 0x31, 0x74, 0xcf, 0xed, 0xf5, 0xbc, 0x6e, 0xdb, 0x3a, 0xf7, 0xb4,
	0x4e, 0xef, 0xf0, 0xc4, 0xe1, 0xd7, 0xc4, 0xcf, 0x66, 0x73, 0xf9, 0x79, 0x17, 0x51, 0x0a, 0x43,
	0x97, 0x90, 0xae, 0x8f, 0x53, 0x2d, 0x26, 0xc1, 0x88, 0x0a, 0x06, 0x78, 0x59, 0xa2, 0x65, 0xf5,
	0xbc, 0xbe, 0xdd, 0xfb, 0xfe, 0xdc, 0xb3, 0x8d, 0xd7, 0xa6, 0xe3, 0x1a, 0xb6, 0xa1, 0xf3, 0x35,
	0x71, 0x7b, 0x36, 0x97, 0x9f, 0x59, 0xa4, 0x9f, 0x90, 0x8b, 0xd4, 0x86, 0x11, 0xa2, 0x0c, 0x26,
	0x30, 0x14, 0xbe, 0x03, 0x7b, 0x65, 0x83, 0x3b, 0x8e, 0xe7, 0x9c, 0x3a, 0x7d, 0xc3, 0xd2, 0x0d,
	0x9d, 0xe7, 0x45, 0x7e, 0x36, 0x97, 0x37, 0x5d, 0x4a, 0x9d, 0x29, 0x9d, 0x40, 0x1c, 0xc2, 0x50,
	0xe8, 0x02, 0xa5, 0x84, 0xbb, 0xf0, 0x37, 0x2d, 0xe7, 0xf4, 0xc8, 0x3c, 0x34, 0x0d, 0xcb, 0xf5,
	0x8e, 0x4e, 0x2d, 0xdd, 0xe1, 0x65, 0x71, 0x67, 0x36, 0x97, 0xb7, 0xb3, 0x10, 0x26, 0xa6, 0xd3,
	0xb7, 0x28, 0x40, 0x10, 0xb3, 0xa3, 0x29, 0x0e, 0xa9, 0x58, 0xfb, 0xf1, 0x67, 0xa9, 0xf2, 0xeb,
	0x2f, 0x52, 0x45, 0x3b, 0xbb, 0xfc, 0x57, 0xaa, 0x5c, 0xde, 0x48, 0xdc, 0xd5, 0x8d, 0xc4, 0xfd,
	0x73, 0x23, 0x71, 0xef, 0x6f, 0xa5, 0xca, 0xd5, 0xad, 0x54, 0xf9, 0xf3, 0x56, 0xaa, 0xbc, 0xf9,
	0x76, 0xe5, 0x91, 0xf7, 0x2f, 0x60, 0xec, 0x27, 0x18, 0xb2, 0x1f, 0x48, 0x32, 0xca, 0xab, 0x56,
	0x40, 0x12, 0xa8, 0x5e, 0xa8, 0x0f, 0xbe, 0xce, 0x83, 0xf5, 0xec, 0x85, 0xfc, 0xfa, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x10, 0x11, 0x8a, 0xc4, 0xb9, 0x05, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShareCount != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ShareCount))
		i--
		dAtA[i] = 0x10
	}
	if m.SDKValidator != nil {
		{
			size, err := m.SDKValidator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Snapshot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CorruptionThreshold != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CorruptionThreshold))
		i--
		dAtA[i] = 0x38
	}
	if m.KeyShareDistributionPolicy != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeyShareDistributionPolicy))
		i--
		dAtA[i] = 0x30
	}
	if m.Counter != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.TotalShareCount.Size()
		i -= size
		if _, err := m.TotalShareCount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SDKValidator != nil {
		l = m.SDKValidator.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ShareCount != 0 {
		n += 1 + sovTypes(uint64(m.ShareCount))
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovTypes(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = m.TotalShareCount.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Counter != 0 {
		n += 1 + sovTypes(uint64(m.Counter))
	}
	if m.KeyShareDistributionPolicy != 0 {
		n += 1 + sovTypes(uint64(m.KeyShareDistributionPolicy))
	}
	if m.CorruptionThreshold != 0 {
		n += 1 + sovTypes(uint64(m.CorruptionThreshold))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SDKValidator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SDKValidator == nil {
				m.SDKValidator = &types.Any{}
			}
			if err := m.SDKValidator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareCount", wireType)
			}
			m.ShareCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShareCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *Snapshot) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShareCount", wireType)
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
			if err := m.TotalShareCount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShareDistributionPolicy", wireType)
			}
			m.KeyShareDistributionPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyShareDistributionPolicy |= exported.KeyShareDistributionPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorruptionThreshold", wireType)
			}
			m.CorruptionThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CorruptionThreshold |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
