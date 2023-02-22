package types

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)


type tDef struct {
	ntype string
	definition string
}


func NewType ( type_name string, type_definition string) string {
	str := fmt.Sprintf(`
		DO $$
			BEGIN
				IF NOT EXISTS (select 1 from pg_type where typname = '%s') THEN
					CREATE TYPE %s AS %s;
			END IF;
		END $$;
	`, type_name, type_name, type_definition)

	return str
}


func CreateTypes (db *gorm.DB) {


	for _, t_def := range []tDef{
		{
			ntype: "component_type",
			definition: "enum ('feat', 'fix', 'docs', 'style', 'refactor', 'perf', 'test', 'build', 'ci', 'chore', 'revert')",
		},
		{
			ntype: "check_point_state",
			definition: "enum ('open', 'closed', 'merged')",
		},
		{
			ntype: "feature_state",
			definition: "enum ('in_backlog', 'in_progress', 'in_review', 'in_testing', 'in_production', 'in_archive', 'bug_found', 'bug_fixed', 'bug_in_progress', 'bug_in_review', 'bug_in_testing', 'bug_in_production', 'bug_in_archive')",
		},
	} {
		value := t_def 

		err := db.Exec(NewType(t_def.ntype, t_def.definition)).Error

		if err != nil {
			log.Fatalf("Unable to create type %s; error:: %s", value.ntype, err)
		}
	}



}

func MigrateTypes (db *gorm.DB) {
	CreateTypes(db)
}