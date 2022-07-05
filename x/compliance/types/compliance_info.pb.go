// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compliance/compliance_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ComplianceInfo struct {
	Vid                                int32                    `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	Pid                                int32                    `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	SoftwareVersion                    uint32                   `protobuf:"varint,3,opt,name=softwareVersion,proto3" json:"softwareVersion,omitempty"`
	CertificationType                  string                   `protobuf:"bytes,4,opt,name=certificationType,proto3" json:"certificationType,omitempty"`
	SoftwareVersionString              string                   `protobuf:"bytes,5,opt,name=softwareVersionString,proto3" json:"softwareVersionString,omitempty"`
	CDVersionNumber                    uint32                   `protobuf:"varint,6,opt,name=cDVersionNumber,proto3" json:"cDVersionNumber,omitempty"`
	SoftwareVersionCertificationStatus uint32                   `protobuf:"varint,7,opt,name=softwareVersionCertificationStatus,proto3" json:"softwareVersionCertificationStatus,omitempty"`
	Date                               string                   `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	Reason                             string                   `protobuf:"bytes,9,opt,name=reason,proto3" json:"reason,omitempty"`
	Owner                              string                   `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
	History                            []*ComplianceHistoryItem `protobuf:"bytes,11,rep,name=history,proto3" json:"history,omitempty"`
	CDCertificateId                    string                   `protobuf:"bytes,12,opt,name=cDCertificateId,proto3" json:"cDCertificateId,omitempty"`
	CertificationRoute                 string                   `protobuf:"bytes,13,opt,name=certificationRoute,proto3" json:"certificationRoute,omitempty"`
	ProgramType                        string                   `protobuf:"bytes,14,opt,name=programType,proto3" json:"programType,omitempty"`
	ProgramTypeVersion                 string                   `protobuf:"bytes,15,opt,name=programTypeVersion,proto3" json:"programTypeVersion,omitempty"`
	CompliantPlatformUsed              string                   `protobuf:"bytes,16,opt,name=compliantPlatformUsed,proto3" json:"compliantPlatformUsed,omitempty"`
	CompliantPlatformVersion           string                   `protobuf:"bytes,17,opt,name=compliantPlatformVersion,proto3" json:"compliantPlatformVersion,omitempty"`
	Transport                          string                   `protobuf:"bytes,18,opt,name=transport,proto3" json:"transport,omitempty"`
	FamilyId                           string                   `protobuf:"bytes,19,opt,name=familyId,proto3" json:"familyId,omitempty"`
	SupportedClusters                  string                   `protobuf:"bytes,20,opt,name=supportedClusters,proto3" json:"supportedClusters,omitempty"`
	OSVersion                          string                   `protobuf:"bytes,21,opt,name=OSVersion,proto3" json:"OSVersion,omitempty"`
	ParentChild                        string                   `protobuf:"bytes,22,opt,name=parentChild,proto3" json:"parentChild,omitempty"`
	CertificationIdOfSoftwareComponent string                   `protobuf:"bytes,23,opt,name=certificationIdOfSoftwareComponent,proto3" json:"certificationIdOfSoftwareComponent,omitempty"`
}

func (m *ComplianceInfo) Reset()         { *m = ComplianceInfo{} }
func (m *ComplianceInfo) String() string { return proto.CompactTextString(m) }
func (*ComplianceInfo) ProtoMessage()    {}
func (*ComplianceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96709118384d66e8, []int{0}
}
func (m *ComplianceInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComplianceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComplianceInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComplianceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceInfo.Merge(m, src)
}
func (m *ComplianceInfo) XXX_Size() int {
	return m.Size()
}
func (m *ComplianceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceInfo proto.InternalMessageInfo

func (m *ComplianceInfo) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *ComplianceInfo) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ComplianceInfo) GetSoftwareVersion() uint32 {
	if m != nil {
		return m.SoftwareVersion
	}
	return 0
}

func (m *ComplianceInfo) GetCertificationType() string {
	if m != nil {
		return m.CertificationType
	}
	return ""
}

func (m *ComplianceInfo) GetSoftwareVersionString() string {
	if m != nil {
		return m.SoftwareVersionString
	}
	return ""
}

