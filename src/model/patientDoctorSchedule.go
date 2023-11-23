package model

import "time"

type PatientDoctorSchedule struct {
	id                int
	appointmentDate   time.Time
	timetable         string
	patientId         int
	doctorId          int
	medicalScheduleId int
	status            bool
	observations      string
}

func NewPatientDoctorSchedule(appointmentDate time.Time, timetable string, patientId int, doctorId int, medicalScheduleId int, status bool, observations string) PatientDoctorSchedule {
	return PatientDoctorSchedule{
		appointmentDate:   appointmentDate,
		timetable:         timetable,
		patientId:         patientId,
		doctorId:          doctorId,
		medicalScheduleId: medicalScheduleId,
		status:            status,
		observations:      observations,
	}
}

func (p PatientDoctorSchedule) GetId() int {
	return p.id
}

func (p *PatientDoctorSchedule) SetId(id int) {
	p.id = id
}

func (p PatientDoctorSchedule) GetAppointmentDate() time.Time {
	return p.appointmentDate
}

func (p *PatientDoctorSchedule) SetAppointmentDate(appointmentDate time.Time) {
	p.appointmentDate = appointmentDate
}

func (p PatientDoctorSchedule) GetTimeTable() string {
	return p.timetable
}

func (p *PatientDoctorSchedule) SetTimeTable(timeTable string) {
	p.timetable = timeTable
}

func (p PatientDoctorSchedule) GetMedicalSchedule() int {
	return p.medicalScheduleId
}

func (p *PatientDoctorSchedule) SetMedicalSchedule(medicalSchedule int) {
	p.medicalScheduleId = medicalSchedule
}

func (p PatientDoctorSchedule) GetStatus() bool {
	return p.status
}

func (p *PatientDoctorSchedule) SetStatus(status bool) {
	p.status = status
}

func (p PatientDoctorSchedule) GetObservations() string {
	return p.observations
}

func (p *PatientDoctorSchedule) SetObservations(observations string) {
	p.observations = observations
}
