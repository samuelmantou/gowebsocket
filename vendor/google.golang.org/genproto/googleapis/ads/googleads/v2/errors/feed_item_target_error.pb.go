// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/feed_item_target_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible feed item target errors.
type FeedItemTargetErrorEnum_FeedItemTargetError int32

const (
	// Enum unspecified.
	FeedItemTargetErrorEnum_UNSPECIFIED FeedItemTargetErrorEnum_FeedItemTargetError = 0
	// The received error code is not known in this version.
	FeedItemTargetErrorEnum_UNKNOWN FeedItemTargetErrorEnum_FeedItemTargetError = 1
	// On CREATE, the FeedItemTarget must have a populated field in the oneof
	// target.
	FeedItemTargetErrorEnum_MUST_SET_TARGET_ONEOF_ON_CREATE FeedItemTargetErrorEnum_FeedItemTargetError = 2
	// The specified feed item target already exists, so it cannot be added.
	FeedItemTargetErrorEnum_FEED_ITEM_TARGET_ALREADY_EXISTS FeedItemTargetErrorEnum_FeedItemTargetError = 3
	// The schedules for a given feed item cannot overlap.
	FeedItemTargetErrorEnum_FEED_ITEM_SCHEDULES_CANNOT_OVERLAP FeedItemTargetErrorEnum_FeedItemTargetError = 4
	// Too many targets of a given type were added for a single feed item.
	FeedItemTargetErrorEnum_TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE FeedItemTargetErrorEnum_FeedItemTargetError = 5
	// Too many AdSchedules are enabled for the feed item for the given day.
	FeedItemTargetErrorEnum_TOO_MANY_SCHEDULES_PER_DAY FeedItemTargetErrorEnum_FeedItemTargetError = 6
	// A feed item may either have an enabled campaign target or an enabled ad
	// group target.
	FeedItemTargetErrorEnum_CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS FeedItemTargetErrorEnum_FeedItemTargetError = 7
	// Duplicate ad schedules aren't allowed.
	FeedItemTargetErrorEnum_DUPLICATE_AD_SCHEDULE FeedItemTargetErrorEnum_FeedItemTargetError = 8
	// Duplicate keywords aren't allowed.
	FeedItemTargetErrorEnum_DUPLICATE_KEYWORD FeedItemTargetErrorEnum_FeedItemTargetError = 9
)

var FeedItemTargetErrorEnum_FeedItemTargetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MUST_SET_TARGET_ONEOF_ON_CREATE",
	3: "FEED_ITEM_TARGET_ALREADY_EXISTS",
	4: "FEED_ITEM_SCHEDULES_CANNOT_OVERLAP",
	5: "TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE",
	6: "TOO_MANY_SCHEDULES_PER_DAY",
	7: "CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS",
	8: "DUPLICATE_AD_SCHEDULE",
	9: "DUPLICATE_KEYWORD",
}

var FeedItemTargetErrorEnum_FeedItemTargetError_value = map[string]int32{
	"UNSPECIFIED":                                               0,
	"UNKNOWN":                                                   1,
	"MUST_SET_TARGET_ONEOF_ON_CREATE":                           2,
	"FEED_ITEM_TARGET_ALREADY_EXISTS":                           3,
	"FEED_ITEM_SCHEDULES_CANNOT_OVERLAP":                        4,
	"TARGET_LIMIT_EXCEEDED_FOR_GIVEN_TYPE":                      5,
	"TOO_MANY_SCHEDULES_PER_DAY":                                6,
	"CANNOT_HAVE_ENABLED_CAMPAIGN_AND_ENABLED_AD_GROUP_TARGETS": 7,
	"DUPLICATE_AD_SCHEDULE":                                     8,
	"DUPLICATE_KEYWORD":                                         9,
}

func (x FeedItemTargetErrorEnum_FeedItemTargetError) String() string {
	return proto.EnumName(FeedItemTargetErrorEnum_FeedItemTargetError_name, int32(x))
}

