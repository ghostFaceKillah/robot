# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: sensor.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='sensor.proto',
  package='msg',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0csensor.proto\x12\x03msg\"F\n\x0cOdom2DUpdate\x12\x11\n\ttimestamp\x18\x01 \x01(\x03\x12\t\n\x01x\x18\x02 \x01(\x01\x12\t\n\x01y\x18\x03 \x01(\x01\x12\r\n\x05theta\x18\x04 \x01(\x01\"~\n\x0bLidar2DScan\x12\x11\n\ttimestamp\x18\x01 \x01(\x03\x12\x0c\n\x04scan\x18\x02 \x03(\x01\x12\x0f\n\x07n_scans\x18\x03 \x01(\r\x12\x11\n\tmin_range\x18\x04 \x01(\x01\x12\x11\n\tmax_range\x18\x05 \x01(\x01\x12\x17\n\x0fscan_resolution\x18\x06 \x01(\x01\"T\n\x0bNikolaiScan\x12\x1f\n\x04odom\x18\x01 \x01(\x0b\x32\x11.msg.Odom2DUpdate\x12$\n\nlidar_scan\x18\x02 \x01(\x0b\x32\x10.msg.Lidar2DScan\"3\n\x10NikolaiManyScans\x12\x1f\n\x05scans\x18\x01 \x03(\x0b\x32\x10.msg.NikolaiScanb\x06proto3')
)




_ODOM2DUPDATE = _descriptor.Descriptor(
  name='Odom2DUpdate',
  full_name='msg.Odom2DUpdate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='msg.Odom2DUpdate.timestamp', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='x', full_name='msg.Odom2DUpdate.x', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='y', full_name='msg.Odom2DUpdate.y', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='theta', full_name='msg.Odom2DUpdate.theta', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=21,
  serialized_end=91,
)


_LIDAR2DSCAN = _descriptor.Descriptor(
  name='Lidar2DScan',
  full_name='msg.Lidar2DScan',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='msg.Lidar2DScan.timestamp', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scan', full_name='msg.Lidar2DScan.scan', index=1,
      number=2, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='n_scans', full_name='msg.Lidar2DScan.n_scans', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='min_range', full_name='msg.Lidar2DScan.min_range', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_range', full_name='msg.Lidar2DScan.max_range', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scan_resolution', full_name='msg.Lidar2DScan.scan_resolution', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=93,
  serialized_end=219,
)


_NIKOLAISCAN = _descriptor.Descriptor(
  name='NikolaiScan',
  full_name='msg.NikolaiScan',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='odom', full_name='msg.NikolaiScan.odom', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lidar_scan', full_name='msg.NikolaiScan.lidar_scan', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=221,
  serialized_end=305,
)


_NIKOLAIMANYSCANS = _descriptor.Descriptor(
  name='NikolaiManyScans',
  full_name='msg.NikolaiManyScans',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='scans', full_name='msg.NikolaiManyScans.scans', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=307,
  serialized_end=358,
)

_NIKOLAISCAN.fields_by_name['odom'].message_type = _ODOM2DUPDATE
_NIKOLAISCAN.fields_by_name['lidar_scan'].message_type = _LIDAR2DSCAN
_NIKOLAIMANYSCANS.fields_by_name['scans'].message_type = _NIKOLAISCAN
DESCRIPTOR.message_types_by_name['Odom2DUpdate'] = _ODOM2DUPDATE
DESCRIPTOR.message_types_by_name['Lidar2DScan'] = _LIDAR2DSCAN
DESCRIPTOR.message_types_by_name['NikolaiScan'] = _NIKOLAISCAN
DESCRIPTOR.message_types_by_name['NikolaiManyScans'] = _NIKOLAIMANYSCANS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Odom2DUpdate = _reflection.GeneratedProtocolMessageType('Odom2DUpdate', (_message.Message,), dict(
  DESCRIPTOR = _ODOM2DUPDATE,
  __module__ = 'sensor_pb2'
  # @@protoc_insertion_point(class_scope:msg.Odom2DUpdate)
  ))
_sym_db.RegisterMessage(Odom2DUpdate)

Lidar2DScan = _reflection.GeneratedProtocolMessageType('Lidar2DScan', (_message.Message,), dict(
  DESCRIPTOR = _LIDAR2DSCAN,
  __module__ = 'sensor_pb2'
  # @@protoc_insertion_point(class_scope:msg.Lidar2DScan)
  ))
_sym_db.RegisterMessage(Lidar2DScan)

NikolaiScan = _reflection.GeneratedProtocolMessageType('NikolaiScan', (_message.Message,), dict(
  DESCRIPTOR = _NIKOLAISCAN,
  __module__ = 'sensor_pb2'
  # @@protoc_insertion_point(class_scope:msg.NikolaiScan)
  ))
_sym_db.RegisterMessage(NikolaiScan)

NikolaiManyScans = _reflection.GeneratedProtocolMessageType('NikolaiManyScans', (_message.Message,), dict(
  DESCRIPTOR = _NIKOLAIMANYSCANS,
  __module__ = 'sensor_pb2'
  # @@protoc_insertion_point(class_scope:msg.NikolaiManyScans)
  ))
_sym_db.RegisterMessage(NikolaiManyScans)


# @@protoc_insertion_point(module_scope)