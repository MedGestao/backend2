package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/response"
	"MedGestao/src/util"
	"database/sql"
	"github.com/paemuri/brdoc"
	"time"
)

func PatientInsert(patient model.Patient) (int, error, response.ErrorResponse) {
	db, err := connection.NewConnection()
	var patientId int
	var errorMessage response.ErrorResponse
	if err != nil {
		return patientId, err, errorMessage
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		println(err)
		return patientId, err, errorMessage
	}

	if brdoc.IsCPF(patient.GetUser().GetCpf()) == false {
		tx.Rollback()
		errorMessage = response.NewErrorResponse("O CPF informado é inválido!")
		return patientId, err, errorMessage
	}

	sql := "select exists(select id from patient where cpf=$1) as exist"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return patientId, err, errorMessage
	}

	rows, err := tx.Query(sql,
		patient.GetUser().GetCpf())
	if err != nil {
		tx.Rollback()
		return patientId, err, errorMessage
	}

	var exist bool
	for rows.Next() {
		err = rows.Scan(&exist)
		if err != nil {
			tx.Rollback()
			return patientId, err, errorMessage
		}
	}

	if exist == true {
		tx.Rollback()
		errorMessage = response.NewErrorResponse("Já existe um cadastro com esse cpf!")
		return patientId, err, errorMessage
	}

	sql = "select exists(select id from patient_authentication_information where patient_email=$1) as exist"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return patientId, err, errorMessage
	}

	rows, err = tx.Query(sql,
		patient.GetUser().GetEmail())
	if err != nil {
		tx.Rollback()
		return patientId, err, errorMessage
	}

	for rows.Next() {
		err = rows.Scan(&exist)
		if err != nil {
			tx.Rollback()
			return patientId, err, errorMessage
		}
	}

	if exist == true {
		tx.Rollback()
		errorMessage = response.NewErrorResponse("Esse e-mail já está em uso!")
		return patientId, err, errorMessage
	}

	sql = "insert into patient(name, birthdate, cpf, sex, address, image_url, active, registration_date)" +
		" values ($1, $2, $3, $4, $5, $6, true, current_timestamp) returning id"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error1: ", err.Error())
		return patientId, err, errorMessage
	}

	if err != nil {
		println("Error geração do hash do password: ", err.Error())
		panic(err)
	}

	err = tx.QueryRow(sql,
		patient.GetUser().GetName(),
		string(patient.GetUser().GetBirthDate().Format("2006-01-02")),
		patient.GetUser().GetCpf(),
		patient.GetUser().GetSex(),
		patient.GetUser().GetAddress(),
		patient.GetUser().GetImageUrl(),
	).Scan(&patientId)

	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return patientId, err, errorMessage
	}

	sql = "insert into cellphone_patient(patient_id, number) values ($1, $2)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do paciente: ", patientId)
		println("Error3: ", err.Error())
		return patientId, err, errorMessage
	}
	_, err = tx.Exec(sql,
		patientId,
		patient.GetUser().GetCellphoneUser().GetNumber(),
	)
	if err != nil {
		tx.Rollback()
		println("Id do paciente: ", patientId)
		println("Error4: ", err.Error())
		return patientId, err, errorMessage
	}

	passwordHashDB, saltDB, err := util.PasswordHash(patient.GetUser().GetPassword())
	sql = "insert into patient_authentication_information(patient_id, patient_email, patient_password, patient_salt)" +
		"values($1, $2, $3, $4)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do paciente: ", patientId)
		println("Error5: ", err.Error())
		return patientId, err, errorMessage
	}

	_, err = tx.Exec(sql,
		patientId,
		patient.GetUser().GetEmail(),
		passwordHashDB,
		saltDB,
	)
	if err != nil {
		println("Id do paciente: ", patientId)
		println("Error6: ", err.Error())
		return patientId, err, errorMessage
	}

	tx.Commit()
	return patientId, err, errorMessage
}

