package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	Name	*string			`gorm:"type:varchar(255);not null;" json:"name"`
	Description	*string		`gorm:"type:varchar(255);not null;" json:"description"`
	OwnerID uuid.UUID		`gorm:"type:uuid;not null;" json:"owner_id"`
	GitRepoID *string		`gorm:"type:text;not null;" json:"git_repo_id"`
	GithubRepoName *string	`gorm:"type:text;not null;" json:"github_repo_name"`
	GitDefaultBranch *string	`gorm:"type:varchar(255);not null;" json:"git_default_branch"`
	AssetModel  []AssetModel `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	Discussion []DiscussionModel `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	Feature []FeatureModel `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	RoleModel []RoleModel `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
	CollaboratorModel	[]CollaboratorModel `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID;"`
}


func MigrateProjectModel ( db *gorm.DB ) {
	db.AutoMigrate(&ProjectModel{})
}