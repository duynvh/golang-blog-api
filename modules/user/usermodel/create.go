package usermodel

import (
	"golang-blog-api/common"
	"strings"
)

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (data *UserCreate) Mask(isAdmin bool) {
	data.GenUID(common.DBTypeUser)
}

func (res *UserCreate) Validate() error {
	res.Email = strings.TrimSpace(res.Email)
	res.Password = strings.TrimSpace(res.Password)
	res.FirstName = strings.TrimSpace(res.FirstName)
	res.LastName = strings.TrimSpace(res.LastName)

	if len(res.Email) == 0 {
		return ErrEmailCannotBeEmpty
	}
	// if err := checkmail.ValidateFormat(res.Email); err != nil {
	// 	return ErrInvalidEmail
	// }
	if len(res.Password) == 0 {
		return ErrPasswordCannotBeEmpty
	}
	if len(res.Password) < 6 || len(res.Password) > 50 {
		return ErrInvalidPassword
	}
	if len(res.FirstName) == 0 {
		return ErrFirstNameCannotBeEmpty
	}
	if len(res.FirstName) > 200 {
		return ErrFirstNameIsTooLong
	}
	if len(res.LastName) == 0 {
		return ErrLastNameCannotBeEmpty
	}
	if len(res.LastName) > 200 {
		return ErrLastNameIsTooLong
	}
	// if res.Avatar == nil {
	// 	return ErrAvatarCannotBeEmpty
	// }

	return nil
}
