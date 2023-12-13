package model

type MedicalSchedule struct {
	id           int
	doctorId     int
	queryValue   float64
	dayOfService string
	period1      string
	period2      string
	active       bool
}

func NewMedicalSchedule(doctorId int, queryValue float64, dayOfService string, startTime string,
	finalTime string) MedicalSchedule {
	return MedicalSchedule{
		doctorId:     doctorId,
		queryValue:   queryValue,
		dayOfService: dayOfService,
		period1:      startTime,
		period2:      finalTime,
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

func (m MedicalSchedule) IsActive() bool {
	return m.active
}

func (m *MedicalSchedule) SetActive(active bool) {
	m.active = active
}
