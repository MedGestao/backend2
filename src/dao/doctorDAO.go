package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"MedGestao/src/response"
	"MedGestao/src/util"
	"database/sql"
	"strings"
	"time"

	"github.com/paemuri/brdoc"
)

var globalRows *sql.Rows

func InsertDoctor(doctor model.Doctor) (int, error, response.ErrorResponse) {
	var doctorId int
	var err error
	var errorMessage response.ErrorResponse
	db, err := connection.NewConnection()
	defer db.Close()
	if err != nil {
		return doctorId, err, errorMessage
	}

	tx, err := db.Begin()
	if err != nil {
		println(err)
		return doctorId, err, errorMessage
	}

	if brdoc.IsCPF(doctor.GetUser().GetCpf()) == false {
		tx.Rollback()
		errorMessage = response.NewErrorResponse("O CPF informado é inválido!")
		return doctorId, err, errorMessage
	}

	sql := "select exists(select id from doctor where cpf=$1) as exist"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return doctorId, err, errorMessage
	}

	rows, err := tx.Query(sql,
		doctor.GetUser().GetCpf())
	if err != nil {
		tx.Rollback()
		return doctorId, err, errorMessage
	}

	var exist bool
	for rows.Next() {
		err = rows.Scan(&exist)
		if err != nil {
			tx.Rollback()
			return doctorId, err, errorMessage
		}
	}

	if exist == true {
		tx.Rollback()
		errorMessage = response.NewErrorResponse("Já existe um cadastro com esse cpf!")
		return doctorId, err, errorMessage
	}

	sql = "insert into doctor(name, birthdate, cpf, sex, address, crm, image_url, active, registration_date)" +
		" values ($1, $2, $3, $4, $5, $6, $7, true, current_timestamp) returning id"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error1: ", err.Error())
		return doctorId, err, errorMessage
	}

	if err != nil {
		println("Error geração do hash do password: ", err.Error())
		return doctorId, err, errorMessage
	}

	var tempDoctorId int
	err = tx.QueryRow(sql,
		doctor.GetUser().GetName(),
		doctor.GetUser().GetBirthDate().Format("2006-01-02"),
		doctor.GetUser().GetCpf(),
		doctor.GetUser().GetSex(),
		doctor.GetUser().GetAddress(),
		doctor.GetCrm(),
		doctor.GetUser().GetImageUrl(),
	).Scan(&tempDoctorId)

	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return doctorId, err, errorMessage
	}

	sql = "insert into cellphone_doctor(doctor_id, number) values ($1, $2)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do médico: ", tempDoctorId)
		println("Error3: ", err.Error())
		return doctorId, err, errorMessage
	}
	_, err = tx.Exec(sql,
		tempDoctorId,
		doctor.GetUser().GetCellphoneUser().GetNumber(),
	)
	if err != nil {
		tx.Rollback()
		println("Id do médico: ", tempDoctorId)
		println("Error4: ", err.Error())
		return doctorId, err, errorMessage
	}

	sql = "insert into medical_specialty(doctor_id, specialty_id) values($1, $2)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error5: ", err.Error())
		return doctorId, err, errorMessage
	}

	_, err = tx.Exec(sql,
		tempDoctorId,
		doctor.GetSpecialty().GetId(),
	)
	if err != nil {
		tx.Rollback()
		println("Error6: ", err.Error())
		return doctorId, err, errorMessage
	}

	sql = "select exists (select id from doctor_authentication_information where doctor_email=$1) as ex"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return doctorId, err, errorMessage
	}

	err = tx.QueryRow(sql,
		doctor.GetUser().GetEmail()).Scan(&exist)
	if err != nil {
		tx.Rollback()
		return doctorId, err, errorMessage
	} else if exist == true {
		tx.Rollback()
		errorMessage = response.NewErrorResponse("Esse e-mail já está em uso!")
		return doctorId, err, errorMessage
	}

	passwordHashDB, saltDB, err := util.PasswordHash(doctor.GetUser().GetPassword())
	sql = "insert into doctor_authentication_information(doctor_id, doctor_email, doctor_password, doctor_salt)" +
		"values($1, $2, $3, $4)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Id do médico: ", tempDoctorId)
		println("Error7: ", err.Error())
		return doctorId, err, errorMessage
	}

	_, err = tx.Exec(sql,
		tempDoctorId,
		doctor.GetUser().GetEmail(),
		passwordHashDB,
		saltDB,
	)
	if err != nil {
		tx.Rollback()
		println("Id do médico: ", tempDoctorId)
		println("Error8: ", err.Error())
		return doctorId, err, errorMessage
	}

	tx.Commit()

	doctorId = tempDoctorId
	return doctorId, err, errorMessage
}

