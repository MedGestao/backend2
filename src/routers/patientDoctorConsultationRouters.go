package routers

import (
	"MedGestao/src/controller"
	"MedGestao/src/request"
	"MedGestao/src/util"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreatePatientDoctorConsutation(w http.ResponseWriter, r *http.Request) {

	var patientDoctorConsultation request.PatientDoctorConsultationRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patientDoctorConsultation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Println("[ROUTER.CreatePatientDoctorConsutation] " + request.LogPatientDoctorConsultationRequest(patientDoctorConsultation))

	var err error
	patientDoctorConsultation.AppointmentDate, err = time.Parse(util.DateFormat, patientDoctorConsultation.AppointmentDate.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	patientDoctorConsultationId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	patientDoctorConsultationList, err := controller.SearchPatientDoctorConsultationAllByDoctor(patientDoctorConsultationId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patientDoctorConsultationList)
	w.WriteHeader(http.StatusOK)
}
