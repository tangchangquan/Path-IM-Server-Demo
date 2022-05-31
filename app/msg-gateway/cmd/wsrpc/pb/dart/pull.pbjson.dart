///
//  Generated code. Do not modify.
//  source: pull.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields,deprecated_member_use_from_same_package

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use pullMessageBySeqListReqDescriptor instead')
const PullMessageBySeqListReq$json = const {
  '1': 'PullMessageBySeqListReq',
  '2': const [
    const {'1': 'userID', '3': 1, '4': 1, '5': 9, '10': 'userID'},
    const {'1': 'seqList', '3': 2, '4': 3, '5': 13, '10': 'seqList'},
  ],
};

/// Descriptor for `PullMessageBySeqListReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pullMessageBySeqListReqDescriptor = $convert.base64Decode('ChdQdWxsTWVzc2FnZUJ5U2VxTGlzdFJlcRIWCgZ1c2VySUQYASABKAlSBnVzZXJJRBIYCgdzZXFMaXN0GAIgAygNUgdzZXFMaXN0');
@$core.Deprecated('Use pullMessageByGroupSeqListReqDescriptor instead')
const PullMessageByGroupSeqListReq$json = const {
  '1': 'PullMessageByGroupSeqListReq',
  '2': const [
    const {'1': 'groupID', '3': 1, '4': 1, '5': 9, '10': 'groupID'},
    const {'1': 'seqList', '3': 2, '4': 3, '5': 13, '10': 'seqList'},
  ],
};

/// Descriptor for `PullMessageByGroupSeqListReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pullMessageByGroupSeqListReqDescriptor = $convert.base64Decode('ChxQdWxsTWVzc2FnZUJ5R3JvdXBTZXFMaXN0UmVxEhgKB2dyb3VwSUQYASABKAlSB2dyb3VwSUQSGAoHc2VxTGlzdBgCIAMoDVIHc2VxTGlzdA==');
@$core.Deprecated('Use pullMessageBySeqListRespDescriptor instead')
const PullMessageBySeqListResp$json = const {
  '1': 'PullMessageBySeqListResp',
  '2': const [
    const {'1': 'list', '3': 1, '4': 3, '5': 11, '6': '.pb.MsgData', '10': 'list'},
  ],
};

/// Descriptor for `PullMessageBySeqListResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pullMessageBySeqListRespDescriptor = $convert.base64Decode('ChhQdWxsTWVzc2FnZUJ5U2VxTGlzdFJlc3ASHwoEbGlzdBgBIAMoCzILLnBiLk1zZ0RhdGFSBGxpc3Q=');
