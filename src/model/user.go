package model

import (
	"time"
)

type User struct {
	id            int
	name          string
	birthDate     time.Time
	cpf           string
	sex           string
	address       string
	email         string
	password      string
	active        bool
	cellphoneUser CellphoneUser
}

func (u User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u User) GetBirthDate() time.Time {
	return u.birthDate
}

func (u *User) SetBirthDate(birthDate time.Time) {
	u.birthDate = birthDate
}

func (u User) GetCpf() string {
	return u.cpf
}

func (u *User) SetCpf(cpf string) {
	u.cpf = cpf
}

func (u User) GetSex() string {
	return u.sex
}

func (u *User) SetSex(sex string) {
	u.sex = sex
}

func (u User) GetAddress() string {
	return u.address
}

func (u *User) SetAddress(address string) {
	u.address = address
}

func (u User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u User) IsActive() bool {
	return u.active
}

func (u *User) SetActive(activity bool) {
	u.active = activity
}

func (u User) GetCellphoneUser() CellphoneUser {
	return u.cellphoneUser
}

func (u *User) SetCellphoneUser(cellphoneUser CellphoneUser) {
	u.cellphoneUser = cellphoneUser
}