func (m *ComplianceInfo) GetCDVersionNumber() uint32 {
	if m != nil {
		return m.CDVersionNumber
	}
	return 0
}

func (m *ComplianceInfo) GetSoftwareVersionCertificationStatus() uint32 {
	if m != nil {
		return m.SoftwareVersionCertificationStatus
	}
	return 0
}

func (m *ComplianceInfo) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *ComplianceInfo) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *ComplianceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ComplianceInfo) GetHistory() []*ComplianceHistoryItem {
	if m != nil {
		return m.History
	}
	return nil
}

func (m *ComplianceInfo) GetCDCertificateId() string {
	if m != nil {
		return m.CDCertificateId
	}
	return ""
}

func (m *ComplianceInfo) GetCertificationRoute() string {
	if m != nil {
		return m.CertificationRoute
	}
	return ""
}

func (m *ComplianceInfo) GetProgramType() string {
	if m != nil {
		return m.ProgramType
	}
	return ""
}

func (m *ComplianceInfo) GetProgramTypeVersion() string {
	if m != nil {
		return m.ProgramTypeVersion
	}
	return ""
}

func (m *ComplianceInfo) GetCompliantPlatformUsed() string {
	if m != nil {
		return m.CompliantPlatformUsed
	}
	return ""
}

func (m *ComplianceInfo) GetCompliantPlatformVersion() string {
	if m != nil {
		return m.CompliantPlatformVersion
	}
	return ""
}

func (m *ComplianceInfo) GetTransport() string {
	if m != nil {
		return m.Transport
	}
	return ""
}

func (m *ComplianceInfo) GetFamilyId() string {
	if m != nil {
		return m.FamilyId
	}
	return ""
}

func (m *ComplianceInfo) GetSupportedClusters() string {
	if m != nil {
		return m.SupportedClusters
	}
	return ""
}

func (m *ComplianceInfo) GetOSVersion() string {
	if m != nil {
		return m.OSVersion
	}
	return ""
}

func (m *ComplianceInfo) GetParentChild() string {
	if m != nil {
		return m.ParentChild
	}
	return ""
}

func (m *ComplianceInfo) GetCertificationIdOfSoftwareComponent() string {
	if m != nil {
		return m.CertificationIdOfSoftwareComponent
	}
	return ""
}

func init() {
	proto.RegisterType((*ComplianceInfo)(nil), "zigbeealliance.distributedcomplianceledger.compliance.ComplianceInfo")
}

func init() { proto.RegisterFile("compliance/compliance_info.proto", fileDescriptor_96709118384d66e8) }

