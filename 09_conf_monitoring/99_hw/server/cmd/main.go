package main

import (
	"fmt"
	prom "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"server/internal/api/middleware"
	"server/internal/metrics"
	"server/internal/pkg/comment/handler"
	commentrepo "server/internal/pkg/comment/repository"
	commentsvc "server/internal/pkg/comment/service"
	"server/internal/pkg/session"
	threadhttp "server/internal/pkg/thread/handler"
	threadrepo "server/internal/pkg/thread/repository"
	threadsvc "server/internal/pkg/thread/service"
)

func main() {
	e := echo.New()
	m := metrics.NewMetrics()
	p := prom.NewPrometheus("echo", nil)

	sessionSvc := session.NewService()

	e.Use(m.MetricsMiddleware)
	e.Use(middleware.AuthEchoMiddleware(sessionSvc))

	threadRepo := threadrepo.NewRepository()
	threadSvc := threadsvc.NewService(threadRepo)
	threadHandler := threadhttp.Handler{ThreadSvc: threadSvc}

	commentRepo := commentrepo.NewRepository()
	commentSvc := commentsvc.NewService(commentRepo, threadRepo)
	commentHandler := handler.Handler{CommentSvc: commentSvc}

	e.GET("/thread/:id", threadHandler.GetThread)
	e.POST("/thread", threadHandler.CreateThread)
	e.POST("/thread/:tid/comment", commentHandler.Create)
	e.POST("/thread/:tid/comment/:cid/like", commentHandler.Like)

	p.Use(e)
	fmt.Print(e.Start("localhost:8080"))
}
