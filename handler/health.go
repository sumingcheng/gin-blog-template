package handler

import (
	"blog/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	checks := gin.H{}
	healthy := true

	// 检查 PostgreSQL
	sqlDB, err := database.GetBlogDBConnection().DB()
	if err != nil {
		checks["postgres"] = "error: " + err.Error()
		healthy = false
	} else {
		pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
		if err := sqlDB.PingContext(pingCtx); err != nil {
			checks["postgres"] = "error: " + err.Error()
			healthy = false
		} else {
			checks["postgres"] = "ok"
		}
	}

	// 检查 Redis
	rdb := database.GetRedisClient()
	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := rdb.Ping(pingCtx).Err(); err != nil {
		checks["redis"] = "error: " + err.Error()
		healthy = false
	} else {
		checks["redis"] = "ok"
	}

	status := http.StatusOK
	if !healthy {
		status = http.StatusServiceUnavailable
	}
	ctx.JSON(status, gin.H{"status": healthy, "checks": checks})
}
