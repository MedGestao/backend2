package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
	"MedGestao/src/request"
	"MedGestao/src/response"
)

func RegisterPatientDoctorConsultation(patientDoctorConsultationRequest request.PatientDoctorConsultationRequest) (bool, error, response.ErrorResponse) {
	var success bool
	var err error
	var errorMessage response.ErrorResponse

	if patientDoctorConsultationRequest == (request.PatientDoctorConsultationRequest{}) {
		println("Nenhum dado recebido")
		return success, err, errorMessage
	}

	patientDoctorConsultation := model.NewPatientDoctorConsultation(patientDoctorConsultationRequest.AppointmentDate,
		patientDoctorConsultationRequest.AppointmentTime, patientDoctorConsultationRequest.PatientId,
		patientDoctorConsultationRequest.DoctorId, patientDoctorConsultationRequest.Value)

	success, err, errorMessage = dao.PatientDoctorConsultationInsert(patientDoctorConsultation)
	if err != nil {
		return success, err, errorMessage
	}

	return success, err, errorMessage
}

func SearchPatientDoctorConsultationAllByDoctor(doctorId int) ([]response.PatientDoctorConsultationByDoctorResponse, error) {
	patientDoctorConsultationByDoctorList, err := dao.PatientDoctorConsultationAllByIdDoctor(doctorId)

	return patientDoctorConsultationByDoctorList, err
}
