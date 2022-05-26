///
//  Generated code. Do not modify.
//  source: seq.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetMaxAndMinSeqReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinSeqReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'UserID', protoName: 'UserID')
    ..hasRequiredFields = false
  ;

  GetMaxAndMinSeqReq._() : super();
  factory GetMaxAndMinSeqReq({
    $core.String? userID,
  }) {
    final _result = create();
    if (userID != null) {
      _result.userID = userID;
    }
    return _result;
  }
  factory GetMaxAndMinSeqReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinSeqReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSeqReq clone() => GetMaxAndMinSeqReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSeqReq copyWith(void Function(GetMaxAndMinSeqReq) updates) => super.copyWith((message) => updates(message as GetMaxAndMinSeqReq)) as GetMaxAndMinSeqReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSeqReq create() => GetMaxAndMinSeqReq._();
  GetMaxAndMinSeqReq createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinSeqReq> createRepeated() => $pb.PbList<GetMaxAndMinSeqReq>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSeqReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinSeqReq>(create);
  static GetMaxAndMinSeqReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userID => $_getSZ(0);
  @$pb.TagNumber(1)
  set userID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserID() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserID() => clearField(1);
}

class GetMaxAndMinSeqResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinSeqResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ErrCode', $pb.PbFieldType.O3, protoName: 'ErrCode')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ErrMsg', protoName: 'ErrMsg')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'MaxSeq', $pb.PbFieldType.OU3, protoName: 'MaxSeq')
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'MinSeq', $pb.PbFieldType.OU3, protoName: 'MinSeq')
    ..hasRequiredFields = false
  ;

  GetMaxAndMinSeqResp._() : super();
  factory GetMaxAndMinSeqResp({
    $core.int? errCode,
    $core.String? errMsg,
    $core.int? maxSeq,
    $core.int? minSeq,
  }) {
    final _result = create();
    if (errCode != null) {
      _result.errCode = errCode;
    }
    if (errMsg != null) {
      _result.errMsg = errMsg;
    }
    if (maxSeq != null) {
      _result.maxSeq = maxSeq;
    }
    if (minSeq != null) {
      _result.minSeq = minSeq;
    }
    return _result;
  }
  factory GetMaxAndMinSeqResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinSeqResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSeqResp clone() => GetMaxAndMinSeqResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSeqResp copyWith(void Function(GetMaxAndMinSeqResp) updates) => super.copyWith((message) => updates(message as GetMaxAndMinSeqResp)) as GetMaxAndMinSeqResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSeqResp create() => GetMaxAndMinSeqResp._();
  GetMaxAndMinSeqResp createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinSeqResp> createRepeated() => $pb.PbList<GetMaxAndMinSeqResp>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSeqResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinSeqResp>(create);
  static GetMaxAndMinSeqResp? _defaultInstance;

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
  $core.int get maxSeq => $_getIZ(2);
  @$pb.TagNumber(3)
  set maxSeq($core.int v) { $_setUnsignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMaxSeq() => $_has(2);
  @$pb.TagNumber(3)
  void clearMaxSeq() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get minSeq => $_getIZ(3);
  @$pb.TagNumber(4)
  set minSeq($core.int v) { $_setUnsignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasMinSeq() => $_has(3);
  @$pb.TagNumber(4)
  void clearMinSeq() => clearField(4);
}

class GetMaxAndMinSuperGroupSeqReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinSuperGroupSeqReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..pPS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SuperGroupIDList', protoName: 'SuperGroupIDList')
    ..hasRequiredFields = false
  ;

  GetMaxAndMinSuperGroupSeqReq._() : super();
  factory GetMaxAndMinSuperGroupSeqReq({
    $core.Iterable<$core.String>? superGroupIDList,
  }) {
    final _result = create();
    if (superGroupIDList != null) {
      _result.superGroupIDList.addAll(superGroupIDList);
    }
    return _result;
  }
  factory GetMaxAndMinSuperGroupSeqReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinSuperGroupSeqReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSuperGroupSeqReq clone() => GetMaxAndMinSuperGroupSeqReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSuperGroupSeqReq copyWith(void Function(GetMaxAndMinSuperGroupSeqReq) updates) => super.copyWith((message) => updates(message as GetMaxAndMinSuperGroupSeqReq)) as GetMaxAndMinSuperGroupSeqReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSuperGroupSeqReq create() => GetMaxAndMinSuperGroupSeqReq._();
  GetMaxAndMinSuperGroupSeqReq createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinSuperGroupSeqReq> createRepeated() => $pb.PbList<GetMaxAndMinSuperGroupSeqReq>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSuperGroupSeqReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinSuperGroupSeqReq>(create);
  static GetMaxAndMinSuperGroupSeqReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get superGroupIDList => $_getList(0);
}

