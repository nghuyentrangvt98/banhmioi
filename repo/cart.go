package repo

import (
	"context"
	"log/slog"

	"github.com/nghuyentrangvt98/banhmioi/config"
	"github.com/nghuyentrangvt98/banhmioi/ent"
	"github.com/nghuyentrangvt98/banhmioi/ent/cart"
	"github.com/nghuyentrangvt98/banhmioi/ent/user"
	"github.com/nghuyentrangvt98/banhmioi/model"
)

func CreateCart(ctx context.Context, cartCreate model.CartCreate) *ent.Cart {
	cart := config.Client.Cart.Create().
		SetProductID(cartCreate.ProductID).
		SetUserID(cartCreate.UserID).
		SetVariation(cartCreate.Variation).
		SetQty(cartCreate.Qty).
		SetPrice(cartCreate.Price).
		SetStatus(cart.Status("processing")).
		SetNote(cartCreate.Note).
		SaveX(ctx)
	return cart
}

func GetMultiCarts(ctx context.Context, user_id int) model.CartList {
	query := config.Client.Cart.Query()
	query = query.Where(cart.StatusEQ(cart.StatusProcessing))
	query = query.Where(cart.HasUserWith(user.ID(user_id)))
	query = query.Order(cart.OrderOption(cart.ByID()))
	carts, err := query.All(ctx)
	if err != nil {
		slog.Info(err.Error())
	}
	total := len(carts)
	return model.CartList{Data: carts, Total: total}
}
