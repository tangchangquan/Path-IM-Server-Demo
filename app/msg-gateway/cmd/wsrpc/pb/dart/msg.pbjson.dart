///
//  Generated code. Do not modify.
//  source: msg.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use sendMsgReqDescriptor instead')
const SendMsgReq$json = const {
  '1': 'SendMsgReq',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '10': 'token'},
    const {'1': 'msgData', '3': 2, '4': 1, '5': 11, '6': '.pb.MsgData', '10': 'msgData'},
  ],
};

/// Descriptor for `SendMsgReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sendMsgReqDescriptor = $convert.base64Decode('CgpTZW5kTXNnUmVxEhQKBXRva2VuGAEgASgJUgV0b2tlbhIlCgdtc2dEYXRhGAIgASgLMgsucGIuTXNnRGF0YVIHbXNnRGF0YQ==');
@$core.Deprecated('Use msgDataDescriptor instead')
const MsgData$json = const {
  '1': 'MsgData',
  '2': const [
    const {'1': 'sendID', '3': 1, '4': 1, '5': 9, '10': 'sendID'},
    const {'1': 'recvID', '3': 2, '4': 1, '5': 9, '10': 'recvID'},
    const {'1': 'groupID', '3': 3, '4': 1, '5': 9, '10': 'groupID'},
    const {'1': 'clientMsgID', '3': 4, '4': 1, '5': 9, '10': 'clientMsgID'},
    const {'1': 'serverMsgID', '3': 5, '4': 1, '5': 9, '10': 'serverMsgID'},
    const {'1': 'senderPlatformID', '3': 6, '4': 1, '5': 5, '10': 'senderPlatformID'},
    const {'1': 'senderNickname', '3': 7, '4': 1, '5': 9, '10': 'senderNickname'},
    const {'1': 'senderFaceURL', '3': 8, '4': 1, '5': 9, '10': 'senderFaceURL'},
    const {'1': 'sessionType', '3': 9, '4': 1, '5': 5, '10': 'sessionType'},
    const {'1': 'msgFrom', '3': 10, '4': 1, '5': 5, '10': 'msgFrom'},
    const {'1': 'contentType', '3': 11, '4': 1, '5': 5, '10': 'contentType'},
    const {'1': 'content', '3': 12, '4': 1, '5': 12, '10': 'content'},
    const {'1': 'seq', '3': 14, '4': 1, '5': 13, '10': 'seq'},
    const {'1': 'serverTime', '3': 15, '4': 1, '5': 3, '10': 'serverTime'},
    const {'1': 'clientTime', '3': 16, '4': 1, '5': 3, '10': 'clientTime'},
    const {'1': 'offlinePushInfo', '3': 17, '4': 1, '5': 11, '6': '.pb.OfflinePushInfo', '10': 'offlinePushInfo'},
    const {'1': 'atUserIDList', '3': 18, '4': 3, '5': 9, '10': 'atUserIDList'},
    const {'1': 'options', '3': 19, '4': 3, '5': 11, '6': '.pb.MsgData.OptionsEntry', '10': 'options'},
  ],
  '3': const [MsgData_OptionsEntry$json],
};

