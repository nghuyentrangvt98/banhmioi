package model

import (
	"github.com/nghuyentrangvt98/banhmioi/ent"
	"github.com/nghuyentrangvt98/banhmioi/ent/cart"
)

type CartCreate struct {
	Variation map[string]interface{} `json:"variation,omitempty"`
	Qty       int32                  `json:"qty,omitempty"`
	Price     int64                  `json:"price,omitempty"`
	Note      string                 `json:"note,omitempty"`
	ProductID int                    `json:"product_id,omitempty"`
	UserID    int                    `json:"user_id,omitempty"`
}

type Cart struct {
	ID        int                    `json:"id,omitempty"`
	Variation map[string]interface{} `json:"variation,omitempty"`
	Qty       int32                  `json:"qty,omitempty"`
	Price     int64                  `json:"price,omitempty"`
	Status    cart.Status            `json:"status,omitempty"`
	Note      string                 `json:"note,omitempty"`
	Product   Product                `json:"product,omitempty"`
	User      UserResponse           `json:"user,omitempty"`
}

type CartList struct {
	Data  []*ent.Cart `json:"data,omitempty"`
	Total int         `json:"total,omitempty"`
}
