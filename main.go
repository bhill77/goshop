package main

import (
	"net/http"

	"github.com/bhill77/goshop/config"
	"github.com/bhill77/goshop/database"
	"github.com/bhill77/goshop/entity"
	"github.com/bhill77/goshop/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.GetConfig()
	db := database.NewConnection(conf)

	db.AutoMigrate(
		entity.User{},
		entity.Category{},
		entity.Product{},
		entity.Order{},
		entity.OrderDetail{},
		entity.Role{},
	)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.SetupRoute(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}
