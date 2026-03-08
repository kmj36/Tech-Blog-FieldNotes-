package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmj36/fieldnotes-tech-blog/internal/model"
)

// 404 에러 핸들링 - 정의되지 않은 라우트 공통 처리, 비즈니스 핸들러 별도 404 에러
func NoRoute() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		ctx.JSON(http.StatusNotFound, model.CommonResponse[any]{
			Status: model.ErrNotFound.Status,
			Code: model.ErrNotFound.Code,
			Detail: ctx.Request.URL.Path + " resource not found.",
			Message: model.ErrNotFound.Message,
			Timestamp: time.Now().UTC(),
			Path: ctx.Request.URL.Path,
		})
	}
}

// 405 에러 핸들링 - 정의되지 않은 메소드 처리
func NoMethod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, model.CommonResponse[any]{
			Status: model.ErrMethodNotAllowed.Status,
			Code: model.ErrMethodNotAllowed.Code,
			Detail: ctx.Request.Method + " method not allowed.",
			Message: model.ErrMethodNotAllowed.Message,
			Timestamp: time.Now().UTC(),
			Path: ctx.Request.URL.Path,
		})
	}
}