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
	GetMachines(context.Context) (machines []Machine, err error)
	IsEmptySlot(context.Context, uint, uint, string) (isEmpty bool)
	AddBooking(context.Context, Booking) (bookingId uint, err error)
	BookSlot(context.Context, Slot) (err error)
	GetBaseCharge(context.Context, uint) (baseCharge uint, err error)
	GenrateInvoice(context.Context, Invoice) (invoiceId uint, err error)
	GetBookedSlot(context.Context, uint) (map[uint]struct{}, error)
}
