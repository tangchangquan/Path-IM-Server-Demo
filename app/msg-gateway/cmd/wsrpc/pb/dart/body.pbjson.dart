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
    const {'1': 'ReqIdentifier', '3': 1, '4': 1, '5': 13, '10': 'ReqIdentifier'},
    const {'1': 'Token', '3': 2, '4': 1, '5': 9, '10': 'Token'},
    const {'1': 'SendID', '3': 3, '4': 1, '5': 9, '10': 'SendID'},
    const {'1': 'MsgIncr', '3': 4, '4': 1, '5': 9, '10': 'MsgIncr'},
    const {'1': 'Data', '3': 5, '4': 1, '5': 12, '10': 'Data'},
  ],
};

/// Descriptor for `BodyReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bodyReqDescriptor = $convert.base64Decode('CgdCb2R5UmVxEiQKDVJlcUlkZW50aWZpZXIYASABKA1SDVJlcUlkZW50aWZpZXISFAoFVG9rZW4YAiABKAlSBVRva2VuEhYKBlNlbmRJRBgDIAEoCVIGU2VuZElEEhgKB01zZ0luY3IYBCABKAlSB01zZ0luY3ISEgoERGF0YRgFIAEoDFIERGF0YQ==');
@$core.Deprecated('Use bodyRespDescriptor instead')
const BodyResp$json = const {
  '1': 'BodyResp',
  '2': const [
    const {'1': 'ReqIdentifier', '3': 1, '4': 1, '5': 13, '10': 'ReqIdentifier'},
    const {'1': 'MsgIncr', '3': 2, '4': 1, '5': 9, '10': 'MsgIncr'},
    const {'1': 'ErrCode', '3': 3, '4': 1, '5': 13, '10': 'ErrCode'},
    const {'1': 'ErrMsg', '3': 4, '4': 1, '5': 9, '10': 'ErrMsg'},
    const {'1': 'Data', '3': 5, '4': 1, '5': 12, '10': 'Data'},
  ],
};

/// Descriptor for `BodyResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List bodyRespDescriptor = $convert.base64Decode('CghCb2R5UmVzcBIkCg1SZXFJZGVudGlmaWVyGAEgASgNUg1SZXFJZGVudGlmaWVyEhgKB01zZ0luY3IYAiABKAlSB01zZ0luY3ISGAoHRXJyQ29kZRgDIAEoDVIHRXJyQ29kZRIWCgZFcnJNc2cYBCABKAlSBkVyck1zZxISCgREYXRhGAUgASgMUgREYXRh');
