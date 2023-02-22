package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComponentModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	Title	*string			`gorm:"type:varchar(255);not null;" json:"title"`
	Message	*string			`gorm:"type:text;not null;" json:"message"`
	Description	*string		`gorm:"type:text;not null;" json:"description"`
	FeatureID uuid.UUID		`gorm:"type:uuid;not null;" json:"feature_id"`
	ComponentType *string   `gorm:"type:component_type;not null;default:'feat'" json:"component_type"`
	GitCommitSha *string	`gorm:"type:varchar(255);not null;" json:"git_commit_sha"`
	DiscussionModel []DiscussionModel `gorm:"foreignKey:ComponentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	AssetModel  []AssetModel `gorm:"foreignKey:ComponentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
}

func MigrateComponentModel(db *gorm.DB) {
	db.AutoMigrate(&ComponentModel{})
} 