var fileDescriptor_96709118384d66e8 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xed, 0x7c, 0xfd, 0x77, 0xbf, 0xfe, 0x99, 0xb6, 0x98, 0x0a, 0x45, 0x51, 0x57, 0x59, 0x90,
	0x19, 0x89, 0x9f, 0x0d, 0x3b, 0x1a, 0x16, 0x44, 0xa0, 0x16, 0x4d, 0x80, 0x05, 0x9b, 0xca, 0x19,
	0xdf, 0x49, 0x2d, 0xcd, 0xd8, 0x23, 0xdb, 0x43, 0x09, 0x4f, 0xc1, 0x8a, 0x27, 0xe1, 0x21, 0x58,
	0x56, 0xac, 0x58, 0xa2, 0xf6, 0x45, 0x90, 0xed, 0x99, 0x4e, 0x9a, 0x04, 0xa9, 0x62, 0xe7, 0x7b,
	0xce, 0xbd, 0xe7, 0x5e, 0xfb, 0x1e, 0x19, 0xb5, 0x13, 0x99, 0x17, 0x19, 0xa7, 0x22, 0x81, 0xa8,
	0x39, 0x9e, 0x71, 0x91, 0xca, 0xb0, 0x50, 0xd2, 0x48, 0xfc, 0xec, 0x0b, 0x1f, 0x0d, 0x01, 0x68,
	0xe6, 0xa9, 0x90, 0x71, 0x6d, 0x14, 0x1f, 0x96, 0x06, 0x58, 0x53, 0x90, 0x01, 0x1b, 0x81, 0x0a,
	0x1b, 0xe0, 0xb0, 0x33, 0x5f, 0xf8, 0x9c, 0x6b, 0x23, 0xd5, 0xf8, 0x8c, 0x1b, 0xc8, 0x7d, 0x83,
	0xc3, 0x07, 0x89, 0xd4, 0xb9, 0xd4, 0x67, 0x2e, 0x8a, 0x7c, 0xe0, 0xa9, 0xa3, 0x6f, 0x6b, 0x68,
	0xab, 0x77, 0x53, 0xdc, 0x17, 0xa9, 0xc4, 0x3b, 0x68, 0xf1, 0x13, 0x67, 0x24, 0x68, 0x07, 0x9d,
	0xe5, 0xd8, 0x1e, 0x2d, 0x52, 0x70, 0x46, 0xfe, 0xf3, 0x48, 0xc1, 0x19, 0xee, 0xa0, 0x6d, 0x2d,
	0x53, 0x73, 0x41, 0x15, 0x7c, 0x00, 0xa5, 0xb9, 0x14, 0x64, 0xb1, 0x1d, 0x74, 0x36, 0xe3, 0x69,
	0x18, 0x3f, 0x42, 0xbb, 0x09, 0x28, 0xc3, 0x53, 0x9e, 0x50, 0xc3, 0xa5, 0x78, 0x37, 0x2e, 0x80,
	0x2c, 0xb5, 0x83, 0xce, 0x7a, 0x3c, 0x4b, 0xe0, 0xa7, 0x68, 0x7f, 0x4a, 0x60, 0x60, 0x14, 0x17,
	0x23, 0xb2, 0xec, 0x2a, 0xe6, 0x93, 0x76, 0x9a, 0xe4, 0x65, 0x05, 0x9d, 0x94, 0xf9, 0x10, 0x14,
	0x59, 0xf1, 0xd3, 0x4c, 0xc1, 0xf8, 0x04, 0x1d, 0x4d, 0x49, 0xf4, 0x26, 0x67, 0x18, 0x18, 0x6a,
	0x4a, 0x4d, 0x56, 0x5d, 0xf1, 0x1d, 0x32, 0x31, 0x46, 0x4b, 0x8c, 0x1a, 0x20, 0x6b, 0x6e, 0x3c,
	0x77, 0xc6, 0x07, 0x68, 0x45, 0x01, 0xd5, 0x52, 0x90, 0x75, 0x87, 0x56, 0x11, 0x0e, 0xd1, 0xb2,
	0xbc, 0x10, 0xa0, 0x08, 0xb2, 0xf0, 0x31, 0xf9, 0xf9, 0xbd, 0xbb, 0x57, 0xed, 0xe2, 0x05, 0x63,
	0x0a, 0xb4, 0xf6, 0xd7, 0x89, 0x7d, 0x1a, 0x4e, 0xd1, 0x6a, 0xb5, 0x4b, 0xb2, 0xd1, 0x5e, 0xec,
	0x6c, 0x3c, 0x7e, 0x13, 0xfe, 0x93, 0x51, 0xc2, 0x66, 0xbf, 0xaf, 0xbc, 0x5e, 0xdf, 0x40, 0x1e,
	0xd7, 0xe2, 0xfe, 0xf5, 0x9a, 0xcb, 0x41, 0x9f, 0x91, 0xff, 0xdd, 0xe0, 0xd3, 0x30, 0x0e, 0x11,
	0xbe, 0xb5, 0xb2, 0x58, 0x96, 0x06, 0xc8, 0xa6, 0x4b, 0x9e, 0xc3, 0xe0, 0x36, 0xda, 0x28, 0x94,
	0x1c, 0x29, 0x9a, 0xbb, 0xad, 0x6f, 0xb9, 0xc4, 0x49, 0xc8, 0x2a, 0x4e, 0x84, 0xb5, 0x95, 0xb6,
	0xbd, 0xe2, 0x2c, 0x63, 0xfd, 0x51, 0x5f, 0xcc, 0xbc, 0xcd, 0xa8, 0x49, 0xa5, 0xca, 0xdf, 0x6b,
	0x60, 0x64, 0xc7, 0xfb, 0x63, 0x2e, 0x89, 0x9f, 0x23, 0x32, 0x43, 0xd4, 0xbd, 0x76, 0x5d, 0xe1,
	0x5f, 0x79, 0xfc, 0x10, 0xad, 0x1b, 0x45, 0x85, 0x2e, 0xa4, 0x32, 0x04, 0xbb, 0xe4, 0x06, 0xc0,
	0x87, 0x68, 0x2d, 0xa5, 0x39, 0xcf, 0xc6, 0x7d, 0x46, 0xee, 0x39, 0xf2, 0x26, 0xb6, 0xce, 0xd7,
	0x65, 0x61, 0xd3, 0x80, 0xf5, 0xb2, 0x52, 0x1b, 0x50, 0x9a, 0xec, 0x79, 0xe7, 0xcf, 0x10, 0xb6,
	0xcf, 0xe9, 0xa0, 0x1e, 0x6a, 0xdf, 0xf7, 0xb9, 0x01, 0xdc, 0x4b, 0x52, 0x05, 0xc2, 0xf4, 0xce,
	0x79, 0xc6, 0xc8, 0x41, 0xf5, 0x92, 0x0d, 0x64, 0x9d, 0x7d, 0x6b, 0x03, 0x7d, 0x76, 0x9a, 0x0e,
	0x2a, 0x03, 0x5b, 0x03, 0x48, 0x01, 0xc2, 0x90, 0xfb, 0xae, 0xf0, 0x0e, 0x99, 0xc7, 0xf0, 0xe3,
	0xaa, 0x15, 0x5c, 0x5e, 0xb5, 0x82, 0xdf, 0x57, 0xad, 0xe0, 0xeb, 0x75, 0x6b, 0xe1, 0xf2, 0xba,
	0xb5, 0xf0, 0xeb, 0xba, 0xb5, 0xf0, 0xf1, 0xf5, 0x88, 0x9b, 0xf3, 0x72, 0x68, 0x1d, 0x16, 0x79,
	0x43, 0x76, 0x6b, 0x47, 0x46, 0x13, 0x8e, 0xec, 0x36, 0x0e, 0xec, 0x7a, 0x4f, 0x46, 0x9f, 0x27,
	0xfe, 0xa9, 0xc8, 0x8c, 0x0b, 0xd0, 0xc3, 0x15, 0xf7, 0x0d, 0x3d, 0xf9, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x04, 0x7c, 0x3a, 0x53, 0x26, 0x05, 0x00, 0x00,
}

