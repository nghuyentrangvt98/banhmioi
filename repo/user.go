package repo

import (
	"context"
	"time"

	"github.com/nghuyentrangvt98/banhmioi/config"
	"github.com/nghuyentrangvt98/banhmioi/ent"
	"github.com/nghuyentrangvt98/banhmioi/ent/user"
	"github.com/nghuyentrangvt98/banhmioi/model"
)

func GetUser(ctx context.Context, id int) *ent.User {
	user, _ := config.Client.User.Get(ctx, id)
	return user
}

func GetUserByUserName(ctx context.Context, username string) *ent.User {
	user, _ := config.Client.User.Query().Where(user.Username(username)).First(ctx)
	return user
}

func CreateUser(ctx context.Context, userCreate model.UserCreate) *ent.User {
	user := config.Client.User.Create().
		SetUsername(userCreate.Username).
		SetFullname(userCreate.Fullname).
		SetGender(user.Gender(userCreate.Gender)).
		SetDob(userCreate.Dob).
		SetEmail(userCreate.Email).
		SetPhone(userCreate.Phone).
		SetPassword(userCreate.Password).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SaveX(ctx)
	return user
}
