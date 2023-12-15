package request

import "time"

type CellphoneResponse struct {
	UserId int    `json:"userId"`
	Number string `json:"number"`
}

type EditCellphoneResponse struct {
	Id int `json:"id"`
}

type UserRequest struct {
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

func LogUserRequest(user UserRequest) string {
	return "password: " + user.Password +
		" email: " + user.Email +
		" name: " + user.Name +
		" imageUrl: " + user.ImageUrl +
		" address: " + user.Address +
		" sex: " + user.Sex
}
