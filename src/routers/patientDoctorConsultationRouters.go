package routers

import (
	"MedGestao/src/controller"
	"MedGestao/src/request"
	"MedGestao/src/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func CreatePatientDoctorConsutation(w http.ResponseWriter, r *http.Request) {

	var patientDoctorConsultation request.PatientDoctorConsultationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patientDoctorConsultation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	patientDoctorConsultation.AppointmentDate, err = time.Parse(util.DateFormat, patientDoctorConsultation.AppointmentDate.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("Id do médico da consulta: ", patientDoctorConsultation.DoctorId)
	println("Valor da consulta: ", patientDoctorConsultation.Value)

	success, err, errorMessage := controller.RegisterPatientDoctorConsultation(patientDoctorConsultation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Consulta marcada com sucesso!")
	} else if errorMessage.Message != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(errorMessage.Message)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível marcar a consulta!")
	}
}

func GetPatientDoctorConsultationAllByIdDoctor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}
	var patientDoctorConsultationRequest request.PatientDoctorConsultationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patientDoctorConsultationRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	DoctorId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	patientDoctorConsultationList, err := controller.SearchPatientDoctorConsultationAllByDoctor(DoctorId, patientDoctorConsultationRequest.AppointmentDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patientDoctorConsultationList)
	w.WriteHeader(http.StatusOK)
}

func GetPatientDoctorConsultationAllByIdPatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
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

	patientDoctorConsultationList, err := controller.SearchPatientDoctorConsultationAllByPatient(patientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patientDoctorConsultationList)
	w.WriteHeader(http.StatusOK)
}

func GetPatientDoctorConsultationById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	patientDoctorConsultationId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	patientDoctorConsultation, err := controller.SearchPatientDoctorConsultationById(patientDoctorConsultationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patientDoctorConsultation)
	w.WriteHeader(http.StatusOK)
}

func EditPatientDoctorConsultation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	var patientDoctorConsultationRequest request.PatientDoctorConsultationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patientDoctorConsultationRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//medicalScheduleIdRequest := dataRequest.MedicalScheduleIdRequest
	patientDoctorConsultationIdRequest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//medicalScheduleRequest := dataRequest.MedicalScheduleRequest

	success, err, errorMessage := controller.EditPatientDoctorConsultation(patientDoctorConsultationIdRequest, patientDoctorConsultationRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Consulta alterada com sucesso!")
	} else if errorMessage.Message != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(errorMessage.Message)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível alterar a consulta!")
	}
}

func CompletePatientDoctorConsultation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	patientDoctorConsultationId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success, err := controller.CompletePatientDoctorConsultation(patientDoctorConsultationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Consulta finalizada com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível finalizar a consulta!")
	}
}

func DeactivatePatientDoctorConsultation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	patientDoctorConsultationId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success, err := controller.DeactivatePatientDoctorConsultation(patientDoctorConsultationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Consulta cancelada com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível cancelar a consulta!")
	}
}
