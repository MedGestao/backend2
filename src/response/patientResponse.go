package response

type PatientResponse struct {
	User UserResponse `json:"user"`
}

type PatientIdResponse struct {
	Id int `json:"id"`
}
