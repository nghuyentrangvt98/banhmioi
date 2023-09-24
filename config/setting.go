package config

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/nghuyentrangvt98/banhmioi/ent"
)

var Client *ent.Client
var err error
var SECRETKEY []byte

func init() {
	Client, err = ent.Open("postgres", "host=localhost port=5432 user=banhmioi dbname=banhmioi password=postgres sslmode=disable", ent.Debug())
	if err != nil {
		panic(err)
	}
	SECRETKEY = []byte("HelloWorld")
	Client.Schema.Create(context.Background())
}
