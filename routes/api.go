package routes

import (
	"github.com/bhill77/goshop/config"
	"github.com/bhill77/goshop/handler"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func SetupRoute(e *echo.Echo, db *gorm.DB) {
	conf := config.GetConfig()
	userHandler := handler.NewUserHandler(db, conf)
	e.GET("/user", userHandler.Index)
	e.POST("/user", userHandler.Create, echojwt.JWT([]byte(conf.JwtSecret)))
	e.PUT("/user/:id", userHandler.Update, echojwt.JWT([]byte(conf.JwtSecret)))
	e.GET("/user/:id", userHandler.Show)
	e.DELETE("/user/:id", userHandler.Delete, echojwt.JWT([]byte(conf.JwtSecret)))
	e.POST("/login", userHandler.Login)

	categoryHandler := handler.NewCategoryHandler(db)
	cat := e.Group("/category")
	cat.GET("", categoryHandler.Index)
	cat.POST("", categoryHandler.Create, echojwt.JWT([]byte(conf.JwtSecret)))

	productHandler := handler.NewProductHandler(db)
	product := e.Group("/product")
	product.GET("", productHandler.Index, middleware.RequestID())
	product.POST("", productHandler.Create, echojwt.JWT([]byte(conf.JwtSecret)))

}
