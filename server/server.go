package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercuryheart123/lmwn-backend-go/covid"
	"github.com/mercuryheart123/lmwn-backend-go/data"
	"go.uber.org/zap"
)

type GinServer struct {
	*gin.Engine
	logger *zap.Logger
	port   string
}

func InitServer(port string, logger *zap.Logger) *GinServer {
	return &GinServer{Engine: gin.Default(), logger: logger, port: port}
}

func (r *GinServer) InitRoutes() {
	r.GET("/covid/summary", func(ctx *gin.Context) {
		r.logger.Info("Get covid summary")
		data, err := data.GetData()
		if err != nil {
			r.logger.Error("Failed to get data", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		summarize, err := covid.GetSummarize(data)
		if err != nil {
			r.logger.Error("Failed to get summarize", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, summarize)

	})
}

func (r *GinServer) RunServer() {
	srv := &http.Server{
		Addr:    ":" + r.port,
		Handler: r,
	}
	r.logger.Info("Server is running on port: " + r.port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		r.logger.Fatal("Failed to listen and serve", zap.Error(err))
	}
}
