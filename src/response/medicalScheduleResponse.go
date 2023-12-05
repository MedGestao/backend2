package response

import "time"

type MedicalScheduleResponse struct {
	DoctorId     DoctorIdResponse `json:"doctorIdResponse"`
	QueryValue   float64          `json:"queryValueResponse"`
	DayOfService string           `json:"dayOfServiceResponse"`
	SpecificDate time.Time        `json:"specificDateResponse"`
	StartTime    string           `json:"startTimeResponse"`
	FinalTime    string           `json:"finalTimeResponse"`
	Year         string           `json:"yearResponse"`
}

type MedicalScheduleIdResponse struct {
	Id int `json:"idMedicalScheduleResponse"`
}
