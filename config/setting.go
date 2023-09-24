package config

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/nghuyentrangvt98/banhmioi/ent"
)

var Client *ent.Client
var err error
var SECRETKEY []byte

// postgres://eitaraui:D3ubi18Q5ghOP7M5z2EpCe7m_kr4hEZy@tiny.db.elephantsql.com/eitaraui
func init() {
	Client, err = ent.Open("postgres", "host=tiny.db.elephantsql.com port=5432 user=eitaraui dbname=eitaraui password=D3ubi18Q5ghOP7M5z2EpCe7m_kr4hEZy sslmode=disable", ent.Debug())
	if err != nil {
		panic(err)
	}
	SECRETKEY = []byte("HelloWorld")
	Client.Schema.Create(context.Background())
}
