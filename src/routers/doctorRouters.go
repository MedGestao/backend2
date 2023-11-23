package routers

import (
	"MedGestao/src/controller"
	"MedGestao/src/request"
	"MedGestao/src/response"
	"encoding/json"
	"net/http"
	"time"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor request.DoctorRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	doctor.User.BirthDate, err = time.Parse(DateFormat, doctor.User.BirthDate.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	success, err := controller.DoctorRegister(doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Médico cadastrado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível cadastrar o médico!")
	}
}

func GetDoctorById(w http.ResponseWriter, r *http.Request) {
	// Decodifica os dados JSON do corpo da solicitação
	var idRequest request.DoctorIdRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&idRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doctorId := idRequest.Id

	// Chama a função que lê o paciente do banco de dados
	doctor, err := controller.DoctorSelectRegisterById(doctorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna os dados do paciente no formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctor)
	w.WriteHeader(http.StatusOK)
}

func EditDoctor(w http.ResponseWriter, r *http.Request) {
	var dataRequest request.EditDoctorRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dataRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	doctorIdRequest := dataRequest.DoctorIdRequest
	doctorEditRequest := dataRequest.DoctorRequest

	success, err := controller.DoctorRegisterEdit(doctorIdRequest, doctorEditRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Cadastro do médico alterado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possivel alterar o cadastro!")
	}
}

func ValidateLoginDoctor(w http.ResponseWriter, r *http.Request) {
	var validateLogin request.DoctorAuthenticatorRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&validateLogin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	doctorEmail := validateLogin.Email
	doctorPassword := validateLogin.Password

	authorized, patientId, err := controller.DoctorAuthenticatorLogin(doctorEmail, doctorPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if authorized == true {
		doctorIdResponse := response.DoctorIdResponse{Id: patientId}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctorIdResponse)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
	}
}

func DeactivateDoctor(w http.ResponseWriter, r *http.Request) {
	// Decodifica os dados JSON do corpo da solicitação
	var idRequest request.DoctorIdRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&idRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doctorId := idRequest.Id

	// Chama a função que lê o paciente do banco de dados
	success, err := controller.DoctorRegisterOff(doctorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Registro do médico desativado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível desativar o médico!")
	}
}
