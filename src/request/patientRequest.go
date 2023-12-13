package request

type PatientRequest struct {
	User UserRequest `json:"user"`
}

type PatientIdRequest struct {
	Id int `json:"id"`
}

//type EditPatientRequest struct {
//	PatientIdRequest PatientIdRequest `json:"patientEditIdRequest"`
//	PatientRequest   PatientRequest   `json:"patientEditRequest"`
//}

type PatientAuthenticatorRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
