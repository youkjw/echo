package main

import (
	"echo/router"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	. "echo/conf"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${level}")

	router.StartRoutes(e)

	e.Logger.Fatal(e.Start(Conf.Server.Addr))
}
