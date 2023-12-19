package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
	"MedGestao/src/request"
	"MedGestao/src/response"
	"log"
	"os"
)

func PatientRegister(patientRequest request.PatientRequest) (int, error, response.ErrorResponse) {

	var patientId int
	var err error
	var errorMessage response.ErrorResponse
	if patientRequest.User.Name == "" {
		return patientId, err, errorMessage
	}
	cellPhoneUser := model.NewCellphoneUser(patientRequest.User.CellphoneUser.Number)
	//name string, birthDate time.Time, cpf string, sex string,
	//	address string, email string, password string, active bool, cellphonePatient CellphoneUser
	patient := model.NewPatient(patientRequest.User.Name, patientRequest.User.BirthDate, patientRequest.User.Cpf,
		patientRequest.User.Sex, patientRequest.User.Address, patientRequest.User.Email, patientRequest.User.Password,
		patientRequest.User.ImageUrl, cellPhoneUser)
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	logger.Println("[CONTROLLER.PatientRegister] " + model.LogUser(patient.GetUser()))

	patientId, err, errorMessage = dao.PatientInsert(patient)
	if err != nil {
		return patientId, err, errorMessage
	}

	return patientId, err, errorMessage
}

func PatientRegisterEdit(idPatientRequest int, patientRequest request.PatientRequest) (bool, error) {

	var success bool
	var err error

	//Adicionar essas condições depois: || patient == nil || patient.User == nil no lugar da que está comparando o nome
	if idPatientRequest == 0 || patientRequest.User.Name == "" {
		return success, err
	}

	cellPhoneUser := model.NewCellphoneUser(patientRequest.User.CellphoneUser.Number)
	//name string, birthDate time.Time, cpf string, sex string,
	//	address string, email string, password string, active bool, cellphonePatient CellphoneUser
	patient := model.NewPatient(patientRequest.User.Name, patientRequest.User.BirthDate, patientRequest.User.Cpf,
		patientRequest.User.Sex, patientRequest.User.Address, patientRequest.User.Email, patientRequest.User.Password,
		patientRequest.User.ImageUrl, cellPhoneUser)

	success, err = dao.PatientEdit(idPatientRequest, patient)
	if err != nil {
		return success, err
	}

	if success == true {
		println("O cadastro foi alterado com sucesso!")
	} else {
		println("Não foi possível alterar o cadastro!")
	}

	return success, err
}

func PatientAuthenticatorLogin(email string, password string) (bool, int, error) {
	var authorized bool
	var patientId int
	var err error
	if email == "" || password == "" {
		return authorized, patientId, err
	}

	authorized, patientId, err = dao.PatientValidateLogin(email, password)
	if err != nil {
		return authorized, patientId, err
	}

	return authorized, patientId, err

}

func PatientSelectByIdRegister(patientId int) (response.PatientResponse, error) {
	var p model.Patient
	var patient response.PatientResponse
	var err error
	if patientId != 0 {
		p, err = dao.PatientSelectById(patientId)
		if err != nil {
			println("Error na busca das informações do paciente: ", err.Error())
			return patient, err
		}

		cellphoneUserResponse := response.CellphoneResponse{
			Number: p.GetUser().GetCellphoneUser().GetNumber(),
		}
		userResponse := response.UserResponse{
			Id:            p.GetUser().GetId(),
			Name:          p.GetUser().GetName(),
			BirthDate:     p.GetUser().GetBirthDate(),
			Cpf:           p.GetUser().GetCpf(),
			Sex:           p.GetUser().GetSex(),
			Address:       p.GetUser().GetAddress(),
			Email:         p.GetUser().GetEmail(),
			ImageUrl:      p.GetUser().GetImageUrl(),
			CellphoneUser: cellphoneUserResponse,
		}
		patient = response.PatientResponse{
			User: userResponse,
		}
	} else {
		println("Informe um id de paciente válido!")
	}
	return patient, err
}

func PatientRegisterOff(patientId int) (bool, error) {
	var success bool
	var err error
	if patientId != 0 {
		success, err = dao.PatientOff(patientId)
		if err != nil {
			println("Error durante o desligamento do registro do paciente: ", err.Error())
			return success, err
		}
	} else {
		println("Informe um id de paciente válido!")
		success = false
		return success, err
	}

	return success, err
}
