package request

type MedicalScheduleRequest struct {
	DoctorId     DoctorIdRequest `json:"doctorId"`
	QueryValue   float64         `json:"queryValue"`
	DayOfService string          `json:"dayOfService"`
	Period1      string          `json:"period1"`
	Period2      string          `json:"period2"`
}

type MedicalScheduleIdRequest struct {
	Id int `json:"idMedicalScheduleSelectRequest"`
}

//type EditMedicalScheduleRequest struct {
//	MedicalScheduleIdRequest MedicalScheduleIdRequest `json:"editMedicalScheduleIdRequest"`
//	MedicalScheduleRequest   MedicalScheduleRequest   `json:"editMedicalScheduleRequest"`
//}
