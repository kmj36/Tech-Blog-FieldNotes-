package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kmj36/fieldnotes-tech-blog/internal/dto"
	"github.com/kmj36/fieldnotes-tech-blog/internal/service"
)

// 관리자 계정 관련 HTTP 요청/응답 처리
type AccountHandler struct {service *service.AccountService}

func NewAccountHandler(service *service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) Create(ctx *gin.Context) {
	var req dto.CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.respondBindError(ctx, err)
        return
	}

	// 계정 추가 처리
	if err := h.service.Create(ctx, &req) ; err != nil {
		h.respondProcessError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseWrapper[any]{
		Status: http.StatusCreated,
		Code: dto.ErrCreated.Code,
		Detail: "Account created successfully.",
		Message: dto.ErrCreated.Message,
		Timestamp: time.Now().UTC(),
		Path: ctx.Request.URL.Path,
	})
}

func (h *AccountHandler) respondBindError(ctx *gin.Context, err error) {
	var ve validator.ValidationErrors
    switch {
    case errors.As(err, &ve):
        fe := ve[0]
        ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper[any]{
            Status:    http.StatusBadRequest,
            Code:      dto.ErrBadRequestType.Code,
            Detail:    fmt.Sprintf("'%s' does not satisfy '%s' condition", fe.Field(), fe.Tag()),
            Message:   dto.ErrBadRequestType.Message,
            Timestamp: time.Now().UTC(),
            Path:      ctx.Request.URL.Path,
        })
    case errors.Is(err, io.EOF):
        ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper[any]{
            Status:    http.StatusBadRequest,
            Code:      dto.ErrBadRequestMissing.Code,
            Detail:    "request body is empty",
            Message:   dto.ErrBadRequestMissing.Message,
            Timestamp: time.Now().UTC(),
            Path:      ctx.Request.URL.Path,
        })
    default:
        ctx.JSON(http.StatusBadRequest, dto.ResponseWrapper[any]{
            Status:    http.StatusBadRequest,
            Code:      dto.ErrBadRequestMissing.Code,
            Detail:    err.Error(),
            Message:   dto.ErrBadRequestMissing.Message,
            Timestamp: time.Now().UTC(),
            Path:      ctx.Request.URL.Path,
        })
    }
}

func (h *AccountHandler) respondProcessError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, dto.ErrAccountAlreadyExists):
		ctx.JSON(http.StatusConflict, dto.ResponseWrapper[any]{
			Status: http.StatusConflict,
			Code: dto.ErrConflict.Code,
			Detail: err.Error(),
			Message: dto.ErrConflict.Message,
			Timestamp: time.Now().UTC(),
			Path: ctx.Request.URL.Path,
		})
	default:
		ctx.JSON(http.StatusInternalServerError, dto.ResponseWrapper[any]{
			Status:    http.StatusInternalServerError,
			Code:      dto.ErrInternal.Code,
			Detail:    err.Error(),
			Message:   dto.ErrInternal.Message,
			Timestamp: time.Now().UTC(),
			Path:      ctx.Request.URL.Path,
		})
	}
}