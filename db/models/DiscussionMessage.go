package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscussionMessageModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	DiscussionID uuid.UUID	`gorm:"type:uuid;not null;" json:"discussion_id"`
	ProfileID uuid.UUID		`gorm:"type:uuid;not null;" json:"user_id"`
	Message *string			`gorm:"type:text;not null;" json:"message"`
	InReplyToID uuid.UUID	`gorm:"type:uuid;default:null;" json:"in_reply_to_id"`
	DiscussionMessageModel *DiscussionMessageModel `gorm:"foreignKey:InReplyToID;references:ID" json:"discussion_message_model"`
	AssetModel	[]AssetModel	`gorm:"foreignKey:DiscussionMessageID;references:ID" json:"asset_model"`
}


func MigrateDiscussionMessageModel(db *gorm.DB) {
	db.AutoMigrate(&DiscussionMessageModel{})
}