package main

import (
	"log"

	"github.com/kmj36/fieldnotes-tech-blog/configs"
	"github.com/kmj36/fieldnotes-tech-blog/internal/app"
	"github.com/kmj36/fieldnotes-tech-blog/pkg/database"
)

func main() {
	// 사전 설정 호출
	cfg := configs.Load()

	// 데이터베이스 연결 정보 DSN
	dsn := database.NewPostgresDSN(cfg.DB)

	// 데이터베이스 객체 생성
	db, err := database.NewPostgresDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	backendApp := app.New(db)

	if err := backendApp.Run(cfg.ServerAddr); err != nil {
		log.Fatal(err)
	}
}
