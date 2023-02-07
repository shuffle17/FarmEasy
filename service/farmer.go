package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"
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

	return
}
