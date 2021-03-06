// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.proto

package pb

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
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

type ChatEntity struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Time                 uint32   `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatEntity) Reset()         { *m = ChatEntity{} }
func (m *ChatEntity) String() string { return proto.CompactTextString(m) }
func (*ChatEntity) ProtoMessage()    {}
func (*ChatEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}
func (m *ChatEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatEntity.Merge(m, src)
}
func (m *ChatEntity) XXX_Size() int {
	return m.Size()
}
func (m *ChatEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatEntity.DiscardUnknown(m)
}

var xxx_messageInfo_ChatEntity proto.InternalMessageInfo

func (m *ChatEntity) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ChatEntity) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ChatEntity) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type GetChatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChatsRequest) Reset()         { *m = GetChatsRequest{} }
func (m *GetChatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetChatsRequest) ProtoMessage()    {}
func (*GetChatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}
func (m *GetChatsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChatsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetChatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChatsRequest.Merge(m, src)
}
func (m *GetChatsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetChatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChatsRequest proto.InternalMessageInfo

type GetChatsResponse struct {
	Chats                []*ChatEntity `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetChatsResponse) Reset()         { *m = GetChatsResponse{} }
func (m *GetChatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetChatsResponse) ProtoMessage()    {}
func (*GetChatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}
func (m *GetChatsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChatsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetChatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChatsResponse.Merge(m, src)
}
func (m *GetChatsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetChatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetChatsResponse proto.InternalMessageInfo

func (m *GetChatsResponse) GetChats() []*ChatEntity {
	if m != nil {
		return m.Chats
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatEntity)(nil), "pb.ChatEntity")
	proto.RegisterType((*GetChatsRequest)(nil), "pb.GetChatsRequest")
	proto.RegisterType((*GetChatsResponse)(nil), "pb.GetChatsResponse")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xc9, 0xcc, 0xff, 0x2b, 0x5e, 0x51, 0xc7, 0xac, 0xc2, 0x20, 0xa5, 0x14, 0x17, 0xdd,
	0xd8, 0x82, 0x6e, 0x5c, 0x2b, 0xe2, 0xce, 0x45, 0xdf, 0x20, 0x19, 0xae, 0x69, 0xc0, 0xe6, 0xc6,
	0xb9, 0xb7, 0x0b, 0xdf, 0xc8, 0x47, 0x71, 0xe9, 0x23, 0x48, 0x9f, 0x44, 0x9a, 0x22, 0xb3, 0x3b,
	0xe7, 0xcb, 0x47, 0x38, 0x17, 0x60, 0xd7, 0x5b, 0x69, 0xd2, 0x9e, 0x84, 0xf4, 0x2a, 0xb9, 0xed,
	0x95, 0x27, 0xf2, 0x6f, 0xd8, 0xda, 0x14, 0x5a, 0x1b, 0x23, 0x89, 0x95, 0x40, 0x91, 0x17, 0x63,
	0x7b, 0xe3, 0x83, 0xf4, 0xa3, 0x6b, 0x76, 0x34, 0xb4, 0x9e, 0x3c, 0xb5, 0x19, 0xbb, 0xf1, 0x35,
	0xb7, 0x5c, 0x72, 0x5a, 0xf4, 0xea, 0x05, 0xe0, 0xb1, 0xb7, 0xf2, 0x14, 0x25, 0xc8, 0x87, 0xd6,
	0xf0, 0x6f, 0x64, 0xdc, 0x1b, 0x55, 0xaa, 0xfa, 0xa4, 0xcb, 0x59, 0x1b, 0x38, 0x1e, 0x90, 0xd9,
	0x7a, 0x34, 0xab, 0x8c, 0xff, 0xea, 0x6c, 0x4b, 0x18, 0xd0, 0xac, 0x4b, 0x55, 0x9f, 0x75, 0x39,
	0x57, 0x97, 0x70, 0xf1, 0x8c, 0x32, 0x7f, 0xc9, 0x1d, 0xbe, 0x8f, 0xc8, 0x52, 0xdd, 0xc3, 0xe6,
	0x80, 0x38, 0x51, 0x64, 0xd4, 0xd7, 0xf0, 0x7f, 0xbe, 0x8a, 0x8d, 0x2a, 0xd7, 0xf5, 0xe9, 0xed,
	0x79, 0x93, 0x5c, 0x73, 0xd8, 0xd1, 0x2d, 0x8f, 0x0f, 0x9b, 0xcf, 0xa9, 0x50, 0x5f, 0x53, 0xa1,
	0xbe, 0xa7, 0x42, 0xfd, 0x4c, 0x85, 0x72, 0x47, 0x79, 0xf5, 0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0x72, 0x86, 0x1a, 0x14, 0x01, 0x00, 0x00,
}

func (this *ChatEntity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChatEntity)
	if !ok {
		that2, ok := that.(ChatEntity)
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
	if this.User != that1.User {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.Time != that1.Time {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetChatsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetChatsRequest)
	if !ok {
		that2, ok := that.(GetChatsRequest)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetChatsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetChatsResponse)
	if !ok {
		that2, ok := that.(GetChatsResponse)
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
	if len(this.Chats) != len(that1.Chats) {
		return false
	}
	for i := range this.Chats {
		if !this.Chats[i].Equal(that1.Chats[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *ChatEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatEntity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatEntity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Time != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintChat(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetChatsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChatsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetChatsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *GetChatsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChatsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetChatsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Chats) > 0 {
		for iNdEx := len(m.Chats) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chats[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintChat(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintChat(dAtA []byte, offset int, v uint64) int {
	offset -= sovChat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChatEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	if m.Time != 0 {
		n += 1 + sovChat(uint64(m.Time))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetChatsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetChatsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Chats) > 0 {
		for _, e := range m.Chats {
			l = e.Size()
			n += 1 + l + sovChat(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovChat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChat(x uint64) (n int) {
	return sovChat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChatEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: ChatEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetChatsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: GetChatsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChatsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetChatsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: GetChatsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChatsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chats = append(m.Chats, &ChatEntity{})
			if err := m.Chats[len(m.Chats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
				return 0, ErrInvalidLengthChat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChat = fmt.Errorf("proto: unexpected end of group")
)
