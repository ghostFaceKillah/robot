// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.7.0
// source: sensor.proto

package msg

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Odom2DUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp float64 `protobuf:"fixed64,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	X         float64 `protobuf:"fixed64,2,opt,name=x,proto3" json:"x,omitempty"`         // orientation: forward
	Y         float64 `protobuf:"fixed64,3,opt,name=y,proto3" json:"y,omitempty"`         // orientation: left
	Theta     float64 `protobuf:"fixed64,4,opt,name=theta,proto3" json:"theta,omitempty"` // orientation: counterclockwise rotation around z
}

func (x *Odom2DUpdate) Reset() {
	*x = Odom2DUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Odom2DUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Odom2DUpdate) ProtoMessage() {}

func (x *Odom2DUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_sensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Odom2DUpdate.ProtoReflect.Descriptor instead.
func (*Odom2DUpdate) Descriptor() ([]byte, []int) {
	return file_sensor_proto_rawDescGZIP(), []int{0}
}

func (x *Odom2DUpdate) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Odom2DUpdate) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Odom2DUpdate) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Odom2DUpdate) GetTheta() float64 {
	if x != nil {
		return x.Theta
	}
	return 0
}

type Lidar2DScan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp      float64   `protobuf:"fixed64,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // using timestamp stuff
	Scan           []float64 `protobuf:"fixed64,2,rep,packed,name=scan,proto3" json:"scan,omitempty"`    // 1 x 1081 in the case of data from
	NScans         uint32    `protobuf:"varint,3,opt,name=n_scans,json=nScans,proto3" json:"n_scans,omitempty"`
	MinRange       float64   `protobuf:"fixed64,4,opt,name=min_range,json=minRange,proto3" json:"min_range,omitempty"`
	MaxRange       float64   `protobuf:"fixed64,5,opt,name=max_range,json=maxRange,proto3" json:"max_range,omitempty"`
	ScanResolution float64   `protobuf:"fixed64,6,opt,name=scan_resolution,json=scanResolution,proto3" json:"scan_resolution,omitempty"`
}

func (x *Lidar2DScan) Reset() {
	*x = Lidar2DScan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lidar2DScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lidar2DScan) ProtoMessage() {}

func (x *Lidar2DScan) ProtoReflect() protoreflect.Message {
	mi := &file_sensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lidar2DScan.ProtoReflect.Descriptor instead.
func (*Lidar2DScan) Descriptor() ([]byte, []int) {
	return file_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *Lidar2DScan) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Lidar2DScan) GetScan() []float64 {
	if x != nil {
		return x.Scan
	}
	return nil
}

func (x *Lidar2DScan) GetNScans() uint32 {
	if x != nil {
		return x.NScans
	}
	return 0
}

func (x *Lidar2DScan) GetMinRange() float64 {
	if x != nil {
		return x.MinRange
	}
	return 0
}

func (x *Lidar2DScan) GetMaxRange() float64 {
	if x != nil {
		return x.MaxRange
	}
	return 0
}

func (x *Lidar2DScan) GetScanResolution() float64 {
	if x != nil {
		return x.ScanResolution
	}
	return 0
}

type NikolaiScan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Odom  *Odom2DUpdate `protobuf:"bytes,1,opt,name=odom,proto3" json:"odom,omitempty"`
	Lidar *Lidar2DScan  `protobuf:"bytes,2,opt,name=lidar,proto3" json:"lidar,omitempty"`
}

func (x *NikolaiScan) Reset() {
	*x = NikolaiScan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sensor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NikolaiScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NikolaiScan) ProtoMessage() {}

func (x *NikolaiScan) ProtoReflect() protoreflect.Message {
	mi := &file_sensor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NikolaiScan.ProtoReflect.Descriptor instead.
func (*NikolaiScan) Descriptor() ([]byte, []int) {
	return file_sensor_proto_rawDescGZIP(), []int{2}
}

func (x *NikolaiScan) GetOdom() *Odom2DUpdate {
	if x != nil {
		return x.Odom
	}
	return nil
}

func (x *NikolaiScan) GetLidar() *Lidar2DScan {
	if x != nil {
		return x.Lidar
	}
	return nil
}

type NikolaiManyScans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scans []*NikolaiScan `protobuf:"bytes,1,rep,name=scans,proto3" json:"scans,omitempty"`
}

