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

func (s *pgStore) RegisterFarmer(ctx context.Context, farmer Farmer) (addedFarmer Farmer, err error) {
	err = s.db.QueryRowContext(ctx, "INSERT INTO farmers (fname, lname, email, phone, address, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", farmer.FirstName, farmer.LastName, farmer.Email, farmer.Phone, farmer.Address, farmer.Password).Scan(&farmer.Id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting farmer")
		return
	}
	// id, err := res.LastInsertId()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting farmer")
		return
	}
	addedFarmer = farmer

	return
}