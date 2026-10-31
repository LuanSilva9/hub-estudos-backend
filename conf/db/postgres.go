package db

import (
	"time"

	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/users"
	"github.com/LuanSilva9/hub-estudos-backend/pkg/modules/workspaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConnection(dsn string, env string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxOpenConns(25)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(time.Hour)
	sqlDb.SetConnMaxIdleTime(10 * time.Minute)

	if env == "dev" {
		if err := db.AutoMigrate(&users.User{}, &workspaces.Workspace{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}
