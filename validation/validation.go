package validation

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode"

	"github.com/rs401/myauth/pb"
)

var (
	ErrEmptyName       = errors.New("name cannot be empty")
	ErrEmptyEmail      = errors.New("email cannot be empty")
	ErrEmptyPassword   = errors.New("password cannot be empty")
	ErrInvalidEmail    = errors.New("email not valid")
	ErrEmailExists     = errors.New("email already exists")
	ErrNameExists      = errors.New("name already exists")
	ErrInvalidPassword error

	MaxPwLen int = 50
	MinPwLen int = 8
)

func IsValidSignUp(user *pb.User) error {
	if IsEmptyString(user.Name) {
		return ErrEmptyName
	}
	if IsEmptyString(user.Email) {
		return ErrEmptyEmail
	}
	if IsEmptyString(user.Password) {
		return ErrEmptyPassword
	}
	if !IsValidEmail(user.Email) {
		return ErrInvalidEmail
	}
	if !IsValidPassword(user.Password) {
		return ErrInvalidPassword
	}

	return nil
}

func IsEmptyString(in string) bool {
	return strings.TrimSpace(in) == ""
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPassword(s string) bool {
	var (
		isMin   bool
		special bool
		number  bool
		upper   bool
		lower   bool
		errStr  string
	)

	// Check length
	if len(s) < MinPwLen || len(s) > MaxPwLen {
		isMin = false
		errStr += fmt.Sprintf("password length must be between %d and %d, ", MinPwLen, MaxPwLen)
	}

	// Check other requirements
	for _, c := range s {
		if special && number && upper && lower && isMin {
			break
		}

		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	// Append error messages
	if !special {
		errStr += "should contain at least a single special character, "
	}
	if !number {
		errStr += "should contain at least a single digit, "
	}
	if !lower {
		errStr += "should contain at least a single lowercase letter, "
	}
	if !upper {
		errStr += "should contain at least single uppercase letter, "
	}

	// If there are any errors
	if len(errStr) > 0 {
		ErrInvalidPassword = errors.New(errStr)
		return false
	}

	// No errors
	return true
}