@$core.Deprecated('Use msgDataDescriptor instead')
const MsgData_OptionsEntry$json = const {
  '1': 'OptionsEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 8, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `MsgData`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List msgDataDescriptor = $convert.base64Decode('CgdNc2dEYXRhEhYKBnNlbmRJRBgBIAEoCVIGc2VuZElEEhYKBnJlY3ZJRBgCIAEoCVIGcmVjdklEEhgKB2dyb3VwSUQYAyABKAlSB2dyb3VwSUQSIAoLY2xpZW50TXNnSUQYBCABKAlSC2NsaWVudE1zZ0lEEiAKC3NlcnZlck1zZ0lEGAUgASgJUgtzZXJ2ZXJNc2dJRBIqChBzZW5kZXJQbGF0Zm9ybUlEGAYgASgFUhBzZW5kZXJQbGF0Zm9ybUlEEiYKDnNlbmRlck5pY2tuYW1lGAcgASgJUg5zZW5kZXJOaWNrbmFtZRIkCg1zZW5kZXJGYWNlVVJMGAggASgJUg1zZW5kZXJGYWNlVVJMEiAKC3Nlc3Npb25UeXBlGAkgASgFUgtzZXNzaW9uVHlwZRIYCgdtc2dGcm9tGAogASgFUgdtc2dGcm9tEiAKC2NvbnRlbnRUeXBlGAsgASgFUgtjb250ZW50VHlwZRIYCgdjb250ZW50GAwgASgMUgdjb250ZW50EhAKA3NlcRgOIAEoDVIDc2VxEh4KCnNlcnZlclRpbWUYDyABKANSCnNlcnZlclRpbWUSHgoKY2xpZW50VGltZRgQIAEoA1IKY2xpZW50VGltZRI9Cg9vZmZsaW5lUHVzaEluZm8YESABKAsyEy5wYi5PZmZsaW5lUHVzaEluZm9SD29mZmxpbmVQdXNoSW5mbxIiCgxhdFVzZXJJRExpc3QYEiADKAlSDGF0VXNlcklETGlzdBIyCgdvcHRpb25zGBMgAygLMhgucGIuTXNnRGF0YS5PcHRpb25zRW50cnlSB29wdGlvbnMaOgoMT3B0aW9uc0VudHJ5EhAKA2tleRgBIAEoCVIDa2V5EhQKBXZhbHVlGAIgASgIUgV2YWx1ZToCOAE=');
@$core.Deprecated('Use offlinePushInfoDescriptor instead')
const OfflinePushInfo$json = const {
  '1': 'OfflinePushInfo',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'desc', '3': 2, '4': 1, '5': 9, '10': 'desc'},
    const {'1': 'ex', '3': 3, '4': 1, '5': 9, '10': 'ex'},
    const {'1': 'iOSPushSound', '3': 4, '4': 1, '5': 9, '10': 'iOSPushSound'},
    const {'1': 'iOSBadgeCount', '3': 5, '4': 1, '5': 8, '10': 'iOSBadgeCount'},
  ],
};

/// Descriptor for `OfflinePushInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List offlinePushInfoDescriptor = $convert.base64Decode('Cg9PZmZsaW5lUHVzaEluZm8SFAoFdGl0bGUYASABKAlSBXRpdGxlEhIKBGRlc2MYAiABKAlSBGRlc2MSDgoCZXgYAyABKAlSAmV4EiIKDGlPU1B1c2hTb3VuZBgEIAEoCVIMaU9TUHVzaFNvdW5kEiQKDWlPU0JhZGdlQ291bnQYBSABKAhSDWlPU0JhZGdlQ291bnQ=');
@$core.Deprecated('Use sendMsgRespDescriptor instead')
const SendMsgResp$json = const {
  '1': 'SendMsgResp',
  '2': const [
    const {'1': 'errCode', '3': 1, '4': 1, '5': 5, '10': 'errCode'},
    const {'1': 'errMsg', '3': 2, '4': 1, '5': 9, '10': 'errMsg'},
    const {'1': 'serverMsgID', '3': 3, '4': 1, '5': 9, '10': 'serverMsgID'},
    const {'1': 'clientMsgID', '3': 4, '4': 1, '5': 9, '10': 'clientMsgID'},
    const {'1': 'serverTime', '3': 5, '4': 1, '5': 3, '10': 'serverTime'},
  ],
};

/// Descriptor for `SendMsgResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sendMsgRespDescriptor = $convert.base64Decode('CgtTZW5kTXNnUmVzcBIYCgdlcnJDb2RlGAEgASgFUgdlcnJDb2RlEhYKBmVyck1zZxgCIAEoCVIGZXJyTXNnEiAKC3NlcnZlck1zZ0lEGAMgASgJUgtzZXJ2ZXJNc2dJRBIgCgtjbGllbnRNc2dJRBgEIAEoCVILY2xpZW50TXNnSUQSHgoKc2VydmVyVGltZRgFIAEoA1IKc2VydmVyVGltZQ==');
