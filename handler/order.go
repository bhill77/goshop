package handler

import (
	"net/http"

	"github.com/bhill77/goshop/entity"
	"github.com/bhill77/goshop/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		db: db,
	}
}

func (h *OrderHandler) Create(c echo.Context) error {
	var payload entity.OrderRequest
	// c.Bind(&payload)
	e := validateRequest(c, &payload)
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userID := claims.ID

	var total float32 = 0
	for _, item := range payload.Details {
		subTotal := float32(item.Quantity) * item.UnitPrice
		total = total + subTotal
	}

	order := entity.Order{
		UserID:  userID,
		Status:  "pending",
		Total:   total,
		Details: payload.Details,
	}

	h.db.Create(&order)

	return c.JSON(200, order)
}

func (h *OrderHandler) Index(c echo.Context) error {
	var orders []entity.Order
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userID := claims.ID

	h.db.Order("id desc").Where("user_id = ?", userID).Find(&orders)
	return c.JSON(http.StatusOK, orders)
}
