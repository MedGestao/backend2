package model

import "time"

type Patient struct {
	user User
}

func NewPatient(name string, birthDate time.Time, cpf string, sex string,
	address string, email string, password string, active bool, cellphonePatient CellphoneUser) Patient {

	u := User{
		name:          name,
		birthDate:     birthDate,
		cpf:           cpf,
		sex:           sex,
		address:       address,
		email:         email,
		password:      password,
		active:        active,
		cellphoneUser: cellphonePatient,
	}

	return Patient{
		user: u,
	}

}

func (p Patient) GetUser() User {
	return p.user
}

func (p *Patient) SetUser(user User) {
	p.user = user
}

func (p *Patient) SetUserId(newId int) {
	p.user.SetId(newId)
}

func (p Patient) SetUserName(newName string) {
	p.user.SetName(newName)
}

func (p Patient) SetUserBirthDate(newBirthDate time.Time) {
	p.user.SetBirthDate(newBirthDate)
}

func (p Patient) SetUserCpf(newCpf string) {
	p.user.SetCpf(newCpf)
}

func (p Patient) SetUserSex(newSex string) {
	p.user.SetSex(newSex)
}

func (p Patient) SetUserAddress(newAddress string) {
	p.user.SetAddress(newAddress)
}

func (p Patient) SetUserEmail(newEmail string) {
	p.user.SetEmail(newEmail)
}

func (p Patient) SetUserPassword(newPassword string) {
	p.user.SetPassword(newPassword)
}

func (p Patient) SetUserActive(newActive bool) {
	p.user.SetActive(newActive)
}

func (p Patient) SetUserCellphoneUser(newCellphoneUser CellphoneUser) {
	p.user.SetCellphoneUser(newCellphoneUser)
}
