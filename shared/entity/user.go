package entity

import (
	"cloud.google.com/go/civil"
	"github.com/RevenueMonster/sqlike/types"
	"github.com/paulmach/orb"
)

type UserGender string

// Gender :
const (
	UserGenderMale   UserGender = "MALE"
	UserGenderFemale UserGender = "FEMALE"
)

type Address struct {
	ID        string
	Name      string
	Email     string
	PhoneNo   string
	Address   string
	PostCode  string
	StateCode string
}

// User :
type User struct {
	cinemaOperators []*types.Key
	Key             *types.Key `sqlike:"$Key"`
	MasterUserKey   *types.Key
	Name            string
	CallingCode     string `sqlike:",size=3"`
	PhoneNumber     string `sqlike:",size=20"`
	Email           string `sqlike:",size=200"`
	NRIC            string `sqlike:",size=20"`
	AvatarURL       string
	Gender          UserGender `sqlike:",enum=MALE|FEMALE"`
	BirthDate       *civil.Date
	Payload         struct {
		Address           string
		PostCode          string
		StateCode         string
		CountryCode       string
		SelectedAddressID *string
		Addresses         []Address
		Favour            struct {
			Cinemas []*types.Key
		}
		HasViewedMovieMoneyDisclaimer bool
		SessionDateTime               string
	}
	GeoPoint    orb.Point
	HasRegister bool
	// DeletedDateTime time.Time
	Model
}
