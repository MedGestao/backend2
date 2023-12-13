package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/response"
	sql2 "database/sql"
	"strconv"
	"strings"
	"time"
)

func MedicalScheduleInsert(medicalSchedule model.MedicalSchedule) (bool, error, response.ErrorResponse) {
	db, err := connection.NewConnection()
	var success bool
	var errorMessage response.ErrorResponse
	if err != nil {
		return success, err, errorMessage
	}
	defer db.Close()

	sql := "select exists(select id from medical_schedule where doctor_id=$1) as exist "
	if err != err {
		return success, err, errorMessage
	}
	_, err = db.Prepare(sql)
	if err != nil {
		return success, err, errorMessage
	}

	rows, err := db.Query(sql,
		medicalSchedule.GetDoctorId())
	if err != nil {
		return success, err, errorMessage
	}

	var existDB bool
	for rows.Next() {
		err = rows.Scan(&existDB)
		if err != nil {
			return success, err, errorMessage
		}
	}

	if existDB == true {
		errorMessage = response.NewErrorResponse("Você já possui uma agenda cadastrada. Caso necessário edite a sua agenda!")
		return success, err, errorMessage
	}

	if medicalSchedule.GetSpecificDate().IsZero() {
		sql = "insert into medical_schedule(doctor_id, day_of_service, period_1, period_2, year, active, " +
			"registration_date, query_value, schedule_limit) values($1, $2, $3, $4, $5, true, current_timestamp, $6, $7)"
		if err != err {
			return success, err, errorMessage
		}
		_, err = db.Prepare(sql)
		if err != nil {
			return success, err, errorMessage
		}

		_, err = db.Exec(sql,
			medicalSchedule.GetDoctorId(),
			medicalSchedule.GetDayOfService(),
			medicalSchedule.GetPeriod1(),
			medicalSchedule.GetPeriod2(),
			medicalSchedule.GetYear(),
			medicalSchedule.GetQueryValue(),
			medicalSchedule.GetScheduleLimit())
		if err != nil {
			return success, err, errorMessage
		}
	} else {
		sql := "insert into medical_schedule(doctor_id, day_of_service, period_1, period_2, specific_date, year, active, " +
			"registration_date, query_value, schedule_limit) values($1, $2, $3, $4, $5, $6, true, current_timestamp, $7, $8)"
		if err != err {
			return success, err, errorMessage
		}
		_, err = db.Prepare(sql)
		if err != nil {
			return success, err, errorMessage
		}

		_, err = db.Exec(sql,
			medicalSchedule.GetDoctorId(),
			medicalSchedule.GetDayOfService(),
			medicalSchedule.GetPeriod1(),
			medicalSchedule.GetPeriod2(),
			medicalSchedule.GetSpecificDate(),
			medicalSchedule.GetYear(),
			medicalSchedule.GetQueryValue(),
			medicalSchedule.GetScheduleLimit())
		if err != nil {
			return success, err, errorMessage
		}
	}

	success = true
	return success, err, errorMessage
}

