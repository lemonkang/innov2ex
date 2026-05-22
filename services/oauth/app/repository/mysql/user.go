package mysql

import (
	"context"
	"demo04/shared/entity"

	"github.com/RevenueMonster/sqlike/sql/expr"
	"github.com/RevenueMonster/sqlike/types"

	"github.com/RevenueMonster/sqlike/sqlike/actions"
	"github.com/RevenueMonster/sqlike/sqlike/options"
)

// FindUser :
func (r *MySQL) FindUser(
	ctx context.Context,
	userKey *types.Key,
) (user entity.User, err error) {
	err = r.db.Table(entity.TableUser).FindOne(
		ctx,
		actions.FindOne().
			Where(
				expr.Equal("$Key", userKey),
			),
		options.FindOne().SetDebug(true),
	).Decode(&user)
	return
}
func (r MySQL) FindMasterUserByEmail(
	ctx context.Context,
	email string,
) (user entity.MasterUser, err error) {
	err = r.db.Table(entity.TableMasterUser).FindOne(
		ctx,
		actions.FindOne().Where(
			expr.Equal("Email", email),
		),
	).Decode(&user)
	return
}
