package model

import "time"

type MedicalSchedule struct {
	id           int
	doctorId     int
	queryValue   float64
	dayOfService string
	specificDate time.Time
	startTime    string
	finalTime    string
	year         string
	active       bool
}

func NewMedicalSchedule(doctorId int, queryValue float64, dayOfService string, specificDate time.Time, startTime string,
	finalTime string, year string) MedicalSchedule {
	return MedicalSchedule{
		doctorId:     doctorId,
		queryValue:   queryValue,
		dayOfService: dayOfService,
		specificDate: specificDate,
		startTime:    startTime,
		finalTime:    finalTime,
		year:         year,
	}
}

func (m MedicalSchedule) GetId() int {
	return m.id
}

func (m *MedicalSchedule) SetId(id int) {
	m.id = id
}

func (m MedicalSchedule) GetDoctorId() int {
	return m.doctorId
}

func (m *MedicalSchedule) SetDoctorId(doctorId int) {
	m.doctorId = doctorId
}

func (m MedicalSchedule) GetQueryValue() float64 {
	return m.queryValue
}

func (m *MedicalSchedule) SetQueryValue(queryValue float64) {
	m.queryValue = queryValue
}

func (m MedicalSchedule) GetDayOfService() string {
	return m.dayOfService
}

func (m *MedicalSchedule) SetDayOfService(dayOfService string) {
	m.dayOfService = dayOfService
}

func (m MedicalSchedule) GetSpecificDate() time.Time {
	return m.specificDate
}

func (m *MedicalSchedule) SetSpecificTime(specificDate time.Time) {
	m.specificDate = specificDate
}

func (m MedicalSchedule) GetStartTime() string {
	return m.startTime
}

func (m *MedicalSchedule) SetStartTime(startTime string) {
	m.startTime = startTime
}

func (m MedicalSchedule) GetFinalTime() string {
	return m.finalTime
}

func (m *MedicalSchedule) SetFinalTime(finalTime string) {
	m.finalTime = finalTime
}

func (m MedicalSchedule) GetYear() string {
	return m.year
}

func (m *MedicalSchedule) SetYear(year string) {
	m.year = year
}

func (m MedicalSchedule) IsActive() bool {
	return m.active
}

func (m *MedicalSchedule) SetActive(active bool) {
	m.active = active
}
