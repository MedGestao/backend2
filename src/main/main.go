package main

import (
	"MedGestao/src/controller"
	"MedGestao/src/model"
	"time"
)

func main() {
	//controller.CellphoneRegister()
	//controller.DoctorRegister()
	//fmt.Println("Cadastrado com sucesso!")

	//success, doctorId := controller.AuthenticatorLoginDoctor("anaPaula3@gmail.com", "Ap#144")
	//if success == true && doctorId != 0 {
	//	println("Seja bem vindo!")
	//	println("Id do usuário: ", doctorId)
	//} else {
	//	println("O email ou a senha estão incorretos!")
	//}

	//birthDate, err := time.Parse("2006-01-02", "1988-09-22")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//cellphonePatient := model.NewCellphoneUser("82994321572")
	//
	//patient := model.NewPatient("TESTE5", birthDate, "222", "M", "Rua Fictícia da Silva",
	//	"rA@gmail.com", "mL#42", true, cellphonePatient)
	//patient.SetUserId(22)
	//println("Id: ", patient.GetUser().GetId())
	//println("Telefone: ", patient.GetUser().GetCellphoneUser().GetNumber())
	//success := controller.PatientRegisterEdit(patient)
	//if success == true {
	//	println("Cadastro editado com sucesso!")
	//} else {
	//	println("Ocorreu um erro ao tentar editar o cadastro!")
	//}

	//patient := controller.PatientSelectRegister(21)
	//println("Nome do paciente: ", patient.GetUser().GetName())
	//println("Data de nascimento: ", patient.GetUser().GetBirthDate().Format("01/02/2006"))
	//println("Sexo: ", patient.GetUser().GetSex())
	//println("CPF: ", patient.GetUser().GetCpf())
	//println("Endereço: ", patient.GetUser().GetAddress())
	//println("Email: ", patient.GetUser().GetEmail())

	//success := controller.PatientRegisterOff(21)
	//if success == true {
	//	println("Paciente desligado com sucesso!")
	//} else {
	//	println("Error ao tentar desligar o paciente!")
	//}

	//doctor := controller.DoctorSelectRegister(6)
	//println("Nome do paciente: ", doctor.GetUser().GetName())
	//println("Data de nascimento: ", doctor.GetUser().GetBirthDate().Format("01/02/2006"))
	//println("Sexo: ", doctor.GetUser().GetSex())
	//println("CPF: ", doctor.GetUser().GetCpf())
	//println("Endereço: ", doctor.GetUser().GetAddress())
	//println("Email: ", doctor.GetUser().GetEmail())
	//println("Cns: ", doctor.GetCns())
	//println("Crm: ", doctor.GetCrm())
	//println("Especialidade: ", doctor.GetSpecialty().GetDescription())

	//success := controller.DoctorRegisterOff(6)
	//if success == true {
	//	println("Médico desligado com sucesso!")
	//} else {
	//	println("Error ao tentar desligar o médico!")
	//}

	birthDate, err := time.Parse("2006-01-02", "1992-07-05")
	if err != nil {
		panic(err)
	}

	cellphoneDoctor := model.NewCellphoneUser("82996426813")

	specialty := model.NewSpecialty("Clínica Geral")

	//success, err := dao.InsertSpecialty(specialty)
	//if err != nil {
	//	panic(err)
	//}

	doctor := model.NewDoctor("Ana Carla", birthDate, "45567599203", "F", "Rua Fictícia dos Santos",
		"anaCarla@gmail.com", cellphoneDoctor, "Ap#144", true, "5940377689", "2342584354", specialty)
	doctor.SetUserId(5)

	success := controller.DoctorRegisterEdit(doctor)
	if success == true {
		println("Cadastro editado com sucesso!")
	} else {
		println("Ocorreu um erro ao tentar editar o cadastro!")
	}

	//server.OpenServer()
}
