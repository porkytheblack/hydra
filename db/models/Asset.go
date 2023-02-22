package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	ProjectID uuid.UUID		`gorm:"type:uuid;not null;" json:"project_id"`
	FeatureID uuid.UUID		`gorm:"type:uuid;not null;" json:"feature_id"`
	ComponentID uuid.UUID	`gorm:"type:uuid;not null;" json:"component_id"`
	DiscussionMessageID uuid.UUID	`gorm:"type:uuid;not null;" json:"discussion_message_id"`
	FileName *string		`gorm:"type:varchar(255);not null;" json:"file_name"`
	FileLocation *string	`gorm:"type:varchar(255);not null;" json:"file_location"`
	FileType *string		`gorm:"type:varchar(255);not null;" json:"file_type"`
}

func MigrateAssetModel(db *gorm.DB) {
	db.AutoMigrate(&AssetModel{})
}