func DoctorValidateLogin(emailLogin string, passwordLogin string) (bool, int, error) {
	db, err := connection.NewConnection()
	authorized := false
	if err != nil {
		return false, 0, err
	}
	defer db.Close()

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

func DoctorSelectAll(doctorName string, specialtyDescription string) ([]response.DoctorResponse, error) {

	db, err := connection.NewConnection()
	var doctors []response.DoctorResponse

	if err != nil {
		return doctors, err
	}
	defer db.Close()

	sql := "select distinct on (d.cpf) d.id, d.name, d.image_url, s.id, s.description as specialty from doctor d " +
		"left join medical_specialty ms on d.id = ms.doctor_id " +
		"left join specialty s on ms.specialty_id = s.id where d.active is true and d.image_url is not null"

	if doctorName != "" && specialtyDescription != "" {
		sql = "select distinct on (d.cpf) d.id, d.name, d.image_url, s.id, s.description as specialty from doctor d " +
			"left join medical_specialty ms on d.id = ms.doctor_id " +
			"left join specialty s on ms.specialty_id = s.id where d.active is true and d.image_url is not null" +
			" and d.name like '%' || $1 || '%' and s.description like '%' || $2 || '%''"
	}
	if doctorName != "" {
		println("Entrou!")
		sql = "select distinct on (d.cpf) d.id, d.name, d.image_url, s.id, s.description as specialty from doctor d " +
			"left join medical_specialty ms on d.id = ms.doctor_id " +
			"left join specialty s on ms.specialty_id = s.id where d.active is true and d.image_url is not null" +
			" and LOWER(d.name) like '%' || $1 || '%'"
	}
	if specialtyDescription != "" {
		sql = "select distinct on (d.cpf) d.id, d.name, d.image_url, s.id, s.description as specialty from doctor d " +
			"left join medical_specialty ms on d.id = ms.doctor_id " +
			"left join specialty s on ms.specialty_id = s.id where d.active is true and d.image_url is not null " +
			"and s.description like '%' || $1 || '%'"
	}
	_, err = db.Prepare(sql)
	if err != nil {
		println("Error3: ", err.Error())
	}

	// Adicionar condições dinamicamente
	//var params []interface{}
	//
	//if doctorName != "" {
	//	sql += "AND d.name LIKE $1 "
	//	params = append(params, "%"+doctorName+"%")
	//}
	//
	//if specialtyDescription != "" {
	//	sql += "AND s.description LIKE $2 "
	//	params = append(params, "%"+specialtyDescription+"%")
	//}

	var doctorNameDB,
		doctorImageUrlDB,
		doctorSpecialtyDB string
	var doctorIdDB, doctorSpecialtyIdDB int

	if doctorName == "" && specialtyDescription == "" {
		globalRows, err = db.Query(sql)
		if err != nil {
			return doctors, err
		}
	} else if doctorName != "" {
		globalRows, err = db.Query(sql,
			strings.ToLower(doctorName))
		if err != nil {
			return doctors, err
		}
	} else if specialtyDescription != "" {
		globalRows, err = db.Query(sql,
			strings.ToLower(specialtyDescription))
		if err != nil {
			return doctors, err
		}
	} else {
		globalRows, err = db.Query(sql,
			strings.ToLower(doctorName),
			strings.ToLower(specialtyDescription))
		if err != nil {
			return doctors, err
		}
	}

	for globalRows.Next() {
		err = globalRows.Scan(
			&doctorIdDB,
			&doctorNameDB,
			&doctorImageUrlDB,
			&doctorSpecialtyIdDB,
			&doctorSpecialtyDB)
		if err != nil {
			println("Error nos dados retornados: ", err.Error())
			return doctors, err
		}
		doctor := response.DoctorResponse{
			User: response.UserResponse{
				Id:       doctorIdDB,
				Name:     doctorNameDB,
				ImageUrl: doctorImageUrlDB,
			},
			Specialty: response.SpecialtyResponse{
				Id:          doctorSpecialtyIdDB,
				Description: doctorSpecialtyDB,
			},
		}
		logger.Println("[DAO.DoctorSelectAll] doctorName: " + doctorName + " - " + response.LogUserResponse(doctor.User))

		doctors = append(doctors, doctor)
	}

	return doctors, err
}

func DoctorSelectById(doctorId int) (model.Doctor, error) {
	db, err := connection.NewConnection()
	if err != nil {
		println("Error1: ", err.Error())
	}
	defer db.Close()

	sql := "select distinct on (d.cpf) d.name, d.birthdate, d.sex, d.cpf, d.address, cd.number, d.crm, d.image_url, dai.doctor_email, s.description as specialty, s.id as specialty_id, d.active from doctor d " +
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
		doctorCrm,
		doctorImageUrlDB,
		doctorEmailDB,
		doctorSpecialtyDB string
	var doctorBirthdateDB time.Time
	var doctorActiveDB bool
	var doctor model.Doctor
	var doctorSpecialtyId int
	rows, err := db.Query(sql, doctorId)
	if err != nil {
		return doctor, err
	}

	for rows.Next() {
		err = rows.Scan(
			&doctorNameDB,
			&doctorBirthdateDB,
			&doctorSexDB,
			&doctorCpfDB,
			&doctorAddressDB,
			&doctorNumberDB,
			&doctorCrm,
			&doctorImageUrlDB,
			&doctorEmailDB,
			&doctorSpecialtyDB,
			&doctorSpecialtyId,
			&doctorActiveDB)
		if err != nil {
			println("Error nos dados retornados: ", err.Error())
			return doctor, err
		}
	}
	doctor = model.NewDoctor(doctorNameDB, doctorBirthdateDB, doctorCpfDB, doctorSexDB, doctorAddressDB, doctorEmailDB, model.NewCellphoneUser(doctorNumberDB), "", doctorImageUrlDB, doctorCrm,
		model.NewSpecialty(doctorSpecialtyId, doctorSpecialtyDB))
	doctor.SetUserActive(doctorActiveDB)

	return doctor, nil

}

func DoctorEdit(doctorId int, doctor model.Doctor) (bool, error) {
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

	sql := "update doctor set name = $1, birthdate = $2, cpf = $3, sex = $4, address = $5, crm = $6, image_url = $7, " +
		"last_modified_date = current_timestamp where id = $8"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error2: ", err.Error())
		return success, err
	}
	println("Id do Médico: ", doctor.GetUser().GetId())
	_, err = tx.Exec(sql,
		doctor.GetUser().GetName(),
		doctor.GetUser().GetBirthDate(),
		doctor.GetUser().GetCpf(),
		doctor.GetUser().GetSex(),
		doctor.GetUser().GetAddress(),
		doctor.GetCrm(),
		doctor.GetUser().GetImageUrl(),
		doctorId)
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

	_, err = tx.Exec(sql,
		doctor.GetUser().GetCellphoneUser().GetNumber(),
		doctorId)
	if err != nil {
		tx.Rollback()
		println("Error5: ", err.Error())
		return success, err
	}

	sql = "update medical_specialty set specialty_id = $1 where doctor_id=$2"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error5: ", err.Error())
		return success, err
	}

	_, err = tx.Exec(sql,
		doctor.GetSpecialty().GetId(),
		doctorId,
	)
	if err != nil {
		tx.Rollback()
		println("Error6: ", err.Error())
		return success, err
	}

	if (doctor.GetUser().GetEmail() == "") || (doctor.GetUser().GetPassword() == "") {
		tx.Commit()
	} else {
		err = DoctorEditLogin(doctor.GetUser().GetEmail(), doctor.GetUser().GetPassword(), doctorId, tx)
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
	var success bool
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
	return success, err
}

