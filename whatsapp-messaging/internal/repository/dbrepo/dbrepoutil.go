package dbrepo

import (
	"context"
	"database/sql"

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

func GetDBConn(ctx context.Context) (*sql.Conn, error){
	if DBConnectionPool == nil{
		panic("no database connection found")
	}

	db,err := DBConnectionPool.Conn(ctx)
	if err != nil {
		return nil,err
	}

	return db, nil

}

func GetDBConnPool() (*sql.DB){
	return DBConnectionPool
}