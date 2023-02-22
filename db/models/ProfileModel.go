package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	FirstName *string			`gorm:"type:varchar(255);not null;" json:"first_name"`
	LastName *string			`gorm:"type:varchar(255);not null;" json:"last_name"`
	Email *string			`gorm:"type:varchar(255);not null;" json:"email"`
	PicUrl *string			`gorm:"type:varchar(255);not null;" json:"pic_url"`
	CollaboratorModel	[]CollaboratorModel	`gorm:"foreignKey:ProfileID;references:ID" json:"collaborator_model"`
	DiscussionMessageModel	[]DiscussionMessageModel	`gorm:"foreignKey:ProfileID;references:ID" json:"discussion_message_model"`
	CheckPoint	[]CheckPoint	`gorm:"foreignKey:CreatorID;references:ID" json:"check_point"`
}

func MigrateProfileModel(db *gorm.DB) {
	db.AutoMigrate(&ProfileModel{})
}