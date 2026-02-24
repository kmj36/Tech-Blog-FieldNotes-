package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kmj36/fieldnotes-tech-blog/internal/handler"
	"github.com/kmj36/fieldnotes-tech-blog/pkg/validator"
	"gorm.io/gorm"
)

type App struct {
	router *gin.Engine
	pingHandler *handler.PingHandler
	db *gorm.DB
}

// app 패키지 생성자(constructor) App 객체 반환
func New(db *gorm.DB) *App {
	return &App{
		router: gin.Default(),
		pingHandler: handler.NewPingHandler(),
		db: db,
	}
}

// App 객체 메소드 - Gin 라우팅 설정
func (app *App) setupRoutes() {
	// 미들웨어 설정
	//app.router.Use(middleware.CORS())

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
func (app *App) Run(addr string) error {
	if err := validator.ValidateAddr(addr); err != nil {
		return err
	}

	app.setupRoutes()

	fmt.Println("Server started")
	app.router.Run(addr)
	return nil
}
