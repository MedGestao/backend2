package routers

import (
	"MedGestao/src/controller"
	"MedGestao/src/request"
	"MedGestao/src/response"
	"MedGestao/src/util"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	// Decodifica os dados JSON do corpo da solicitação
	//var patient model.Patient
	var patient request.PatientRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//birthDate, err := time.Parse("2006-01-02", "1992-07-05")
	//patient.User.BirthDate = patient.User.BirthDate.Format(DateFormat)
	var err error
	patient.User.BirthDate, err = time.Parse(util.DateFormat, patient.User.BirthDate.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("Email do paciente: ", patient.User.Email)
	println("Telefone do paciente: ", patient.User.CellphoneUser.Number)
	println("Data de nascimento:", patient.User.BirthDate.String())

	// Aqui, você pode realizar a lógica de negócios para criar o paciente no banco de dados
	// e, em seguida, retornar uma resposta adequada

	// Por exemplo, retornar o paciente criado em formato JSON
	success, err := controller.PatientRegister(patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Paciente cadastrado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível cadastrar o paciente!")
	}
}

func GetPatientById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Decodifica os dados JSON do corpo da solicitação
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	patientId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Chama a função que lê o paciente do banco de dados
	patient, err := controller.PatientSelectByIdRegister(patientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna os dados do paciente no formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
	w.WriteHeader(http.StatusOK)
}

func EditPatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var patientEditRequest request.PatientRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patientEditRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	patientIdRequest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//patientEditRequest := dataRequest.PatientRequest

	success, err := controller.PatientRegisterEdit(patientIdRequest, patientEditRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Cadastro do paciente alterado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível editar o cadastro do paciente!")
	}
}

func ValidateLoginPatient(w http.ResponseWriter, r *http.Request) {
	var validateLogin request.PatientAuthenticatorRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&validateLogin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	patientEmail := validateLogin.Email
	patientPassword := validateLogin.Password

	authorized, patientId, err := controller.PatientAuthenticatorLogin(patientEmail, patientPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if authorized == true {
		patientIdResponse := response.PatientIdResponse{Id: patientId}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(patientIdResponse)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeactivatePatient(w http.ResponseWriter, r *http.Request) {
	// Decodifica os dados JSON do corpo da solicitação
	var idRequest request.PatientIdRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&idRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//vars := mux.Vars(r)
	//idRequest := vars["id"]
	//patientId, err := strconv.Atoi(idRequest)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	patientId := idRequest.Id

	// Chama a função que lê o paciente do banco de dados
	success, err := controller.PatientRegisterOff(patientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Registro de paciente desativado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível desativar o paciente!")
	}
}
