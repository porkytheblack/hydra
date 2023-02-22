package models

import "gorm.io/gorm"


func MigrateModels( db *gorm.DB ) {
	MigrateAssetModel(db)
	MigrateCollaboratorModel(db)
	MigrateComponentModel(db)
	MigrateDiscussionMessageModel(db)
	MigrateFeatureModel(db)
	MigrateProjectModel(db)
	MigrateProfileModel(db)
	MigrateRoleModel(db)
	MigrateCheckPoint(db)
	
} 