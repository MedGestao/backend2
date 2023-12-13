package request

import "time"

type PatientDoctorConsultationRequest struct {
	AppointmentDate time.Time `json:"appointmentDate"`
	AppointmentTime string    `json:"appointmentTime"`
	PatientId       int       `json:"patientId"`
	DoctorId        int       `json:"doctorId"`
	Value           float64   `json:"value"`
}

type PatientDoctorConsultationIdRequest struct {
	Id int `json:"id"`
}
