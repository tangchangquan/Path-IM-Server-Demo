///
//  Generated code. Do not modify.
//  source: body.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use bodyReqDescriptor instead')
const BodyReq$json = const {
  '1': 'BodyReq',
  '2': const [
    const {'1': 'reqIdentifier', '3': 1, '4': 1, '5': 13, '10': 'reqIdentifier'},
    const {'1': 'token', '3': 2, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'sendID', '3': 3, '4': 1, '5': 9, '10': 'sendID'},
    const {'1': 'data', '3': 4, '4': 1, '5': 12, '10': 'data'},
  ],
};

/// Descriptor for `BodyReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bodyReqDescriptor = $convert.base64Decode('CgdCb2R5UmVxEiQKDXJlcUlkZW50aWZpZXIYASABKA1SDXJlcUlkZW50aWZpZXISFAoFdG9rZW4YAiABKAlSBXRva2VuEhYKBnNlbmRJRBgDIAEoCVIGc2VuZElEEhIKBGRhdGEYBCABKAxSBGRhdGE=');
@$core.Deprecated('Use bodyRespDescriptor instead')
const BodyResp$json = const {
  '1': 'BodyResp',
  '2': const [
    const {'1': 'reqIdentifier', '3': 1, '4': 1, '5': 13, '10': 'reqIdentifier'},
    const {'1': 'errCode', '3': 2, '4': 1, '5': 13, '10': 'errCode'},
    const {'1': 'errMsg', '3': 3, '4': 1, '5': 9, '10': 'errMsg'},
    const {'1': 'data', '3': 4, '4': 1, '5': 12, '10': 'data'},
  ],
};

/// Descriptor for `BodyResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bodyRespDescriptor = $convert.base64Decode('CghCb2R5UmVzcBIkCg1yZXFJZGVudGlmaWVyGAEgASgNUg1yZXFJZGVudGlmaWVyEhgKB2VyckNvZGUYAiABKA1SB2VyckNvZGUSFgoGZXJyTXNnGAMgASgJUgZlcnJNc2cSEgoEZGF0YRgEIAEoDFIEZGF0YQ==');
