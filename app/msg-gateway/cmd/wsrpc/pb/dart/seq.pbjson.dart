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
};

/// Descriptor for `GetMaxAndMinSeqReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinSeqReqDescriptor = $convert.base64Decode('ChJHZXRNYXhBbmRNaW5TZXFSZXE=');
@$core.Deprecated('Use getMaxAndMinSeqRespDescriptor instead')
const GetMaxAndMinSeqResp$json = const {
  '1': 'GetMaxAndMinSeqResp',
  '2': const [
    const {'1': 'maxSeq', '3': 1, '4': 1, '5': 13, '10': 'maxSeq'},
    const {'1': 'minSeq', '3': 2, '4': 1, '5': 13, '10': 'minSeq'},
  ],
};

/// Descriptor for `GetMaxAndMinSeqResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinSeqRespDescriptor = $convert.base64Decode('ChNHZXRNYXhBbmRNaW5TZXFSZXNwEhYKBm1heFNlcRgBIAEoDVIGbWF4U2VxEhYKBm1pblNlcRgCIAEoDVIGbWluU2Vx');
@$core.Deprecated('Use getMaxAndMinGroupSeqReqDescriptor instead')
const GetMaxAndMinGroupSeqReq$json = const {
  '1': 'GetMaxAndMinGroupSeqReq',
  '2': const [
    const {'1': 'groupIDList', '3': 1, '4': 3, '5': 9, '10': 'groupIDList'},
  ],
};

/// Descriptor for `GetMaxAndMinGroupSeqReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinGroupSeqReqDescriptor = $convert.base64Decode('ChdHZXRNYXhBbmRNaW5Hcm91cFNlcVJlcRIgCgtncm91cElETGlzdBgBIAMoCVILZ3JvdXBJRExpc3Q=');
@$core.Deprecated('Use getMaxAndMinGroupSeqRespDescriptor instead')
const GetMaxAndMinGroupSeqResp$json = const {
  '1': 'GetMaxAndMinGroupSeqResp',
  '2': const [
    const {'1': 'groupSeqList', '3': 1, '4': 3, '5': 11, '6': '.pb.GetMaxAndMinGroupSeqRespItem', '10': 'groupSeqList'},
  ],
};

/// Descriptor for `GetMaxAndMinGroupSeqResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinGroupSeqRespDescriptor = $convert.base64Decode('ChhHZXRNYXhBbmRNaW5Hcm91cFNlcVJlc3ASRAoMZ3JvdXBTZXFMaXN0GAEgAygLMiAucGIuR2V0TWF4QW5kTWluR3JvdXBTZXFSZXNwSXRlbVIMZ3JvdXBTZXFMaXN0');
@$core.Deprecated('Use getMaxAndMinGroupSeqRespItemDescriptor instead')
const GetMaxAndMinGroupSeqRespItem$json = const {
  '1': 'GetMaxAndMinGroupSeqRespItem',
  '2': const [
    const {'1': 'groupID', '3': 1, '4': 1, '5': 9, '10': 'groupID'},
    const {'1': 'maxSeq', '3': 2, '4': 1, '5': 13, '10': 'maxSeq'},
    const {'1': 'minSeq', '3': 3, '4': 1, '5': 13, '10': 'minSeq'},
  ],
};

/// Descriptor for `GetMaxAndMinGroupSeqRespItem`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMaxAndMinGroupSeqRespItemDescriptor = $convert.base64Decode('ChxHZXRNYXhBbmRNaW5Hcm91cFNlcVJlc3BJdGVtEhgKB2dyb3VwSUQYASABKAlSB2dyb3VwSUQSFgoGbWF4U2VxGAIgASgNUgZtYXhTZXESFgoGbWluU2VxGAMgASgNUgZtaW5TZXE=');
