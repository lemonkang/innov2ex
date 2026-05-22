package mysql

import (
	"github.com/RevenueMonster/sqlike/sqlike"
)

type MySQL struct {
	db *sqlike.Database
}

func New(db *sqlike.Database) *MySQL {
	return &MySQL{db: db}
}
