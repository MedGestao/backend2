package response

type PatientResponse struct {
	User UserResponse `json:"userResponse"`
}

type PatientIdResponse struct {
	Id int `json:"idPatientSelectResponse"`
}
