package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
	"time"
)

func DoctorRegister() bool {
	birthDate, err := time.Parse("2006-01-02", "1992-07-05")

	cellphoneDoctor := model.NewCellphoneUser("82996426813")

	specialty := model.NewSpecialty("Clínica Geral")

	success, err := dao.InsertSpecialty(specialty)
	if err != nil {
		panic(err)
	}

	doctor := model.NewDoctor("Ana Paula", birthDate, "45567599203", "F", "Rua Fictícia dos Santos",
		"anaPaula3@gmail.com", cellphoneDoctor, "Ap#144", true, "5940377689", "2342584354", specialty)

	success, err = dao.InsertDoctor(doctor)
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

func AuthenticatorLoginDoctor(email string, password string) (bool, int) {
	if email == "" || password == "" {
		return false, 0
	}

	Authorized, doctorId, err := dao.ValidateLoginDoctor(email, password)
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

	if success == true {
		println("O cadastro foi alterado com sucesso!")
	} else {
		println("Não foi possível alterar o cadastro!")
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
