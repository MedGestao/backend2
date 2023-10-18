package main

import "MedGestao/src/controller"

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

	success := controller.PatientRegisterEdit()
	if success == true {
		println("Seja bem vindo!")
	} else {
		println("O email ou a senha estão incorretos!")
	}
	//server.OpenServer()
}
