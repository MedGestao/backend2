package response

import "time"

type PatientDoctorConsultationResponse struct {
	Id              int       `json:"id"`
	PatientId       int       `json:"patientId"`
	DoctorId        int       `json:"doctorId"`
	AppointmentDate time.Time `json:"appointmentDate"`
	AppointmentTime string    `json:"appointmentTime"`
	Value           float64   `json:"value"`
	Status          string    `json:"status"`
}

type PatientDoctorConsultationByDoctorResponse struct {
	Name                              string                            `json:"name"`
	Specialty                         string                            `json:"specialty"`
	BirthDate                         time.Time                         `json:"birthDate"`
	Age                               string                            `json:"age"`
	ImageUrl                          string                            `json:"imageUrl"`
	PatientDoctorConsultationResponse PatientDoctorConsultationResponse `json:"patientDoctorConsultationResponse"`
}
