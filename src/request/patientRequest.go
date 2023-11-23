package request

type PatientRequest struct {
	User UserRequest `json:"userPatientRequest"`
}

type PatientIdRequest struct {
	Id int `json:"idPatientSelectRequest"`
}

type EditPatientRequest struct {
	PatientIdRequest PatientIdRequest `json:"patientEditIdRequest"`
	PatientRequest   PatientRequest   `json:"patientEditRequest"`
}

type PatientAuthenticatorRequest struct {
	Email    string `json:"emailAuthenticatorRequest"`
	Password string `json:"passwordAuthenticatorRequest"`
}
