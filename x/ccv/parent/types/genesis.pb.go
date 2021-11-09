// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/parent/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/interchain-security/x/ccv/types"
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

// GenesisState defines the CCV parent chain genesis state
type GenesisState struct {
	ChildStates []ChildState `protobuf:"bytes,1,rep,name=child_states,json=childStates,proto3" json:"child_states" yaml:"child_states"`
	Params      Params       `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ffb7fd2386920f, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetChildStates() []ChildState {
	if m != nil {
		return m.ChildStates
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// ChildState defines the state that the parent chain stores for each child chain
type ChildState struct {
	ChainId   string       `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	ChannelId string       `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	Status    types.Status `protobuf:"varint,3,opt,name=status,proto3,enum=interchain_security.ccv.v1.Status" json:"status,omitempty"`
}

func (m *ChildState) Reset()         { *m = ChildState{} }
func (m *ChildState) String() string { return proto.CompactTextString(m) }
func (*ChildState) ProtoMessage()    {}
func (*ChildState) Descriptor() ([]byte, []int) {
	return fileDescriptor_87ffb7fd2386920f, []int{1}
}
func (m *ChildState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChildState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChildState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChildState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChildState.Merge(m, src)
}
func (m *ChildState) XXX_Size() int {
	return m.Size()
}
func (m *ChildState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChildState.DiscardUnknown(m)
}

var xxx_messageInfo_ChildState proto.InternalMessageInfo

func (m *ChildState) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChildState) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChildState) GetStatus() types.Status {
	if m != nil {
		return m.Status
	}
	return types.UNINITIALIZED
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "interchain_security.ccv.parent.v1.GenesisState")
	proto.RegisterType((*ChildState)(nil), "interchain_security.ccv.parent.v1.ChildState")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/parent/v1/genesis.proto", fileDescriptor_87ffb7fd2386920f)
}

var fileDescriptor_87ffb7fd2386920f = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x6e, 0xe2, 0x40,
	0x14, 0xc6, 0x3d, 0xb0, 0x62, 0x97, 0x01, 0xed, 0x6a, 0xcd, 0xae, 0x64, 0xb1, 0x92, 0xed, 0xb5,
	0x52, 0x38, 0x05, 0x33, 0x32, 0x49, 0x11, 0x51, 0x3a, 0x05, 0xa2, 0x8b, 0x4c, 0x97, 0x06, 0x99,
	0xf1, 0xc8, 0xb6, 0x84, 0xff, 0xc8, 0x33, 0x58, 0xe1, 0x16, 0x39, 0x49, 0x8e, 0x11, 0x51, 0x52,
	0xa6, 0xb2, 0x22, 0xb8, 0x01, 0x27, 0x88, 0x3c, 0x76, 0x08, 0x45, 0x10, 0xe9, 0xde, 0xd3, 0xfb,
	0x7e, 0xef, 0xfb, 0x46, 0x6f, 0x20, 0x0e, 0x63, 0x4e, 0x33, 0x12, 0xb8, 0x61, 0x3c, 0x63, 0x94,
	0x2c, 0xb3, 0x90, 0xaf, 0x30, 0x21, 0x39, 0x4e, 0xdd, 0x8c, 0xc6, 0x1c, 0xe7, 0x16, 0xf6, 0x69,
	0x4c, 0x59, 0xc8, 0x50, 0x9a, 0x25, 0x3c, 0x91, 0xff, 0x7f, 0x02, 0x20, 0x42, 0x72, 0x54, 0x01,
	0x28, 0xb7, 0xfa, 0x7f, 0xfc, 0xc4, 0x4f, 0x84, 0x1a, 0x97, 0x55, 0x05, 0xf6, 0x2f, 0x4e, 0x39,
	0xe5, 0x16, 0x16, 0xbc, 0x50, 0xa1, 0xf3, 0x79, 0x6a, 0x23, 0xa1, 0x37, 0x9e, 0x01, 0xec, 0x8e,
	0xab, 0x80, 0x53, 0xee, 0x72, 0x2a, 0x47, 0xb0, 0x4b, 0x82, 0x70, 0xe1, 0xcd, 0x58, 0xd9, 0x32,
	0x05, 0xe8, 0x4d, 0xb3, 0x33, 0x1c, 0xa0, 0xb3, 0xb1, 0xd1, 0x6d, 0x89, 0x89, 0x25, 0xf6, 0xbf,
	0x75, 0xa1, 0x49, 0xfb, 0x42, 0xeb, 0xad, 0xdc, 0x68, 0x31, 0x32, 0x8e, 0x17, 0x1a, 0x4e, 0x87,
	0x1c, 0x84, 0x4c, 0x1e, 0xc3, 0x56, 0xea, 0x66, 0x6e, 0xc4, 0x94, 0x86, 0x0e, 0xcc, 0xce, 0xf0,
	0xf2, 0x0b, 0x46, 0x77, 0x02, 0xb0, 0xbf, 0x95, 0x26, 0x4e, 0x8d, 0x1b, 0x4f, 0x00, 0xc2, 0x8f,
	0x04, 0x32, 0x82, 0x3f, 0xaa, 0x1d, 0xa1, 0xa7, 0x00, 0x1d, 0x98, 0x6d, 0xbb, 0xb7, 0x2f, 0xb4,
	0x5f, 0xef, 0x79, 0xaa, 0x89, 0xe1, 0x7c, 0x17, 0xe5, 0xc4, 0x93, 0xaf, 0x21, 0x24, 0x81, 0x1b,
	0xc7, 0x74, 0x51, 0x12, 0x0d, 0x41, 0xfc, 0xdd, 0x17, 0xda, 0xef, 0x03, 0x51, 0xcf, 0x0c, 0xa7,
	0x5d, 0x37, 0x13, 0x4f, 0x1e, 0xc1, 0x56, 0xf9, 0xaa, 0x25, 0x53, 0x9a, 0x3a, 0x30, 0x7f, 0x0e,
	0x8d, 0x93, 0xe9, 0x73, 0x0b, 0x4d, 0x85, 0xd2, 0xa9, 0x09, 0xdb, 0x59, 0x6f, 0x55, 0xb0, 0xd9,
	0xaa, 0xe0, 0x75, 0xab, 0x82, 0xc7, 0x9d, 0x2a, 0x6d, 0x76, 0xaa, 0xf4, 0xb2, 0x53, 0xa5, 0xfb,
	0x1b, 0x3f, 0xe4, 0xc1, 0x72, 0x8e, 0x48, 0x12, 0x61, 0x92, 0xb0, 0x28, 0x61, 0x47, 0xbf, 0x6c,
	0x70, 0xb8, 0xea, 0xc3, 0xf1, 0x5d, 0xf9, 0x2a, 0xa5, 0x6c, 0xde, 0x12, 0x47, 0xbd, 0x7a, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xce, 0x43, 0xb4, 0xd0, 0x96, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ChildStates) > 0 {
		for iNdEx := len(m.ChildStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChildStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChildState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChildState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChildState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChildStates) > 0 {
		for _, e := range m.ChildStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ChildState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovGenesis(uint64(m.Status))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildStates = append(m.ChildStates, ChildState{})
			if err := m.ChildStates[len(m.ChildStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ChildState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ChildState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChildState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
