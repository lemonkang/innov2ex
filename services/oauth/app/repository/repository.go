package repository

import (
	"demo04/services/oauth/app/repository/mysql"
	"errors"

	"github.com/RevenueMonster/sqlike/sqlike"
)

type Repository struct {
	*mysql.MySQL
}

func New(
	db *sqlike.Database,
) (*Repository, error) {
	if db == nil {
		return nil, errors.New("missing db")
	}
	repo := new(Repository)
	repo.MySQL = mysql.New(db)
	return repo, nil
}
