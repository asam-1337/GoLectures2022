package handler

import (
	"server/internal/metrics"
	"server/internal/pkg/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	ThreadSvc domain.ThreadService
}

func (h Handler) GetThread(ctx echo.Context) error {
	var err error
	tid := ctx.Param("tid")

	m := ctx.Get(metrics.KeyMetrics).(*metrics.Metrics)

	defer func() {
		status := "200"
		if err != nil {
			status = "500"
		}
		m.Hits.WithLabelValues(status, ctx.Path()).Inc()
	}()

	t, err := h.ThreadSvc.Get(tid)
	if err != nil {
		return err
	}

	m.Hits.WithLabelValues()
	return ctx.JSON(200, t)
}

func (h Handler) CreateThread(ctx echo.Context) error {
	var thread domain.Thread

	err := ctx.Bind(&thread)
	if err != nil {
		return err
	}

	err = h.ThreadSvc.Create(thread)
	if err != nil {
		return err
	}

	return ctx.NoContent(200)
}
