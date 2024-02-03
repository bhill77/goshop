package handler

import (
	"net/http"

	"github.com/bhill77/goshop/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductHandler struct {
	db *gorm.DB
}

func NewProductHandler(db *gorm.DB) ProductHandler {
	return ProductHandler{
		db: db,
	}
}

func (h ProductHandler) Create(c echo.Context) error {
	var product entity.Product
	c.Bind(&product)

	h.db.Create(&product)

	return c.JSON(http.StatusCreated, product)
}

func (h ProductHandler) Index(c echo.Context) error {
	var products []entity.Product

	h.db.Preload("Category").Find(&products)

	return c.JSON(http.StatusCreated, products)
}
