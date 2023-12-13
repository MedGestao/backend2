package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
	"MedGestao/src/request"
	"MedGestao/src/response"
)

func RegisterMedicalSchedule(medicalScheduleRequest request.MedicalScheduleRequest) (bool, error, response.ErrorResponse) {
	var success bool
	var err error
	var errorMessage response.ErrorResponse

	if medicalScheduleRequest == (request.MedicalScheduleRequest{}) {
		println("Nenhum dado recebido")
		return success, err, errorMessage
	}

	medicalSchedule := model.NewMedicalSchedule(medicalScheduleRequest.DoctorId.Id, medicalScheduleRequest.QueryValue,
		medicalScheduleRequest.DayOfService, medicalScheduleRequest.SpecificDate, medicalScheduleRequest.Period1,
		medicalScheduleRequest.Period2, medicalScheduleRequest.Year, medicalScheduleRequest.ScheduleLimit)

	success, err, errorMessage = dao.MedicalScheduleInsert(medicalSchedule)
	if err != nil {
		return success, err, errorMessage
	}

	return success, err, errorMessage
}

func SearchAllMedicalScheduleByIdDoctor(doctorId int) ([]response.MedicalScheduleResponse, error) {
	var m []model.MedicalSchedule
	var medicalScheduleList []response.MedicalScheduleResponse
	var err error

	if doctorId == 0 {
		return medicalScheduleList, err
	}
	m, err = dao.MedicalScheduleSelectAllByIdDoctor(doctorId)
	if err != nil {
		return medicalScheduleList, err
	}

	for _, medicalScheduleDB := range m {
		doctorId := response.DoctorIdResponse{
			Id: medicalScheduleDB.GetDoctorId(),
		}
		medicalSchedule := response.MedicalScheduleResponse{
			Id:            medicalScheduleDB.GetId(),
			DoctorId:      doctorId,
			QueryValue:    medicalScheduleDB.GetQueryValue(),
			DayOfService:  medicalScheduleDB.GetDayOfService(),
			SpecificDate:  medicalScheduleDB.GetSpecificDate(),
			Period1:       medicalScheduleDB.GetPeriod1(),
			Period2:       medicalScheduleDB.GetPeriod2(),
			Year:          medicalScheduleDB.GetYear(),
			ScheduleLimit: medicalScheduleDB.GetScheduleLimit(),
		}

		medicalScheduleList = append(medicalScheduleList, medicalSchedule)
	}

	return medicalScheduleList, err

}

func SearchByIdMedicalSchedule(id int) (response.MedicalScheduleResponse, error) {
	var m model.MedicalSchedule
	var medicalSchedule response.MedicalScheduleResponse
	var err error

	if id == 0 {
		return medicalSchedule, err
	}

	m, err = dao.MedicalScheduleSelectById(id)
	if err != nil {
		return medicalSchedule, err
	}

	//DoctorId     DoctorIdResponse `json:"doctorIdResponse"`
	//QueryValue   float64          `json:"queryValueResponse"`
	//DayOfService string           `json:"dayOfServiceResponse"`
	//SpecificDate time.Time        `json:"specificDateResponse"`
	//Period1Period1    string      `json:"startTimeResponse"`
	//Period2    string           	`json:"finalTimeResponse"`
	//Year         string           `json:"yearResponse"`
	//ScheduleLimit int 			`json:"schedule_limit"`

	doctorId := response.DoctorIdResponse{
		Id: m.GetDoctorId(),
	}

	medicalSchedule = response.MedicalScheduleResponse{
		Id:            m.GetId(),
		DoctorId:      doctorId,
		QueryValue:    m.GetQueryValue(),
		DayOfService:  m.GetDayOfService(),
		SpecificDate:  m.GetSpecificDate(),
		Period1:       m.GetPeriod1(),
		Period2:       m.GetPeriod2(),
		Year:          m.GetYear(),
		ScheduleLimit: m.GetScheduleLimit(),
	}

	return medicalSchedule, err
}

func EditMedicalSchedule(medicalScheduleIdRequest int, medicalScheduleRequest request.MedicalScheduleRequest) (bool, error) {
	var success bool
	var err error
	if medicalScheduleRequest == (request.MedicalScheduleRequest{}) {
		return success, err
	}

	medicalSchedule := model.NewMedicalSchedule(medicalScheduleRequest.DoctorId.Id, medicalScheduleRequest.QueryValue,
		medicalScheduleRequest.DayOfService, medicalScheduleRequest.SpecificDate, medicalScheduleRequest.Period1,
		medicalScheduleRequest.Period2, medicalScheduleRequest.Year, medicalScheduleRequest.ScheduleLimit)
	medicalSchedule.SetId(medicalScheduleIdRequest)
	success, err = dao.MedicalScheduleEdit(medicalSchedule)
	if err != nil {
		return success, err
	}

	return success, err
}

func OffMedicalSchedule(id int) (bool, error) {
	var success bool
	var err error
	if id == 0 {
		return success, err
	}

	success, err = dao.MedicalScheduleDelete(id)
	if err != nil {
		return success, err
	}

	return success, err
}
