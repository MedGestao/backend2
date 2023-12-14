package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
	"MedGestao/src/request"
	"MedGestao/src/response"
)

func DoctorRegister(doctorRequest request.DoctorRequest) (int, error, response.ErrorResponse) {

	var doctorId int
	var err error
	var errorMessage response.ErrorResponse
	if doctorRequest.User.Name == "" {
		return doctorId, err, errorMessage
	}

	cellPhoneUser := model.NewCellphoneUser(doctorRequest.User.CellphoneUser.Number)
	//specialty := model.NewSpecialty(doctorRequest.Specialty.Description)

	doctor := model.NewDoctor(doctorRequest.User.Name, doctorRequest.User.BirthDate, doctorRequest.User.Cpf,
		doctorRequest.User.Sex, doctorRequest.User.Address, doctorRequest.User.Email, cellPhoneUser, doctorRequest.User.Password,
		doctorRequest.User.ImageUrl, doctorRequest.Crm, model.Specialty{})
	doctor.SetSpecialtyId(doctorRequest.Specialty.Id)

	doctorId, err, errorMessage = dao.InsertDoctor(doctor)
	if err != nil {
		return doctorId, err, errorMessage
	}

	return doctorId, err, errorMessage
}

func DoctorAuthenticatorLogin(email string, password string) (bool, int, error) {
	var authorized bool
	var doctorId int
	var err error
	if email == "" || password == "" {
		return authorized, doctorId, err
	}

	authorized, doctorId, err = dao.DoctorValidateLogin(email, password)
	if err != nil {
		return authorized, doctorId, err
	}

	return authorized, doctorId, err
}

func DoctorSelectRegisterAll(doctorName string, specialtyName string) ([]response.DoctorResponse, error) {
	doctors, err := dao.DoctorSelectAll(doctorName, specialtyName)

	return doctors, err
}

func DoctorSelectRegisterById(doctorId int) (response.DoctorResponse, error) {
	var d model.Doctor
	var doctor response.DoctorResponse
	var err error
	if doctorId != 0 {
		d, err = dao.DoctorSelectById(doctorId)
		if err != nil {
			println("Error na busca das informações do paciente: ", err.Error())
			return doctor, err
		}

		cellphoneUserResponse := response.CellphoneResponse{
			Number: d.GetUser().GetCellphoneUser().GetNumber(),
		}
		userResponse := response.UserResponse{
			Name:          d.GetUser().GetName(),
			BirthDate:     d.GetUser().GetBirthDate(),
			Cpf:           d.GetUser().GetCpf(),
			Sex:           d.GetUser().GetSex(),
			Address:       d.GetUser().GetAddress(),
			Email:         d.GetUser().GetEmail(),
			ImageUrl:      d.GetUser().GetImageUrl(),
			CellphoneUser: cellphoneUserResponse,
		}

		specialtyUserResponse := response.SpecialtyResponse{
			Id:          d.GetSpecialty().GetId(),
			Description: d.GetSpecialty().GetDescription()}
		doctor = response.DoctorResponse{
			User:      userResponse,
			Crm:       d.GetCrm(),
			Specialty: specialtyUserResponse,
		}
	} else {
		println("Informe um id de paciente válido!")
	}
	return doctor, err
}

func DoctorRegisterEdit(idDoctorRequest int, doctorRequest request.DoctorRequest) (bool, error) {

	var success bool
	var err error

	//Adicionar essas condições depois: || patient == nil || patient.User == nil no lugar da que está comparando o nome
	if idDoctorRequest == 0 || doctorRequest.User.Name == "" {
		return success, err
	}

	cellPhoneUser := model.NewCellphoneUser(doctorRequest.User.CellphoneUser.Number)
	//specialty := model.NewSpecialty(doctorRequest.Specialty.Description)

	doctor := model.NewDoctor(doctorRequest.User.Name, doctorRequest.User.BirthDate, doctorRequest.User.Cpf,
		doctorRequest.User.Sex, doctorRequest.User.Address, doctorRequest.User.Email, cellPhoneUser, doctorRequest.User.Password,
		doctorRequest.User.ImageUrl, doctorRequest.Crm, model.Specialty{})
	doctor.SetSpecialtyId(doctorRequest.Specialty.Id)

	success, err = dao.DoctorEdit(idDoctorRequest, doctor)
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

func DoctorRegisterOff(doctorId int) (bool, error) {
	var success bool
	var err error
	if doctorId != 0 {
		success, err = dao.DoctorOff(doctorId)
		if err != nil {
			println("Error durante o desligamento do registro do médico: ", err.Error())
			return success, err
		}
	} else {
		println("Informe um id de médico válido!")
		success = false
		return success, err
	}

	return success, err
}
