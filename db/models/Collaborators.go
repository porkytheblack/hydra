package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type CollaboratorModel struct {
	ID		uuid.UUID		`gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4();" json:"id"`
	ProjectID uuid.UUID     `gorm:"type:uuid;not null;" json:"project_id"`
	ProfileID uuid.UUID        `gorm:"type:uuid;not null;" json:"user_id"`
	RoleID uuid.UUID        `gorm:"type:uuid;not null;" json:"role_id"`
}

func MigrateCollaboratorModel (db *gorm.DB) {
	db.AutoMigrate(&CollaboratorModel{})
} 