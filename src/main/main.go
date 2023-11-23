package main

import "MedGestao/src/server"

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
	//cellphonePatient := model.NewCellphoneUser("829943219211")
	//
	//patient := model.NewPatient("Martinho Lutero", birthDate, "11642462533", "M", "Rua Fictícia da Silva",
	//	"martinho@gmail.com", "l2#Pm12", true, cellphonePatient)
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
	//success, patientId := controller.PatientAuthenticatorLogin("martinho@gmail.com", "l2#Pm12")
	//if success == true && patientId != 0 {
	//	println("Seja bem vindo!")
	//	println("Id do usuário: ", patientId)
	//} else {
	//	println("O email ou a senha estão incorretos!")
	//}
	//===================================================================================================

	//Buscar dados do paciente pelo id
	//patient := controller.PatientSelectByIdRegister(24)
	//println("Nome do paciente: ", patient.GetUser().GetName())
	//println("Data de nascimento: ", patient.GetUser().GetBirthDate().Format("02/01/2006"))
	//println("Sexo: ", patient.GetUser().GetSex())
	//println("CPF: ", patient.GetUser().GetCpf())
	//println("Endereço: ", patient.GetUser().GetAddress())
	//println("Email: ", patient.GetUser().GetEmail())
	//=================================================================================================

	//Editar Paciente
	//birthDate, err := time.Parse("2006-01-02", "1998-08-04")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//cellphonePatient := model.NewCellphoneUser("829971442")
	//
	//patient := model.NewPatient("João Guilherme", birthDate, "93239797435", "M", "Rua Dom Carlos",
	//	"joaoGuilherme@gmail.com", "l7#T11", true, cellphonePatient)
	//patient.SetUserId(24)
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
	//success := controller.PatientRegisterOff(25)
	//if success == true {
	//	println("Paciente desligado com sucesso!")
	//} else {
	//	println("Erro ao tentar desligar o paciente!")
	//}
	//==============================================================================================

	//CONFIGURAR AGENDA DO MÉDICO

	//Inserir Agenda
	//medicalSchedule := model.NewMedicalSchedule(7, "03", "", "09:00", "17:00",
	//	"2023")
	//success, err := controller.RegisterMedicalSchedule(medicalSchedule)
	//if err != nil {
	//	println("Excessão lançada. Erro gerado:", err.Error())
	//}
	//
	//if success == true {
	//	println("Agenda cadastrada com sucesso!")
	//} else {
	//	println("Erro ao cadastrar agenda!")
	//}

	//Buscar Agendas
	//medicalScheduleList, err := controller.SearchAllMedicalSchedule()
	//if err != nil {
	//	println("Excessão lançada. Erro gerado:", err.Error())
	//}else {
	//	for i, medicalSchedule := range medicalScheduleList {
	//		println("Agenda", i+1)
	//		println("Horário inicial do atendimento:", medicalSchedule.GetStartTime())
	//		println("Horário final do atendimento:", medicalSchedule.GetFinalTime())
	//		println("Dia de serviço:", medicalSchedule.GetDayOfService())
	//		println("Ano do atendimento:", medicalSchedule.GetYear())
	//		println("")
	//	}
	//}

	//Buscar agenda pelo id
	//medicalSchedule, err := controller.SearchByIdMedicalSchedule(1)
	//if err != nil {
	//	println("Excessão lançada. Error gerado:", err.Error())
	//} else if medicalSchedule == (model.MedicalSchedule{}) {
	//	println("Agenda não encontrada!")
	//} else {
	//	println("Agenda")
	//	println("Horário inicial do atendimento:", medicalSchedule.GetStartTime())
	//	println("Horário final do atendimento:", medicalSchedule.GetFinalTime())
	//	println("Dia de serviço:", medicalSchedule.GetDayOfService())
	//	println("Ano do atendimento:", medicalSchedule.GetYear())
	//}

	//Editar agenda
	//medicalSchedule := model.NewMedicalSchedule(7, "05", "", "08:00", "15:00",
	//	"2023")
	//medicalSchedule.SetId(2)
	//success, err := controller.EditMedicalSchedule(medicalSchedule)
	//if err != nil {
	//	println("Excessão lançada. Erro gerado:", err.Error())
	//}
	//
	//if success == true {
	//	println("Agenda editada com sucesso!")
	//} else {
	//	println("Erro ao editar agenda!")
	//}

	//Apagar Agenda
	//success, err := controller.OffMedicalSchedule(1)
	//if err != nil {
	//	println("Excessão lançada:", err.Error())
	//}
	//
	//if success == true {
	//	println("Agenda desligada com sucesso!")
	//} else {
	//	println("Erro ao desligar agenda!")
	//}

	//server.OpenServer()

	server.OpenServerTest()
}
