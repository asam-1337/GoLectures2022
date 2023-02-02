package middleware

import (
	"log"
	"runtime/debug"
	"server/internal/pkg/domain"
	"server/internal/pkg/errors"

	"github.com/labstack/echo/v4"
)

func AuthEchoMiddleware(service domain.SessionService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			_, err := service.CheckSession(context.Request().Header)
			if err != nil {
				return context.NoContent(401)
			}

			return next(context)
		}
	}
}

func PanicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				errors.HttpErrorResponse(500, "something went wrong")
				log.Panic(debug.Stack())
			}
		}()
		return next(c)
	}
}
