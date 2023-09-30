package repo

import (
	"context"
	"log/slog"

	"github.com/nghuyentrangvt98/banhmioi/config"
	"github.com/nghuyentrangvt98/banhmioi/ent"
	"github.com/nghuyentrangvt98/banhmioi/ent/cart"
	"github.com/nghuyentrangvt98/banhmioi/ent/order"
	"github.com/nghuyentrangvt98/banhmioi/ent/user"
	"github.com/nghuyentrangvt98/banhmioi/model"
)

func CreateOrder(ctx context.Context, orderCreate model.OrderCreate, user_id int) *ent.Order {
	config.Client.Cart.Update().SetStatus(cart.StatusDone).Where(cart.IDIn(orderCreate.CartIds...)).SaveX(ctx)
	order := config.Client.Order.Create().
		SetAddress(orderCreate.Address).
		SetDiscount(orderCreate.Discount).
		SetNote(orderCreate.Note).
		SetPhone(orderCreate.Phone).
		SetTotal(orderCreate.Total).
		SetUserID(user_id).
		AddCartIDs(orderCreate.CartIds...).
		SaveX(ctx)
	return order
}

func GetMultiOrders(ctx context.Context, user_id int) model.OrderList {
	query := config.Client.Order.Query()
	query = query.Where(order.HasUserWith(user.ID(user_id)))
	query = query.Order(order.OrderOption(order.ByID()))
	orders, err := query.All(ctx)
	if err != nil {
		slog.Info(err.Error())
	}
	total := len(orders)
	return model.OrderList{Data: orders, Total: total}
}
