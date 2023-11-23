package model

type MedicalSchedule struct {
	id           int
	doctorId     int
	dayOfService string
	specificTime string
	startTime    string
	finalTime    string
	year         string
	active       bool
}

func NewMedicalSchedule(doctorId int, dayOfService string, specificTime string, startTime string,
	finalTime string, year string) MedicalSchedule {
	return MedicalSchedule{
		doctorId:     doctorId,
		dayOfService: dayOfService,
		specificTime: specificTime,
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

func (a *MedicalSchedule) SetDoctorId(doctorId int) {
	a.doctorId = doctorId
}

func (a MedicalSchedule) GetDayOfService() string {
	return a.dayOfService
}

func (a *MedicalSchedule) SetDayOfService(dayOfService string) {
	a.dayOfService = dayOfService
}

func (a MedicalSchedule) GetSpecificTime() string {
	return a.specificTime
}

func (a *MedicalSchedule) SetSpecificTime(specificTime string) {
	a.specificTime = specificTime
}

func (a MedicalSchedule) GetStartTime() string {
	return a.startTime
}

func (a *MedicalSchedule) SetStartTime(startTime string) {
	a.startTime = startTime
}

func (a MedicalSchedule) GetFinalTime() string {
	return a.finalTime
}

func (a *MedicalSchedule) SetFinalTime(finalTime string) {
	a.finalTime = finalTime
}

func (a MedicalSchedule) GetYear() string {
	return a.year
}

func (a *MedicalSchedule) SetYear(year string) {
	a.year = year
}

func (a MedicalSchedule) IsActive() bool {
	return a.active
}

func (a *MedicalSchedule) SetActive(active bool) {
	a.active = active
}
