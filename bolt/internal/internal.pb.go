// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Source
	Dashboard
	DashboardCell
	DashboardRange
	Template
	TemplateValue
	TemplateQuery
	Server
	Layout
	Cell
	Query
	Range
	AlertRule
	User
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Source struct {
	ID                 int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type               string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Username           string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	URL                string `protobuf:"bytes,6,opt,name=URL,proto3" json:"URL,omitempty"`
	Default            bool   `protobuf:"varint,7,opt,name=Default,proto3" json:"Default,omitempty"`
	Telegraf           string `protobuf:"bytes,8,opt,name=Telegraf,proto3" json:"Telegraf,omitempty"`
	InsecureSkipVerify bool   `protobuf:"varint,9,opt,name=InsecureSkipVerify,proto3" json:"InsecureSkipVerify,omitempty"`
	MetaURL            string `protobuf:"bytes,10,opt,name=MetaURL,proto3" json:"MetaURL,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

type Dashboard struct {
	ID        int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string           `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Cells     []*DashboardCell `protobuf:"bytes,3,rep,name=cells" json:"cells,omitempty"`
	Templates []*Template      `protobuf:"bytes,4,rep,name=templates" json:"templates,omitempty"`
}

func (m *Dashboard) Reset()                    { *m = Dashboard{} }
func (m *Dashboard) String() string            { return proto.CompactTextString(m) }
func (*Dashboard) ProtoMessage()               {}
func (*Dashboard) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Dashboard) GetCells() []*DashboardCell {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Dashboard) GetTemplates() []*Template {
	if m != nil {
		return m.Templates
	}
	return nil
}

type DashboardCell struct {
	X       int32                      `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32                      `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32                      `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32                      `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query                   `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	Name    string                     `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Type    string                     `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	ID      string                     `protobuf:"bytes,8,opt,name=ID,proto3" json:"ID,omitempty"`
	Axes    map[string]*DashboardRange `protobuf:"bytes,9,rep,name=axes" json:"axes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DashboardCell) Reset()                    { *m = DashboardCell{} }
func (m *DashboardCell) String() string            { return proto.CompactTextString(m) }
func (*DashboardCell) ProtoMessage()               {}
func (*DashboardCell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *DashboardCell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *DashboardCell) GetAxes() map[string]*DashboardRange {
	if m != nil {
		return m.Axes
	}
	return nil
}

type DashboardRange struct {
	Bounds []int32 `protobuf:"varint,1,rep,name=bounds" json:"bounds,omitempty"`
}

func (m *DashboardRange) Reset()                    { *m = DashboardRange{} }
func (m *DashboardRange) String() string            { return proto.CompactTextString(m) }
func (*DashboardRange) ProtoMessage()               {}
func (*DashboardRange) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

type Template struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TempVar string           `protobuf:"bytes,2,opt,name=temp_var,json=tempVar,proto3" json:"temp_var,omitempty"`
	Values  []*TemplateValue `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	Type    string           `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Label   string           `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Query   *TemplateQuery   `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Template) GetValues() []*TemplateValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Template) GetQuery() *TemplateQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type TemplateValue struct {
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Selected bool   `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
}

func (m *TemplateValue) Reset()                    { *m = TemplateValue{} }
func (m *TemplateValue) String() string            { return proto.CompactTextString(m) }
func (*TemplateValue) ProtoMessage()               {}
func (*TemplateValue) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

type TemplateQuery struct {
	Command     string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Db          string `protobuf:"bytes,2,opt,name=db,proto3" json:"db,omitempty"`
	Rp          string `protobuf:"bytes,3,opt,name=rp,proto3" json:"rp,omitempty"`
	Measurement string `protobuf:"bytes,4,opt,name=measurement,proto3" json:"measurement,omitempty"`
	TagKey      string `protobuf:"bytes,5,opt,name=tag_key,json=tagKey,proto3" json:"tag_key,omitempty"`
	FieldKey    string `protobuf:"bytes,6,opt,name=field_key,json=fieldKey,proto3" json:"field_key,omitempty"`
}

func (m *TemplateQuery) Reset()                    { *m = TemplateQuery{} }
func (m *TemplateQuery) String() string            { return proto.CompactTextString(m) }
func (*TemplateQuery) ProtoMessage()               {}
func (*TemplateQuery) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

type Server struct {
	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	URL      string `protobuf:"bytes,5,opt,name=URL,proto3" json:"URL,omitempty"`
	SrcID    int64  `protobuf:"varint,6,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	Active   bool   `protobuf:"varint,7,opt,name=Active,proto3" json:"Active,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

type Layout struct {
	ID          string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Application string  `protobuf:"bytes,2,opt,name=Application,proto3" json:"Application,omitempty"`
	Measurement string  `protobuf:"bytes,3,opt,name=Measurement,proto3" json:"Measurement,omitempty"`
	Cells       []*Cell `protobuf:"bytes,4,rep,name=Cells" json:"Cells,omitempty"`
	Autoflow    bool    `protobuf:"varint,5,opt,name=Autoflow,proto3" json:"Autoflow,omitempty"`
}

func (m *Layout) Reset()                    { *m = Layout{} }
func (m *Layout) String() string            { return proto.CompactTextString(m) }
func (*Layout) ProtoMessage()               {}
func (*Layout) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{8} }

