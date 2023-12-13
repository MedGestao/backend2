package response

import "time"

type MedicalScheduleResponse struct {
	Id            int              `json:"id"`
	DoctorId      DoctorIdResponse `json:"doctorId"`
	QueryValue    float64          `json:"queryValue"`
	DayOfService  string           `json:"dayOfService"`
	SpecificDate  time.Time        `json:"specificDate"`
	Period1       string           `json:"period1"`
	Period2       string           `json:"period2"`
	Year          string           `json:"year"`
	ScheduleLimit int              `json:"scheduleLimit"`
}

type MedicalScheduleIdResponse struct {
	Id int `json:"idMedicalScheduleResponse"`
}
