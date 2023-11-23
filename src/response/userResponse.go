package response

import "time"

type CellphoneResponse struct {
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
	ImageUrl      string            `json:"imageUrlResponse"`
	CellphoneUser CellphoneResponse `json:"cellphoneUserResponse"`
}
