package entity

import "github.com/thedevsaddam/govalidator"

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (Category) Rules() govalidator.MapData {
	return govalidator.MapData{
		"name": []string{"required"},
	}
}
