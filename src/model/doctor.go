package model

import "time"

type Doctor struct {
	user      User
	cns       string
	crm       string
	specialty Specialty
}

func NewDoctor(name string, birthDate time.Time, cpf string, sex string,
	address string, email string, cellphoneDoctor CellphoneUser, password string, active bool, cns string, crm string, specialty Specialty) Doctor {
	u := User{
		name:          name,
		birthDate:     birthDate,
		cpf:           cpf,
		sex:           sex,
		address:       address,
		email:         email,
		cellphoneUser: cellphoneDoctor,
		password:      password,
		active:        active,
	}

	return Doctor{
		user:      u,
		cns:       cns,
		crm:       crm,
		specialty: specialty,
	}
}

func (d Doctor) GetUser() User {
	return d.user
}

//func (d Doctor) SetUser(person User) {
//	d.user = person
//}

func (p *Doctor) SetUserId(newId int) {
	p.user.SetId(newId)
}

func (p *Doctor) SetUserName(newName string) {
	p.user.SetName(newName)
}

func (p *Doctor) SetUserBirthDate(newBirthDate time.Time) {
	p.user.SetBirthDate(newBirthDate)
}

func (p *Doctor) SetUserCpf(newCpf string) {
	p.user.SetCpf(newCpf)
}

func (p *Doctor) SetUserSex(newSex string) {
	p.user.SetSex(newSex)
}

func (p *Doctor) SetUserAddress(newAddress string) {
	p.user.SetAddress(newAddress)
}

func (p *Doctor) SetUserEmail(newEmail string) {
	p.user.SetEmail(newEmail)
}

func (p *Doctor) SetUserPassword(newPassword string) {
	p.user.SetPassword(newPassword)
}

func (p *Doctor) SetUserActive(newActive bool) {
	p.user.SetActive(newActive)
}

func (p *Doctor) SetUserCellphoneUser(newCellphoneUser CellphoneUser) {
	p.user.SetCellphoneUser(newCellphoneUser)
}

func (d Doctor) GetCns() string {
	return d.cns
}

func (d *Doctor) SetCns(cns string) {
	d.cns = cns
}

func (d Doctor) GetCrm() string {
	return d.crm
}

func (d *Doctor) SetCrm(crm string) {
	d.crm = crm
}

func (d Doctor) GetSpecialty() Specialty {
	return d.specialty
}

func (d *Doctor) SetSpecialty(specialty Specialty) {
	d.specialty = specialty
}
