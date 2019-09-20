package router

import (
	"echo/router/web"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	. "echo/conf"
	"github.com/labstack/gommon/log"
)

var (
	confFilePath string
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	// Hosts
	hosts := make(map[string]*Host)

	hosts[Conf.Server.DomainWeb] = &Host{web.Routers()}

	return hosts
}

func StartRoutes(e *echo.Echo) {

	// 配置初始化
	if err := InitConfig(confFilePath); err != nil {
		log.Panic(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo!")
	})

	hosts := InitRoutes()

	e.Any("/*", func(c echo.Context) (err error) {

		req := c.Request()
		res := c.Response()

		u, _err := url.Parse(c.Scheme() + "://" + req.Host)
		if _err != nil {
			e.Logger.Errorf("Request URL parse error:%v", _err)
		}

		host := hosts[u.Hostname()]

		if host == nil {
			e.Logger.Info("Host not found")
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

}
