package main

import (
	"context"
	"log"
	"todo-golang/config"
	"todo-golang/di"
	"todo-golang/infrastructure/database"
	"todo-golang/infrastructure/server/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()

	// デバッグ用：接続文字列を出力
	databaseURL := config.PostgresURL()
	log.Printf("Database URL: %s", databaseURL)

	// dbの初期設定
	db, err := database.NewDatabase(ctx, databaseURL)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// テーブルの自動マイグレーション
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	apiContainer := di.InitializeAPIContainer(db)

	routes.RegisterRoutes(e, apiContainer)

	// サーバー起動
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