func (m *Layout) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type Cell struct {
	X       int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32    `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32    `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	I       string   `protobuf:"bytes,6,opt,name=i,proto3" json:"i,omitempty"`
	Name    string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Yranges []int64  `protobuf:"varint,8,rep,name=yranges" json:"yranges,omitempty"`
	Ylabels []string `protobuf:"bytes,9,rep,name=ylabels" json:"ylabels,omitempty"`
	Type    string   `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *Cell) Reset()                    { *m = Cell{} }
func (m *Cell) String() string            { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()               {}
func (*Cell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{9} }

func (m *Cell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Query struct {
	Command  string   `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	DB       string   `protobuf:"bytes,2,opt,name=DB,proto3" json:"DB,omitempty"`
	RP       string   `protobuf:"bytes,3,opt,name=RP,proto3" json:"RP,omitempty"`
	GroupBys []string `protobuf:"bytes,4,rep,name=GroupBys" json:"GroupBys,omitempty"`
	Wheres   []string `protobuf:"bytes,5,rep,name=Wheres" json:"Wheres,omitempty"`
	Label    string   `protobuf:"bytes,6,opt,name=Label,proto3" json:"Label,omitempty"`
	Range    *Range   `protobuf:"bytes,7,opt,name=Range" json:"Range,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{10} }

func (m *Query) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

type Range struct {
	Upper int64 `protobuf:"varint,1,opt,name=Upper,proto3" json:"Upper,omitempty"`
	Lower int64 `protobuf:"varint,2,opt,name=Lower,proto3" json:"Lower,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{11} }

type AlertRule struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	JSON   string `protobuf:"bytes,2,opt,name=JSON,proto3" json:"JSON,omitempty"`
	SrcID  int64  `protobuf:"varint,3,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	KapaID int64  `protobuf:"varint,4,opt,name=KapaID,proto3" json:"KapaID,omitempty"`
}

func (m *AlertRule) Reset()                    { *m = AlertRule{} }
func (m *AlertRule) String() string            { return proto.CompactTextString(m) }
func (*AlertRule) ProtoMessage()               {}
func (*AlertRule) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{12} }

type User struct {
	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{13} }

func init() {
	proto.RegisterType((*Source)(nil), "internal.Source")
	proto.RegisterType((*Dashboard)(nil), "internal.Dashboard")
	proto.RegisterType((*DashboardCell)(nil), "internal.DashboardCell")
	proto.RegisterType((*DashboardRange)(nil), "internal.DashboardRange")
	proto.RegisterType((*Template)(nil), "internal.Template")
	proto.RegisterType((*TemplateValue)(nil), "internal.TemplateValue")
	proto.RegisterType((*TemplateQuery)(nil), "internal.TemplateQuery")
	proto.RegisterType((*Server)(nil), "internal.Server")
	proto.RegisterType((*Layout)(nil), "internal.Layout")
	proto.RegisterType((*Cell)(nil), "internal.Cell")
	proto.RegisterType((*Query)(nil), "internal.Query")
	proto.RegisterType((*Range)(nil), "internal.Range")
	proto.RegisterType((*AlertRule)(nil), "internal.AlertRule")
	proto.RegisterType((*User)(nil), "internal.User")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x51, 0x8f, 0xdb, 0x44,
	0x10, 0xd6, 0xc6, 0x76, 0x12, 0xcf, 0xb5, 0x07, 0x5a, 0x55, 0x74, 0x29, 0x2f, 0xc1, 0x02, 0x29,
	0x20, 0x11, 0x50, 0x2b, 0x24, 0xc4, 0x5b, 0x7a, 0x41, 0x28, 0xdc, 0xb5, 0x5c, 0x37, 0x77, 0xc7,
	0x13, 0xaa, 0x36, 0xc9, 0xe4, 0xce, 0xaa, 0x13, 0x9b, 0xb5, 0x7d, 0x39, 0xff, 0x05, 0x9e, 0xf8,
	0x05, 0x48, 0x48, 0xfc, 0x02, 0xc4, 0x3b, 0x3f, 0x82, 0x3f, 0x84, 0x66, 0x77, 0x6d, 0x27, 0x34,
	0x45, 0x7d, 0xe2, 0xe9, 0xf6, 0x9b, 0xd9, 0xcc, 0x78, 0xbf, 0xf9, 0xe6, 0xd3, 0xc1, 0x71, 0xbc,
	0x29, 0x50, 0x6f, 0x54, 0x32, 0xca, 0x74, 0x5a, 0xa4, 0xbc, 0x5f, 0xe3, 0xe8, 0xe7, 0x0e, 0x74,
	0x67, 0x69, 0xa9, 0x17, 0xc8, 0x8f, 0xa1, 0x33, 0x9d, 0x08, 0x36, 0x60, 0x43, 0x4f, 0x76, 0xa6,
	0x13, 0xce, 0xc1, 0x7f, 0xae, 0xd6, 0x28, 0x3a, 0x03, 0x36, 0x0c, 0xa5, 0x39, 0x53, 0xec, 0xa2,
	0xca, 0x50, 0x78, 0x36, 0x46, 0x67, 0xfe, 0x08, 0xfa, 0x97, 0x39, 0x55, 0x5b, 0xa3, 0xf0, 0x4d,
	0xbc, 0xc1, 0x94, 0x3b, 0x57, 0x79, 0xbe, 0x4d, 0xf5, 0x52, 0x04, 0x36, 0x57, 0x63, 0xfe, 0x2e,
	0x78, 0x97, 0xf2, 0x4c, 0x74, 0x4d, 0x98, 0x8e, 0x5c, 0x40, 0x6f, 0x82, 0x2b, 0x55, 0x26, 0x85,
	0xe8, 0x0d, 0xd8, 0xb0, 0x2f, 0x6b, 0x48, 0x75, 0x2e, 0x30, 0xc1, 0x6b, 0xad, 0x56, 0xa2, 0x6f,
	0xeb, 0xd4, 0x98, 0x8f, 0x80, 0x4f, 0x37, 0x39, 0x2e, 0x4a, 0x8d, 0xb3, 0x57, 0x71, 0x76, 0x85,
	0x3a, 0x5e, 0x55, 0x22, 0x34, 0x05, 0x0e, 0x64, 0xa8, 0xcb, 0x33, 0x2c, 0x14, 0xf5, 0x06, 0x53,
	0xaa, 0x86, 0xd1, 0x2f, 0x0c, 0xc2, 0x89, 0xca, 0x6f, 0xe6, 0xa9, 0xd2, 0xcb, 0xb7, 0xe2, 0xe3,
	0x33, 0x08, 0x16, 0x98, 0x24, 0xb9, 0xf0, 0x06, 0xde, 0xf0, 0xe8, 0xf1, 0xc3, 0x51, 0x43, 0x74,
	0x53, 0xe7, 0x04, 0x93, 0x44, 0xda, 0x5b, 0xfc, 0x0b, 0x08, 0x0b, 0x5c, 0x67, 0x89, 0x2a, 0x30,
	0x17, 0xbe, 0xf9, 0x09, 0x6f, 0x7f, 0x72, 0xe1, 0x52, 0xb2, 0xbd, 0x14, 0xfd, 0xd9, 0x81, 0xfb,
	0x7b, 0xa5, 0xf8, 0x3d, 0x60, 0x77, 0xe6, 0xab, 0x02, 0xc9, 0xee, 0x08, 0x55, 0xe6, 0x8b, 0x02,
	0xc9, 0x2a, 0x42, 0x5b, 0x33, 0x9b, 0x40, 0xb2, 0x2d, 0xa1, 0x1b, 0x33, 0x91, 0x40, 0xb2, 0x1b,
	0xfe, 0x09, 0xf4, 0x7e, 0x2a, 0x51, 0xc7, 0x98, 0x8b, 0xc0, 0x74, 0x7e, 0xa7, 0xed, 0xfc, 0xa2,
	0x44, 0x5d, 0xc9, 0x3a, 0x4f, 0x2f, 0x35, 0xd3, 0xb4, 0xa3, 0x31, 0x67, 0x8a, 0x15, 0x34, 0xf9,
	0x9e, 0x8d, 0xd1, 0xd9, 0x31, 0x64, 0xe7, 0x41, 0x0c, 0x7d, 0x09, 0xbe, 0xba, 0xc3, 0x5c, 0x84,
	0xa6, 0xfe, 0x87, 0x6f, 0x20, 0x63, 0x34, 0xbe, 0xc3, 0xfc, 0x9b, 0x4d, 0xa1, 0x2b, 0x69, 0xae,
	0x3f, 0x7a, 0x01, 0x61, 0x13, 0x22, 0x55, 0xbc, 0xc2, 0xca, 0x3c, 0x30, 0x94, 0x74, 0xe4, 0x23,
	0x08, 0x6e, 0x55, 0x52, 0x5a, 0xe2, 0x8f, 0x1e, 0x8b, 0x03, 0x65, 0xa5, 0xda, 0x5c, 0xa3, 0xb4,
	0xd7, 0xbe, 0xee, 0x7c, 0xc5, 0xa2, 0x21, 0x1c, 0xef, 0x27, 0xf9, 0x7b, 0xd0, 0x9d, 0xa7, 0xe5,
	0x66, 0x99, 0x0b, 0x36, 0xf0, 0x86, 0x81, 0x74, 0x28, 0xfa, 0x8b, 0x91, 0xb4, 0x2c, 0xdd, 0x3b,
	0x23, 0xb7, 0x0f, 0x7a, 0x1f, 0xfa, 0x34, 0x8a, 0x97, 0xb7, 0x4a, 0xbb, 0xb1, 0xf7, 0x08, 0x5f,
	0x29, 0xcd, 0x3f, 0x87, 0xae, 0x69, 0x77, 0x60, 0xf4, 0x75, 0xb9, 0x2b, 0xca, 0x4b, 0x77, 0xad,
	0x21, 0xd0, 0xdf, 0x21, 0xf0, 0x01, 0x04, 0x89, 0x9a, 0x63, 0xe2, 0x76, 0xc3, 0x02, 0x12, 0x15,
	0x4d, 0xa2, 0x32, 0xfc, 0x1f, 0xac, 0x6c, 0xe7, 0x65, 0x6f, 0x45, 0x97, 0x70, 0x7f, 0xaf, 0x63,
	0xd3, 0x89, 0xed, 0x77, 0x6a, 0x49, 0x0c, 0x1d, 0x55, 0xb4, 0x56, 0x39, 0x26, 0xb8, 0x28, 0x70,
	0x69, 0x64, 0xd3, 0x97, 0x0d, 0x8e, 0x7e, 0x63, 0x6d, 0x5d, 0xd3, 0x8f, 0x16, 0x67, 0x91, 0xae,
	0xd7, 0x6a, 0xb3, 0x74, 0xa5, 0x6b, 0x48, 0xbc, 0x2d, 0xe7, 0xae, 0x74, 0x67, 0x39, 0x27, 0xac,
	0x33, 0x67, 0x12, 0x1d, 0x9d, 0xf1, 0x01, 0x1c, 0xad, 0x51, 0xe5, 0xa5, 0xc6, 0x35, 0x6e, 0x0a,
	0x47, 0xc1, 0x6e, 0x88, 0x3f, 0x84, 0x5e, 0xa1, 0xae, 0x5f, 0xd2, 0xe8, 0x2d, 0x17, 0xdd, 0x42,
	0x5d, 0x9f, 0x62, 0xc5, 0x3f, 0x80, 0x70, 0x15, 0x63, 0xb2, 0x34, 0x29, 0x2b, 0xc8, 0xbe, 0x09,
	0x9c, 0x62, 0x15, 0xfd, 0xce, 0xa0, 0x3b, 0x43, 0x7d, 0x8b, 0xfa, 0xad, 0xb6, 0x75, 0xd7, 0xa9,
	0xbc, 0xff, 0x70, 0x2a, 0xff, 0xb0, 0x53, 0x05, 0xad, 0x53, 0x3d, 0x80, 0x60, 0xa6, 0x17, 0xd3,
	0x89, 0xf9, 0x22, 0x4f, 0x5a, 0x40, 0x1a, 0x1b, 0x2f, 0x8a, 0xf8, 0x16, 0x9d, 0x7d, 0x39, 0x14,
	0xfd, 0xca, 0xa0, 0x7b, 0xa6, 0xaa, 0xb4, 0x2c, 0x5e, 0x53, 0xd8, 0x00, 0x8e, 0xc6, 0x59, 0x96,
	0xc4, 0x0b, 0x55, 0xc4, 0xe9, 0xc6, 0x7d, 0xed, 0x6e, 0x88, 0x6e, 0x3c, 0xdb, 0xe1, 0xce, 0x7e,
	0xf7, 0x6e, 0x88, 0x7f, 0x04, 0xc1, 0x89, 0x31, 0x21, 0xeb, 0x28, 0xc7, 0xad, 0x5e, 0xac, 0xf7,
	0x98, 0x24, 0x3d, 0x70, 0x5c, 0x16, 0xe9, 0x2a, 0x49, 0xb7, 0xe6, 0x25, 0x7d, 0xd9, 0xe0, 0xe8,
	0x6f, 0x06, 0xfe, 0xff, 0x65, 0x2e, 0xf7, 0x80, 0xc5, 0x6e, 0x90, 0x2c, 0x6e, 0xac, 0xa6, 0xb7,
	0x63, 0x35, 0x02, 0x7a, 0x95, 0xa6, 0xa5, 0xcd, 0x45, 0x7f, 0xe0, 0x0d, 0x3d, 0x59, 0x43, 0x93,
	0x31, 0x3b, 0x62, 0x3d, 0x26, 0x94, 0x35, 0x6c, 0x34, 0x0f, 0xad, 0xe6, 0xa3, 0x3f, 0x18, 0x04,
	0x8d, 0x72, 0x4f, 0xf6, 0x95, 0x7b, 0xd2, 0x2a, 0x77, 0xf2, 0xb4, 0x56, 0xee, 0xe4, 0x29, 0x61,
	0x79, 0x5e, 0x2b, 0x57, 0x9e, 0x13, 0x6b, 0xdf, 0xea, 0xb4, 0xcc, 0x9e, 0x56, 0x96, 0xde, 0x50,
	0x36, 0x98, 0xc6, 0xfd, 0xc3, 0x0d, 0x6a, 0xf7, 0xe6, 0x50, 0x3a, 0x44, 0xe2, 0x38, 0x33, 0x5b,
	0x6d, 0x5f, 0x69, 0x01, 0xff, 0x18, 0x02, 0xe3, 0x44, 0xe6, 0xa9, 0x7b, 0x04, 0x39, 0xf7, 0x32,
	0x7f, 0xa2, 0x27, 0xee, 0x1a, 0x55, 0xb9, 0xcc, 0x32, 0xd4, 0x4e, 0xd3, 0x16, 0x98, 0xda, 0xe9,
	0x16, 0xad, 0x1d, 0x79, 0xd2, 0x82, 0xe8, 0x47, 0x08, 0xc7, 0x09, 0xea, 0x42, 0x96, 0xc9, 0xeb,
	0x26, 0xc6, 0xc1, 0xff, 0x6e, 0xf6, 0xfd, 0xf3, 0x7a, 0x13, 0xe8, 0xdc, 0xea, 0xd7, 0xfb, 0x97,
	0x7e, 0x4f, 0x55, 0xa6, 0xa6, 0x13, 0x33, 0x58, 0x4f, 0x3a, 0x14, 0x7d, 0x0a, 0x3e, 0xed, 0xc9,
	0x4e, 0x65, 0xff, 0x4d, 0x3b, 0x36, 0xef, 0x9a, 0xff, 0x30, 0x9e, 0xfc, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0xc4, 0x18, 0xb5, 0x73, 0x08, 0x00, 0x00,
}
