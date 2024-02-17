package entity

import (
	"time"

	"github.com/thedevsaddam/govalidator"
)

type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	RoleID      int       `json:"role_id"`
	CreatedAt   time.Time `json:"created_at"`
	Role        Role      `json:"role"`
}

func (User) Rules() govalidator.MapData {
	return govalidator.MapData{
		"name":     []string{"required"},
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}
}
