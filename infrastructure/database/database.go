package database

import (
	"context"
	"fmt"
	"todo-golang/domain/entity"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type DBClient struct {
// 	gormDB *gorm.DB
// }

func NewDatabase(ctx context.Context, databaseUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm db: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql DB from gorm: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}

// AutoMigrate テーブルの自動マイグレーション
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		// &entity.User{},
		&entity.Todo{},
	)
}
