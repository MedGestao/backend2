package response

import "time"

type CellphoneResponse struct {
	//UserId int    `json:"userId,omitempty"`
	Number string `json:"numberResponse"`
}

type UserResponse struct {
	Name          string            `json:"nameResponse"`
	BirthDate     time.Time         `json:"birthDateResponse"`
	Cpf           string            `json:"cpfResponse"`
	Sex           string            `json:"sexResponse"`
	Address       string            `json:"addressResponse"`
	Email         string            `json:"emailResponse"`
	Password      string            `json:"passwordResponse"`
	CellphoneUser CellphoneResponse `json:"cellphoneUserResponse"`
}

type PatientResponse struct {
	User UserResponse `json:"userResponse"`
}

type PatientIdResponse struct {
	Id int `json:"idPatientSelectRequest"`
}
