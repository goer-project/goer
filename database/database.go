package database

import (
	"errors"
	"fmt"

	"github.com/goer-project/goer/config"
	"gorm.io/gorm"
)

func CurrentDatabase() string {
	return DB.Migrator().CurrentDatabase()
}

func TableName(obj interface{}) string {
	statement := &gorm.Statement{DB: DB}
	err := statement.Parse(obj)
	if err != nil {
		return ""
	}

	return statement.Schema.Table
}

func DeleteAllTables() error {
	var err error

	switch config.NewConfig.Database.Connection {
	case "mysql":
		err = deleteMysqlDatabase()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteMysqlDatabase() error {
	dbname := CurrentDatabase()

	sql := fmt.Sprintf("DROP DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("CREATE DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("USE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func deleteAllSqliteTables() error {
	var tables []string

	DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'")
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			continue
		}
	}
	return nil
}
