package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/response"
	"strconv"
	"strings"
	"time"
)

func PatientDoctorConsultationInsert(patientDoctorConsultation model.PatientDoctorConsultation) (bool, error, response.ErrorResponse) {
	db, err := connection.NewConnection()
	var success bool
	var errorMessage response.ErrorResponse
	if err != nil {
		return success, err, errorMessage
	}
	defer db.Close()

	sql := "select count(*) from patient_doctor_consultation where doctor_id=$1 and appointment_date=$2 and appointment_time=$3"
	if err != err {
		return success, err, errorMessage
	}
	_, err = db.Prepare(sql)
	if err != nil {
		return success, err, errorMessage
	}

	rows, err := db.Query(sql,
		patientDoctorConsultation.GetDoctorId(),
		patientDoctorConsultation.GetAppointmentDate(),
		patientDoctorConsultation.GetAppointmentTime())
	if err != nil {
		return success, err, errorMessage
	}

	var quantityDB int
	for rows.Next() {
		err = rows.Scan(&quantityDB)
		if err != nil {
			return success, err, errorMessage
		}
	}

	sql = "select schedule_limit from medical_schedule where doctor_id=$1"
	_, err = db.Prepare(sql)
	if err != nil {
		return success, err, errorMessage
	}

	rows, err = db.Query(sql,
		patientDoctorConsultation.GetDoctorId())
	if err != nil {
		return success, err, errorMessage
	}

	var limitDB int
	for rows.Next() {
		err = rows.Scan(&limitDB)
		if err != nil {
			return success, err, errorMessage
		}
	}

	if quantityDB >= limitDB {
		errorMessage = response.NewErrorResponse("Não a mais disponibilidade de vaga para a esse horário!")
		return success, err, errorMessage
	}

	sql = "insert into patient_doctor_consultation(patient_id, doctor_id, appointment_date, appointment_time, status, value)" +
		"values($1, $2, $3, $4, true, $5)"
	_, err = db.Prepare(sql)
	if err != nil {
		return success, err, errorMessage
	}

	_, err = db.Exec(sql,
		patientDoctorConsultation.GetPatientId(),
		patientDoctorConsultation.GetDoctorId(),
		patientDoctorConsultation.GetAppointmentDate(),
		patientDoctorConsultation.GetAppointmentTime(),
		patientDoctorConsultation.GetValue(),
	)
	if err != nil {
		return success, err, errorMessage
	}

	success = true
	return success, err, errorMessage
}

func PatientDoctorConsultationAllByIdDoctor(doctorId int) ([]response.PatientDoctorConsultationByDoctorResponse, error) {
	var patientDoctorConsultationByDoctorResponseList []response.PatientDoctorConsultationByDoctorResponse
	db, err := connection.NewConnection()
	if err != nil {
		return patientDoctorConsultationByDoctorResponseList, err
	}
	defer db.Close()
	sql := "select pdc.id, p.name, pdc.appointment_date, pdc.appointment_time, pdc.value from patient_doctor_consultation pdc " +
		"inner join patient p on pdc.patient_id = p.id " +
		"inner join doctor d on pdc.doctor_id = d.id " +
		"where d.id = $1 and d.active is true and pdc.status is true"

	_, err = db.Prepare(sql)
	if err != nil {
		return patientDoctorConsultationByDoctorResponseList, err
	}

	rows, err := db.Query(sql,
		doctorId)
	if err != nil {
		return patientDoctorConsultationByDoctorResponseList, err
	}

	var patientDoctorConsultationIdDB int
	var patientNameDB, appointmentTimeDB string
	var appointmentDateDB time.Time
	var queryValueDB float64
	var strValue string
	for rows.Next() {
		err = rows.Scan(&patientDoctorConsultationIdDB, &patientNameDB, &appointmentDateDB, &appointmentTimeDB, &strValue)
		if err != nil {
			return patientDoctorConsultationByDoctorResponseList, err
		}

		strValue = strings.ReplaceAll(strings.ReplaceAll(strValue, ",", ""), "$", "")
		queryValueDB, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return patientDoctorConsultationByDoctorResponseList, err
		}

		pdc := response.PatientDoctorConsultationByDoctorResponse{
			PatientName: patientNameDB,
			PatientDoctorConsultationResponse: response.PatientDoctorConsultationResponse{
				Id:              patientDoctorConsultationIdDB,
				AppointmentDate: appointmentDateDB,
				AppointmentTime: appointmentTimeDB,
				Value:           queryValueDB,
			},
		}

		patientDoctorConsultationByDoctorResponseList = append(patientDoctorConsultationByDoctorResponseList, pdc)
	}

	return patientDoctorConsultationByDoctorResponseList, err
}
