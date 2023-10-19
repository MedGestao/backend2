package main

func main() {

	//MÉDICO

	//Cadastrar Médico
	//birthDate, err := time.Parse("2006-01-02", "1992-07-05")
	//if err != nil {
	//	println(err.Error())
	//}
	//
	//cellphoneDoctor := model.NewCellphoneUser("82996426813")
	//
	//specialty := model.NewSpecialty("Clínica Geral")
	////
	////success, err := dao.InsertSpecialty(specialty)
	////if err != nil {
	////	panic(err)
	////}
	//
	//doctor := model.NewDoctor("Carlos Henrique", birthDate, "45567599203", "M", "Rua Fictícia dos Santos",
	//	"carlosHenrique@gmail.com", cellphoneDoctor, "Ap#142", true, "2342584354", specialty)
	//success := controller.DoctorRegister(doctor)
	//
	//if success == true {
	//	println("O cadastro foi realizado com sucesso!")
	//} else {
	//	println("O cadastro não foi realizado!")
	//}
	//=====================================================================

	//Autenticar Médico
	//success, doctorId := controller.DoctorAuthenticatorLogin("carlosHenrique@gmail.com", "Ap#142")
	//if success == true && doctorId != 0 {
	//	println("Seja bem vindo!")
	//	println("Id do usuário: ", doctorId)
	//} else {
	//	println("O email ou a senha estão incorretos!")
	//}
	//========================================================================

	//Buscar dados do médico pelo id
	//doctor := controller.DoctorSelectRegister(7)
	//println("Nome do paciente: ", doctor.GetUser().GetName())
	//println("Data de nascimento: ", doctor.GetUser().GetBirthDate().Format("01/02/2006"))
	//println("Sexo: ", doctor.GetUser().GetSex())
	//println("CPF: ", doctor.GetUser().GetCpf())
	//println("Endereço: ", doctor.GetUser().GetAddress())
	//println("Email: ", doctor.GetUser().GetEmail())
	//println("Crm: ", doctor.GetCrm())
	//println("Especialidade: ", doctor.GetSpecialty().GetDescription())
	//===============================================================================

	//Editar Médico
	//birthDate, err := time.Parse("2006-01-02", "1992-07-05")
	//if err != nil {
	//	panic(err)
	//}
	//
	//cellphoneDoctor := model.NewCellphoneUser("82996426813")
	//
	//specialty := model.NewSpecialty("Clínica Geral")
	//
	////success, err := dao.InsertSpecialty(specialty)
	////if err != nil {
	////	panic(err)
	////}
	//
	//doctor := model.NewDoctor("Ana", birthDate, "45567599203", "F", "Rua Fictícia dos Santos",
	//	"ana@gmail.com", cellphoneDoctor, "Ap#146", true, "2342584354", specialty)
	//doctor.SetUserId(5)
	//
	//success := controller.DoctorRegisterEdit(doctor)
	//if success == true {
	//	println("Cadastro editado com sucesso!")
	//} else {
	//	println("Ocorreu um erro ao tentar editar o cadastro!")
	//}
	//===================================================================================================

	//Desligar Médico
	//success := controller.DoctorRegisterOff(7)
	//if success == true {
	//	println("Médico desligado com sucesso!")
	//} else {
	//	println("Error ao tentar desligar o médico!")
	//}
	//===================================================================================================

	//==============================================================================================================

	//PACIENTE

	//Cadastrar Paciente
	//birthDate, err := time.Parse("2006-01-02", "1982-09-22")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//cellphonePatient := model.NewCellphoneUser("82994321880")
	//
	//patient := model.NewPatient("Samara Ferreira", birthDate, "11642462533", "F", "Rua Fictícia da Silva",
	//	"samara@gmail.com", "33333", true, cellphonePatient)
	//
	//success := controller.PatientRegister(patient)
	//
	//if success == true {
	//	println("O cadastro foi realizado com sucesso!")
	//} else {
	//	println("O cadastro não foi realizado!")
	//}
	//=================================================================================================

	//Autenticar Paciente
	//success, patientId := controller.PatientAuthenticatorLogin("samara@gmail.com", "33333")
	//if success == true && patientId != 0 {
	//	println("Seja bem vindo!")
	//	println("Id do usuário: ", patientId)
	//} else {
	//	println("O email ou a senha estão incorretos!")
	//}
	//===================================================================================================

	//Buscar dados do paciente pelo id
	//patient := controller.PatientSelectRegister(23)
	//println("Nome do paciente: ", patient.GetUser().GetName())
	//println("Data de nascimento: ", patient.GetUser().GetBirthDate().Format("01/02/2006"))
	//println("Sexo: ", patient.GetUser().GetSex())
	//println("CPF: ", patient.GetUser().GetCpf())
	//println("Endereço: ", patient.GetUser().GetAddress())
	//println("Email: ", patient.GetUser().GetEmail())
	//=================================================================================================

	//Editar Paciente
	//birthDate, err := time.Parse("2006-01-02", "1988-09-22")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//cellphonePatient := model.NewCellphoneUser("82994321572")
	//
	//patient := model.NewPatient("TESTE6", birthDate, "44444444433", "M", "Rua Fictícia da Silva",
	//	"teste6@gmail.com", "t6#87", true, cellphonePatient)
	//patient.SetUserId(22)
	//println("Id: ", patient.GetUser().GetId())
	//println("Telefone: ", patient.GetUser().GetCellphoneUser().GetNumber())
	//success := controller.PatientRegisterEdit(patient)
	//if success == true {
	//	println("Cadastro editado com sucesso!")
	//} else {
	//	println("Ocorreu um erro ao tentar editar o cadastro!")
	//}
	//==============================================================================================

	//Desligar Paciente
	//success := controller.PatientRegisterOff(23)
	//if success == true {
	//	println("Paciente desligado com sucesso!")
	//} else {
	//	println("Error ao tentar desligar o paciente!")
	//}
	//==============================================================================================

	//server.OpenServer()
}
