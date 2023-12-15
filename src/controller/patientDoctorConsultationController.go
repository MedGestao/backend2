package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
	"MedGestao/src/request"
	"MedGestao/src/response"
	"time"
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

func SearchPatientDoctorConsultationAllByDoctor(doctorId int, appointmentDate time.Time) ([]response.PatientDoctorConsultationByDoctorResponse, error) {
	var patientDoctorConsultationByDoctorList []response.PatientDoctorConsultationByDoctorResponse
	var err error
	if doctorId == 0 {
		return patientDoctorConsultationByDoctorList, err
	}
	patientDoctorConsultationByDoctorList, err = dao.PatientDoctorConsultationAllByIdDoctor(doctorId, appointmentDate)

	return patientDoctorConsultationByDoctorList, err
}

func SearchPatientDoctorConsultationAllByPatient(patientId int) ([]response.PatientDoctorConsultationByDoctorResponse, error) {
	var patientDoctorConsultationByDoctorList []response.PatientDoctorConsultationByDoctorResponse
	var err error
	if patientId == 0 {
		return patientDoctorConsultationByDoctorList, err
	}

	patientDoctorConsultationByDoctorList, err = dao.PatientDoctorConsultationAllByIdPatient(patientId)

	return patientDoctorConsultationByDoctorList, err
}

func SearchPatientDoctorConsultationById(patientDoctorConsultationId int) (response.PatientDoctorConsultationResponse, error) {
	patientDoctorConsultation, err := dao.PatientDoctorConsultationById(patientDoctorConsultationId)

	return patientDoctorConsultation, err
}

func EditPatientDoctorConsultation(patientDoctorConsultationRequestId int, patientDoctorConsultationRequest request.PatientDoctorConsultationRequest) (bool, error, response.ErrorResponse) {
	var success bool
	var err error
	var errorMessage response.ErrorResponse

	if patientDoctorConsultationRequestId == 0 || patientDoctorConsultationRequest == (request.PatientDoctorConsultationRequest{}) {
		println("Nenhum dado recebido")
		return success, err, errorMessage
	}

	patientDoctorConsultation := model.NewPatientDoctorConsultation(patientDoctorConsultationRequest.AppointmentDate,
		patientDoctorConsultationRequest.AppointmentTime, patientDoctorConsultationRequest.PatientId,
		patientDoctorConsultationRequest.DoctorId, patientDoctorConsultationRequest.Value)

	success, err, errorMessage = dao.PatientDoctorConsultationEdit(patientDoctorConsultationRequestId, patientDoctorConsultation)

	return success, err, errorMessage
}

func CompletePatientDoctorConsultation(patientDoctorConsultationId int) (bool, error) {
	var success bool
	var err error
	if patientDoctorConsultationId == 0 {
		return success, err
	}

	success, err = dao.EndPatientDoctorConsultation(patientDoctorConsultationId)

	return success, err
}

func DeactivatePatientDoctorConsultation(patientDoctorConsultationId int) (bool, error) {
	var success bool
	var err error
	if patientDoctorConsultationId == 0 {
		return success, err
	}

	success, err = dao.PatientDoctorConsultationOff(patientDoctorConsultationId)

	return success, err
}
