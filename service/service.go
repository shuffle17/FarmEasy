package service

import (
	"FarmEasy/db"
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("I'mGoingToBeAGolangDeveloper")

type Service interface {
	Register(ctx context.Context, farmer NewFarmer) (addedFarmer db.Farmer, err error)
	Login(ctx context.Context, fAuth LoginRequest) (token string, err error)
}

type FarmService struct {
	store db.Storer
}

func NewFarmService(s db.Storer) Service {
	return &FarmService{
		store: s,
	}
}

func (s *FarmService) Register(ctx context.Context, farmer NewFarmer) (addedFarmer db.Farmer, err error) {

	newFarmer := db.Farmer{
		FirstName: farmer.FirstName,
		LastName:  farmer.LastName,
		Email:     farmer.Email,
		Phone:     farmer.Phone,
		Address:   farmer.Address,
		Password:  farmer.Password,
	}

	newFarmer.Password = Hash_password(newFarmer.Password)

	addedFarmer, err = s.store.RegisterFarmer(ctx, newFarmer)
	return
}
func generateJWT(farmerId uint) (token string, err error) {
	tokenExpirationTime := time.Now().Add(time.Hour * 24)
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": farmerId,
		"exp":     tokenExpirationTime.Unix(),
	})
	token, err = tokenObject.SignedString(secretKey)
	return
}

func (s *FarmService) Login(ctx context.Context, fAuth LoginRequest) (token string, err error) {
	var farmerId uint
	fAuth.Password = Hash_password(fAuth.Password)
	farmerId, err = s.store.LoginFarmer(ctx, fAuth.Email, fAuth.Password)
	if err != nil {
		return
	}
	token, err = generateJWT(farmerId)
	if err != nil {
		return
	}
	return
}
