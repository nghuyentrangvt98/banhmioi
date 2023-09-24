package model

import "github.com/nghuyentrangvt98/banhmioi/ent"

type Product struct {
	ID              int                    `json:"id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Description     string                 `json:"description,omitempty"`
	ListPrice       int64                  `json:"list_price,omitempty"`
	SalePrice       int64                  `json:"sale_price,omitempty"`
	Variation       map[string]interface{} `json:"variation,omitempty"`
	StockQty        int32                  `json:"stock_qty,omitempty"`
	ProductURL      string                 `json:"product_url,omitempty"`
	MeasurementUnit string                 `json:"measurement_unit,omitempty"`
}

type ProductList struct {
	Data  []*ent.Product `json:"data,omitempty"`
	Total int            `json:"total,omitempty"`
}

type ProductQuery struct {
	Category string `query:"category"`
	Keyword  string `query:"keyword"`
}
