package bootstrap

import (
	"context"
	"demo04/services/oauth/env"

	"github.com/RevenueMonster/sqlike/sqlike"
	"github.com/RevenueMonster/sqlike/sqlike/options"
	_ "github.com/go-sql-driver/mysql"
)

func (bs *Bootstrap) initMySQL(ctx context.Context) *Bootstrap {
	var (
		client *sqlike.Client
		driver = "mysql"
	)
	client = sqlike.MustConnect(
		ctx,
		driver,
		options.Connect().
			SetUsername(env.Config.MySQL.Username).
			SetPassword(env.Config.MySQL.Password).
			SetHost(env.Config.MySQL.Host).
			SetPort(env.Config.MySQL.Port),
	)
	bs.Mysql = client

	return bs

}
