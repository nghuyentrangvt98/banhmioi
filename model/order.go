package model

import (
	"github.com/nghuyentrangvt98/banhmioi/ent"
)

type OrderCreate struct {
	Name     string  `json:"name,omitempty"`
	Phone    string  `json:"phone,omitempty"`
	Address  string  `json:"address,omitempty"`
	Note     string  `json:"note,omitempty"`
	Total    float32 `json:"total,omitempty"`
	Discount float32 `json:"discount,omitempty"`
	CartIds  []int   `json:"cart_ids,omitempty"`
}

type Order struct {
	Name     string  `json:"name,omitempty"`
	Phone    string  `json:"phone,omitempty"`
	Address  string  `json:"address,omitempty"`
	Note     string  `json:"note,omitempty"`
	Total    float32 `json:"total,omitempty"`
	Discount float32 `json:"discount,omitempty"`
	Carts    []Cart  `json:"carts,omitempty"`
}

type OrderList struct {
	Data  []*ent.Order `json:"data,omitempty"`
	Total int          `json:"total,omitempty"`
}
