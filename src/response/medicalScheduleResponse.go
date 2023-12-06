package response

import "time"

type MedicalScheduleResponse struct {
	Id           int              `json:"idResponse"`
	DoctorId     DoctorIdResponse `json:"doctorIdResponse"`
	QueryValue   float64          `json:"queryValueResponse"`
	DayOfService string           `json:"dayOfServiceResponse"`
	SpecificDate time.Time        `json:"specificDateResponse"`
	Period1      string           `json:"period1"`
	Period2      string           `json:"period2"`
	Year         string           `json:"yearResponse"`
}

type MedicalScheduleIdResponse struct {
	Id int `json:"idMedicalScheduleResponse"`
}
