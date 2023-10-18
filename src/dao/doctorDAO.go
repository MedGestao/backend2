package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/util"
)

func InsertDoctor(doctor model.Doctor) (bool, error) {
	db, err := connection.NewConnection()
	defer db.Close()
	if err != nil {
		return false, err
	}

	tx, err := db.Begin()
	if err != nil {
		println(err)
		return false, err
	}

	sql := "insert into doctor(name, birthdate, cpf, sex, address, cns, crm, active)" +
		" values ($1, $2, $3, $4, $5, $6, $7, $8) returning id"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error1: ", err.Error())
		return false, err
	}

	if err != nil {
		println("Error geração do hash do password: ", err.Error())
		panic(err)
	}

	println("Ativo: ", doctor.GetUser().IsActive())
	var tempDoctorId int
	err = tx.QueryRow(sql,
		doctor.GetUser().GetName(),
		doctor.GetUser().GetBirthDate().Format("2006-01-02"),
		doctor.GetUser().GetCpf(),
		doctor.GetUser().GetSex(),
		doctor.GetUser().GetAddress(),
		doctor.GetCns(),
		doctor.GetCrm(),
		doctor.GetUser().IsActive(),
	).Scan(&tempDoctorId)

	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return false, err
	}

	sql = "insert into cellphone_doctor(doctor_id, number) values ($1, $2)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do médico: ", tempDoctorId)
		println("Error3: ", err.Error())
		return false, err
	}
	_, err = tx.Exec(sql,
		tempDoctorId,
		doctor.GetUser().GetCellphoneUser().GetNumber(),
	)
	if err != nil {
		tx.Rollback()
		println("Id do médico: ", tempDoctorId)
		println("Error4: ", err.Error())
		return false, err
	}

	sql = "insert into medical_specialty(doctor_id, specialty_id) values($1, $2)"
	_, err = tx.Prepare(sql)
	if err != nil {
		println("Error5: ", err.Error())
		return false, err
	}

	_, err = tx.Exec(sql,
		tempDoctorId,
		//doctor.GetSpecialty().GetId()
		1,
	)
	if err != nil {
		println("Error6: ", err.Error())
		return false, err
	}

	passwordHashDB, saltDB, err := util.PasswordHash(doctor.GetUser().GetPassword())
	sql = "insert into doctor_authentication_information(doctor_id, doctor_email, doctor_password, doctor_salt)" +
		"values($1, $2, $3, $4)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do médico: ", tempDoctorId)
		println("Error7: ", err.Error())
		return false, err
	}

	_, err = tx.Exec(sql,
		tempDoctorId,
		doctor.GetUser().GetEmail(),
		passwordHashDB,
		saltDB,
	)
	if err != nil {
		println("Id do médico: ", tempDoctorId)
		println("Error8: ", err.Error())
		return false, err
	}

	tx.Commit()
	return true, nil
}

func ValidateLoginDoctor(emailLogin string, passwordLogin string) (bool, int, error) {
	db, err := connection.NewConnection()
	defer db.Close()
	if err != nil {
		return false, 0, err
	}

	authorized := false

	sql := "select doctor_password, doctor_salt, doctor_id from doctor_authentication_information where doctor_email = $1"

	db.Prepare(sql)
	rows, err := db.Query(sql, emailLogin)
	if err != nil {
		println("Error consulta sql: ", err.Error())
		return authorized, 0, err
	}

	var passwordDB []byte
	var saltDB []byte
	var doctorIdDB int
	for rows.Next() {
		err = rows.Scan(&passwordDB, &saltDB, &doctorIdDB)
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

	return authorized, doctorIdDB, err
}
