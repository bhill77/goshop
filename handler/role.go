package handler

import (
	"net/http"

	"github.com/bhill77/goshop/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoleHandler struct {
	db *gorm.DB
}

func NewRoleHandler(db *gorm.DB) *RoleHandler {
	return &RoleHandler{
		db: db,
	}
}

func (h *RoleHandler) Index(c echo.Context) error {
	var roles []entity.Role
	h.db.Find(&roles)
	return c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) Create(c echo.Context) error {
	var role entity.Role
	c.Bind(&role)

	h.db.Create(&role)
	return c.JSON(http.StatusOK, role)
}
