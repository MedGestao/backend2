package request

type SpecialtyRequest struct {
	Description string `json:"specialtyDescription,omitempty"`
}

//type EditSpecialtyRequest struct {
//	Id          int    `json:"id,omitempty"`
//	Description string `json:"description,omitempty"`
//}

type DoctorRequest struct {
	User      UserRequest      `json:"userDoctorRequest"`
	Crm       string           `json:"crmDoctorRequest"`
	Specialty SpecialtyRequest `json:"specialtyDoctorRequest"`
}

type DoctorIdRequest struct {
	Id int `json:"idDoctorSelectRequest"`
}

type EditDoctorRequest struct {
	DoctorIdRequest DoctorIdRequest `json:"doctorEditIdRequest"`
	DoctorRequest   DoctorRequest   `json:"doctorEditRequest"`
}

type DoctorAuthenticatorRequest struct {
	Email    string `json:"emailAuthenticatorRequest"`
	Password string `json:"passwordAuthenticatorRequest"`
}
