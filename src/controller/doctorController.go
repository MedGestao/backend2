package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
)

func DoctorRegister(doctor model.Doctor) bool {

	success, err := dao.InsertDoctor(doctor)
	if err != nil {
		panic(err)
	}

	return success
}

func DoctorAuthenticatorLogin(email string, password string) (bool, int) {
	if email == "" || password == "" {
		return false, 0
	}

	Authorized, doctorId, err := dao.DoctorValidateLogin(email, password)
	if err != nil {
		panic(err)
	}

	if Authorized == true {
		println("Entrou :)")
		return true, doctorId
	} else {
		println("Não entrou :(")
		return false, 0
	}

}

func DoctorSelectRegister(doctorId int) model.Doctor {
	var doctor model.Doctor
	var err error
	if doctorId != 0 {
		doctor, err = dao.DoctorSelectById(doctorId)
		if err != nil {
			println("Error na busca das informações do paciente: ", err.Error())
			panic(err)
		}
	} else {
		println("Informe um id de paciente válido!")
		return doctor
	}
	return doctor
}

func DoctorRegisterEdit(doctor model.Doctor) bool {

	success, err := dao.DoctorEdit(doctor)
	if err != nil {
		panic(err)
	}

	return success
}

func DoctorRegisterOff(doctorId int) bool {
	var success bool
	var err error
	if doctorId != 0 {
		success, err = dao.DoctorOff(doctorId)
		if err != nil {
			println("Error durante o desligamento do registro do médico: ", err.Error())
			panic(err)
		}
	} else {
		println("Informe um id de médico válido!")
		success = false
		return success
	}

	return success
}