func PatientValidateLogin(emailLogin string, passwordLogin string) (bool, int, error) {
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

func PatientSelectById(patientId int) (model.Patient, error) {
	db, err := connection.NewConnection()
	if err != nil {
		println("Error1: ", err.Error())
	}
	defer db.Close()

	sql := "select distinct on (p.cpf) p.name, p.birthdate, p.sex, p.cpf, p.address, cp.number, pai.patient_email, p.image_url, p.active from patient p " +
		"left join cellphone_patient cp on p.id = cp.patient_id " +
		"left join patient_authentication_information pai on p.id = pai.patient_id " +
		"where p.id = $1 and p.active is true"
	_, err = db.Prepare(sql)
	if err != nil {
		println("Error3: ", err.Error())
	}

	var patientNameDB, patientSexDB, patientCpfDB, patientAddressDB, patientNumberDB, patientEmailDB, patientImageUrlDB string
	var patientBirthdateDB time.Time
	var patientActiveDB bool
	var patient model.Patient
	rows, err := db.Query(sql, patientId)
	for rows.Next() {
		err = rows.Scan(&patientNameDB, &patientBirthdateDB, &patientSexDB, &patientCpfDB, &patientAddressDB, &patientNumberDB, &patientEmailDB, &patientImageUrlDB, &patientActiveDB)
		if err != nil {
			println("Error nos dados retornados: ", err.Error())
			return patient, err
		}
	}
	patient = model.NewPatient(patientNameDB, patientBirthdateDB, patientCpfDB, patientSexDB, patientAddressDB, patientEmailDB, "", patientImageUrlDB, model.NewCellphoneUser(patientNumberDB))
	patient.SetUserActive(patientActiveDB)
	patient.SetUserId(patientId)

	return patient, nil

}

func PatientEdit(idPatient int, patient model.Patient) (bool, error) {
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

	sql := "update patient set name = $1, birthdate = $2, cpf = $3, sex = $4, address = $5, image_url = $6, " +
		"active = true, last_modified_date = current_timestamp where id = $7"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return success, err
	}
	println("Id do paciente: ", idPatient)
	_, err = tx.Exec(sql,
		patient.GetUser().GetName(),
		patient.GetUser().GetBirthDate(),
		patient.GetUser().GetCpf(),
		patient.GetUser().GetSex(),
		patient.GetUser().GetAddress(),
		patient.GetUser().GetImageUrl(),
		idPatient)
	if err != nil {
		tx.Rollback()
		println("Error3: ", err.Error())
		return success, err
	}

	sql = "update cellphone_patient set number = $1 where patient_id = $2"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error4: ", err.Error())
		return success, err
	}

	_, err = tx.Exec(sql,
		patient.GetUser().GetCellphoneUser().GetNumber(),
		idPatient)
	if err != nil {
		tx.Rollback()
		println("Error5: ", err.Error())
		return success, err
	}

	if (patient.GetUser().GetEmail() == "") || (patient.GetUser().GetPassword() == "") {
		tx.Commit()
	} else {
		err = PatientEditLogin(patient.GetUser().GetEmail(), patient.GetUser().GetPassword(), idPatient, tx)
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

func PatientEditLogin(email string, password string, patientId int, tx *sql.Tx) error {

	passwordHashDB, saltDB, err := util.PasswordHash(password)
	sql := "update patient_authentication_information set patient_email = $1, " +
		"patient_password = $2, patient_salt = $3 where patient_id = $4"

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
		patientId)
	if err != nil {
		tx.Rollback()
		println("Error de execução do sql de update de login: ", err.Error())
		return err
	}

	return nil
}

func PatientOff(patientId int) (bool, error) {
	var success bool
	db, err := connection.NewConnection()
	if err != nil {
		println("ErrorDesligamento1: ", err.Error())
		return false, err
	}

	sql := "update patient set active = false where id = $1"
	_, err = db.Prepare(sql)
	if err != nil {
		println("ErrorDesligamento2: ", err.Error())
		return success, err
	}

	_, err = db.Exec(sql, patientId)
	if err != nil {
		println("ErrorDesligamento3: ", err.Error())
		return success, err
	}

	success = true
	return success, err
}
