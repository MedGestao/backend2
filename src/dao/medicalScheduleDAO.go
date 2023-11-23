package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
)

func MedicalScheduleInsert(medicalSchedule model.MedicalSchedule) (bool, error) {
	db, err := connection.NewConnection()
	var success bool
	if err != nil {
		return success, err
	}
	defer db.Close()

	sql := "insert into medical_schedule(doctor_id, day_of_service, start_time, final_time, specific_time, year, active, " +
		"registration_date) values($1, $2, $3, $4, $5, $6, true, current_timestamp)"
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
		medicalSchedule.GetSpecificTime(),
		medicalSchedule.GetYear())
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
	sql := "select id, doctor_id, day_of_service, start_time, final_time, specific_time, year from medical_schedule where active is true"

	_, err = db.Prepare(sql)
	if err != nil {
		return medicalScheduleList, err
	}

	rows, err := db.Query(sql)
	if err != nil {
		return medicalScheduleList, err
	}

	var idDB, doctorIdDB int
	var dayOfServiceDB, startTimeDB, finalTimeDB, specificTimeDB, yearDB string
	for rows.Next() {
		err = rows.Scan(&idDB, &doctorIdDB, &dayOfServiceDB, &startTimeDB, &finalTimeDB, &specificTimeDB, &yearDB)
		if err != nil {
			return medicalScheduleList, err
		}
		medicalSchedule = model.NewMedicalSchedule(doctorIdDB, dayOfServiceDB, specificTimeDB, startTimeDB, finalTimeDB, yearDB)
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

	sql := "select doctor_id, day_of_service, start_time, final_time, specific_time, year from medical_schedule " +
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
	var dayOfServiceDB, startTimeDB, finalTimeDB, specificTimeDB, yearDB string
	for rows.Next() {
		err = rows.Scan(&doctorIdDB, &dayOfServiceDB, &startTimeDB, &finalTimeDB, &specificTimeDB, &yearDB)
		if err != nil {
			return medicalSchedule, err
		}
	}

	medicalSchedule = model.NewMedicalSchedule(doctorIdDB, dayOfServiceDB, specificTimeDB, startTimeDB, finalTimeDB, yearDB)
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

	sql := "update medical_schedule set doctor_id=$1, day_of_service=$2, start_time=$3, final_time=$4, specific_time=$5, " +
		"year=$6, active=true, last_modified_date=current_timestamp where id=$7"

	_, err = db.Prepare(sql)
	if err != nil {
		return success, err
	}

	_, err = db.Exec(sql,
		medicalSchedule.GetDoctorId(),
		medicalSchedule.GetDayOfService(),
		medicalSchedule.GetStartTime(),
		medicalSchedule.GetFinalTime(),
		medicalSchedule.GetSpecificTime(),
		medicalSchedule.GetYear(),
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