func MedicalScheduleSelectAllByIdDoctor(doctorId int) ([]model.MedicalSchedule, error) {
	var medicalSchedule model.MedicalSchedule
	var medicalScheduleList []model.MedicalSchedule
	db, err := connection.NewConnection()
	if err != nil {
		return medicalScheduleList, err
	}
	defer db.Close()
	sql := "select id, doctor_id, day_of_service, period_1, period_2, specific_date, year, query_value, schedule_limit from medical_schedule where doctor_id = $1 and active is true"

	_, err = db.Prepare(sql)
	if err != nil {
		return medicalScheduleList, err
	}

	rows, err := db.Query(sql,
		doctorId)
	if err != nil {
		return medicalScheduleList, err
	}

	var idDB, doctorIdDB, scheduleLimitDB int
	var dayOfServiceDB, period1DB, period2DB, yearDB string
	var specificDateDB time.Time
	var queryValueDB float64
	var strValue string
	var specificDateNull sql2.NullTime
	for rows.Next() {
		err = rows.Scan(&idDB, &doctorIdDB, &dayOfServiceDB, &period1DB, &period2DB, &specificDateNull, &yearDB, &strValue, &scheduleLimitDB)
		if err != nil {
			return medicalScheduleList, err
		}

		strValue = strings.ReplaceAll(strings.ReplaceAll(strValue, ",", ""), "$", "")
		queryValueDB, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return medicalScheduleList, err
		}

		if specificDateNull.Valid {
			specificDateDB = specificDateNull.Time
			medicalSchedule = model.NewMedicalSchedule(doctorIdDB, queryValueDB, dayOfServiceDB, specificDateDB, period1DB, period2DB, yearDB, scheduleLimitDB)
			if medicalSchedule != (model.MedicalSchedule{}) {
				medicalSchedule.SetId(idDB)
				medicalScheduleList = append(medicalScheduleList, medicalSchedule)
			}
		} else {
			medicalSchedule = model.NewMedicalSchedule(doctorIdDB, queryValueDB, dayOfServiceDB, time.Time{}, period1DB, period2DB, yearDB, scheduleLimitDB)
			if medicalSchedule != (model.MedicalSchedule{}) {
				medicalSchedule.SetId(idDB)
				medicalScheduleList = append(medicalScheduleList, medicalSchedule)
			}
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

	sql := "select doctor_id, day_of_service, period_1, period_2, specific_date, year, query_value, schedule_limit from medical_schedule " +
		"where id = $1 and active is true"
	_, err = db.Prepare(sql)
	if err != nil {
		return medicalSchedule, err
	}

	rows, err := db.Query(sql, medicalScheduleId)
	if err != nil {
		return medicalSchedule, err
	}

	var doctorIdDB, scheduleLimitDB int
	var dayOfServiceDB, period1DB, period2DB, yearDB string
	var specificDateDB time.Time
	var queryValueDB float64
	var strValue string
	var specificDateNull sql2.NullTime
	for rows.Next() {
		err = rows.Scan(&doctorIdDB, &dayOfServiceDB, &period1DB, &period2DB, &specificDateNull, &yearDB, &strValue, &scheduleLimitDB)
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
		medicalSchedule = model.NewMedicalSchedule(doctorIdDB, queryValueDB, dayOfServiceDB, specificDateDB, period1DB, period2DB, yearDB, scheduleLimitDB)
	} else {
		medicalSchedule = model.NewMedicalSchedule(doctorIdDB, queryValueDB, dayOfServiceDB, time.Time{}, period1DB, period2DB, yearDB, scheduleLimitDB)
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

	if medicalSchedule.GetSpecificDate().IsZero() {
		sql := "update medical_schedule set doctor_id=$1, day_of_service=$2, period_1=$3, period_2=$4, " +
			"year=$5, active=true, last_modified_date=current_timestamp, query_value=$6, schedule_limit=$7 where id=$8"

		_, err = db.Prepare(sql)
		if err != nil {
			return success, err
		}

		_, err = db.Exec(sql,
			medicalSchedule.GetDoctorId(),
			medicalSchedule.GetDayOfService(),
			medicalSchedule.GetPeriod1(),
			medicalSchedule.GetPeriod2(),
			medicalSchedule.GetYear(),
			medicalSchedule.GetQueryValue(),
			medicalSchedule.GetScheduleLimit(),
			medicalSchedule.GetId())
		if err != nil {
			return success, err
		}
	} else {
		sql := "update medical_schedule set doctor_id=$1, day_of_service=$2, period_1=$3, period_2=$4, specific_date=$5, " +
			"year=$6, active=true, last_modified_date=current_timestamp, query_value=$7, schedule_limit=$8 where id=$9"

		_, err = db.Prepare(sql)
		if err != nil {
			return success, err
		}

		_, err = db.Exec(sql,
			medicalSchedule.GetDoctorId(),
			medicalSchedule.GetDayOfService(),
			medicalSchedule.GetPeriod1(),
			medicalSchedule.GetPeriod2(),
			medicalSchedule.GetSpecificDate(),
			medicalSchedule.GetYear(),
			medicalSchedule.GetQueryValue(),
			medicalSchedule.GetScheduleLimit(),
			medicalSchedule.GetId())
		if err != nil {
			return success, err
		}
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
