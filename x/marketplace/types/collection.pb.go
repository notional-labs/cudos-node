// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cudoventures/cudosnode/marketplace/collection.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Collection listed for sale in the marketplace
type Collection struct {
	Id              uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DenomId         string    `protobuf:"bytes,2,opt,name=denomId,proto3" json:"denomId,omitempty"`
	MintRoyalties   []Royalty `protobuf:"bytes,3,rep,name=mintRoyalties,proto3" json:"mintRoyalties"`
	ResaleRoyalties []Royalty `protobuf:"bytes,4,rep,name=resaleRoyalties,proto3" json:"resaleRoyalties"`
	Verified        bool      `protobuf:"varint,5,opt,name=verified,proto3" json:"verified,omitempty"`
	Owner           string    `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdd7b78f471f5cb8, []int{0}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Collection) GetDenomId() string {
	if m != nil {
		return m.DenomId
	}
	return ""
}

func (m *Collection) GetMintRoyalties() []Royalty {
	if m != nil {
		return m.MintRoyalties
	}
	return nil
}

func (m *Collection) GetResaleRoyalties() []Royalty {
	if m != nil {
		return m.ResaleRoyalties
	}
	return nil
}

func (m *Collection) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Collection) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Collection)(nil), "cudoventures.cudosnode.marketplace.Collection")
}

func init() {
	proto.RegisterFile("cudoventures/cudosnode/marketplace/collection.proto", fileDescriptor_cdd7b78f471f5cb8)
}

var fileDescriptor_cdd7b78f471f5cb8 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe3, 0xf4, 0xe7, 0xeb, 0x67, 0x04, 0x48, 0x56, 0x87, 0xa8, 0x83, 0x89, 0x3a, 0x45,
	0x42, 0x38, 0x88, 0x4a, 0x5c, 0x40, 0x3b, 0xb1, 0xa1, 0x0c, 0x20, 0xc1, 0x94, 0xc6, 0x87, 0x62,
	0x91, 0xf8, 0x54, 0x8e, 0x53, 0xc8, 0x5d, 0x70, 0x05, 0x5c, 0x4f, 0xc7, 0x8e, 0x4c, 0x08, 0xb5,
	0x37, 0x82, 0x9a, 0x88, 0xd2, 0xb2, 0x80, 0xd8, 0xce, 0x6b, 0xe9, 0x79, 0x7c, 0xec, 0x97, 0x0e,
	0x92, 0x42, 0xe2, 0x0c, 0xb4, 0x2d, 0x0c, 0xe4, 0xe1, 0x3a, 0xe4, 0x1a, 0x25, 0x84, 0x59, 0x6c,
	0x1e, 0xc0, 0x4e, 0xd3, 0x38, 0x81, 0x30, 0xc1, 0x34, 0x85, 0xc4, 0x2a, 0xd4, 0x62, 0x6a, 0xd0,
	0x22, 0xeb, 0x6f, 0x43, 0x62, 0x03, 0x89, 0x2d, 0xa8, 0x77, 0xfa, 0x0b, 0xb1, 0xc1, 0x32, 0x4e,
	0x6d, 0x59, 0x5b, 0x7b, 0xdd, 0x09, 0x4e, 0xb0, 0x1a, 0xc3, 0xf5, 0x54, 0x9f, 0xf6, 0x5f, 0x5c,
	0x4a, 0x47, 0x9b, 0x05, 0xd8, 0x01, 0x75, 0x95, 0xf4, 0x88, 0x4f, 0x82, 0x66, 0xe4, 0x2a, 0xc9,
	0x3c, 0xfa, 0x4f, 0x82, 0xc6, 0xec, 0x42, 0x7a, 0xae, 0x4f, 0x82, 0xff, 0xd1, 0x67, 0x64, 0xd7,
	0x74, 0x3f, 0x53, 0xda, 0x46, 0xd5, 0x1d, 0x0a, 0x72, 0xaf, 0xe1, 0x37, 0x82, 0xbd, 0xb3, 0x63,
	0xf1, 0xf3, 0xf2, 0xa2, 0x86, 0xca, 0x61, 0x73, 0xfe, 0x76, 0xe4, 0x44, 0xbb, 0x1e, 0x76, 0x4b,
	0x0f, 0x0d, 0xe4, 0x71, 0x0a, 0x5f, 0xea, 0xe6, 0x5f, 0xd5, 0xdf, 0x4d, 0xac, 0x47, 0x3b, 0x33,
	0x30, 0xea, 0x4e, 0x81, 0xf4, 0x5a, 0x3e, 0x09, 0x3a, 0xd1, 0x26, 0xb3, 0x2e, 0x6d, 0xe1, 0xa3,
	0x06, 0xe3, 0xb5, 0xab, 0x97, 0xd6, 0x61, 0x78, 0x39, 0x5f, 0x72, 0xb2, 0x58, 0x72, 0xf2, 0xbe,
	0xe4, 0xe4, 0x79, 0xc5, 0x9d, 0xc5, 0x8a, 0x3b, 0xaf, 0x2b, 0xee, 0xdc, 0x9c, 0x4f, 0x94, 0xbd,
	0x2f, 0xc6, 0x22, 0xc1, 0x2c, 0x1c, 0x15, 0x12, 0xaf, 0x76, 0xda, 0x38, 0xa9, 0xea, 0x78, 0xda,
	0x29, 0xc4, 0x96, 0x53, 0xc8, 0xc7, 0xed, 0xea, 0xe7, 0x07, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x45, 0xf8, 0x83, 0xdb, 0x1c, 0x02, 0x00, 0x00,
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x32
	}
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.ResaleRoyalties) > 0 {
		for iNdEx := len(m.ResaleRoyalties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResaleRoyalties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MintRoyalties) > 0 {
		for iNdEx := len(m.MintRoyalties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintRoyalties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCollection(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DenomId) > 0 {
		i -= len(m.DenomId)
		copy(dAtA[i:], m.DenomId)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.DenomId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCollection(uint64(m.Id))
	}
	l = len(m.DenomId)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.MintRoyalties) > 0 {
		for _, e := range m.MintRoyalties {
			l = e.Size()
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	if len(m.ResaleRoyalties) > 0 {
		for _, e := range m.ResaleRoyalties {
			l = e.Size()
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	if m.Verified {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	return n
}

func sovCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollection(x uint64) (n int) {
	return sovCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollection
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
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRoyalties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRoyalties = append(m.MintRoyalties, Royalty{})
			if err := m.MintRoyalties[len(m.MintRoyalties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResaleRoyalties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResaleRoyalties = append(m.ResaleRoyalties, Royalty{})
			if err := m.ResaleRoyalties[len(m.ResaleRoyalties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
			m.Verified = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollection
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
func skipCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
				return 0, ErrInvalidLengthCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollection = fmt.Errorf("proto: unexpected end of group")
)
