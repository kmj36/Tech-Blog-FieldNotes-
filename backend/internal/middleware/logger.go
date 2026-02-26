package middleware

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerNew() *zap.Logger {
	var logger *zap.Logger

	if gin.Mode() == gin.ReleaseMode {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	return logger
}

func ZapLoggerHandler() gin.HandlerFunc {
    logger := LoggerNew()
    
    return ginzap.Ginzap(logger, time.RFC3339, true)
}

func ZapRecoveryHandler() gin.HandlerFunc {
    logger := LoggerNew()

    return ginzap.RecoveryWithZap(logger, true)
}