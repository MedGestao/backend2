package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	sql2 "database/sql"
	"strconv"
	"strings"
	"time"
)

func MedicalScheduleInsert(medicalSchedule model.MedicalSchedule) (bool, error) {
	db, err := connection.NewConnection()
	var success bool
	if err != nil {
		return success, err
	}
	defer db.Close()

	sql := "insert into medical_schedule(doctor_id, day_of_service, start_time, final_time, specific_date, year, active, " +
		"registration_date, query_value) values($1, $2, $3, $4, $5, $6, true, current_timestamp, $7)"
	if err != err {
		return success, err
	}
	_, err = db.Prepare(sql)
	if err != nil {
		return success, err
	}

	_, err = db.Exec(sql,
		medicalSchedule.GetDoctorId(),
		medicalSchedule.GetDayOfService(),
		medicalSchedule.GetStartTime(),
		medicalSchedule.GetFinalTime(),
		medicalSchedule.GetSpecificDate(),
		medicalSchedule.GetYear(),
		medicalSchedule.GetQueryValue())
	if err != nil {
		return success, err
	}

	success = true
	return success, err
}

func MedicalScheduleSelectAll() ([]model.MedicalSchedule, error) {
	var medicalSchedule model.MedicalSchedule
	var medicalScheduleList []model.MedicalSchedule
	db, err := connection.NewConnection()
	if err != nil {
		return medicalScheduleList, err
	}
	defer db.Close()
	sql := "select id, doctor_id, day_of_service, start_time, final_time, specific_date, year, query_value from medical_schedule where active is true"

	_, err = db.Prepare(sql)
	if err != nil {
		return medicalScheduleList, err
	}

	rows, err := db.Query(sql)
	if err != nil {
		return medicalScheduleList, err
	}

	var idDB, doctorIdDB int
	var dayOfServiceDB, startTimeDB, finalTimeDB, yearDB string
	var specificDateDB time.Time
	var queryValueDB float64
	var strValue string
	for rows.Next() {
		err = rows.Scan(&idDB, &doctorIdDB, &dayOfServiceDB, &startTimeDB, &finalTimeDB, &specificDateDB, &yearDB, &strValue)
		if err != nil {
			return medicalScheduleList, err
		}

		strValue = strings.ReplaceAll(strings.ReplaceAll(strValue, ",", ""), "$", "")
		queryValueDB, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return medicalScheduleList, err
		}
		medicalSchedule = model.NewMedicalSchedule(doctorIdDB, queryValueDB, dayOfServiceDB, specificDateDB, startTimeDB, finalTimeDB, yearDB)
		if medicalSchedule != (model.MedicalSchedule{}) {
			medicalSchedule.SetId(idDB)
			medicalScheduleList = append(medicalScheduleList, medicalSchedule)
		}
	}

	return medicalScheduleList, err
}

func MedicalScheduleSelectById(medicalScheduleId int) (model.MedicalSchedule, error) {
	var medicalSchedule model.MedicalSchedule
	db, err := connection.NewConnection()
	if err != nil {
		return medicalSchedule, err
	}
	defer db.Close()

	sql := "select doctor_id, day_of_service, start_time, final_time, specific_date, year, query_value from medical_schedule " +
		"where id = $1 and active is true"
	_, err = db.Prepare(sql)
	if err != nil {
		return medicalSchedule, err
	}

	rows, err := db.Query(sql, medicalScheduleId)
	if err != nil {
		return medicalSchedule, err
	}

	var doctorIdDB int
	var dayOfServiceDB, startTimeDB, finalTimeDB, yearDB string
	var specificDateDB time.Time
	var queryValueDB float64
	var strValue string
	var specificDateNull sql2.NullTime
	for rows.Next() {
		err = rows.Scan(&doctorIdDB, &dayOfServiceDB, &startTimeDB, &finalTimeDB, &specificDateNull, &yearDB, &strValue)
		if err != nil {
			return medicalSchedule, err
		}
	}

	strValue = strings.ReplaceAll(strings.ReplaceAll(strValue, ",", ""), "$", "")
	queryValueDB, err = strconv.ParseFloat(strValue, 64)
	if err != nil {
		return medicalSchedule, err
	}

	if specificDateNull.Valid {
		specificDateDB = specificDateNull.Time
		medicalSchedule = model.NewMedicalSchedule(doctorIdDB, queryValueDB, dayOfServiceDB, specificDateDB, startTimeDB, finalTimeDB, yearDB)
	} else {
		medicalSchedule = model.NewMedicalSchedule(doctorIdDB, queryValueDB, dayOfServiceDB, time.Time{}, startTimeDB, finalTimeDB, yearDB)
	}
	if medicalSchedule != (model.MedicalSchedule{}) {
		medicalSchedule.SetId(medicalScheduleId)
	}

	return medicalSchedule, err
}

func MedicalScheduleEdit(medicalSchedule model.MedicalSchedule) (bool, error) {
	var success bool
	db, err := connection.NewConnection()
	if err != nil {
		return success, err
	}
	defer db.Close()

	sql := "update medical_schedule set doctor_id=$1, day_of_service=$2, start_time=$3, final_time=$4, specific_date=$5, " +
		"year=$6, active=true, last_modified_date=current_timestamp, query_value=$7 where id=$8"

	_, err = db.Prepare(sql)
	if err != nil {
		return success, err
	}

	_, err = db.Exec(sql,
		medicalSchedule.GetDoctorId(),
		medicalSchedule.GetDayOfService(),
		medicalSchedule.GetStartTime(),
		medicalSchedule.GetFinalTime(),
		medicalSchedule.GetSpecificDate(),
		medicalSchedule.GetYear(),
		medicalSchedule.GetQueryValue(),
		medicalSchedule.GetId())
	if err != nil {
		return success, err
	}

	success = true
	return success, err
}

func MedicalScheduleDelete(medicalScheduleId int) (bool, error) {
	var success bool
	db, err := connection.NewConnection()
	if err != nil {
		return success, err
	}
	defer db.Close()

	sql := "update medical_schedule set active = false, last_modified_date = current_timestamp where id = $1"

	_, err = db.Prepare(sql)
	if err != nil {
		return success, err
	}

	_, err = db.Exec(sql, medicalScheduleId)
	if err != nil {
		return success, err
	}

	success = true
	return success, err
}
