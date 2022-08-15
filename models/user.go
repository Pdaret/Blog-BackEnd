package models

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
	Phone     string `json:"phone"`
}

func (u *User) HashPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = hash
}

func IsEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func IsPhone(phone string) bool {
	Re := regexp.MustCompile(`^[0-9]{10}$`)
	return Re.MatchString(phone)
}

func (u *User) CompareHashAndPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		return false
	}
	return true
}
