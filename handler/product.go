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

	e := validateRequest(c, &product)
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}

	h.db.Create(&product)
	h.db.Preload("Category").First(&product)

	return c.JSON(http.StatusCreated, product)
}

func (h ProductHandler) Index(c echo.Context) error {
	var products []entity.Product

	h.db.Preload("Category").Find(&products)

	return c.JSON(http.StatusCreated, products)
}

func (h ProductHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var product entity.Product
	h.db.First(&product, id)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "record not found")
	}

	var payload entity.Product
	e := validateRequest(c, &payload)
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}

	product.Name = payload.Name
	product.Description = payload.Description
	product.Price = payload.Price
	product.Stock = payload.Stock
	product.CategoryID = payload.CategoryID

	h.db.Updates(&product)

	return c.JSON(http.StatusOK, product)
}

func (h ProductHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	var product entity.Product
	h.db.First(&product, id)
	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, "record not found")
	}

	h.db.Delete(&product)
	return c.JSON(http.StatusOK, "product deleted")
}
