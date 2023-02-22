package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscussionModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	ProjectID uuid.UUID		`gorm:"type:uuid;not null;" json:"project_id"`
	FeatureID uuid.UUID		`gorm:"type:uuid;not null;" json:"feature_id"`
	ComponentID uuid.UUID	`gorm:"type:uuid;not null;" json:"component_id"`
	GitHubCommentID *int	`gorm:"type:int;not null;" json:"github_comment_id"`
	CheckPointID uuid.UUID	`gorm:"type:uuid;not null;" json:"checkpoint_id"`
	DiscussionMessageModel	[]DiscussionMessageModel `gorm:"foreignKey:DiscussionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
}

func MigrateDiscussionModel(db *gorm.DB) {
	db.AutoMigrate(&DiscussionModel{})
}