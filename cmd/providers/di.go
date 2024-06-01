package providers

import (
	"github.com/Shepherd-Go/Nlj-Internal-Email-Service/internal/infra/api/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(router.New)

	return Container
}