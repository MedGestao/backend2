package request

import (
	"strconv"
	"time"
)

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

func LogPatientDoctorConsultationRequest(c PatientDoctorConsultationRequest) string {
	return "appointmentDate: " + c.AppointmentDate.String() +
		" appointmentTime: " + c.AppointmentTime +
		" patientId: " + strconv.Itoa(c.PatientId) +
		" doctorId: " + strconv.Itoa(c.DoctorId) +
		" value: " + strconv.FormatFloat(c.Value, 'f', -1, 64)
}
