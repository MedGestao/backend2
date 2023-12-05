package routers

import (
	"MedGestao/src/controller"
	"MedGestao/src/request"
	"MedGestao/src/util"
	"encoding/json"
	"net/http"
	"time"
)

const DateFormatMedicalSchedule = "2006-01-02 15:04:05 -0700 MST"

func CreateMedicalSchedule(w http.ResponseWriter, r *http.Request) {

	var medicalSchedule request.MedicalScheduleRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&medicalSchedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	medicalSchedule.SpecificDate, err = time.Parse(util.DateFormat, medicalSchedule.SpecificDate.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	println("Id do médico da consulta: ", medicalSchedule.DoctorId.Id)
	println("Valor da consulta: ", medicalSchedule.QueryValue)

	success, err := controller.RegisterMedicalSchedule(medicalSchedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success == true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Agendamento configurado com sucesso!")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("Não foi possível configurar o agendamento!")
	}
}
