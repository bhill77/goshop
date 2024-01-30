package routes

import (
	"github.com/bhill77/goshop/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoute(e *echo.Echo, db *gorm.DB) {
	userHandler := handler.NewUserHandler(db)
	e.GET("/user", userHandler.Index)
	e.POST("/user", userHandler.Create)

}
