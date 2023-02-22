package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	Name	*string			`gorm:"type:varchar(255);not null;" json:"name"`
	Description	*string		`gorm:"type:text;not null;" json:"description"`
	ProjectID uuid.UUID		`gorm:"type:uuid;default:null;" json:"project_id"`
	CollaboratorModel	[]CollaboratorModel	`gorm:"foreignKey:RoleID;references:ID" json:"collaborator_model"`
}

func MigrateRoleModel(db *gorm.DB) {
	db.AutoMigrate(&RoleModel{})
}