package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CheckPoint struct {
	ID		  uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	BaseFeatureID uuid.UUID		`gorm:"type:uuid;not null;" json:"base_feature_id"`
	FeatureID uuid.UUID		`gorm:"type:uuid;not null;" json:"feature_id"`
	BaseBranch *string		`gorm:"type:varchar(255);not null;" json:"base_branch"`
	Branch *string		`gorm:"type:varchar(255);not null;" json:"branch"`
	Description	*string		`gorm:"type:text;not null;" json:"description"`
	State *string   `gorm:"type:check_point_state;not null;default:'open'" json:"state"`
	GitPullRequestID *string		`gorm:"type:varchar(255);not null;" json:"git_pull_request_id"`
	ProjectPullRequestNumber *int		`gorm:"type:int;not null;" json:"project_pull_request_number"`
	CreatorID uuid.UUID		`gorm:"type:uuid;not null;" json:"creator_id"`
	DiscussionModel []DiscussionModel `gorm:"foreignKey:CheckPointID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
}

func MigrateCheckPoint(db *gorm.DB) {
	db.AutoMigrate(&CheckPoint{})
}
