package bootstrap

import (
	"context"
	"demo04/services/oauth/app/repository"
	"demo04/services/oauth/env"

	"github.com/RevenueMonster/sqlike/sqlike"
)

type Bootstrap struct {
	Mysql      *sqlike.Client
	Repository *repository.Repository
}

// New : Creates a new instance of Bootstrap and
// returns a pointer to this instance.
func New(ctx context.Context) *Bootstrap {
	bs := new(Bootstrap)

	bs.initMySQL(ctx)
	repo, err := repository.New(bs.Mysql.Database(env.Config.MySQL.Database))
	if err != nil {
		// handle error
		panic(err)
	}
	bs.Repository = repo

	return bs
}
func (bs *Bootstrap) Release() {

	if bs.Mysql != nil {
		bs.Mysql.Close()
	}
}
