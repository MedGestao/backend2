package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/util"
	"database/sql"
	"time"
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

func DoctorSelectById(doctorId int) (model.Doctor, error) {
	db, err := connection.NewConnection()
	if err != nil {
		println("Error1: ", err.Error())
	}
	defer db.Close()

	sql := "select distinct on (d.cpf) d.name, d.birthdate, d.sex, d.cpf, d.address, cd.number, d.cns, d.crm, dai.doctor_email, s.description as specialty, d.active from doctor d " +
		"left join cellphone_doctor cd on d.id = cd.doctor_id " +
		"left join doctor_authentication_information dai on d.id = dai.doctor_id " +
		"left join medical_specialty ms on d.id = ms.doctor_id " +
		"left join specialty s on ms.specialty_id = s.id where d.id = $1 and d.active is true "
	_, err = db.Prepare(sql)
	if err != nil {
		println("Error3: ", err.Error())
	}

	var doctorNameDB,
		doctorSexDB,
		doctorCpfDB,
		doctorAddressDB,
		doctorNumberDB,
		doctorEmailDB,
		doctorCns,
		doctorCrm,
		doctorSpecialty string
	var doctorBirthdateDB time.Time
	var doctorActiveDB bool
	var doctor model.Doctor
	rows, err := db.Query(sql, doctorId)

	for rows.Next() {
		err = rows.Scan(
			&doctorNameDB,
			&doctorBirthdateDB,
			&doctorSexDB,
			&doctorCpfDB,
			&doctorAddressDB,
			&doctorNumberDB,
			&doctorCns,
			&doctorCrm,
			&doctorSpecialty,
			&doctorEmailDB,
			&doctorActiveDB)
		if err != nil {
			println("Error nos dados retornados: ", err.Error())
			return doctor, err
		}
	}
	doctor = model.NewDoctor(doctorNameDB, doctorBirthdateDB, doctorCpfDB, doctorSexDB, doctorAddressDB, doctorEmailDB, model.NewCellphoneUser(doctorNumberDB), "", doctorActiveDB, doctorCns, doctorCrm, model.NewSpecialty(doctorSpecialty))

	return doctor, nil

}

func DoctorEdit(doctor model.Doctor) (bool, error) {
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

	sql := "update doctor set name = $1, birthdate = $2, cpf = $3, sex = $4, address = $5, cns = $6, crm = $7, " +
		"active = $8, last_modified_date = current_timestamp where id = $9 returning id"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return success, err
	}
	println("Id do Médico: ", doctor.GetUser().GetId())
	var tempDoctorId int
	err = tx.QueryRow(sql,
		doctor.GetUser().GetName(),
		doctor.GetUser().GetBirthDate(),
		doctor.GetUser().GetCpf(),
		doctor.GetUser().GetSex(),
		doctor.GetUser().GetAddress(),
		doctor.GetCns(),
		doctor.GetCrm(),
		doctor.GetUser().IsActive(),
		doctor.GetUser().GetId()).Scan(&tempDoctorId)
	if err != nil {
		tx.Rollback()
		println("Error3: ", err.Error())
		return success, err
	}

	sql = "update cellphone_doctor set number = $1 where doctor_id = $2"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error4: ", err.Error())
		return success, err
	}

	println("id temporário: ", tempDoctorId)
	_, err = tx.Exec(sql,
		doctor.GetUser().GetCellphoneUser().GetNumber(),
		tempDoctorId)
	if err != nil {
		tx.Rollback()
		println("Error5: ", err.Error())
		return success, err
	}

	if (doctor.GetUser().GetEmail() == "") || (doctor.GetUser().GetPassword() == "") {
		tx.Commit()
	} else {
		err = DoctorEditLogin(doctor.GetUser().GetEmail(), doctor.GetUser().GetPassword(), tempDoctorId, tx)
		if err != nil {
			tx.Rollback()
			println("Error 6: ", err.Error())
			return success, err
		}
		tx.Commit()
	}

	success = true
	return success, nil
}

func DoctorEditLogin(email string, password string, doctortId int, tx *sql.Tx) error {

	passwordHashDB, saltDB, err := util.PasswordHash(password)
	sql := "update doctor_authentication_information set doctor_email = $1, " +
		"doctor_password = $2, doctor_salt = $3 where doctor_id = $4"

	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error de preparação de sql de update de login: ", err.Error())
		return err
	}

	_, err = tx.Exec(sql,
		email,
		passwordHashDB,
		saltDB,
		doctortId)
	if err != nil {
		tx.Rollback()
		println("Error de execução do sql de update de login: ", err.Error())
		return err
	}

	return nil
}

func DoctorOff(doctorId int) (bool, error) {
	success := false
	db, err := connection.NewConnection()
	if err != nil {
		println("ErrorDesligamento1: ", err.Error())
		return false, err
	}

	sql := "update doctor set active = false where id = $1"
	_, err = db.Prepare(sql)
	if err != nil {
		println("ErrorDesligamento2: ", err.Error())
		return success, err
	}

	_, err = db.Exec(sql, doctorId)
	if err != nil {
		println("ErrorDesligamento3: ", err.Error())
		return success, err
	}

	success = true
	return success, nil
}
