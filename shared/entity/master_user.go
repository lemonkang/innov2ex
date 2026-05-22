package entity

import (
	"log"
	"time"

	"github.com/RevenueMonster/sqlike/types"
	"golang.org/x/crypto/bcrypt"
)

// MasterUser :
type MasterUser struct {
	Key               *types.Key `sqlike:"$Key"`
	ID                string     `sqlike:",primary_key"`
	Name              string
	Email             string `sqlike:",size=200,unique"`
	PhoneNumber       string
	PasswordHash      []byte
	PasswordResetCode *string
	VerifyDateTime    *time.Time
	IsVerified        bool
	TIN               string
	SST               string
	IDType            string
	IDValue           string
	Address           struct {
		Line1       string `json:"line1"`
		Line2       string `json:"line2"`
		Line3       string `json:"line3"`
		City        string `json:"city"`
		State       string `json:"state"`
		Postcode    string `json:"postcode"`
		CountryCode string `json:"countryCode"` // MYS
	} `json:"address"`
	Model
}

// VerifyPassword :
func (u MasterUser) VerifyPassword(pwd string) bool {
	log.Printf("正在验证密码，输入密码：%s，存储的密码哈希：%s\n", pwd, string(u.PasswordHash))
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