func (x *NikolaiManyScans) Reset() {
	*x = NikolaiManyScans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sensor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NikolaiManyScans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NikolaiManyScans) ProtoMessage() {}

func (x *NikolaiManyScans) ProtoReflect() protoreflect.Message {
	mi := &file_sensor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NikolaiManyScans.ProtoReflect.Descriptor instead.
func (*NikolaiManyScans) Descriptor() ([]byte, []int) {
	return file_sensor_proto_rawDescGZIP(), []int{3}
}

func (x *NikolaiManyScans) GetScans() []*NikolaiScan {
	if x != nil {
		return x.Scans
	}
	return nil
}

var File_sensor_proto protoreflect.FileDescriptor

var file_sensor_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x5e, 0x0a, 0x0c, 0x4f, 0x64, 0x6f, 0x6d, 0x32, 0x44, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x68, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x22, 0xbb, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x64, 0x61, 0x72, 0x32, 0x44, 0x53,
	0x63, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x63, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x04, 0x73, 0x63, 0x61, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x53, 0x63, 0x61, 0x6e, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6d, 0x61, 0x78, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x5c, 0x0a, 0x0b, 0x4e, 0x69, 0x6b, 0x6f, 0x6c, 0x61, 0x69, 0x53, 0x63, 0x61, 0x6e,
	0x12, 0x25, 0x0a, 0x04, 0x6f, 0x64, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4f, 0x64, 0x6f, 0x6d, 0x32, 0x44, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x04, 0x6f, 0x64, 0x6f, 0x6d, 0x12, 0x26, 0x0a, 0x05, 0x6c, 0x69, 0x64, 0x61, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4c, 0x69, 0x64,
	0x61, 0x72, 0x32, 0x44, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x05, 0x6c, 0x69, 0x64, 0x61, 0x72, 0x22,
	0x3a, 0x0a, 0x10, 0x4e, 0x69, 0x6b, 0x6f, 0x6c, 0x61, 0x69, 0x4d, 0x61, 0x6e, 0x79, 0x53, 0x63,
	0x61, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4e, 0x69, 0x6b, 0x6f, 0x6c, 0x61, 0x69,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sensor_proto_rawDescOnce sync.Once
	file_sensor_proto_rawDescData = file_sensor_proto_rawDesc
)

func file_sensor_proto_rawDescGZIP() []byte {
	file_sensor_proto_rawDescOnce.Do(func() {
		file_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_sensor_proto_rawDescData)
	})
	return file_sensor_proto_rawDescData
}

var file_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sensor_proto_goTypes = []interface{}{
	(*Odom2DUpdate)(nil),     // 0: msg.Odom2DUpdate
	(*Lidar2DScan)(nil),      // 1: msg.Lidar2DScan
	(*NikolaiScan)(nil),      // 2: msg.NikolaiScan
	(*NikolaiManyScans)(nil), // 3: msg.NikolaiManyScans
}
var file_sensor_proto_depIdxs = []int32{
	0, // 0: msg.NikolaiScan.odom:type_name -> msg.Odom2DUpdate
	1, // 1: msg.NikolaiScan.lidar:type_name -> msg.Lidar2DScan
	2, // 2: msg.NikolaiManyScans.scans:type_name -> msg.NikolaiScan
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sensor_proto_init() }
func file_sensor_proto_init() {
	if File_sensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sensor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Odom2DUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sensor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lidar2DScan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sensor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NikolaiScan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sensor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NikolaiManyScans); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sensor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sensor_proto_goTypes,
		DependencyIndexes: file_sensor_proto_depIdxs,
		MessageInfos:      file_sensor_proto_msgTypes,
	}.Build()
	File_sensor_proto = out.File
	file_sensor_proto_rawDesc = nil
	file_sensor_proto_goTypes = nil
	file_sensor_proto_depIdxs = nil
}