func SelectSpecialties() ([]response.SpecialtyResponse, error) {
	db, err := connection.NewConnection()
	if err != nil {
		println("Error1: ", err.Error())
	}
	defer db.Close()

	sql := "select id, description from specialty"
	_, err = db.Prepare(sql)
	if err != nil {
		println("Error3: ", err.Error())
	}

	var id int
	var description string

	var specialties []response.SpecialtyResponse
	rows, err := db.Query(sql)
	if err != nil {
		println("Error nos dados retornados: ", err.Error())
		return specialties, err
	}

	for rows.Next() {
		err = rows.Scan(&id, &description)
		if err != nil {
			println("Error nos dados retornados: ", err.Error())
			return specialties, err
		}
		specialty := response.SpecialtyResponse{
			Id:          id,
			Description: description,
		}
		specialties = append(specialties, specialty)
	}

	return specialties, nil

}

func ValidateEmailDoctor(email string) (bool, error) {
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

	sql := "select exists (select id from doctor_authentication_information where doctor_email=$1) as ex"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	var exists bool
	err = tx.QueryRow(sql, email).Scan(&exists)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	if exists == true {
		tx.Rollback()
		return false, err
	}

	return true, err
}

func ValidateCPFDoctor(cpf string) (bool, error) {
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

	sql := "select exists (select id from doctor where cpf=$1) as ex"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	var exists bool
	err = tx.QueryRow(sql, cpf).Scan(&exists)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	if exists == true {
		tx.Rollback()
		return false, err
	}

	return true, err
}
