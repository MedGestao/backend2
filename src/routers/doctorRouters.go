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

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor request.DoctorRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&doctor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	doctor.User.BirthDate, err = time.Parse(util.DateFormat, doctor.User.BirthDate.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doctorId, err, errorMessage := controller.DoctorRegister(doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errorMessage.Message != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
	} else {
		doctorIdResponse := response.DoctorIdResponse{Id: doctorId}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(doctorIdResponse)
	}

	//if doctorId != 0 {
	//	doctorIdResponse := response.DoctorIdResponse{Id: doctorId}
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusCreated)
	//	json.NewEncoder(w).Encode(doctorIdResponse)
	//} else {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusBadRequest)
	//	json.NewEncoder(w).Encode(response.NewErrorResponse("Não foi possível cadastrar o médico!"))
	//}
}

func GetDoctorsAll(w http.ResponseWriter, r *http.Request) {
	var doctorFilterParameters request.DoctorFilterParameters
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&doctorFilterParameters); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doctors, err := controller.DoctorSelectRegisterAll(doctorFilterParameters.DoctorName, doctorFilterParameters.SpecialtyName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna os dados do paciente no formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctors)
	w.WriteHeader(http.StatusOK)
}

func GetDoctorById(w http.ResponseWriter, r *http.Request) {
	/* 	params := mux.Vars(r)
	   	// Decodifica os dados JSON do corpo da solicitação
	   	id := params["id"]
	   	if id == "" {
	   		http.Error(w, "Id não foi informado", http.StatusBadRequest)
	   		return
	   	}

	   	doctorId, err := strconv.Atoi(id)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Chama a função que lê o paciente do banco de dados
	   	doctor, err := controller.DoctorSelectRegisterById(doctorId)
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusInternalServerError)
	   		return
	   	}

	   	// Retorna os dados do paciente no formato JSON
	   	w.Header().Set("Content-Type", "application/json")
	   	json.NewEncoder(w).Encode(doctor)
	   	w.WriteHeader(http.StatusOK) */
	var doctor response.DoctorResponse

	cellphoneUserResponse := response.CellphoneResponse{
		Number: "1223",
	}
	userResponse := response.UserResponse{
		Name:          "Test",
		ImageUrl:      "http://192.168.0.164:3001/public/upload-3517911352.png",
		CellphoneUser: cellphoneUserResponse,
	}

	specialtyUserResponse := response.SpecialtyResponse{Description: "sdhjfghjad"}
	doctor = response.DoctorResponse{
		User:      userResponse,
		Crm:       "122",
		Specialty: specialtyUserResponse,
	}

	// Retorna os dados do paciente no formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctor)
	w.WriteHeader(http.StatusOK)
}

func EditDoctor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var doctorEditRequest request.DoctorRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&doctorEditRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	doctorIdRequest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//doctorEditRequest := dataRequest.DoctorRequest

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
		json.NewEncoder(w).Encode(response.NewErrorResponse("Não foi possivel alterar o cadastro!"))
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

	authorized, doctorId, err := controller.DoctorAuthenticatorLogin(doctorEmail, doctorPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if authorized == true {
		doctorIdResponse := response.DoctorIdResponse{Id: doctorId}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(doctorIdResponse)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func ValidateEmailDoctor(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]

	isValid := controller.ValidateEmailDoctor(email)

	if isValid == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
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
		json.NewEncoder(w).Encode(response.NewErrorResponse("Não foi possível desativar o médico!"))
	}
}

func GetSpecialty(w http.ResponseWriter, r *http.Request) {

	var specialties []response.SpecialtyResponse

	specialties, err := controller.SelectSpecialties()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(specialties)
	w.WriteHeader(http.StatusOK)
}
