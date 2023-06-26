package storage

import (
	"github.com/yansuan/go-admin-backend/pkg/ormx"

	"gorm.io/gorm"
)

func New(cfg ormx.DBConfig) (*gorm.DB, error) {
	db, err := ormx.New(cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}
