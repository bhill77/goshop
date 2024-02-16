package entity

import (
	"time"

	"github.com/thedevsaddam/govalidator"
)

type Order struct {
	ID        int           `json:"id"`
	UserID    int           `json:"user_id"`
	Total     float32       `json:"total"`
	Status    string        `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Details   []OrderDetail `json:"details"`
}

type OrderDetail struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float32 `json:"unit_price"`
}

type OrderRequest struct {
	Details []OrderDetail `json:"details"`
}

func (OrderRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"details": []string{"required"},
	}
}
