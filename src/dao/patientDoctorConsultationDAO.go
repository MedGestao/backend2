package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/response"
	"MedGestao/src/util"
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
		"values($1, $2, $3, $4, 'A', $5)"
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

func PatientDoctorConsultationAllByIdDoctor(doctorId int, appointmentDate time.Time) ([]response.PatientDoctorConsultationByDoctorResponse, error) {
	var patientDoctorConsultationByDoctorResponseList []response.PatientDoctorConsultationByDoctorResponse
	db, err := connection.NewConnection()
	if err != nil {
		return patientDoctorConsultationByDoctorResponseList, err
	}
	defer db.Close()
	sql := "select pdc.id, p.name, p.birthdate, extract(year from age(current_date, p.birthdate)) as age, pdc.appointment_date, pdc.appointment_time, pdc.value, pdc.status from patient_doctor_consultation pdc " +
		"inner join patient p on pdc.patient_id = p.id " +
		"inner join doctor d on pdc.doctor_id = d.id " +
		"where d.id = $1 and d.active is true and pdc.status != 'C' and p.active is true "

	if !appointmentDate.IsZero() {
		sql += "and pdc.appointment_date = $2"
	}
	_, err = db.Prepare(sql)
	if err != nil {
		return patientDoctorConsultationByDoctorResponseList, err
	}

	if appointmentDate.IsZero() {
		util.GlobalRows, err = db.Query(sql,
			doctorId)
		if err != nil {
			return patientDoctorConsultationByDoctorResponseList, err
		}
	} else {
		util.GlobalRows, err = db.Query(sql,
			doctorId,
			appointmentDate)
		if err != nil {
			return patientDoctorConsultationByDoctorResponseList, err
		}
	}

	var patientDoctorConsultationIdDB int
	var patientNameDB, ageDB, appointmentTimeDB, statusDB string
	var birthdatePatientDB, appointmentDateDB time.Time
	var queryValueDB float64
	var strValue string
	for util.GlobalRows.Next() {
		err = util.GlobalRows.Scan(&patientDoctorConsultationIdDB, &patientNameDB, &birthdatePatientDB, &ageDB, &appointmentDateDB, &appointmentTimeDB, &strValue, &statusDB)
		if err != nil {
			return patientDoctorConsultationByDoctorResponseList, err
		}

		strValue = strings.ReplaceAll(strings.ReplaceAll(strValue, ",", ""), "$", "")
		queryValueDB, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return patientDoctorConsultationByDoctorResponseList, err
		}

		pdc := response.PatientDoctorConsultationByDoctorResponse{
			Name:      patientNameDB,
			BirthDate: birthdatePatientDB,
			Age:       ageDB,
			PatientDoctorConsultationResponse: response.PatientDoctorConsultationResponse{
				Id:              patientDoctorConsultationIdDB,
				AppointmentDate: appointmentDateDB,
				AppointmentTime: appointmentTimeDB,
				Value:           queryValueDB,
				Status:          statusDB,
			},
		}

		patientDoctorConsultationByDoctorResponseList = append(patientDoctorConsultationByDoctorResponseList, pdc)
	}

	return patientDoctorConsultationByDoctorResponseList, err
}

func PatientDoctorConsultationAllByIdPatient(patientId int) ([]response.PatientDoctorConsultationByDoctorResponse, error) {
	var patientDoctorConsultationByPatientResponseList []response.PatientDoctorConsultationByDoctorResponse
	db, err := connection.NewConnection()
	if err != nil {
		return patientDoctorConsultationByPatientResponseList, err
	}
	defer db.Close()
	sql := "select pdc.id, d.name, d.image_url, sp.description, pdc.appointment_date, pdc.appointment_time, pdc.value from patient_doctor_consultation pdc " +
		"inner join patient p on pdc.patient_id = p.id " +
		"inner join doctor d on pdc.doctor_id = d.id " +
		"inner join medical_specialty m on m.doctor_id = d.id " +
		"inner join specialty sp on m.specialty_id = sp.id " +
		"where p.id = $1 and d.active is true and p.active is true"

	_, err = db.Prepare(sql)
	if err != nil {
		return patientDoctorConsultationByPatientResponseList, err
	}

	rows, err := db.Query(sql,
		patientId)
	if err != nil {
		return patientDoctorConsultationByPatientResponseList, err
	}

	var patientDoctorConsultationIdDB int
	var doctorNameDB, doctorSpecialtyDB, doctorImgDB, appointmentTimeDB string
	var appointmentDateDB time.Time
	var queryValueDB float64
	var strValue string
	for rows.Next() {
		err = rows.Scan(&patientDoctorConsultationIdDB, &doctorNameDB, &doctorImgDB, &doctorSpecialtyDB, &appointmentDateDB, &appointmentTimeDB, &strValue)
		if err != nil {
			return patientDoctorConsultationByPatientResponseList, err
		}

		strValue = strings.ReplaceAll(strings.ReplaceAll(strValue, ",", ""), "$", "")
		queryValueDB, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return patientDoctorConsultationByPatientResponseList, err
		}

		pdc := response.PatientDoctorConsultationByDoctorResponse{
			Name:      doctorNameDB,
			Specialty: doctorSpecialtyDB,
			ImageUrl:  doctorImgDB,
			PatientDoctorConsultationResponse: response.PatientDoctorConsultationResponse{
				Id:              patientDoctorConsultationIdDB,
				AppointmentDate: appointmentDateDB,
				AppointmentTime: appointmentTimeDB,
				Value:           queryValueDB,
			},
		}

		patientDoctorConsultationByPatientResponseList = append(patientDoctorConsultationByPatientResponseList, pdc)
	}

	return patientDoctorConsultationByPatientResponseList, err
}

