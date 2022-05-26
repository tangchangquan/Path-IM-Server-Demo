///
//  Generated code. Do not modify.
//  source: param.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use reqDescriptor instead')
const Req$json = const {
  '1': 'Req',
  '2': const [
    const {'1': 'ReqIdentifier', '3': 1, '4': 1, '5': 13, '10': 'ReqIdentifier'},
    const {'1': 'Token', '3': 2, '4': 1, '5': 9, '10': 'Token'},
    const {'1': 'SendID', '3': 3, '4': 1, '5': 9, '10': 'SendID'},
    const {'1': 'MsgIncr', '3': 4, '4': 1, '5': 9, '10': 'MsgIncr'},
    const {'1': 'Data', '3': 5, '4': 1, '5': 12, '10': 'Data'},
  ],
};

/// Descriptor for `Req`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List reqDescriptor = $convert.base64Decode('CgNSZXESJAoNUmVxSWRlbnRpZmllchgBIAEoDVINUmVxSWRlbnRpZmllchIUCgVUb2tlbhgCIAEoCVIFVG9rZW4SFgoGU2VuZElEGAMgASgJUgZTZW5kSUQSGAoHTXNnSW5jchgEIAEoCVIHTXNnSW5jchISCgREYXRhGAUgASgMUgREYXRh');
@$core.Deprecated('Use respDescriptor instead')
const Resp$json = const {
  '1': 'Resp',
  '2': const [
    const {'1': 'ReqIdentifier', '3': 1, '4': 1, '5': 13, '10': 'ReqIdentifier'},
    const {'1': 'MsgIncr', '3': 2, '4': 1, '5': 9, '10': 'MsgIncr'},
    const {'1': 'ErrCode', '3': 3, '4': 1, '5': 13, '10': 'ErrCode'},
    const {'1': 'ErrMsg', '3': 4, '4': 1, '5': 9, '10': 'ErrMsg'},
    const {'1': 'Data', '3': 5, '4': 1, '5': 12, '10': 'Data'},
  ],
};

/// Descriptor for `Resp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List respDescriptor = $convert.base64Decode('CgRSZXNwEiQKDVJlcUlkZW50aWZpZXIYASABKA1SDVJlcUlkZW50aWZpZXISGAoHTXNnSW5jchgCIAEoCVIHTXNnSW5jchIYCgdFcnJDb2RlGAMgASgNUgdFcnJDb2RlEhYKBkVyck1zZxgEIAEoCVIGRXJyTXNnEhIKBERhdGEYBSABKAxSBERhdGE=');
