package dbrepo

import (
	"context"

	"gorm.io/gorm"
)

func GetGormConn(ctx context.Context) (*gorm.DB, error) {
	if GormConnectionPool == nil {
		panic("no database connection found")
	}

	db := GormConnectionPool.WithContext(ctx)
	if db.Error != nil {
		return nil, db.Error
	}

	if GormLog {
		db = db.Session(&gorm.Session{Logger: GormDebugLogger})
	}

	return db, nil
}
