package response

type SpecialtyResponse struct {
	Description string `json:"description"`
}

type DoctorResponse struct {
	User      UserResponse      `json:"user"`
	Crm       string            `json:"crm"`
	Specialty SpecialtyResponse `json:"specialty"`
}

type DoctorIdResponse struct {
	Id int `json:"id"`
}
