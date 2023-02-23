package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeatureModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	Name	*string			`gorm:"type:varchar(255);not null;" json:"name"`
	Description	*string		`gorm:"type:text;not null;" json:"description"`
	ProjectID uuid.UUID		`gorm:"type:uuid;not null;" json:"project_id"`
	CreatorID uuid.UUID		`gorm:"type:uuid;not null;" json:"creator_id"`
	GitBranchName *string	`gorm:"type:varchar(255);not null;" json:"git_branch_name"`
	GitCommitID *string		`gorm:"type:varchar(255);not null;" json:"git_commit_id"`
	CurrentState *string	`gorm:"type:feature_state;not null;default:'in_backlog';" json:"current_state"`
	AssetModel  []AssetModel `gorm:"foreignKey:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	CheckPoint	[]CheckPoint `gorm:"foreignKey:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	Component []ComponentModel `gorm:"foreignKey:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	Discussion []DiscussionModel `gorm:"foreignKey:FeatureID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	Base		[]CheckPoint `gorm:"foreignKey:BaseFeatureID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
}

func MigrateFeatureModel(db *gorm.DB) {
	db.AutoMigrate(&FeatureModel{})
}