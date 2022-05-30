///
//  Generated code. Do not modify.
//  source: pull.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'msg.pb.dart' as $0;

class PullMessageBySeqListReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PullMessageBySeqListReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'userID', protoName: 'userID')
    ..p<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'seqList', $pb.PbFieldType.PU3, protoName: 'seqList')
    ..hasRequiredFields = false
  ;

  PullMessageBySeqListReq._() : super();
  factory PullMessageBySeqListReq({
    $core.String? userID,
    $core.Iterable<$core.int>? seqList,
  }) {
    final _result = create();
    if (userID != null) {
      _result.userID = userID;
    }
    if (seqList != null) {
      _result.seqList.addAll(seqList);
    }
    return _result;
  }
  factory PullMessageBySeqListReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PullMessageBySeqListReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PullMessageBySeqListReq clone() => PullMessageBySeqListReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PullMessageBySeqListReq copyWith(void Function(PullMessageBySeqListReq) updates) => super.copyWith((message) => updates(message as PullMessageBySeqListReq)) as PullMessageBySeqListReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PullMessageBySeqListReq create() => PullMessageBySeqListReq._();
  PullMessageBySeqListReq createEmptyInstance() => create();
  static $pb.PbList<PullMessageBySeqListReq> createRepeated() => $pb.PbList<PullMessageBySeqListReq>();
  @$core.pragma('dart2js:noInline')
  static PullMessageBySeqListReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PullMessageBySeqListReq>(create);
  static PullMessageBySeqListReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userID => $_getSZ(0);
  @$pb.TagNumber(1)
  set userID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserID() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserID() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get seqList => $_getList(1);
}

class PullMessageByGroupSeqListReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PullMessageByGroupSeqListReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'groupID', protoName: 'groupID')
    ..p<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'seqList', $pb.PbFieldType.PU3, protoName: 'seqList')
    ..hasRequiredFields = false
  ;

  PullMessageByGroupSeqListReq._() : super();
  factory PullMessageByGroupSeqListReq({
    $core.String? groupID,
    $core.Iterable<$core.int>? seqList,
  }) {
    final _result = create();
    if (groupID != null) {
      _result.groupID = groupID;
    }
    if (seqList != null) {
      _result.seqList.addAll(seqList);
    }
    return _result;
  }
  factory PullMessageByGroupSeqListReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PullMessageByGroupSeqListReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PullMessageByGroupSeqListReq clone() => PullMessageByGroupSeqListReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PullMessageByGroupSeqListReq copyWith(void Function(PullMessageByGroupSeqListReq) updates) => super.copyWith((message) => updates(message as PullMessageByGroupSeqListReq)) as PullMessageByGroupSeqListReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PullMessageByGroupSeqListReq create() => PullMessageByGroupSeqListReq._();
  PullMessageByGroupSeqListReq createEmptyInstance() => create();
  static $pb.PbList<PullMessageByGroupSeqListReq> createRepeated() => $pb.PbList<PullMessageByGroupSeqListReq>();
  @$core.pragma('dart2js:noInline')
  static PullMessageByGroupSeqListReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PullMessageByGroupSeqListReq>(create);
  static PullMessageByGroupSeqListReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get groupID => $_getSZ(0);
  @$pb.TagNumber(1)
  set groupID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasGroupID() => $_has(0);
  @$pb.TagNumber(1)
  void clearGroupID() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get seqList => $_getList(1);
}

class PullMessageBySeqListResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PullMessageBySeqListResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errCode', $pb.PbFieldType.O3, protoName: 'errCode')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errMsg', protoName: 'errMsg')
    ..pc<$0.MsgData>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'list', $pb.PbFieldType.PM, subBuilder: $0.MsgData.create)
    ..hasRequiredFields = false
  ;

  PullMessageBySeqListResp._() : super();
  factory PullMessageBySeqListResp({
    $core.int? errCode,
    $core.String? errMsg,
    $core.Iterable<$0.MsgData>? list,
  }) {
    final _result = create();
    if (errCode != null) {
      _result.errCode = errCode;
    }
    if (errMsg != null) {
      _result.errMsg = errMsg;
    }
    if (list != null) {
      _result.list.addAll(list);
    }
    return _result;
  }
  factory PullMessageBySeqListResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PullMessageBySeqListResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PullMessageBySeqListResp clone() => PullMessageBySeqListResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PullMessageBySeqListResp copyWith(void Function(PullMessageBySeqListResp) updates) => super.copyWith((message) => updates(message as PullMessageBySeqListResp)) as PullMessageBySeqListResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PullMessageBySeqListResp create() => PullMessageBySeqListResp._();
  PullMessageBySeqListResp createEmptyInstance() => create();
  static $pb.PbList<PullMessageBySeqListResp> createRepeated() => $pb.PbList<PullMessageBySeqListResp>();
  @$core.pragma('dart2js:noInline')
  static PullMessageBySeqListResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PullMessageBySeqListResp>(create);
  static PullMessageBySeqListResp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get errCode => $_getIZ(0);
  @$pb.TagNumber(1)
  set errCode($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasErrCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearErrCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get errMsg => $_getSZ(1);
  @$pb.TagNumber(2)
  set errMsg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrMsg() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$0.MsgData> get list => $_getList(2);
}