class GetMaxAndMinSuperGroupSeqRespItem extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinSuperGroupSeqRespItem', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SuperGroupID', protoName: 'SuperGroupID')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'MaxSeq', $pb.PbFieldType.OU3, protoName: 'MaxSeq')
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'MinSeq', $pb.PbFieldType.OU3, protoName: 'MinSeq')
    ..hasRequiredFields = false
  ;

  GetMaxAndMinSuperGroupSeqRespItem._() : super();
  factory GetMaxAndMinSuperGroupSeqRespItem({
    $core.String? superGroupID,
    $core.int? maxSeq,
    $core.int? minSeq,
  }) {
    final _result = create();
    if (superGroupID != null) {
      _result.superGroupID = superGroupID;
    }
    if (maxSeq != null) {
      _result.maxSeq = maxSeq;
    }
    if (minSeq != null) {
      _result.minSeq = minSeq;
    }
    return _result;
  }
  factory GetMaxAndMinSuperGroupSeqRespItem.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinSuperGroupSeqRespItem.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSuperGroupSeqRespItem clone() => GetMaxAndMinSuperGroupSeqRespItem()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSuperGroupSeqRespItem copyWith(void Function(GetMaxAndMinSuperGroupSeqRespItem) updates) => super.copyWith((message) => updates(message as GetMaxAndMinSuperGroupSeqRespItem)) as GetMaxAndMinSuperGroupSeqRespItem; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSuperGroupSeqRespItem create() => GetMaxAndMinSuperGroupSeqRespItem._();
  GetMaxAndMinSuperGroupSeqRespItem createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinSuperGroupSeqRespItem> createRepeated() => $pb.PbList<GetMaxAndMinSuperGroupSeqRespItem>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSuperGroupSeqRespItem getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinSuperGroupSeqRespItem>(create);
  static GetMaxAndMinSuperGroupSeqRespItem? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get superGroupID => $_getSZ(0);
  @$pb.TagNumber(1)
  set superGroupID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSuperGroupID() => $_has(0);
  @$pb.TagNumber(1)
  void clearSuperGroupID() => clearField(1);

  @$pb.TagNumber(3)
  $core.int get maxSeq => $_getIZ(1);
  @$pb.TagNumber(3)
  set maxSeq($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasMaxSeq() => $_has(1);
  @$pb.TagNumber(3)
  void clearMaxSeq() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get minSeq => $_getIZ(2);
  @$pb.TagNumber(4)
  set minSeq($core.int v) { $_setUnsignedInt32(2, v); }
  @$pb.TagNumber(4)
  $core.bool hasMinSeq() => $_has(2);
  @$pb.TagNumber(4)
  void clearMinSeq() => clearField(4);
}

class GetMaxAndMinSuperGroupSeqResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinSuperGroupSeqResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ErrCode', $pb.PbFieldType.O3, protoName: 'ErrCode')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ErrMsg', protoName: 'ErrMsg')
    ..pc<GetMaxAndMinSuperGroupSeqRespItem>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SuperGroupSeqList', $pb.PbFieldType.PM, protoName: 'SuperGroupSeqList', subBuilder: GetMaxAndMinSuperGroupSeqRespItem.create)
    ..hasRequiredFields = false
  ;

  GetMaxAndMinSuperGroupSeqResp._() : super();
  factory GetMaxAndMinSuperGroupSeqResp({
    $core.int? errCode,
    $core.String? errMsg,
    $core.Iterable<GetMaxAndMinSuperGroupSeqRespItem>? superGroupSeqList,
  }) {
    final _result = create();
    if (errCode != null) {
      _result.errCode = errCode;
    }
    if (errMsg != null) {
      _result.errMsg = errMsg;
    }
    if (superGroupSeqList != null) {
      _result.superGroupSeqList.addAll(superGroupSeqList);
    }
    return _result;
  }
  factory GetMaxAndMinSuperGroupSeqResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinSuperGroupSeqResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSuperGroupSeqResp clone() => GetMaxAndMinSuperGroupSeqResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinSuperGroupSeqResp copyWith(void Function(GetMaxAndMinSuperGroupSeqResp) updates) => super.copyWith((message) => updates(message as GetMaxAndMinSuperGroupSeqResp)) as GetMaxAndMinSuperGroupSeqResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSuperGroupSeqResp create() => GetMaxAndMinSuperGroupSeqResp._();
  GetMaxAndMinSuperGroupSeqResp createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinSuperGroupSeqResp> createRepeated() => $pb.PbList<GetMaxAndMinSuperGroupSeqResp>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinSuperGroupSeqResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinSuperGroupSeqResp>(create);
  static GetMaxAndMinSuperGroupSeqResp? _defaultInstance;

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
  $core.List<GetMaxAndMinSuperGroupSeqRespItem> get superGroupSeqList => $_getList(2);
}

