///
//  Generated code. Do not modify.
//  source: seq.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use getMaxAndMinSeqReqDescriptor instead')
const GetMaxAndMinSeqReq$json = const {
  '1': 'GetMaxAndMinSeqReq',
  '2': const [
    const {'1': 'UserID', '3': 1, '4': 1, '5': 9, '10': 'UserID'},
  ],
};

/// Descriptor for `GetMaxAndMinSeqReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinSeqReqDescriptor = $convert.base64Decode('ChJHZXRNYXhBbmRNaW5TZXFSZXESFgoGVXNlcklEGAEgASgJUgZVc2VySUQ=');
@$core.Deprecated('Use getMaxAndMinSeqRespDescriptor instead')
const GetMaxAndMinSeqResp$json = const {
  '1': 'GetMaxAndMinSeqResp',
  '2': const [
    const {'1': 'ErrCode', '3': 1, '4': 1, '5': 5, '10': 'ErrCode'},
    const {'1': 'ErrMsg', '3': 2, '4': 1, '5': 9, '10': 'ErrMsg'},
    const {'1': 'MaxSeq', '3': 3, '4': 1, '5': 13, '10': 'MaxSeq'},
    const {'1': 'MinSeq', '3': 4, '4': 1, '5': 13, '10': 'MinSeq'},
  ],
};

/// Descriptor for `GetMaxAndMinSeqResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinSeqRespDescriptor = $convert.base64Decode('ChNHZXRNYXhBbmRNaW5TZXFSZXNwEhgKB0VyckNvZGUYASABKAVSB0VyckNvZGUSFgoGRXJyTXNnGAIgASgJUgZFcnJNc2cSFgoGTWF4U2VxGAMgASgNUgZNYXhTZXESFgoGTWluU2VxGAQgASgNUgZNaW5TZXE=');
@$core.Deprecated('Use getMaxAndMinGroupSeqReqDescriptor instead')
const GetMaxAndMinGroupSeqReq$json = const {
  '1': 'GetMaxAndMinGroupSeqReq',
  '2': const [
    const {'1': 'GroupIDList', '3': 1, '4': 3, '5': 9, '10': 'GroupIDList'},
  ],
};

/// Descriptor for `GetMaxAndMinGroupSeqReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinGroupSeqReqDescriptor = $convert.base64Decode('ChdHZXRNYXhBbmRNaW5Hcm91cFNlcVJlcRIgCgtHcm91cElETGlzdBgBIAMoCVILR3JvdXBJRExpc3Q=');
@$core.Deprecated('Use getMaxAndMinGroupSeqRespDescriptor instead')
const GetMaxAndMinGroupSeqResp$json = const {
  '1': 'GetMaxAndMinGroupSeqResp',
  '2': const [
    const {'1': 'ErrCode', '3': 1, '4': 1, '5': 5, '10': 'ErrCode'},
    const {'1': 'ErrMsg', '3': 2, '4': 1, '5': 9, '10': 'ErrMsg'},
    const {'1': 'GroupSeqList', '3': 3, '4': 3, '5': 11, '6': '.pb.GetMaxAndMinGroupSeqRespItem', '10': 'GroupSeqList'},
  ],
};

/// Descriptor for `GetMaxAndMinGroupSeqResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinGroupSeqRespDescriptor = $convert.base64Decode('ChhHZXRNYXhBbmRNaW5Hcm91cFNlcVJlc3ASGAoHRXJyQ29kZRgBIAEoBVIHRXJyQ29kZRIWCgZFcnJNc2cYAiABKAlSBkVyck1zZxJECgxHcm91cFNlcUxpc3QYAyADKAsyIC5wYi5HZXRNYXhBbmRNaW5Hcm91cFNlcVJlc3BJdGVtUgxHcm91cFNlcUxpc3Q=');
@$core.Deprecated('Use getMaxAndMinGroupSeqRespItemDescriptor instead')
const GetMaxAndMinGroupSeqRespItem$json = const {
  '1': 'GetMaxAndMinGroupSeqRespItem',
  '2': const [
    const {'1': 'GroupID', '3': 1, '4': 1, '5': 9, '10': 'GroupID'},
    const {'1': 'MaxSeq', '3': 2, '4': 1, '5': 13, '10': 'MaxSeq'},
    const {'1': 'MinSeq', '3': 3, '4': 1, '5': 13, '10': 'MinSeq'},
  ],
};

/// Descriptor for `GetMaxAndMinGroupSeqRespItem`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinGroupSeqRespItemDescriptor = $convert.base64Decode('ChxHZXRNYXhBbmRNaW5Hcm91cFNlcVJlc3BJdGVtEhgKB0dyb3VwSUQYASABKAlSB0dyb3VwSUQSFgoGTWF4U2VxGAIgASgNUgZNYXhTZXESFgoGTWluU2VxGAMgASgNUgZNaW5TZXE=');
