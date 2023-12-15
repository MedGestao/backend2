package routers

import (
	"MedGestao/src/controller"
	"MedGestao/src/request"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const DateFormatMedicalSchedule = "2006-01-02 15:04:05 -0700 MST"

func CreateMedicalSchedule(w http.ResponseWriter, r *http.Request) {

	var medicalSchedule request.MedicalScheduleRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&medicalSchedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/* 	var err error
	   	medicalSchedule.SpecificDate, err = time.Parse(util.DateFormat, medicalSchedule.SpecificDate.String())
	   	if err != nil {
	   		http.Error(w, err.Error(), http.StatusBadRequest)
	   		return
	   	} */

	success, err, errorMessage := controller.RegisterMedicalSchedule(medicalSchedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Agendamento configurado com sucesso!")
	} else if errorMessage.Message != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(errorMessage.Message)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível configurar o agendamento!")
	}
}

func GetMedicalScheduleAllByIdDoctor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	selectedDateString := mux.Vars(r)["selectedDate"]
	selectedDay := mux.Vars(r)["selectedDay"]

	medicalScheduleId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	selectedDate, err := time.Parse("2006-01-02", selectedDateString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	medicalScheduleList, err := controller.SearchAllMedicalScheduleByIdDoctor(medicalScheduleId, selectedDate, selectedDay)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicalScheduleList)
	w.WriteHeader(http.StatusOK)
}

func GetMedicalScheduleById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "Id não foi informado", http.StatusBadRequest)
		return
	}

	medicalScheduleId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	medicalSchedule, err := controller.SearchByIdMedicalSchedule(medicalScheduleId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicalSchedule)
	w.WriteHeader(http.StatusOK)
}

func EditMedicalSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var medicalScheduleRequest request.MedicalScheduleRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&medicalScheduleRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//medicalScheduleIdRequest := dataRequest.MedicalScheduleIdRequest
	medicalScheduleIdRequest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//medicalScheduleRequest := dataRequest.MedicalScheduleRequest

	success, err := controller.EditMedicalSchedule(medicalScheduleIdRequest, medicalScheduleRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Configuração de agenda alterada com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível editar a configuração da agenda!")
	}
}
