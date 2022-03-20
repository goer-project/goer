package migration

import (
	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator)

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

var migrationFiles []MigrationFile

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

func GetMigrationFile(name string) MigrationFile {
	for _, migrationFile := range migrationFiles {
		if name == migrationFile.FileName {
			return migrationFile
		}
	}

	return MigrationFile{}
}
