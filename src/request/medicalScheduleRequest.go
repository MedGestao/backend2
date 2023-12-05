package request

import "time"

type MedicalScheduleRequest struct {
	DoctorId     DoctorIdRequest `json:"doctorId"`
	QueryValue   float64         `json:"queryValue"`
	DayOfService string          `json:"dayOfService"`
	SpecificDate time.Time       `json:"specificDate"`
	StartTime    string          `json:"startTime"`
	FinalTime    string          `json:"finalTime"`
	Year         string          `json:"year"`
}

type MedicalScheduleIdRequest struct {
	Id int `json:"idMedicalScheduleSelectRequest"`
}

type EditMedicalScheduleRequest struct {
	MedicalScheduleIdRequest MedicalScheduleIdRequest `json:"editMedicalScheduleIdRequest"`
	MedicalScheduleRequest   MedicalScheduleRequest   `json:"editMedicalScheduleRequest"`
}
