package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmj36/fieldnotes-tech-blog/internal/dto"
	"github.com/kmj36/fieldnotes-tech-blog/pkg/cryption"
)

func JWTAuthMiddleware(jwtmanager *cryption.JWTManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwtmanager.ValidateJWT(tokenString)

		if authHeader == "" || !token.Valid || err != nil {
			ctx.JSON(http.StatusUnauthorized, dto.CommonResponse[any]{
				Status: http.StatusUnauthorized,
				Code: "E401_001",
				Message: "인증에 실패하였습니다.",
				Detail: "Invalid admin credentials.",
				Timestamp: time.Now(),
				Path: ctx.Request.URL.Path,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}