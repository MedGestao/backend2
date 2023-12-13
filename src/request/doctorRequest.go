package request

type SpecialtyRequest struct {
	Description string `json:"description,omitempty"`
}

//type EditSpecialtyRequest struct {
//	Id          int    `json:"id,omitempty"`
//	Description string `json:"description,omitempty"`
//}

type DoctorRequest struct {
	User      UserRequest      `json:"user"`
	Crm       string           `json:"crm"`
	Specialty SpecialtyRequest `json:"specialty"`
}

type DoctorIdRequest struct {
	Id int `json:"id"`
}

//type EditDoctorRequest struct {
//	DoctorIdRequest DoctorIdRequest `json:"doctorEditIdRequest"`
//	DoctorRequest   DoctorRequest   `json:"doctorEditRequest"`
//}

type DoctorAuthenticatorRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DoctorFilterParameters struct {
	DoctorName    string `json:"doctorName"`
	SpecialtyName string `json:"specialtyName"`
}
