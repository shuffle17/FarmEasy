package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"

	"github.com/sirupsen/logrus"
)

func ValidateFarmerPhone(phone string) (err error) {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	if !re.MatchString(phone) {
		err = errors.New("invalid phone number")
	}
	return
}

func ValidateFarmerEmail(email string) (err error) {
	re := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	if !re.MatchString(email) {
		err = errors.New("invalid email")
	}
	return
}

func Hash_password(password string) (hash string) {

	hsha := sha256.New()
	hsha.Write([]byte(password))
	hash = base64.URLEncoding.EncodeToString(hsha.Sum(nil))
	logrus.Info(password, " -> ", hash)
	return
}

func ValidateBookingslots(slots []uint) (err error) {
	if len(slots) == 0 {
		err = errors.New("no slots selected")
	}
	for _, v := range slots {
		if v > 24 || v < 1 {
			err = errors.New("invalid slot selected")
		}
	}
	return
}

func ValidateBookingDate(date string) (err error) {
	re := regexp.MustCompile(`^([0-9]{4})-([0-9]{2})-([0-9]{2})$`)
	if !re.MatchString(date) {
		err = errors.New("invalid date")
	}
	return
}
