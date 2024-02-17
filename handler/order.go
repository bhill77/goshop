package handler

import (
	"errors"
	"fmt"
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
	var err error
	for i, item := range payload.Details {
		var product entity.Product
		h.db.First(&product, item.ProductID)
		if product.ID == 0 {
			err = fmt.Errorf("invalid product id: %d", item.ProductID)
			break
		}
		product.Stock = product.Stock - item.Quantity
		h.db.Save(&product)

		subTotal := float32(item.Quantity) * product.Price
		total = total + subTotal

		payload.Details[i].UnitPrice = product.Price
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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

	h.db.Preload("Details").Order("id desc").Where("user_id = ?", userID).Find(&orders)
	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) Update(c echo.Context) error {
	id := c.Param("id")

	var order entity.Order
	err := h.db.First(&order, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "Invalid id")
	}

	var payload = map[string]string{}
	c.Bind(&payload)

	status, ok := payload["status"]
	if !ok {
		return c.JSON(http.StatusBadRequest, "status is required")
	}
	order.Status = status
	h.db.Save(&order)
	return c.JSON(http.StatusOK, order)
}