func (FeedItemTargetErrorEnum_FeedItemTargetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8914d01526a778d9, []int{0, 0}
}

// Container for enum describing possible feed item target errors.
type FeedItemTargetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemTargetErrorEnum) Reset()         { *m = FeedItemTargetErrorEnum{} }
func (m *FeedItemTargetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemTargetErrorEnum) ProtoMessage()    {}
func (*FeedItemTargetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8914d01526a778d9, []int{0}
}

func (m *FeedItemTargetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTargetErrorEnum.Unmarshal(m, b)
}
func (m *FeedItemTargetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTargetErrorEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemTargetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTargetErrorEnum.Merge(m, src)
}
func (m *FeedItemTargetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemTargetErrorEnum.Size(m)
}
func (m *FeedItemTargetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTargetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTargetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.FeedItemTargetErrorEnum_FeedItemTargetError", FeedItemTargetErrorEnum_FeedItemTargetError_name, FeedItemTargetErrorEnum_FeedItemTargetError_value)
	proto.RegisterType((*FeedItemTargetErrorEnum)(nil), "google.ads.googleads.v2.errors.FeedItemTargetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/feed_item_target_error.proto", fileDescriptor_8914d01526a778d9)
}

var fileDescriptor_8914d01526a778d9 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0xa9, 0x0b, 0x2d, 0x6c, 0x0f, 0x18, 0xa3, 0x0a, 0xa8, 0x50, 0x90, 0x02, 0x42, 0x9c,
	0x6c, 0x29, 0x9c, 0x70, 0xc5, 0x61, 0xe2, 0x9d, 0xb8, 0x56, 0x9d, 0x5d, 0xcb, 0x5e, 0xbb, 0x0d,
	0x8a, 0x34, 0x32, 0xd8, 0x58, 0x91, 0x1a, 0x3b, 0xb2, 0x4d, 0x7f, 0x10, 0x47, 0x7e, 0x0a, 0xbf,
	0x82, 0x33, 0x57, 0xae, 0x1c, 0x90, 0xe3, 0x24, 0xe5, 0x50, 0x38, 0xed, 0xd3, 0xcc, 0xf7, 0xde,
	0xae, 0xf4, 0x96, 0x9d, 0x16, 0x55, 0x55, 0x5c, 0xe5, 0x56, 0x9a, 0x35, 0x56, 0x2f, 0x3b, 0x75,
	0x3d, 0xb2, 0xf2, 0xba, 0xae, 0xea, 0xc6, 0xfa, 0x9c, 0xe7, 0x19, 0x2d, 0xda, 0x7c, 0x49, 0x6d,
	0x5a, 0x17, 0x79, 0x4b, 0xeb, 0xb9, 0xb9, 0xaa, 0xab, 0xb6, 0x32, 0x06, 0xbd, 0xc3, 0x4c, 0xb3,
	0xc6, 0xdc, 0x99, 0xcd, 0xeb, 0x91, 0xd9, 0x9b, 0x4f, 0x9e, 0x6f, 0xc3, 0x57, 0x0b, 0x2b, 0x2d,
	0xcb, 0xaa, 0x4d, 0xdb, 0x45, 0x55, 0x36, 0xbd, 0x7b, 0xf8, 0x5b, 0x63, 0x4f, 0x26, 0x79, 0x9e,
	0x79, 0x6d, 0xbe, 0x54, 0xeb, 0x70, 0xec, 0x6c, 0x58, 0x7e, 0x59, 0x0e, 0x7f, 0x68, 0xec, 0xf1,
	0x2d, 0x3b, 0xe3, 0x21, 0x3b, 0x8a, 0x45, 0x14, 0xa0, 0xe3, 0x4d, 0x3c, 0xe4, 0xfa, 0x1d, 0xe3,
	0x88, 0x1d, 0xc6, 0xe2, 0x5c, 0xc8, 0x0b, 0xa1, 0xef, 0x19, 0x2f, 0xd9, 0x8b, 0x69, 0x1c, 0x29,
	0x8a, 0x50, 0x91, 0x82, 0xd0, 0x45, 0x45, 0x52, 0xa0, 0x9c, 0x90, 0x14, 0xe4, 0x84, 0x08, 0x0a,
	0x75, 0xad, 0x83, 0x26, 0x88, 0x9c, 0x3c, 0x85, 0xd3, 0x2d, 0x05, 0x7e, 0x88, 0xc0, 0x67, 0x84,
	0x97, 0x5e, 0xa4, 0x22, 0x7d, 0xdf, 0x78, 0xcd, 0x86, 0x37, 0x50, 0xe4, 0x9c, 0x21, 0x8f, 0x7d,
	0x8c, 0xc8, 0x01, 0x21, 0xa4, 0x22, 0x99, 0x60, 0xe8, 0x43, 0xa0, 0xdf, 0x35, 0xde, 0xb0, 0x57,
	0x9b, 0x08, 0xdf, 0x9b, 0x7a, 0x8a, 0xf0, 0xd2, 0x41, 0xe4, 0xc8, 0x69, 0x22, 0x43, 0x72, 0xbd,
	0x04, 0x05, 0xa9, 0x59, 0x80, 0xfa, 0x3d, 0x63, 0xc0, 0x4e, 0x94, 0x94, 0x34, 0x05, 0x31, 0xfb,
	0x2b, 0x30, 0xc0, 0x90, 0x38, 0xcc, 0xf4, 0x03, 0xe3, 0x3d, 0x7b, 0xb7, 0x49, 0x3f, 0x83, 0x04,
	0x09, 0x05, 0x8c, 0x7d, 0xe4, 0xe4, 0xc0, 0x34, 0x00, 0xcf, 0x15, 0x04, 0x82, 0xef, 0x86, 0xc0,
	0xc9, 0x0d, 0x65, 0x1c, 0x6c, 0x9e, 0x1f, 0xe9, 0x87, 0xc6, 0x33, 0x76, 0xcc, 0xe3, 0xc0, 0xf7,
	0x1c, 0x50, 0xd8, 0xed, 0xb7, 0x57, 0xe8, 0xf7, 0x8d, 0x63, 0xf6, 0xe8, 0x66, 0x75, 0x8e, 0xb3,
	0x0b, 0x19, 0x72, 0xfd, 0xc1, 0xf8, 0xd7, 0x1e, 0x1b, 0x7e, 0xaa, 0x96, 0xe6, 0xff, 0x3b, 0x1c,
	0x3f, 0xbd, 0xa5, 0x86, 0xa0, 0xeb, 0x2f, 0xd8, 0xfb, 0xc0, 0x37, 0xde, 0xa2, 0xba, 0x4a, 0xcb,
	0xc2, 0xac, 0xea, 0xc2, 0x2a, 0xf2, 0x72, 0xdd, 0xee, 0xf6, 0x33, 0xad, 0x16, 0xcd, 0xbf, 0xfe,
	0xd6, 0x69, 0x7f, 0x7c, 0xd5, 0xf6, 0x5d, 0x80, 0x6f, 0xda, 0xc0, 0xed, 0xc3, 0x20, 0x6b, 0xcc,
	0x5e, 0x76, 0x2a, 0x19, 0x99, 0xeb, 0x2b, 0x9b, 0xef, 0x5b, 0x60, 0x0e, 0x59, 0x33, 0xdf, 0x01,
	0xf3, 0x64, 0x34, 0xef, 0x81, 0x9f, 0xda, 0xb0, 0x9f, 0xda, 0x36, 0x64, 0x8d, 0x6d, 0xef, 0x10,
	0xdb, 0x4e, 0x46, 0xb6, 0xdd, 0x43, 0x1f, 0x0f, 0xd6, 0xaf, 0x7b, 0xfb, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x8c, 0x2f, 0xa6, 0x5e, 0xf8, 0x02, 0x00, 0x00,
}
