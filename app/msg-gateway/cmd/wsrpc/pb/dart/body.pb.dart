///
//  Generated code. Do not modify.
//  source: body.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class BodyReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BodyReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'reqIdentifier', $pb.PbFieldType.OU3, protoName: 'reqIdentifier')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sendID', protoName: 'sendID')
    ..a<$core.List<$core.int>>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'data', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  BodyReq._() : super();
  factory BodyReq({
    $core.int? reqIdentifier,
    $core.String? token,
    $core.String? sendID,
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
    if (data != null) {
      _result.data = data;
    }
    return _result;
  }
  factory BodyReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BodyReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BodyReq clone() => BodyReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BodyReq copyWith(void Function(BodyReq) updates) => super.copyWith((message) => updates(message as BodyReq)) as BodyReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BodyReq create() => BodyReq._();
  BodyReq createEmptyInstance() => create();
  static $pb.PbList<BodyReq> createRepeated() => $pb.PbList<BodyReq>();
  @$core.pragma('dart2js:noInline')
  static BodyReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BodyReq>(create);
  static BodyReq? _defaultInstance;

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
  $core.List<$core.int> get data => $_getN(3);
  @$pb.TagNumber(4)
  set data($core.List<$core.int> v) { $_setBytes(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasData() => $_has(3);
  @$pb.TagNumber(4)
  void clearData() => clearField(4);
}

class BodyResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'BodyResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'reqIdentifier', $pb.PbFieldType.OU3, protoName: 'reqIdentifier')
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errCode', $pb.PbFieldType.OU3, protoName: 'errCode')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errMsg', protoName: 'errMsg')
    ..a<$core.List<$core.int>>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'data', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  BodyResp._() : super();
  factory BodyResp({
    $core.int? reqIdentifier,
    $core.int? errCode,
    $core.String? errMsg,
    $core.List<$core.int>? data,
  }) {
    final _result = create();
    if (reqIdentifier != null) {
      _result.reqIdentifier = reqIdentifier;
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
  factory BodyResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BodyResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  BodyResp clone() => BodyResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  BodyResp copyWith(void Function(BodyResp) updates) => super.copyWith((message) => updates(message as BodyResp)) as BodyResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BodyResp create() => BodyResp._();
  BodyResp createEmptyInstance() => create();
  static $pb.PbList<BodyResp> createRepeated() => $pb.PbList<BodyResp>();
  @$core.pragma('dart2js:noInline')
  static BodyResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BodyResp>(create);
  static BodyResp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get reqIdentifier => $_getIZ(0);
  @$pb.TagNumber(1)
  set reqIdentifier($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasReqIdentifier() => $_has(0);
  @$pb.TagNumber(1)
  void clearReqIdentifier() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get errCode => $_getIZ(1);
  @$pb.TagNumber(2)
  set errCode($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get errMsg => $_getSZ(2);
  @$pb.TagNumber(3)
  set errMsg($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrMsg() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrMsg() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$core.int> get data => $_getN(3);
  @$pb.TagNumber(4)
  set data($core.List<$core.int> v) { $_setBytes(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasData() => $_has(3);
  @$pb.TagNumber(4)
  void clearData() => clearField(4);
}

