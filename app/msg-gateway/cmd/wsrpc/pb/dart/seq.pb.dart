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
    ..hasRequiredFields = false
  ;

  GetMaxAndMinSeqReq._() : super();
  factory GetMaxAndMinSeqReq() => create();
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
}

class GetMaxAndMinSeqResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinSeqResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'maxSeq', $pb.PbFieldType.OU3, protoName: 'maxSeq')
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'minSeq', $pb.PbFieldType.OU3, protoName: 'minSeq')
    ..hasRequiredFields = false
  ;

  GetMaxAndMinSeqResp._() : super();
  factory GetMaxAndMinSeqResp({
    $core.int? maxSeq,
    $core.int? minSeq,
  }) {
    final _result = create();
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
  $core.int get maxSeq => $_getIZ(0);
  @$pb.TagNumber(1)
  set maxSeq($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMaxSeq() => $_has(0);
  @$pb.TagNumber(1)
  void clearMaxSeq() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get minSeq => $_getIZ(1);
  @$pb.TagNumber(2)
  set minSeq($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMinSeq() => $_has(1);
  @$pb.TagNumber(2)
  void clearMinSeq() => clearField(2);
}

class GetMaxAndMinGroupSeqReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinGroupSeqReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..pPS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'groupIDList', protoName: 'groupIDList')
    ..hasRequiredFields = false
  ;

  GetMaxAndMinGroupSeqReq._() : super();
  factory GetMaxAndMinGroupSeqReq({
    $core.Iterable<$core.String>? groupIDList,
  }) {
    final _result = create();
    if (groupIDList != null) {
      _result.groupIDList.addAll(groupIDList);
    }
    return _result;
  }
  factory GetMaxAndMinGroupSeqReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinGroupSeqReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinGroupSeqReq clone() => GetMaxAndMinGroupSeqReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinGroupSeqReq copyWith(void Function(GetMaxAndMinGroupSeqReq) updates) => super.copyWith((message) => updates(message as GetMaxAndMinGroupSeqReq)) as GetMaxAndMinGroupSeqReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinGroupSeqReq create() => GetMaxAndMinGroupSeqReq._();
  GetMaxAndMinGroupSeqReq createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinGroupSeqReq> createRepeated() => $pb.PbList<GetMaxAndMinGroupSeqReq>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinGroupSeqReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinGroupSeqReq>(create);
  static GetMaxAndMinGroupSeqReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get groupIDList => $_getList(0);
}

class GetMaxAndMinGroupSeqResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinGroupSeqResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..pc<GetMaxAndMinGroupSeqRespItem>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'groupSeqList', $pb.PbFieldType.PM, protoName: 'groupSeqList', subBuilder: GetMaxAndMinGroupSeqRespItem.create)
    ..hasRequiredFields = false
  ;

  GetMaxAndMinGroupSeqResp._() : super();
  factory GetMaxAndMinGroupSeqResp({
    $core.Iterable<GetMaxAndMinGroupSeqRespItem>? groupSeqList,
  }) {
    final _result = create();
    if (groupSeqList != null) {
      _result.groupSeqList.addAll(groupSeqList);
    }
    return _result;
  }
  factory GetMaxAndMinGroupSeqResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinGroupSeqResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinGroupSeqResp clone() => GetMaxAndMinGroupSeqResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinGroupSeqResp copyWith(void Function(GetMaxAndMinGroupSeqResp) updates) => super.copyWith((message) => updates(message as GetMaxAndMinGroupSeqResp)) as GetMaxAndMinGroupSeqResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinGroupSeqResp create() => GetMaxAndMinGroupSeqResp._();
  GetMaxAndMinGroupSeqResp createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinGroupSeqResp> createRepeated() => $pb.PbList<GetMaxAndMinGroupSeqResp>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinGroupSeqResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinGroupSeqResp>(create);
  static GetMaxAndMinGroupSeqResp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<GetMaxAndMinGroupSeqRespItem> get groupSeqList => $_getList(0);
}

class GetMaxAndMinGroupSeqRespItem extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetMaxAndMinGroupSeqRespItem', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'groupID', protoName: 'groupID')
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'maxSeq', $pb.PbFieldType.OU3, protoName: 'maxSeq')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'minSeq', $pb.PbFieldType.OU3, protoName: 'minSeq')
    ..hasRequiredFields = false
  ;

  GetMaxAndMinGroupSeqRespItem._() : super();
  factory GetMaxAndMinGroupSeqRespItem({
    $core.String? groupID,
    $core.int? maxSeq,
    $core.int? minSeq,
  }) {
    final _result = create();
    if (groupID != null) {
      _result.groupID = groupID;
    }
    if (maxSeq != null) {
      _result.maxSeq = maxSeq;
    }
    if (minSeq != null) {
      _result.minSeq = minSeq;
    }
    return _result;
  }
  factory GetMaxAndMinGroupSeqRespItem.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetMaxAndMinGroupSeqRespItem.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetMaxAndMinGroupSeqRespItem clone() => GetMaxAndMinGroupSeqRespItem()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetMaxAndMinGroupSeqRespItem copyWith(void Function(GetMaxAndMinGroupSeqRespItem) updates) => super.copyWith((message) => updates(message as GetMaxAndMinGroupSeqRespItem)) as GetMaxAndMinGroupSeqRespItem; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinGroupSeqRespItem create() => GetMaxAndMinGroupSeqRespItem._();
  GetMaxAndMinGroupSeqRespItem createEmptyInstance() => create();
  static $pb.PbList<GetMaxAndMinGroupSeqRespItem> createRepeated() => $pb.PbList<GetMaxAndMinGroupSeqRespItem>();
  @$core.pragma('dart2js:noInline')
  static GetMaxAndMinGroupSeqRespItem getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetMaxAndMinGroupSeqRespItem>(create);
  static GetMaxAndMinGroupSeqRespItem? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get groupID => $_getSZ(0);
  @$pb.TagNumber(1)
  set groupID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasGroupID() => $_has(0);
  @$pb.TagNumber(1)
  void clearGroupID() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get maxSeq => $_getIZ(1);
  @$pb.TagNumber(2)
  set maxSeq($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMaxSeq() => $_has(1);
  @$pb.TagNumber(2)
  void clearMaxSeq() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get minSeq => $_getIZ(2);
  @$pb.TagNumber(3)
  set minSeq($core.int v) { $_setUnsignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMinSeq() => $_has(2);
  @$pb.TagNumber(3)
  void clearMinSeq() => clearField(3);
}

