///
//  Generated code. Do not modify.
//  source: msg.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class SendMsgReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SendMsgReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'token')
    ..aOM<MsgData>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'msgData', protoName: 'msgData', subBuilder: MsgData.create)
    ..hasRequiredFields = false
  ;

  SendMsgReq._() : super();
  factory SendMsgReq({
    $core.String? token,
    MsgData? msgData,
  }) {
    final _result = create();
    if (token != null) {
      _result.token = token;
    }
    if (msgData != null) {
      _result.msgData = msgData;
    }
    return _result;
  }
  factory SendMsgReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SendMsgReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SendMsgReq clone() => SendMsgReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SendMsgReq copyWith(void Function(SendMsgReq) updates) => super.copyWith((message) => updates(message as SendMsgReq)) as SendMsgReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SendMsgReq create() => SendMsgReq._();
  SendMsgReq createEmptyInstance() => create();
  static $pb.PbList<SendMsgReq> createRepeated() => $pb.PbList<SendMsgReq>();
  @$core.pragma('dart2js:noInline')
  static SendMsgReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SendMsgReq>(create);
  static SendMsgReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get token => $_getSZ(0);
  @$pb.TagNumber(1)
  set token($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearToken() => clearField(1);

  @$pb.TagNumber(2)
  MsgData get msgData => $_getN(1);
  @$pb.TagNumber(2)
  set msgData(MsgData v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsgData() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsgData() => clearField(2);
  @$pb.TagNumber(2)
  MsgData ensureMsgData() => $_ensure(1);
}

class MsgData extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgData', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sendID', protoName: 'sendID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'recvID', protoName: 'recvID')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'groupID', protoName: 'groupID')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientMsgID', protoName: 'clientMsgID')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serverMsgID', protoName: 'serverMsgID')
    ..a<$core.int>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'senderPlatformID', $pb.PbFieldType.O3, protoName: 'senderPlatformID')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'senderNickname', protoName: 'senderNickname')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'senderFaceURL', protoName: 'senderFaceURL')
    ..a<$core.int>(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sessionType', $pb.PbFieldType.O3, protoName: 'sessionType')
    ..a<$core.int>(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'msgFrom', $pb.PbFieldType.O3, protoName: 'msgFrom')
    ..a<$core.int>(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'contentType', $pb.PbFieldType.O3, protoName: 'contentType')
    ..a<$core.List<$core.int>>(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'content', $pb.PbFieldType.OY)
    ..a<$core.int>(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'seq', $pb.PbFieldType.OU3)
    ..aInt64(15, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serverTime', protoName: 'serverTime')
    ..aInt64(16, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientTime', protoName: 'clientTime')
    ..aOM<OfflinePushInfo>(17, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'offlinePushInfo', protoName: 'offlinePushInfo', subBuilder: OfflinePushInfo.create)
    ..pPS(18, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'atUserIDList', protoName: 'atUserIDList')
    ..m<$core.String, $core.bool>(19, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'options', entryClassName: 'MsgData.OptionsEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OB, packageName: const $pb.PackageName('pb'))
    ..hasRequiredFields = false
  ;

  MsgData._() : super();
  factory MsgData({
    $core.String? sendID,
    $core.String? recvID,
    $core.String? groupID,
    $core.String? clientMsgID,
    $core.String? serverMsgID,
    $core.int? senderPlatformID,
    $core.String? senderNickname,
    $core.String? senderFaceURL,
    $core.int? sessionType,
    $core.int? msgFrom,
    $core.int? contentType,
    $core.List<$core.int>? content,
    $core.int? seq,
    $fixnum.Int64? serverTime,
    $fixnum.Int64? clientTime,
    OfflinePushInfo? offlinePushInfo,
    $core.Iterable<$core.String>? atUserIDList,
    $core.Map<$core.String, $core.bool>? options,
  }) {
    final _result = create();
    if (sendID != null) {
      _result.sendID = sendID;
    }
    if (recvID != null) {
      _result.recvID = recvID;
    }
    if (groupID != null) {
      _result.groupID = groupID;
    }
    if (clientMsgID != null) {
      _result.clientMsgID = clientMsgID;
    }
    if (serverMsgID != null) {
      _result.serverMsgID = serverMsgID;
    }
    if (senderPlatformID != null) {
      _result.senderPlatformID = senderPlatformID;
    }
    if (senderNickname != null) {
      _result.senderNickname = senderNickname;
    }
    if (senderFaceURL != null) {
      _result.senderFaceURL = senderFaceURL;
    }
    if (sessionType != null) {
      _result.sessionType = sessionType;
    }
    if (msgFrom != null) {
      _result.msgFrom = msgFrom;
    }
    if (contentType != null) {
      _result.contentType = contentType;
    }
    if (content != null) {
      _result.content = content;
    }
    if (seq != null) {
      _result.seq = seq;
    }
    if (serverTime != null) {
      _result.serverTime = serverTime;
    }
    if (clientTime != null) {
      _result.clientTime = clientTime;
    }
    if (offlinePushInfo != null) {
      _result.offlinePushInfo = offlinePushInfo;
    }
    if (atUserIDList != null) {
      _result.atUserIDList.addAll(atUserIDList);
    }
    if (options != null) {
      _result.options.addAll(options);
    }
    return _result;
  }
  factory MsgData.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgData.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgData clone() => MsgData()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgData copyWith(void Function(MsgData) updates) => super.copyWith((message) => updates(message as MsgData)) as MsgData; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgData create() => MsgData._();
  MsgData createEmptyInstance() => create();
  static $pb.PbList<MsgData> createRepeated() => $pb.PbList<MsgData>();
  @$core.pragma('dart2js:noInline')
  static MsgData getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgData>(create);
  static MsgData? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sendID => $_getSZ(0);
  @$pb.TagNumber(1)
  set sendID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSendID() => $_has(0);
  @$pb.TagNumber(1)
  void clearSendID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get recvID => $_getSZ(1);
  @$pb.TagNumber(2)
  set recvID($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRecvID() => $_has(1);
  @$pb.TagNumber(2)
  void clearRecvID() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get groupID => $_getSZ(2);
  @$pb.TagNumber(3)
  set groupID($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasGroupID() => $_has(2);
  @$pb.TagNumber(3)
  void clearGroupID() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get clientMsgID => $_getSZ(3);
  @$pb.TagNumber(4)
  set clientMsgID($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasClientMsgID() => $_has(3);
  @$pb.TagNumber(4)
  void clearClientMsgID() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get serverMsgID => $_getSZ(4);
  @$pb.TagNumber(5)
  set serverMsgID($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasServerMsgID() => $_has(4);
  @$pb.TagNumber(5)
  void clearServerMsgID() => clearField(5);

  @$pb.TagNumber(6)
  $core.int get senderPlatformID => $_getIZ(5);
  @$pb.TagNumber(6)
  set senderPlatformID($core.int v) { $_setSignedInt32(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasSenderPlatformID() => $_has(5);
  @$pb.TagNumber(6)
  void clearSenderPlatformID() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get senderNickname => $_getSZ(6);
  @$pb.TagNumber(7)
  set senderNickname($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasSenderNickname() => $_has(6);
  @$pb.TagNumber(7)
  void clearSenderNickname() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get senderFaceURL => $_getSZ(7);
  @$pb.TagNumber(8)
  set senderFaceURL($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasSenderFaceURL() => $_has(7);
  @$pb.TagNumber(8)
  void clearSenderFaceURL() => clearField(8);

  @$pb.TagNumber(9)
  $core.int get sessionType => $_getIZ(8);
  @$pb.TagNumber(9)
  set sessionType($core.int v) { $_setSignedInt32(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasSessionType() => $_has(8);
  @$pb.TagNumber(9)
  void clearSessionType() => clearField(9);

  @$pb.TagNumber(10)
  $core.int get msgFrom => $_getIZ(9);
  @$pb.TagNumber(10)
  set msgFrom($core.int v) { $_setSignedInt32(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasMsgFrom() => $_has(9);
  @$pb.TagNumber(10)
  void clearMsgFrom() => clearField(10);

  @$pb.TagNumber(11)
  $core.int get contentType => $_getIZ(10);
  @$pb.TagNumber(11)
  set contentType($core.int v) { $_setSignedInt32(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasContentType() => $_has(10);
  @$pb.TagNumber(11)
  void clearContentType() => clearField(11);

  @$pb.TagNumber(12)
  $core.List<$core.int> get content => $_getN(11);
  @$pb.TagNumber(12)
  set content($core.List<$core.int> v) { $_setBytes(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasContent() => $_has(11);
  @$pb.TagNumber(12)
  void clearContent() => clearField(12);

  @$pb.TagNumber(14)
  $core.int get seq => $_getIZ(12);
  @$pb.TagNumber(14)
  set seq($core.int v) { $_setUnsignedInt32(12, v); }
  @$pb.TagNumber(14)
  $core.bool hasSeq() => $_has(12);
  @$pb.TagNumber(14)
  void clearSeq() => clearField(14);

  @$pb.TagNumber(15)
  $fixnum.Int64 get serverTime => $_getI64(13);
  @$pb.TagNumber(15)
  set serverTime($fixnum.Int64 v) { $_setInt64(13, v); }
  @$pb.TagNumber(15)
  $core.bool hasServerTime() => $_has(13);
  @$pb.TagNumber(15)
  void clearServerTime() => clearField(15);

  @$pb.TagNumber(16)
  $fixnum.Int64 get clientTime => $_getI64(14);
  @$pb.TagNumber(16)
  set clientTime($fixnum.Int64 v) { $_setInt64(14, v); }
  @$pb.TagNumber(16)
  $core.bool hasClientTime() => $_has(14);
  @$pb.TagNumber(16)
  void clearClientTime() => clearField(16);

  @$pb.TagNumber(17)
  OfflinePushInfo get offlinePushInfo => $_getN(15);
  @$pb.TagNumber(17)
  set offlinePushInfo(OfflinePushInfo v) { setField(17, v); }
  @$pb.TagNumber(17)
  $core.bool hasOfflinePushInfo() => $_has(15);
  @$pb.TagNumber(17)
  void clearOfflinePushInfo() => clearField(17);
  @$pb.TagNumber(17)
  OfflinePushInfo ensureOfflinePushInfo() => $_ensure(15);

  @$pb.TagNumber(18)
  $core.List<$core.String> get atUserIDList => $_getList(16);

  @$pb.TagNumber(19)
  $core.Map<$core.String, $core.bool> get options => $_getMap(17);
}

class OfflinePushInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'OfflinePushInfo', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'title')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'desc')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ex')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'iOSPushSound', protoName: 'iOSPushSound')
    ..aOB(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'iOSBadgeCount', protoName: 'iOSBadgeCount')
    ..hasRequiredFields = false
  ;

  OfflinePushInfo._() : super();
  factory OfflinePushInfo({
    $core.String? title,
    $core.String? desc,
    $core.String? ex,
    $core.String? iOSPushSound,
    $core.bool? iOSBadgeCount,
  }) {
    final _result = create();
    if (title != null) {
      _result.title = title;
    }
    if (desc != null) {
      _result.desc = desc;
    }
    if (ex != null) {
      _result.ex = ex;
    }
    if (iOSPushSound != null) {
      _result.iOSPushSound = iOSPushSound;
    }
    if (iOSBadgeCount != null) {
      _result.iOSBadgeCount = iOSBadgeCount;
    }
    return _result;
  }
  factory OfflinePushInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OfflinePushInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  OfflinePushInfo clone() => OfflinePushInfo()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  OfflinePushInfo copyWith(void Function(OfflinePushInfo) updates) => super.copyWith((message) => updates(message as OfflinePushInfo)) as OfflinePushInfo; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static OfflinePushInfo create() => OfflinePushInfo._();
  OfflinePushInfo createEmptyInstance() => create();
  static $pb.PbList<OfflinePushInfo> createRepeated() => $pb.PbList<OfflinePushInfo>();
  @$core.pragma('dart2js:noInline')
  static OfflinePushInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OfflinePushInfo>(create);
  static OfflinePushInfo? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get title => $_getSZ(0);
  @$pb.TagNumber(1)
  set title($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTitle() => $_has(0);
  @$pb.TagNumber(1)
  void clearTitle() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get desc => $_getSZ(1);
  @$pb.TagNumber(2)
  set desc($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDesc() => $_has(1);
  @$pb.TagNumber(2)
  void clearDesc() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get ex => $_getSZ(2);
  @$pb.TagNumber(3)
  set ex($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasEx() => $_has(2);
  @$pb.TagNumber(3)
  void clearEx() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get iOSPushSound => $_getSZ(3);
  @$pb.TagNumber(4)
  set iOSPushSound($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasIOSPushSound() => $_has(3);
  @$pb.TagNumber(4)
  void clearIOSPushSound() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get iOSBadgeCount => $_getBF(4);
  @$pb.TagNumber(5)
  set iOSBadgeCount($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasIOSBadgeCount() => $_has(4);
  @$pb.TagNumber(5)
  void clearIOSBadgeCount() => clearField(5);
}

class SendMsgResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'SendMsgResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errCode', $pb.PbFieldType.O3, protoName: 'errCode')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errMsg', protoName: 'errMsg')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serverMsgID', protoName: 'serverMsgID')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'clientMsgID', protoName: 'clientMsgID')
    ..aInt64(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'serverTime', protoName: 'serverTime')
    ..hasRequiredFields = false
  ;

  SendMsgResp._() : super();
  factory SendMsgResp({
    $core.int? errCode,
    $core.String? errMsg,
    $core.String? serverMsgID,
    $core.String? clientMsgID,
    $fixnum.Int64? serverTime,
  }) {
    final _result = create();
    if (errCode != null) {
      _result.errCode = errCode;
    }
    if (errMsg != null) {
      _result.errMsg = errMsg;
    }
    if (serverMsgID != null) {
      _result.serverMsgID = serverMsgID;
    }
    if (clientMsgID != null) {
      _result.clientMsgID = clientMsgID;
    }
    if (serverTime != null) {
      _result.serverTime = serverTime;
    }
    return _result;
  }
  factory SendMsgResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SendMsgResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SendMsgResp clone() => SendMsgResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SendMsgResp copyWith(void Function(SendMsgResp) updates) => super.copyWith((message) => updates(message as SendMsgResp)) as SendMsgResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SendMsgResp create() => SendMsgResp._();
  SendMsgResp createEmptyInstance() => create();
  static $pb.PbList<SendMsgResp> createRepeated() => $pb.PbList<SendMsgResp>();
  @$core.pragma('dart2js:noInline')
  static SendMsgResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SendMsgResp>(create);
  static SendMsgResp? _defaultInstance;

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
  $core.String get serverMsgID => $_getSZ(2);
  @$pb.TagNumber(3)
  set serverMsgID($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasServerMsgID() => $_has(2);
  @$pb.TagNumber(3)
  void clearServerMsgID() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get clientMsgID => $_getSZ(3);
  @$pb.TagNumber(4)
  set clientMsgID($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasClientMsgID() => $_has(3);
  @$pb.TagNumber(4)
  void clearClientMsgID() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get serverTime => $_getI64(4);
  @$pb.TagNumber(5)
  set serverTime($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasServerTime() => $_has(4);
  @$pb.TagNumber(5)
  void clearServerTime() => clearField(5);
}

