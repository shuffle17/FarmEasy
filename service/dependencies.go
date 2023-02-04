package service

import "FarmEasy/db"

type Dependencies struct {
	Store db.Storer
	// define other service dependencies
}
