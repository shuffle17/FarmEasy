package db

import (
	"context"
)

type Storer interface {

	//Create(context.Context, User) error
	//GetUser(context.Context) (User, error)
	//Delete(context.Context, string) error
	RegisterFarmer(context.Context, Farmer) (addedFarmer Farmer, err error)
	LoginFarmer(context.Context, string, string) (farmerId uint, err error)
	AddMachine(context.Context, Machine) (addedMachine Machine, err error)
}
