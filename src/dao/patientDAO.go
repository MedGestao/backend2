package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/util"
)

func PatientInsert(patient model.Patient) (bool, error) {
	db, err := connection.NewConnection()
	var success bool
	if err != nil {
		return success, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		println(err)
		return success, err
	}

	sql := "insert into patient(name, birthdate, cpf, sex, address, active)" +
		" values ($1, $2, $3, $4, $5, $6) returning id"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error1: ", err.Error())
		return success, err
	}

	println(sql)
	println(patient.GetUser().GetName())
	println(patient.GetUser().GetBirthDate().Date())
	println(patient.GetUser().GetCpf())
	println(patient.GetUser().GetSex())
	println(patient.GetUser().GetAddress())
	println(patient.GetUser().GetEmail())
	println("Telefone: ", patient.GetUser().GetCellphoneUser().GetNumber())

	if err != nil {
		println("Error geração do hash do password: ", err.Error())
		panic(err)
	}

	var tempPatientId int
	err = tx.QueryRow(sql,
		patient.GetUser().GetName(),
		string(patient.GetUser().GetBirthDate().Format("2006-01-02")),
		patient.GetUser().GetCpf(),
		patient.GetUser().GetSex(),
		patient.GetUser().GetAddress(),
		patient.GetUser().IsActive(),
	).Scan(&tempPatientId)

	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return success, err
	}

	sql = "insert into cellphone_patient(patient_id, number) values ($1, $2)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do paciente: ", tempPatientId)
		println("Error3: ", err.Error())
		return success, err
	}
	_, err = tx.Exec(sql,
		tempPatientId,
		patient.GetUser().GetCellphoneUser().GetNumber(),
	)
	if err != nil {
		tx.Rollback()
		println("Id do paciente: ", tempPatientId)
		println("Error4: ", err.Error())
		return success, err
	}

	passwordHashDB, saltDB, err := util.PasswordHash(patient.GetUser().GetPassword())
	sql = "insert into patient_authentication_information(patient_id, patient_email, patient_password, patient_salt)" +
		"values($1, $2, $3, $4)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do paciente: ", tempPatientId)
		println("Error5: ", err.Error())
		return success, err
	}

	_, err = tx.Exec(sql,
		tempPatientId,
		patient.GetUser().GetEmail(),
		passwordHashDB,
		saltDB,
	)
	if err != nil {
		println("Id do paciente: ", tempPatientId)
		println("Error6: ", err.Error())
		return success, err
	}

	tx.Commit()
	success = true
	return success, nil
}

func PatientEdit(patient model.Patient) (bool, error) {
	db, err := connection.NewConnection()
	success := false
	if err != nil {
		return success, err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		println("Error1: ", err.Error())
		return success, err
	}

	sql := "update patient set name = $1, birthdate = $2, cpf = $3, sex = $4, address = $5, active = $6 where id = $7"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return success, err
	}
	println("Id do paciente: ", patient.GetUser().GetId())
	_, err = tx.Exec(sql,
		patient.GetUser().GetName(),
		patient.GetUser().GetBirthDate(),
		patient.GetUser().GetCpf(),
		patient.GetUser().GetSex(),
		patient.GetUser().GetAddress(),
		patient.GetUser().IsActive(),
		patient.GetUser().GetId())
	if err != nil {
		tx.Rollback()
		println("Error3: ", err.Error())
		return success, err
	}

	sql = "update cellphone_patient set number = $1 where id = $2 and patient_id = $3"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error4: ", err.Error())
		return success, err
	}

	_, err = tx.Exec(sql,
		patient.GetUser().GetCellphoneUser().GetNumber(),
		patient.GetUser().GetCellphoneUser().GetId(),
		patient.GetUser().GetCellphoneUser().GetUserId())
	if err != nil {
		tx.Rollback()
		println("Error5: ", err.Error())
		return success, err
	}

	passwordHashDB, saltDB, err := util.PasswordHash(patient.GetUser().GetPassword())
	sql = "update patient_authentication_information set patient_email = $1, " +
		"patient_password = $2, patient_salt = $3 where patient_id = $4"

	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error6: ", err.Error())
		return success, err
	}

	_, err = tx.Exec(sql,
		patient.GetUser().GetEmail(),
		passwordHashDB,
		saltDB,
		patient.GetUser().GetId())
	if err != nil {
		tx.Rollback()
		println("Error7: ", err.Error())
		return success, err
	}

	tx.Commit()
	success = true
	return success, nil
}

func ValidateLoginPatient(emailLogin string, passwordLogin string) (bool, int, error) {
	db, err := connection.NewConnection()
	if err != nil {
		return false, 0, err
	}
	defer db.Close()

	authorized := false

	sql := "select patient_password, patient_salt, patient_id from patient_authentication_information where patient_email = $1"

	db.Prepare(sql)
	rows, err := db.Query(sql, emailLogin)
	if err != nil {
		println("Error consulta sql: ", err.Error())
		return authorized, 0, err
	}

	var passwordDB []byte
	var saltDB []byte
	var patientIdDB int
	for rows.Next() {
		err = rows.Scan(&passwordDB, &saltDB, &patientIdDB)
		if err != nil {
			println("Error nos dados retornados: ", err.Error())
			return false, 0, err
		}
	}

	authorized, err = util.ValidatePasswordHash(passwordLogin, passwordDB, saltDB)
	if err != nil {
		println("Error na validação: ", err.Error())
		return authorized, 0, err
	}

	return authorized, patientIdDB, err
}
