package model

import "time"

type PatientDoctorConsultation struct {
	id                int
	appointmentDate   time.Time
	timetable         string
	patientId         int
	doctorId          int
	medicalScheduleId int
	status            bool
	value             float64
}

func NewPatientDoctorSchedule(appointmentDate time.Time, timetable string, patientId int, doctorId int, medicalScheduleId int, status bool, value float64) PatientDoctorConsultation {
	return PatientDoctorConsultation{
		appointmentDate:   appointmentDate,
		timetable:         timetable,
		patientId:         patientId,
		doctorId:          doctorId,
		medicalScheduleId: medicalScheduleId,
		status:            status,
		value:             value,
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

func (p PatientDoctorConsultation) GetTimeTable() string {
	return p.timetable
}

func (p *PatientDoctorConsultation) SetTimeTable(timeTable string) {
	p.timetable = timeTable
}

func (p PatientDoctorConsultation) GetMedicalSchedule() int {
	return p.medicalScheduleId
}

func (p *PatientDoctorConsultation) SetMedicalSchedule(medicalSchedule int) {
	p.medicalScheduleId = medicalSchedule
}

func (p PatientDoctorConsultation) GetStatus() bool {
	return p.status
}

func (p *PatientDoctorConsultation) SetStatus(status bool) {
	p.status = status
}

func (p PatientDoctorConsultation) GetValue() float64 {
	return p.value
}

func (p *PatientDoctorConsultation) SetValue(value bool) {
	p.status = value
}
