// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planet/blog/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the blog module's genesis state.
type GenesisState struct {
	Params            Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId            string         `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	PostList          []Post         `protobuf:"bytes,3,rep,name=postList,proto3" json:"postList"`
	PostCount         uint64         `protobuf:"varint,4,opt,name=postCount,proto3" json:"postCount,omitempty"`
	SentPostList      []SentPost     `protobuf:"bytes,5,rep,name=sentPostList,proto3" json:"sentPostList"`
	SentPostCount     uint64         `protobuf:"varint,6,opt,name=sentPostCount,proto3" json:"sentPostCount,omitempty"`
	TimedoutPostList  []TimedoutPost `protobuf:"bytes,7,rep,name=timedoutPostList,proto3" json:"timedoutPostList"`
	TimedoutPostCount uint64         `protobuf:"varint,8,opt,name=timedoutPostCount,proto3" json:"timedoutPostCount,omitempty"`
	BridgeStatus      *BridgeStatus  `protobuf:"bytes,9,opt,name=bridgeStatus,proto3" json:"bridgeStatus,omitempty"`
	SentActionList    []SentAction   `protobuf:"bytes,10,rep,name=sentActionList,proto3" json:"sentActionList"`
	SentActionCount   uint64         `protobuf:"varint,11,opt,name=sentActionCount,proto3" json:"sentActionCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6c02ff89d106757, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetPostList() []Post {
	if m != nil {
		return m.PostList
	}
	return nil
}

func (m *GenesisState) GetPostCount() uint64 {
	if m != nil {
		return m.PostCount
	}
	return 0
}

func (m *GenesisState) GetSentPostList() []SentPost {
	if m != nil {
		return m.SentPostList
	}
	return nil
}

func (m *GenesisState) GetSentPostCount() uint64 {
	if m != nil {
		return m.SentPostCount
	}
	return 0
}

func (m *GenesisState) GetTimedoutPostList() []TimedoutPost {
	if m != nil {
		return m.TimedoutPostList
	}
	return nil
}

func (m *GenesisState) GetTimedoutPostCount() uint64 {
	if m != nil {
		return m.TimedoutPostCount
	}
	return 0
}

func (m *GenesisState) GetBridgeStatus() *BridgeStatus {
	if m != nil {
		return m.BridgeStatus
	}
	return nil
}

func (m *GenesisState) GetSentActionList() []SentAction {
	if m != nil {
		return m.SentActionList
	}
	return nil
}

func (m *GenesisState) GetSentActionCount() uint64 {
	if m != nil {
		return m.SentActionCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "planet.blog.GenesisState")
}

func init() { proto.RegisterFile("planet/blog/genesis.proto", fileDescriptor_c6c02ff89d106757) }

