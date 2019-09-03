package web

import (
	"github.com/labstack/echo"
)

func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	e.Logger.SetPrefix("web")

	// Routers
	e.GET("/home", handler(HomeHandler))

	return e
}

type (
	HandlerFunc func(*echo.Context) error
)

/**
 * 自定义Context的Handler
 */
func handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := &c
		return h(ctx)
	}
}