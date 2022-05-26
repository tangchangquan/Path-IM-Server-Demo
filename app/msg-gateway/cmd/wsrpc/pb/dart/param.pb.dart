///
//  Generated code. Do not modify.
//  source: param.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Req extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Req', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ReqIdentifier', $pb.PbFieldType.OU3, protoName: 'ReqIdentifier')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Token', protoName: 'Token')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SendID', protoName: 'SendID')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'MsgIncr', protoName: 'MsgIncr')
    ..a<$core.List<$core.int>>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Data', $pb.PbFieldType.OY, protoName: 'Data')
    ..hasRequiredFields = false
  ;

  Req._() : super();
  factory Req({
    $core.int? reqIdentifier,
    $core.String? token,
    $core.String? sendID,
    $core.String? msgIncr,
    $core.List<$core.int>? data,
  }) {
    final _result = create();
    if (reqIdentifier != null) {
      _result.reqIdentifier = reqIdentifier;
    }
    if (token != null) {
      _result.token = token;
    }
    if (sendID != null) {
      _result.sendID = sendID;
    }
    if (msgIncr != null) {
      _result.msgIncr = msgIncr;
    }
    if (data != null) {
      _result.data = data;
    }
    return _result;
  }
  factory Req.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Req.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Req clone() => Req()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Req copyWith(void Function(Req) updates) => super.copyWith((message) => updates(message as Req)) as Req; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Req create() => Req._();
  Req createEmptyInstance() => create();
  static $pb.PbList<Req> createRepeated() => $pb.PbList<Req>();
  @$core.pragma('dart2js:noInline')
  static Req getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Req>(create);
  static Req? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get reqIdentifier => $_getIZ(0);
  @$pb.TagNumber(1)
  set reqIdentifier($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasReqIdentifier() => $_has(0);
  @$pb.TagNumber(1)
  void clearReqIdentifier() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get token => $_getSZ(1);
  @$pb.TagNumber(2)
  set token($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasToken() => $_has(1);
  @$pb.TagNumber(2)
  void clearToken() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get sendID => $_getSZ(2);
  @$pb.TagNumber(3)
  set sendID($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasSendID() => $_has(2);
  @$pb.TagNumber(3)
  void clearSendID() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get msgIncr => $_getSZ(3);
  @$pb.TagNumber(4)
  set msgIncr($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasMsgIncr() => $_has(3);
  @$pb.TagNumber(4)
  void clearMsgIncr() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get data => $_getN(4);
  @$pb.TagNumber(5)
  set data($core.List<$core.int> v) { $_setBytes(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasData() => $_has(4);
  @$pb.TagNumber(5)
  void clearData() => clearField(5);
}

class Resp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Resp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ReqIdentifier', $pb.PbFieldType.OU3, protoName: 'ReqIdentifier')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'MsgIncr', protoName: 'MsgIncr')
    ..a<$core.int>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ErrCode', $pb.PbFieldType.OU3, protoName: 'ErrCode')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ErrMsg', protoName: 'ErrMsg')
    ..a<$core.List<$core.int>>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Data', $pb.PbFieldType.OY, protoName: 'Data')
    ..hasRequiredFields = false
  ;

  Resp._() : super();
  factory Resp({
    $core.int? reqIdentifier,
    $core.String? msgIncr,
    $core.int? errCode,
    $core.String? errMsg,
    $core.List<$core.int>? data,
  }) {
    final _result = create();
    if (reqIdentifier != null) {
      _result.reqIdentifier = reqIdentifier;
    }
    if (msgIncr != null) {
      _result.msgIncr = msgIncr;
    }
    if (errCode != null) {
      _result.errCode = errCode;
    }
    if (errMsg != null) {
      _result.errMsg = errMsg;
    }
    if (data != null) {
      _result.data = data;
    }
    return _result;
  }
  factory Resp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Resp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Resp clone() => Resp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Resp copyWith(void Function(Resp) updates) => super.copyWith((message) => updates(message as Resp)) as Resp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Resp create() => Resp._();
  Resp createEmptyInstance() => create();
  static $pb.PbList<Resp> createRepeated() => $pb.PbList<Resp>();
  @$core.pragma('dart2js:noInline')
  static Resp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Resp>(create);
  static Resp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get reqIdentifier => $_getIZ(0);
  @$pb.TagNumber(1)
  set reqIdentifier($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasReqIdentifier() => $_has(0);
  @$pb.TagNumber(1)
  void clearReqIdentifier() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msgIncr => $_getSZ(1);
  @$pb.TagNumber(2)
  set msgIncr($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsgIncr() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsgIncr() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get errCode => $_getIZ(2);
  @$pb.TagNumber(3)
  set errCode($core.int v) { $_setUnsignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrCode() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrCode() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get errMsg => $_getSZ(3);
  @$pb.TagNumber(4)
  set errMsg($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasErrMsg() => $_has(3);
  @$pb.TagNumber(4)
  void clearErrMsg() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get data => $_getN(4);
  @$pb.TagNumber(5)
  set data($core.List<$core.int> v) { $_setBytes(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasData() => $_has(4);
  @$pb.TagNumber(5)
  void clearData() => clearField(5);
}
