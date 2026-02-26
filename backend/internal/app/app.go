package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kmj36/fieldnotes-tech-blog/configs"
	"github.com/kmj36/fieldnotes-tech-blog/internal/handler"
	"github.com/kmj36/fieldnotes-tech-blog/internal/middleware"
	"github.com/kmj36/fieldnotes-tech-blog/pkg/validator"
	"gorm.io/gorm"
)

// gin app 패키지
type App struct {
	router *gin.Engine
	pingHandler *handler.PingHandler
	db *gorm.DB
}

// app 패키지 생성자(constructor) App 객체 반환
func New(db *gorm.DB, runMode string) *App {
	gin.SetMode(runMode) // 서버 다중 실행 구조로 변경 시 main.go로 이동

	return &App{
		router: gin.Default(),
		pingHandler: handler.NewPingHandler(),
		db: db,
	}
}

// App 객체 메소드 - 미들웨어 핸들러 등록
func (app *App) setupMiddleware(cfg *configs.Config) {
	// CORS 미들웨어 핸들러
	app.router.Use(middleware.CORSHandler(cfg))
}

// App 객체 메소드 - Gin 라우팅 설정
func (app *App) setupRoutes() {
	// 기본 라우트
	app.router.GET("/ping", app.pingHandler.Ping)

	/*api := app.router.Group("/api/v1")
	{
		userRepo := repository.NewUserRepository(app.db)
		userService := service.NewUserService(userRepo)
        userHandler := handler.NewUserHandler(userService)
        api.GET("/users/:id", userHandler.GetUser)
	}*/
}

// App 객체 메소드 - Gin 실행 함수
func (app *App) Run(cfg *configs.Config) error {
	if err := validator.ValidateAddr(cfg.ServerAddr); err != nil {
		return err
	}

	app.setupMiddleware(cfg)
	app.setupRoutes()

	fmt.Println("Server started")
	app.router.Run(cfg.ServerAddr)
	return nil // [TODO] return app.router.Run(cfg.ServerAddr)
}
