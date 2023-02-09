package service

import (
	"FarmEasy/db"
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("I'mGoingToBeAGolangDeveloper")

type Service interface {
	Register(ctx context.Context, farmer NewFarmer) (addedFarmer db.Farmer, err error)
	Login(ctx context.Context, fAuth LoginRequest) (token string, err error)
	AddMachine(ctx context.Context, machine NewMachine) (addedMachine db.Machine, err error)
	GetMachines(ctx context.Context) (machines []db.Machine, err error)
	BookMachine(ctx context.Context, booking NewBooking) (db.Invoice, error)
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

func ValidateJWT(tokenString string) (farmerId uint, err error) {
	tokenObject, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return secretKey, nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenObject.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	farmerId = uint(claims["user_id"].(float64))
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
func (s *FarmService) AddMachine(ctx context.Context, machine NewMachine) (addedMachine db.Machine, err error) {
	newMachine := db.Machine{
		Name:             machine.Name,
		Description:      machine.Description,
		BaseHourlyCharge: machine.BaseHourlyCharge,
		OwnerId:          machine.OwnerId,
	}
	addedMachine, err = s.store.AddMachine(ctx, newMachine)
	return
}

func (s *FarmService) GetMachines(ctx context.Context) (machines []db.Machine, err error) {
	machines, err = s.store.GetMachines(ctx)
	return
}

func (s *FarmService) BookMachine(ctx context.Context, booking NewBooking) (invoice db.Invoice, err error) {

	for _, slot := range booking.Slots {
		empty := s.store.IsEmptySlot(ctx, booking.MachineId, slot, booking.Date)
		if !empty {

			err = errors.New("slot not empty")
			return
		}
	}

	newBooking := db.Booking{
		MachineId: booking.MachineId,
		FarmerId:  booking.FarmerId,
	}
	newBooking.Id, err = s.store.AddBooking(ctx, newBooking)
	if err != nil {
		return
	}
	for _, slot := range booking.Slots {
		newSlot := db.Slot{
			BookingId: newBooking.Id,
			SlotId:    slot,
			Date:      booking.Date,
		}
		err = s.store.BookSlot(ctx, newSlot)
		if err != nil {
			return
		}
	}
	baseCharge, err := s.store.GetBaseCharge(ctx, booking.MachineId)
	if err != nil {
		return
	}
	totalAmount := uint(len(booking.Slots)) * baseCharge
	newInvoice := db.Invoice{
		BookingId:    newBooking.Id,
		DateGenrated: time.Now().Format("2006-01-02"),
		Amount:       totalAmount,
	}
	newInvoice.Id, err = s.store.GenrateInvoice(ctx, newInvoice)

	invoice = newInvoice

	return
}
