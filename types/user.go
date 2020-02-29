package types

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//TempUser - the temp user for creating a new user
type TempUser struct {
	FirstName       string `json:"first_name" validate:"required,gte=4"`
	LastName        string `json:"last_name" validate:"required,gte=4"`
	Email           string `json:"email" xorm:"email" validate:"required,contains=@"`
	Password        string `json:"_" validate:"required,gte=8"`
	ConfirmPassword string `json:"_" validate:"required,gte=8"`
}

//User - the user in the system
type User struct {
	ID        int64  `json:"id" xorm:"'id' autoincr" schema:"id"`
	FirstName string `json:"first_name" xorm:"first_name" schema:"first_name" validate:"required,gte=4"`
	LastName  string `json:"last_name" xorm:"last_name" schema:"last_name" validate:"required,gte=4"`
	Email     string `json:"email" xorm:"email" schema:"email" validate:"required,contains=@"`
	Password  string `json:"_" xorm:"password" schema:"password" validate:"required,gte=8 "`
	Visible   bool   `json:"visible" xorm:"visible" schema:"visible"`
}

//Table name - the table when using xomr
func (u *User) TableName() string {
	return "user"
}

func NewUser(newUser *TempUser) (user *User, err error) {
	if newUser.Password != newUser.ConfirmPassword {
		err = fmt.Errorf("password and confirm password do not match")
		return nil, err
	}

	user = &User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Visible:   true,
	}

	if err := user.SetPassword(newUser.Password); err != nil {
		return nil, err
	}

	return user, nil
}

//Set password
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

//Authenticate a password against the stored hash
func (u *User) Authenticate(passwprd string) error {
	if !u.Visible {
		return fmt.Errorf("user is inactive")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwprd)); err != nil {
		return err
	}

	return nil
}
