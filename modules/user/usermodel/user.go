package usermodel

import (
	"errors"
	"golang-blog-api/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (data *User) Mask(isAdmin bool) {
	data.GenUID(common.DBTypeUser)
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
	ErrEmailCannotBeEmpty     = common.NewCustomError(nil, "email can't be blank", "ErrEmailCannotBeEmpty")
	ErrEmailExisted           = common.NewCustomError(nil, "email has already existed", "ErrEmailExisted")
	ErrInvalidEmail           = common.NewCustomError(nil, "invalid email", "ErrInvalidEmail")
	ErrEmailIsTooLong         = common.NewCustomError(nil, "email is too long", "ErrEmailIsTooLong")
	ErrEmailOrPasswordInvalid = common.NewCustomError(nil, "email or password invalid", "ErrEmailOrPasswordInvalid")

	ErrInvalidPassword       = common.NewCustomError(nil, "invalid password", "ErrInvalidPassword")
	ErrPasswordCannotBeEmpty = common.NewCustomError(nil, "password can't be blank", "ErrPasswordCannotBeEmpty")

	ErrLastNameCannotBeEmpty = common.NewCustomError(nil, "last name can't be blank", "ErrLastNameCannotBeEmpty")
	ErrLastNameIsTooLong     = common.NewCustomError(nil, "last name is too long", "ErrLastNameIsTooLong")

	ErrFirstNameCannotBeEmpty = common.NewCustomError(nil, "first name can't be blank", "ErrFirstNameCannotBeEmpty")
	ErrFirstNameIsTooLong     = common.NewCustomError(nil, "first name is too long", "ErrFirstNameIsTooLong")

	ErrAvatarCannotBeEmpty = common.NewCustomError(nil, "avatar can't be blank", "ErrAvatarCannotBeEmpty")
)
