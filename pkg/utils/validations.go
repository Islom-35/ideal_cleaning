package utils

import (
	"example.com/m/internal/user/errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

func ValidateUserInfoForSignUp(fullName, email, phone, password string) error {

	if strings.TrimSpace(fullName) == "" {
		return errors.ErrEmptyName
	}
	if ValidateEmail(email) == errors.ErrInvalidEmailFormat {
		return errors.ErrEmptyMail
	}
	if !isValidPhoneNumber(phone) {
		return errors.ErrInvalidPhoneNumber
	}
	if ValidatePassword(password) == errors.ErrInvalidPassword {
		return errors.ErrInvalidPassword
	}
	return nil
}

func ValidateUserInfoForSignIn(email, password string) error {
	if strings.TrimSpace(email) == "" {
		return errors.ErrBadCredentials
	}
	if ValidatePassword(password) == errors.ErrInvalidPassword {
		return errors.ErrBadCredentials
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.ErrInvalidPassword
	}

	var (
		upperCase = regexp.MustCompile(`[A-Z]`)
		lowerCase = regexp.MustCompile(`[a-z]`)
		digit     = regexp.MustCompile(`[0-9]`)
	)
	if !upperCase.MatchString(password) || !lowerCase.MatchString(password) {
		return errors.ErrInvalidPassword
	}
	if !digit.MatchString(password) {
		return errors.ErrInvalidPassword
	}
	return nil
}

func isValidPhoneNumber(phone string) bool {
	if len(phone) < 12 {
		fmt.Sprint("error: phone number count is more than 12")
		return false
	}
	phoneRegex := `^\+?[1-9]\d{1,14}$|^[1-9]\d{0,14}$|^[1-9]\d{0,14}-[1-9]\d{0,14}$`

	regExp := regexp.MustCompile(phoneRegex)

	return regExp.MatchString(phone)
}

func ValidateEmail(email string) error {
	emailReg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailReg.MatchString(email) {
		return errors.ErrInvalidEmailFormat
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
