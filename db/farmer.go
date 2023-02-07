package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
)

type Farmer struct {
	Id        uint   `db:"id" json:"id"`
	FirstName string `db:"fname" json:"fname"`
	LastName  string `db:"lname" json:"lname"`
	Email     string `db:"email" json:"email"`
	Phone     string `db:"phone" json:"phone"`
	Address   string `db:"address" json:"address"`
	Password  string `db:"password" json:"-"`
}

type Machine struct {
	Id               uint   `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	Description      string `db:"description" json:"description"`
	BaseHourlyCharge uint   `db:"base_hourly_charge" json:"base_hourly_charge"`
	OwnerId          uint   `db:"owner_id" json:"owner_id"`
}

func (s *pgStore) RegisterFarmer(ctx context.Context, farmer Farmer) (addedFarmer Farmer, err error) {
	err = s.db.QueryRowContext(ctx, "INSERT INTO farmers (fname, lname, email, phone, address, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", farmer.FirstName, farmer.LastName, farmer.Email, farmer.Phone, farmer.Address, farmer.Password).Scan(&farmer.Id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting farmer")
		return
	}

	addedFarmer = farmer

	return
}

func (s *pgStore) LoginFarmer(ctx context.Context, email string, password string) (farmerId uint, err error) {

	err = s.db.QueryRowContext(ctx, "SELECT id FROM farmers WHERE email = $1 and password = $2", email, password).Scan(&farmerId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error incorrect email or password")
		return
	}

	return
}

func (s *pgStore) AddMachine(ctx context.Context, machine Machine) (addedMachine Machine, err error) {

	err = s.db.QueryRowContext(ctx, "INSERT INTO machines (name, description, base_hourly_charge, owner_id) VALUES ($1, $2, $3, $4) RETURNING id", machine.Name, machine.Description, machine.BaseHourlyCharge, machine.OwnerId).Scan(&machine.Id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting machine")
		return
	}

	addedMachine = machine

	return

}
