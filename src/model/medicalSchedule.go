package model

import "time"

type MedicalSchedule struct {
	id            int
	doctorId      int
	queryValue    float64
	dayOfService  string
	specificDate  time.Time
	period1       string
	period2       string
	year          string
	scheduleLimit int
	active        bool
}

func NewMedicalSchedule(doctorId int, queryValue float64, dayOfService string, specificDate time.Time, startTime string,
	finalTime string, year string, scheduleLimit int) MedicalSchedule {
	return MedicalSchedule{
		doctorId:      doctorId,
		queryValue:    queryValue,
		dayOfService:  dayOfService,
		specificDate:  specificDate,
		period1:       startTime,
		period2:       finalTime,
		year:          year,
		scheduleLimit: scheduleLimit,
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

func (m MedicalSchedule) GetPeriod1() string {
	return m.period1
}

func (m *MedicalSchedule) SetPeriod1(period1 string) {
	m.period1 = period1
}

func (m MedicalSchedule) GetPeriod2() string {
	return m.period2
}

func (m *MedicalSchedule) SetPeriod2(period2 string) {
	m.period2 = period2
}

func (m MedicalSchedule) GetYear() string {
	return m.year
}

func (m *MedicalSchedule) SetYear(year string) {
	m.year = year
}

func (m MedicalSchedule) GetScheduleLimit() int {
	return m.scheduleLimit
}

func (m *MedicalSchedule) SetScheduleLimit(scheduleLimit int) {
	m.scheduleLimit = scheduleLimit
}

func (m MedicalSchedule) IsActive() bool {
	return m.active
}

func (m *MedicalSchedule) SetActive(active bool) {
	m.active = active
}
