package service

import (
	"FarmEasy/db"
	"context"
)

type Service interface {
	Register(ctx context.Context, farmer db.Farmer) (addedFarmer db.Farmer, err error)
}

type FarmService struct {
	store db.Storer
}

func NewFarmService(s db.Storer) Service {
	return &FarmService{
		store: s,
	}
}

func (s *FarmService) Register(ctx context.Context, farmer db.Farmer) (addedFarmer db.Farmer, err error) {
	farmer.Password = Hash_password(farmer.Password)
	addedFarmer, err = s.store.RegisterFarmer(ctx, farmer)
	return
}
