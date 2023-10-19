package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
	"fmt"
	"time"
)

func PatientRegister() bool {
	birthDate, err := time.Parse("2006-01-02", "1982-09-22")

	if err != nil {
		fmt.Println(err)
	}

	cellphonePatient := model.NewCellphoneUser("82994321525")

	patient := model.NewPatient("Robson Alves", birthDate, "11642401202", "M", "Rua Fictícia da Silva",
		"robson@gmail.com", "33333", true, cellphonePatient)

	success, err := dao.PatientInsert(patient)
	if err != nil {
		panic(err)
	}

	if success == true {
		println("O cadastro foi realizado com sucesso!")
	} else {
		println("O cadastro não foi realizado!")
	}

	return success
}

func PatientRegisterEdit(patient model.Patient) bool {

	success, err := dao.PatientEdit(patient)
	if err != nil {
		panic(err)
	}

	if success == true {
		println("O cadastro foi alterado com sucesso!")
	} else {
		println("Não foi possível alterar o cadastro!")
	}

	return success
}

func PatientAuthenticatorLogin(email string, password string) (bool, int) {
	if email == "" || password == "" {
		return false, 0
	}

	Authorized, doctorId, err := dao.PatientValidateLogin(email, password)
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

func PatientSelectRegister(patientId int) model.Patient {
	var patient model.Patient
	var err error
	if patientId != 0 {
		patient, err = dao.PatientSelectById(patientId)
		if err != nil {
			println("Error na busca das informações do paciente: ", err.Error())
			panic(err)
		}
	} else {
		println("Informe um id de paciente válido!")
		return patient
	}
	return patient
}

func PatientRegisterOff(patientId int) bool {
	var success bool
	var err error
	if patientId != 0 {
		success, err = dao.PatientOff(patientId)
		if err != nil {
			println("Error durante o desligamento do registro do paciente: ", err.Error())
			panic(err)
		}
	} else {
		println("Informe um id de paciente válido!")
		success = false
		return success
	}

	return success
}
