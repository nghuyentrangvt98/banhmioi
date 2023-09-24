package repo

import (
	"context"
	"log/slog"

	"github.com/nghuyentrangvt98/banhmioi/config"
	"github.com/nghuyentrangvt98/banhmioi/ent"
	"github.com/nghuyentrangvt98/banhmioi/ent/product"
	"github.com/nghuyentrangvt98/banhmioi/model"
)

func GetProduct(ctx context.Context, id int) *ent.Product {
	product, _ := config.Client.Product.Get(ctx, id)
	return product
}

func GetMultiProduct(ctx context.Context, productQuery model.ProductQuery) model.ProductList {
	query := config.Client.Product.Query()
	if productQuery.Category != "" {
		query = query.Where(product.Category(productQuery.Category))
	}
	if productQuery.Keyword != "" {
		query = query.Where(product.Or(
			product.NameContains(productQuery.Keyword),
			product.DescriptionContains(productQuery.Keyword)))
	}
	query = query.Order(product.OrderOption(product.ByID()))
	products, err := query.All(ctx)
	if err != nil {
		slog.Info(err.Error())
	}
	total := len(products)
	return model.ProductList{Data: products, Total: total}
}
