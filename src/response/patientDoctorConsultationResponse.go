package response

import "time"

type PatientDoctorConsultationResponse struct {
	Id              int       `json:"id"`
	AppointmentDate time.Time `json:"appointmentDate"`
	AppointmentTime string    `json:"appointmentTime"`
	Value           float64   `json:"value"`
}

type PatientDoctorConsultationByDoctorResponse struct {
	PatientName                       string                            `json:"patientName"`
	PatientDoctorConsultationResponse PatientDoctorConsultationResponse `json:"patientDoctorConsultationResponse"`
}
