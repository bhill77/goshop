package routes

import (
	"github.com/bhill77/goshop/config"
	"github.com/bhill77/goshop/handler"
	appMiddleware "github.com/bhill77/goshop/middleware"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoute(e *echo.Echo, db *gorm.DB) {
	conf := config.GetConfig()

	// jwt config
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(appMiddleware.JwtCustomClaims)
		},
		SigningKey: []byte(conf.JwtSecret),
	}

	userHandler := handler.NewUserHandler(db, conf)
	e.GET("/user", userHandler.Index)
	e.POST("/user", userHandler.Create, echojwt.WithConfig(jwtConfig))
	e.PUT("/user/:id", userHandler.Update, echojwt.WithConfig(jwtConfig))
	e.GET("/user/:id", userHandler.Show)
	e.DELETE("/user/:id", userHandler.Delete, echojwt.WithConfig(jwtConfig))
	e.POST("/login", userHandler.Login)

	categoryHandler := handler.NewCategoryHandler(db)
	cat := e.Group("/category")
	cat.GET("", categoryHandler.Index)
	cat.POST("", categoryHandler.Create, echojwt.WithConfig(jwtConfig))

	productHandler := handler.NewProductHandler(db)
	product := e.Group("/product")
	product.GET("", productHandler.Index)
	product.POST("", productHandler.Create, echojwt.WithConfig(jwtConfig))

	orderHandler := handler.NewOrderHandler(db)
	order := e.Group("order", echojwt.WithConfig(jwtConfig))
	order.GET("", orderHandler.Index)
	order.POST("", orderHandler.Create)
	order.PUT("/:id", orderHandler.Update)

}
