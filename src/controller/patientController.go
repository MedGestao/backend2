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

func PatientRegisterEdit() bool {
	birthDate, err := time.Parse("2006-01-02", "1988-09-22")

	if err != nil {
		fmt.Println(err)
	}

	cellphonePatient := model.NewCellphoneUser("82994321567")

	patient := model.NewPatient("Martinho Lutero", birthDate, "11642401202", "M", "Rua Fictícia da Silva",
		"martinhoLutero@gmail.com", "mL#42", true, cellphonePatient)
	//patient.SetUser(pa)
	println("Id: ", patient.GetUser().GetId())

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

func AuthenticatorLoginPatient(email string, password string) (bool, int) {
	if email == "" || password == "" {
		return false, 0
	}

	Authorized, doctorId, err := dao.ValidateLoginPatient(email, password)
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
