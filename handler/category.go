package handler

import (
	"net/http"

	"github.com/bhill77/goshop/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) CategoryHandler {
	return CategoryHandler{
		db: db,
	}
}

func (h CategoryHandler) Index(c echo.Context) error {
	var categories []entity.Category

	h.db.Find(&categories)
	return c.JSON(http.StatusCreated, categories)
}

func (h CategoryHandler) Create(c echo.Context) error {
	var category entity.Category
	e := validateRequest(c, &category)
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}
	h.db.Create(&category)

	return c.JSON(http.StatusCreated, category)
}
