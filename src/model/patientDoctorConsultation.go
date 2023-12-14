package model

import "time"

type PatientDoctorConsultation struct {
	id              int
	appointmentDate time.Time
	appointmentTime string
	patientId       int
	doctorId        int
	status          string
	value           float64
}

func NewPatientDoctorConsultation(appointmentDate time.Time, appointmentTime string, patientId int, doctorId int, value float64) PatientDoctorConsultation {
	return PatientDoctorConsultation{
		appointmentDate: appointmentDate,
		appointmentTime: appointmentTime,
		patientId:       patientId,
		doctorId:        doctorId,
		value:           value,
	}
}

func (p PatientDoctorConsultation) GetId() int {
	return p.id
}

func (p *PatientDoctorConsultation) SetId(id int) {
	p.id = id
}

func (p PatientDoctorConsultation) GetAppointmentDate() time.Time {
	return p.appointmentDate
}

func (p *PatientDoctorConsultation) SetAppointmentDate(appointmentDate time.Time) {
	p.appointmentDate = appointmentDate
}

func (p PatientDoctorConsultation) GetAppointmentTime() string {
	return p.appointmentTime
}

func (p *PatientDoctorConsultation) SetAppointmentTime(timeTable string) {
	p.appointmentTime = timeTable
}

func (p PatientDoctorConsultation) GetPatientId() int {
	return p.patientId
}

func (p *PatientDoctorConsultation) SetPatientId(patientId int) {
	p.patientId = patientId
}

func (p PatientDoctorConsultation) GetDoctorId() int {
	return p.doctorId
}

func (p *PatientDoctorConsultation) SetDoctorId(doctorId int) {
	p.doctorId = doctorId
}

func (p PatientDoctorConsultation) GetStatus() string {
	return p.status
}

func (p *PatientDoctorConsultation) SetStatus(status string) {
	p.status = status
}

func (p PatientDoctorConsultation) GetValue() float64 {
	return p.value
}

func (p *PatientDoctorConsultation) SetValue(value float64) {
	p.value = value
}
