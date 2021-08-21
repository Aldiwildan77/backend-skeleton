package registry

import "github.com/labstack/echo"

type RouteProto interface {
	Register(e *echo.Echo)
}

var _routers []RouterFactory

type RouterFactory func() RouteProto

func RegisterRouter(router RouterFactory) {
	_routers = append(_routers, router)
}

func LoadRouter() []RouterFactory {
	return _routers
}