func (m *ComplianceInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplianceInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComplianceInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CertificationIdOfSoftwareComponent) > 0 {
		i -= len(m.CertificationIdOfSoftwareComponent)
		copy(dAtA[i:], m.CertificationIdOfSoftwareComponent)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CertificationIdOfSoftwareComponent)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xba
	}
	if len(m.ParentChild) > 0 {
		i -= len(m.ParentChild)
		copy(dAtA[i:], m.ParentChild)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.ParentChild)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb2
	}
	if len(m.OSVersion) > 0 {
		i -= len(m.OSVersion)
		copy(dAtA[i:], m.OSVersion)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.OSVersion)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if len(m.SupportedClusters) > 0 {
		i -= len(m.SupportedClusters)
		copy(dAtA[i:], m.SupportedClusters)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.SupportedClusters)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	if len(m.FamilyId) > 0 {
		i -= len(m.FamilyId)
		copy(dAtA[i:], m.FamilyId)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.FamilyId)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if len(m.Transport) > 0 {
		i -= len(m.Transport)
		copy(dAtA[i:], m.Transport)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.Transport)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.CompliantPlatformVersion) > 0 {
		i -= len(m.CompliantPlatformVersion)
		copy(dAtA[i:], m.CompliantPlatformVersion)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CompliantPlatformVersion)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.CompliantPlatformUsed) > 0 {
		i -= len(m.CompliantPlatformUsed)
		copy(dAtA[i:], m.CompliantPlatformUsed)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CompliantPlatformUsed)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.ProgramTypeVersion) > 0 {
		i -= len(m.ProgramTypeVersion)
		copy(dAtA[i:], m.ProgramTypeVersion)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.ProgramTypeVersion)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.ProgramType) > 0 {
		i -= len(m.ProgramType)
		copy(dAtA[i:], m.ProgramType)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.ProgramType)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.CertificationRoute) > 0 {
		i -= len(m.CertificationRoute)
		copy(dAtA[i:], m.CertificationRoute)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CertificationRoute)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.CDCertificateId) > 0 {
		i -= len(m.CDCertificateId)
		copy(dAtA[i:], m.CDCertificateId)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CDCertificateId)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.History) > 0 {
		for iNdEx := len(m.History) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.History[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintComplianceInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x42
	}
	if m.SoftwareVersionCertificationStatus != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.SoftwareVersionCertificationStatus))
		i--
		dAtA[i] = 0x38
	}
	if m.CDVersionNumber != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.CDVersionNumber))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SoftwareVersionString) > 0 {
		i -= len(m.SoftwareVersionString)
		copy(dAtA[i:], m.SoftwareVersionString)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.SoftwareVersionString)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CertificationType) > 0 {
		i -= len(m.CertificationType)
		copy(dAtA[i:], m.CertificationType)
		i = encodeVarintComplianceInfo(dAtA, i, uint64(len(m.CertificationType)))
		i--
		dAtA[i] = 0x22
	}
	if m.SoftwareVersion != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.SoftwareVersion))
		i--
		dAtA[i] = 0x18
	}
	if m.Pid != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x10
	}
	if m.Vid != 0 {
		i = encodeVarintComplianceInfo(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComplianceInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovComplianceInfo(uint64(m.Vid))
	}
	if m.Pid != 0 {
		n += 1 + sovComplianceInfo(uint64(m.Pid))
	}
	if m.SoftwareVersion != 0 {
		n += 1 + sovComplianceInfo(uint64(m.SoftwareVersion))
	}
	l = len(m.CertificationType)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.SoftwareVersionString)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	if m.CDVersionNumber != 0 {
		n += 1 + sovComplianceInfo(uint64(m.CDVersionNumber))
	}
	if m.SoftwareVersionCertificationStatus != 0 {
		n += 1 + sovComplianceInfo(uint64(m.SoftwareVersionCertificationStatus))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	if len(m.History) > 0 {
		for _, e := range m.History {
			l = e.Size()
			n += 1 + l + sovComplianceInfo(uint64(l))
		}
	}
	l = len(m.CDCertificateId)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.CertificationRoute)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.ProgramType)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.ProgramTypeVersion)
	if l > 0 {
		n += 1 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.CompliantPlatformUsed)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.CompliantPlatformVersion)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.Transport)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.FamilyId)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.SupportedClusters)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.OSVersion)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.ParentChild)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	l = len(m.CertificationIdOfSoftwareComponent)
	if l > 0 {
		n += 2 + l + sovComplianceInfo(uint64(l))
	}
	return n
}

func sovComplianceInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceInfo(x uint64) (n int) {
	return sovComplianceInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComplianceInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceInfo
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
			return fmt.Errorf("proto: ComplianceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplianceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersion", wireType)
			}
			m.SoftwareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SoftwareVersionString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CDVersionNumber", wireType)
			}
			m.CDVersionNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CDVersionNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionCertificationStatus", wireType)
			}
			m.SoftwareVersionCertificationStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersionCertificationStatus |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field History", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.History = append(m.History, &ComplianceHistoryItem{})
			if err := m.History[len(m.History)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CDCertificateId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CDCertificateId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationRoute", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationRoute = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgramType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProgramType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgramTypeVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProgramTypeVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompliantPlatformUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompliantPlatformUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompliantPlatformVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompliantPlatformVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transport = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FamilyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FamilyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportedClusters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportedClusters = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OSVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentChild", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentChild = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationIdOfSoftwareComponent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceInfo
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
				return ErrInvalidLengthComplianceInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationIdOfSoftwareComponent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceInfo
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
func skipComplianceInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceInfo
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
					return 0, ErrIntOverflowComplianceInfo
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
					return 0, ErrIntOverflowComplianceInfo
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
				return 0, ErrInvalidLengthComplianceInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceInfo = fmt.Errorf("proto: unexpected end of group")
)
