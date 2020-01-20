package cache

import (
	"time"
	"github.com/labstack/echo"
	. "echo/conf"
	ec "echo/middleware/cache"
)

const (
	DefaultExpiration = 3600
	DEFAULT           = time.Duration(0)
	FOREVER           = time.Duration(-1)
	DefaultKey        = "echo/modules/cache"
)

func Cache() echo.MiddlewareFunc {
	var store ec.CacheStore

	switch Conf.CacheStore {
		case REDIS:
		default:
			store = ec.NewRedisCache(Conf.Redis.Server, Conf.Redis.Pwd, DefaultExpiration)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DefaultKey, store)

			return next(c)
		}
	}
}

// shortcut to get Cache
func Default(c echo.Context) ec.CacheStore {
	// return c.MustGet(DefaultKey).(ec.CacheStore)
	return c.Get(DefaultKey).(ec.CacheStore)
}
