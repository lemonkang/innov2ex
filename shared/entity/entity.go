package entity

import "time"

const (
	TableUser       = "User"
	TableMasterUser = "MasterUser"
)

// Model :
type Model struct {
	CreatedDateTime time.Time `bson:"CreatedDateTime"`
	UpdatedDateTime time.Time `bson:"UpdatedDateTime"`
}
