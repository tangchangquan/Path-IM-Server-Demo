package model

import (
	"database/sql/driver"
	"encoding/json"
	chatpb "github.com/Path-IM/Path-IM-Server-Demo/app/msg/cmd/rpc/pb"
	xormerr "github.com/Path-IM/Path-IM-Server-Demo/common/xorm/err"
	"github.com/Path-IM/Path-IM-Server-Demo/common/xorm/global"
	"gorm.io/gorm"
	"time"
)

type MysqlSingleChat struct {
	ServerMsgID string `gorm:"column:server_msg_id;primary_key;"`

	SendID           string               `gorm:"column:send_id;type:char(32);default:'';not null;"`
	RecvID           string               `gorm:"column:recv_id;type:char(32);default:'';not null;"`
	GroupID          string               `gorm:"column:group_id;type:char(32);default:'';not null;"`
	ClientMsgID      string               `gorm:"column:client_msg_id;type:char(32);default:'';not null;"`
	SenderPlatformID int32                `gorm:"column:sender_platform_id;type:tinyint(3);default:0;not null;"`
	SenderNickname   string               `gorm:"column:sender_nickname;type:varchar(32);default:'';not null;"`
	SenderFaceURL    string               `gorm:"column:sender_face_url;type:varchar(255);default:'';not null;"`
	SessionType      int32                `gorm:"column:session_type;type:tinyint(3);default:0;not null;"`
	MsgFrom          int32                `gorm:"column:msg_from;type:int(10);default:0;not null;"`
	ContentType      int32                `gorm:"column:content_type;type:int(10);default:0;not null;"`
	Content          string               `gorm:"column:content;type:json;"`
	Seq              uint32               `gorm:"column:seq;type:bigint(19);default:0;not null;"`
	SendTime         int64                `gorm:"column:send_time;index;type:bigint(13);default:0;not null;"`
	CreateTime       int64                `gorm:"column:create_time;type:bigint(13);default:0;not null;"`
	OfflinePushInfo  *OfflinePushInfo     `gorm:"column:offline_push_info;type:json;"`
	AtUserIDList     global.SliceString   `gorm:"column:at_user_id_list;type:json;"`
	Options          global.MapStringBool `gorm:"column:options;type:json;"`
}

func (s *MysqlSingleChat) TableName() string {
	return "single_msg_" + s.GetConversationID()
}

func (s *MysqlSingleChat) GetConversationID() string {
	// recvid, sendid 排序
	if s.RecvID < s.SendID {
		return s.RecvID + "-" + s.SendID
	} else {
		return s.SendID + "-" + s.RecvID
	}
}

type OfflinePushInfo struct {
	Title         string
	Desc          string
	Ex            string
	IOSPushSound  string
	IOSBadgeCount bool
}

func (o OfflinePushInfo) Value() (driver.Value, error) {
	return json.Marshal(o)
}
func (o *OfflinePushInfo) Scan(input interface{}) error {
	s := string(input.([]byte))
	err := json.Unmarshal([]byte(s), o)
	if err != nil {

	}
	return nil
}
func (s *MysqlSingleChat) Insert(tx *gorm.DB) error {
	tableName := s.TableName()
	err := tx.Table(tableName).Create(s).Error
	if err != nil {
		if xormerr.TableNotFound(err) {
			err = tx.Table(tableName).AutoMigrate(s)
			if err != nil {
				return err
			} else {
				return tx.Table(tableName).Create(s).Error
			}
		} else {
			return err
		}
	}
	return nil
}
func NewOfflinePushInfo(info *chatpb.OfflinePushInfo) *OfflinePushInfo {
	if info == nil {
		return &OfflinePushInfo{}
	}
	return &OfflinePushInfo{
		Title:         info.Title,
		Desc:          info.Desc,
		Ex:            info.Ex,
		IOSPushSound:  info.IOSPushSound,
		IOSBadgeCount: info.IOSBadgeCount,
	}
}

type MysqlGroupChat struct {
	ServerMsgID string `gorm:"column:server_msg_id;primary_key;"`

	SendID           string               `gorm:"column:send_id;type:char(32);default:'';not null;"`
	RecvID           string               `gorm:"column:recv_id;type:char(32);default:'';not null;"`
	GroupID          string               `gorm:"column:group_id;type:char(32);default:'';not null;"`
	ClientMsgID      string               `gorm:"column:client_msg_id;type:char(32);default:'';not null;"`
	SenderPlatformID int32                `gorm:"column:sender_platform_id;type:tinyint(3);default:0;not null;"`
	SenderNickname   string               `gorm:"column:sender_nickname;type:varchar(32);default:'';not null;"`
	SenderFaceURL    string               `gorm:"column:sender_face_url;type:varchar(255);default:'';not null;"`
	SessionType      int32                `gorm:"column:session_type;type:tinyint(3);default:0;not null;"`
	MsgFrom          int32                `gorm:"column:msg_from;type:int(10);default:0;not null;"`
	ContentType      int32                `gorm:"column:content_type;type:int(10);default:0;not null;"`
	Content          string               `gorm:"column:content;type:json;"`
	Seq              uint32               `gorm:"column:seq;type:bigint(19);default:0;not null;"`
	SendTime         int64                `gorm:"column:send_time;index;type:bigint(13);default:0;not null;"`
	CreateTime       int64                `gorm:"column:create_time;type:bigint(13);default:0;not null;"`
	OfflinePushInfo  *OfflinePushInfo     `gorm:"column:offline_push_info;type:json;"`
	AtUserIDList     global.SliceString   `gorm:"column:at_user_id_list;type:json;"`
	Options          global.MapStringBool `gorm:"column:options;type:json;"`
}

func (g *MysqlGroupChat) TableName() string {
	return "single_msg_" + g.GroupID + "_" + time.Now().Format("200601")
}

func (g *MysqlGroupChat) Insert(tx *gorm.DB) error {
	tableName := g.TableName()
	err := tx.Table(tableName).Create(g).Error
	if err != nil {
		if xormerr.TableNotFound(err) {
			err = tx.Table(tableName).AutoMigrate(g)
			if err != nil {
				return err
			} else {
				return tx.Table(tableName).Create(g).Error
			}
		} else {
			return err
		}
	}
	return nil
}
