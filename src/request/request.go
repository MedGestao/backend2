package request

import (
	"time"
)

type CellphoneRequest struct {
	//UserId int    `json:"userId,omitempty"`
	Number string `json:"number"`
}

type EditCellphoneRequest struct {
	Id int `json:"id"`
}

type SpecialtyRequest struct {
	Description string `json:"description,omitempty"`
}

type EditSpecialtyRequest struct {
	Id int `json:"id"`
}

type UserRequest struct {
	Name          string           `json:"name"`
	BirthDate     time.Time        `json:"birthDate"`
	Cpf           string           `json:"cpf"`
	Sex           string           `json:"sex"`
	Address       string           `json:"address"`
	Email         string           `json:"email"`
	Password      string           `json:"password"`
	CellphoneUser CellphoneRequest `json:"cellphoneUser"`
}

type PatientRequest struct {
	User UserRequest `json:"userPatientRequest"`
}

type PatientIdRequest struct {
	Id int `json:"idPatientSelectRequest"`
}

type EditPatientRequest struct {
	PatientIdRequest PatientIdRequest `json:"patientEditIdRequest"`
	PatientRequest   PatientRequest   `json:"patientEditRequest"`
}

type PatientAuthenticatorRequest struct {
	Email    string `json:"emailAuthenticatorRequest"`
	Password string `json:"passwordAuthenticatorRequest"`
}

type DoctorResponse struct {
	User      UserRequest      `json:"user"`
	Crm       string           `json:"crm,omitempty"`
	Specialty SpecialtyRequest `json:"specialty,omitempty"`
}

type EditDoctorRequest struct {
	Id int `json:"id"`
}
