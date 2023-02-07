package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
)

type Farmer struct {
	FirstName string `db:"fname" json:"fname"`
	LastName  string `db:"lname" json:"lname"`
	Email     string `db:"email" json:"email"`
	Phone     string `db:"phone" json:"phone"`
	Address   string `db:"address" json:"address"`
	Password  string `db:"password" json:"password"`
}

// func (s *pgStore) ListUsers(ctx context.Context) (users []User, err error) {
// 	err = s.db.Select(&users, "SELECT * FROM users ORDER BY name ASC")
// 	if err != nil {
// 		logger.WithField("err", err.Error()).Error("Error listing users")
// 		return
// 	}

// 	return
// }

func (s *pgStore) RegisterFarmer(ctx context.Context, farmer Farmer) (err error) {
	_, err = s.db.NamedExecContext(ctx, "INSERT INTO farmers (fname, lname, email, phone, address, password) VALUES (:fname, :lname, :email, :phone, :address, :password)", farmer)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting farmer")
		return
	}

	return
}
