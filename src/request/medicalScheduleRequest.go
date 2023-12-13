package request

import "time"

type MedicalScheduleRequest struct {
	DoctorId      DoctorIdRequest `json:"doctorId"`
	QueryValue    float64         `json:"queryValue"`
	DayOfService  string          `json:"dayOfService"`
	SpecificDate  time.Time       `json:"specificDate"`
	Period1       string          `json:"period1"`
	Period2       string          `json:"period2"`
	Year          string          `json:"year"`
	ScheduleLimit int             `json:"scheduleLimit"`
}

type MedicalScheduleIdRequest struct {
	Id int `json:"id"`
}

//type EditMedicalScheduleRequest struct {
//	MedicalScheduleIdRequest MedicalScheduleIdRequest `json:"editMedicalScheduleIdRequest"`
//	MedicalScheduleRequest   MedicalScheduleRequest   `json:"editMedicalScheduleRequest"`
//}