var fileDescriptor_c6c02ff89d106757 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcf, 0x6f, 0xda, 0x30,
	0x18, 0x4d, 0x46, 0x16, 0xc0, 0x61, 0x3f, 0x30, 0xdb, 0x08, 0x6c, 0x0b, 0xd1, 0xb4, 0x43, 0x0e,
	0x5b, 0xd0, 0xe0, 0x3c, 0x4d, 0x63, 0x9a, 0xa6, 0xa9, 0x3d, 0xa0, 0xd0, 0x53, 0x2f, 0x28, 0x34,
	0x56, 0x14, 0x09, 0xe2, 0x28, 0x36, 0x52, 0xfb, 0x5f, 0xf4, 0x3f, 0xea, 0x95, 0x23, 0xc7, 0x9e,
	0xaa, 0x0a, 0xfe, 0x91, 0x2a, 0xb6, 0x03, 0x36, 0x51, 0x6f, 0xf6, 0x7b, 0xef, 0xfb, 0xde, 0xfb,
	0xfc, 0x19, 0xf4, 0xb2, 0x65, 0x98, 0x22, 0x3a, 0x5c, 0x2c, 0x71, 0x3c, 0x8c, 0x51, 0x8a, 0x48,
	0x42, 0xfc, 0x2c, 0xc7, 0x14, 0x43, 0x8b, 0x53, 0x7e, 0x41, 0xf5, 0xdf, 0xc5, 0x38, 0xc6, 0x0c,
	0x1f, 0x16, 0x27, 0x2e, 0xe9, 0xdb, 0x72, 0x75, 0x16, 0xe6, 0xe1, 0x4a, 0x14, 0xf7, 0x3f, 0x28,
	0x0c, 0x26, 0x54, 0xe0, 0x1f, 0x65, 0x9c, 0xa0, 0x94, 0xce, 0x25, 0x72, 0x20, 0x93, 0x34, 0x59,
	0xa1, 0x08, 0xaf, 0x9f, 0x17, 0x2c, 0xf2, 0x24, 0x8a, 0xd1, 0x9c, 0xd0, 0x90, 0xae, 0x4b, 0xdb,
	0xcf, 0x95, 0xf6, 0xe1, 0x15, 0x4d, 0x70, 0xca, 0xe9, 0x2f, 0x77, 0x06, 0x68, 0xfd, 0xe3, 0x43,
	0xce, 0x68, 0x48, 0x11, 0xfc, 0x01, 0x4c, 0x1e, 0xdb, 0xd6, 0x5d, 0xdd, 0xb3, 0x46, 0x1d, 0x5f,
	0x1a, 0xda, 0x9f, 0x32, 0x6a, 0x62, 0x6c, 0x1e, 0x06, 0x5a, 0x20, 0x84, 0xb0, 0x0b, 0xea, 0x19,
	0xce, 0xe9, 0x3c, 0x89, 0xec, 0x17, 0xae, 0xee, 0x35, 0x03, 0xb3, 0xb8, 0xfe, 0x8f, 0xe0, 0x18,
	0x34, 0x8a, 0xa8, 0xe7, 0x09, 0xa1, 0x76, 0xcd, 0xad, 0x79, 0xd6, 0xa8, 0xad, 0x76, 0xc3, 0x84,
	0x8a, 0x5e, 0x07, 0x21, 0xfc, 0x04, 0x9a, 0xc5, 0xf9, 0x0f, 0x5e, 0xa7, 0xd4, 0x36, 0x5c, 0xdd,
	0x33, 0x82, 0x23, 0x00, 0x7f, 0x81, 0x56, 0x31, 0xc4, 0xb4, 0x6c, 0xfb, 0x92, 0xb5, 0x7d, 0xaf,
	0xb4, 0x9d, 0x09, 0x81, 0x68, 0xad, 0x14, 0xc0, 0xaf, 0xe0, 0x55, 0x79, 0xe7, 0x16, 0x26, 0xb3,
	0x50, 0x41, 0x78, 0x06, 0xde, 0x96, 0xaf, 0x7d, 0xb0, 0xaa, 0x33, 0xab, 0x9e, 0x62, 0x75, 0x21,
	0x89, 0x84, 0x5d, 0xa5, 0x10, 0x7e, 0x03, 0x6d, 0x19, 0xe3, 0xb6, 0x0d, 0x66, 0x5b, 0x25, 0xe0,
	0x4f, 0xd0, 0xe2, 0x7b, 0x9c, 0xb1, 0x35, 0xda, 0x4d, 0xb6, 0x06, 0xd5, 0x76, 0x22, 0x09, 0x02,
	0x45, 0x0e, 0xff, 0x82, 0xd7, 0xc5, 0x28, 0xbf, 0xd9, 0x92, 0x59, 0x6e, 0xc0, 0x72, 0x77, 0x2b,
	0x4f, 0xc4, 0x25, 0x22, 0xf5, 0x49, 0x11, 0xf4, 0xc0, 0x9b, 0x23, 0xc2, 0x13, 0x5b, 0x2c, 0xf1,
	0x29, 0x3c, 0xf9, 0xbe, 0xd9, 0x39, 0xfa, 0x76, 0xe7, 0xe8, 0x8f, 0x3b, 0x47, 0xbf, 0xdd, 0x3b,
	0xda, 0x76, 0xef, 0x68, 0xf7, 0x7b, 0x47, 0xbb, 0xec, 0x88, 0xaf, 0x77, 0x2d, 0xbe, 0xef, 0x4d,
	0x86, 0xc8, 0xc2, 0x64, 0xff, 0x6e, 0xfc, 0x14, 0x00, 0x00, 0xff, 0xff, 0x00, 0xf8, 0x30, 0xf4,
	0x67, 0x03, 0x00, 0x00,
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
	if m.SentActionCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SentActionCount))
		i--
		dAtA[i] = 0x58
	}
	if len(m.SentActionList) > 0 {
		for iNdEx := len(m.SentActionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SentActionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.BridgeStatus != nil {
		{
			size, err := m.BridgeStatus.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.TimedoutPostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TimedoutPostCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.TimedoutPostList) > 0 {
		for iNdEx := len(m.TimedoutPostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimedoutPostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.SentPostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SentPostCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SentPostList) > 0 {
		for iNdEx := len(m.SentPostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SentPostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PostCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PostCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.PostList) > 0 {
		for iNdEx := len(m.PostList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.PostList) > 0 {
		for _, e := range m.PostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PostCount != 0 {
		n += 1 + sovGenesis(uint64(m.PostCount))
	}
	if len(m.SentPostList) > 0 {
		for _, e := range m.SentPostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SentPostCount != 0 {
		n += 1 + sovGenesis(uint64(m.SentPostCount))
	}
	if len(m.TimedoutPostList) > 0 {
		for _, e := range m.TimedoutPostList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TimedoutPostCount != 0 {
		n += 1 + sovGenesis(uint64(m.TimedoutPostCount))
	}
	if m.BridgeStatus != nil {
		l = m.BridgeStatus.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.SentActionList) > 0 {
		for _, e := range m.SentActionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SentActionCount != 0 {
		n += 1 + sovGenesis(uint64(m.SentActionCount))
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostList", wireType)
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
			m.PostList = append(m.PostList, Post{})
			if err := m.PostList[len(m.PostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostCount", wireType)
			}
			m.PostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentPostList", wireType)
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
			m.SentPostList = append(m.SentPostList, SentPost{})
			if err := m.SentPostList[len(m.SentPostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentPostCount", wireType)
			}
			m.SentPostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SentPostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimedoutPostList", wireType)
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
			m.TimedoutPostList = append(m.TimedoutPostList, TimedoutPost{})
			if err := m.TimedoutPostList[len(m.TimedoutPostList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimedoutPostCount", wireType)
			}
			m.TimedoutPostCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimedoutPostCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BridgeStatus", wireType)
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
			if m.BridgeStatus == nil {
				m.BridgeStatus = &BridgeStatus{}
			}
			if err := m.BridgeStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentActionList", wireType)
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
			m.SentActionList = append(m.SentActionList, SentAction{})
			if err := m.SentActionList[len(m.SentActionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentActionCount", wireType)
			}
			m.SentActionCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SentActionCount |= uint64(b&0x7F) << shift
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