func PatientDoctorConsultationById(patientDoctorConsultationId int) (response.PatientDoctorConsultationResponse, error) {
	var patientDoctorConsultationResponse response.PatientDoctorConsultationResponse
	db, err := connection.NewConnection()
	if err != nil {
		return patientDoctorConsultationResponse, err
	}
	defer db.Close()

	sql := "select pdc.id, pdc.patient_id, pdc.doctor_id, pdc.appointment_date, pdc.appointment_time, pdc.value from patient_doctor_consultation pdc " +
		"inner join patient p on pdc.patient_id = p.id " +
		"inner join doctor d on pdc.doctor_id = d.id " +
		"where pdc.id = $1 and d.active is true and pdc.status != 'C' and p.active is true"

	_, err = db.Prepare(sql)
	if err != nil {
		return patientDoctorConsultationResponse, err
	}

	rows, err := db.Query(sql,
		patientDoctorConsultationId)
	if err != nil {
		return patientDoctorConsultationResponse, err
	}

	var patientDoctorConsultationIdDB, patientIdDB, doctorIdDB int
	var appointmentTimeDB string
	var appointmentDateDB time.Time
	var queryValueDB float64
	var strValue string
	for rows.Next() {
		err = rows.Scan(&patientDoctorConsultationIdDB, &patientIdDB, &doctorIdDB, &appointmentDateDB, &appointmentTimeDB, &strValue)
		if err != nil {
			return patientDoctorConsultationResponse, err
		}

		strValue = strings.ReplaceAll(strings.ReplaceAll(strValue, ",", ""), "$", "")
		queryValueDB, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return patientDoctorConsultationResponse, err
		}

		patientDoctorConsultationResponse = response.PatientDoctorConsultationResponse{
			Id:              patientDoctorConsultationIdDB,
			PatientId:       patientIdDB,
			DoctorId:        doctorIdDB,
			AppointmentDate: appointmentDateDB,
			AppointmentTime: appointmentTimeDB,
			Value:           queryValueDB,
		}
	}

	return patientDoctorConsultationResponse, err
}

func PatientDoctorConsultationEdit(patientDoctorConsultationId int, patientDoctorConsultation model.PatientDoctorConsultation) (bool, error, response.ErrorResponse) {
	var success bool
	var errorMessage response.ErrorResponse
	db, err := connection.NewConnection()
	if err != nil {
		return success, err, errorMessage
	}
	defer db.Close()

	sql := "select count(id) from patient_doctor_consultation " +
		"where doctor_id=$1 and appointment_date=$2 and appointment_time=$3"
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

	sql = "update patient_doctor_consultation set appointment_date=$1, appointment_time=$2 where id=$3"

	_, err = db.Prepare(sql)
	if err != nil {
		return success, err, errorMessage
	}

	_, err = db.Exec(sql,
		patientDoctorConsultation.GetAppointmentDate(),
		patientDoctorConsultation.GetAppointmentTime(),
		patientDoctorConsultationId)
	if err != nil {
		return success, err, errorMessage
	}

	success = true
	return success, err, errorMessage
}

func EndPatientDoctorConsultation(patientDoctorConsultationId int) (bool, error) {
	var success bool
	db, err := connection.NewConnection()
	if err != nil {
		return success, err
	}
	defer db.Close()

	sql := "update patient_doctor_consultation set status = 'F' where id=$1"

	_, err = db.Prepare(sql)
	if err != nil {
		return success, err
	}

	_, err = db.Exec(sql,
		patientDoctorConsultationId)
	if err != nil {
		return success, err
	}

	success = true
	return success, err
}

func PatientDoctorConsultationOff(patientDoctorConsultationId int) (bool, error) {
	var success bool
	db, err := connection.NewConnection()
	if err != nil {
		return success, err
	}
	defer db.Close()

	sql := "update patient_doctor_consultation set status = 'C' where id=$1"

	_, err = db.Prepare(sql)
	if err != nil {
		return success, err
	}

	_, err = db.Exec(sql,
		patientDoctorConsultationId)
	if err != nil {
		return success, err
	}

	success = true
	return success, err
}
