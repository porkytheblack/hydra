package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	Auth0ID *string			`gorm:"type:varchar(255);not null;" json:"auth0_id"`
	FirstName *string			`gorm:"type:varchar(255);not null;" json:"first_name"`
	LastName *string			`gorm:"type:varchar(255);not null;" json:"last_name"`
	Email *string			`gorm:"type:varchar(255);not null;" json:"email"`
	PicUrl *string			`gorm:"type:varchar(255);not null;" json:"pic_url"`
	CollaboratorModel	[]CollaboratorModel	`gorm:"foreignKey:ProfileID;references:ID" json:"collaborator_model"`
	DiscussionMessageModel	[]DiscussionMessageModel	`gorm:"foreignKey:ProfileID;references:ID" json:"discussion_message_model"`
	CheckPoint	[]CheckPoint	`gorm:"foreignKey:CreatorID;references:ID" json:"check_point"`
	FeatureModel	[]FeatureModel	`gorm:"foreignKey:CreatorID;references:ID" json:"feature_model"`
	ProjectModel	[]ProjectModel	`gorm:"foreignKey:OwnerID;references:ID" json:"project_model"`
	ComponentModel	[]ComponentModel	`gorm:"foreignKey:CreatorID;references:ID" json:"component_model"`
}

func MigrateProfileModel(db *gorm.DB) {
	db.AutoMigrate(&ProfileModel{})
}