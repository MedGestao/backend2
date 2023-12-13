package response

import "time"

type CellphoneResponse struct {
	Number string `json:"number"`
}

type UserResponse struct {
	Id            int               `json:"id"`
	Name          string            `json:"name"`
	BirthDate     time.Time         `json:"birthDate"`
	Cpf           string            `json:"cpf"`
	Sex           string            `json:"sex"`
	Address       string            `json:"address"`
	Email         string            `json:"email"`
	Password      string            `json:"password"`
	ImageUrl      string            `json:"imageUrl"`
	CellphoneUser CellphoneResponse `json:"cellphoneUser"`
}
