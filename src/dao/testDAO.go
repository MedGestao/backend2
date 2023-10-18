package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
	"fmt"
)

func InsertTest(patient model.Patient) error {
	db, err := connection.NewConnection()
	defer db.Close()
	if err != nil {
		println(err.Error())
		panic(err)
	}

	//tx, err := db.Begin()

	sql := "insert into patient (name, birthdate, cpf, sex, address, email) values ($1, $2, $3, $4, $5, $6) returning id"
	//_, err = db.Prepare(sql)
	//if err != nil {
	//	println("Error1: ", err.Error())
	//	return err
	//}
	//
	var tempPatientId int
	//err = db.QueryRow(
	//	patient.GetUser().GetName(),
	//	string(patient.GetUser().GetBirthDate().Format("2006-01-02")),
	//	patient.GetUser().GetCpf(),
	//	patient.GetUser().GetSex(),
	//	patient.GetUser().GetAddress(),
	//	patient.GetUser().GetEmail(),
	//).Scan(&tempPatientId)
	//
	//if err != nil {
	//	println("Error2: ", err.Error())
	//	panic(err)
	//}

	tx, err := db.Begin()
	//name := "TESTE9"
	//email := "teste9Email@gmail.com"
	//birthDate, err := time.Parse("2006-01-02", "1969-02-20")
	number := "82999978864"
	if err != nil {
		tx.Rollback()
		fmt.Println("Error1", err.Error())
		panic(err)
	}

	//sql := "insert into test (name, email, birth_date) values ($1, $2, $3) returning id"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error2", err.Error())
		panic(err)
	}
	//_, err = db.Exec("INSERT INTO test (name) VALUES ($1)", name)
	err = tx.QueryRow(sql,
		patient.GetUser().GetName(),
		patient.GetUser().GetBirthDate().Format("2006-01-02"),
		patient.GetUser().GetCpf(),
		patient.GetUser().GetSex(),
		patient.GetUser().GetAddress(),
		patient.GetUser().GetEmail()).Scan(&tempPatientId)
	println("Id: ", tempPatientId)
	if err != nil {
		tx.Rollback()
		println("Error3: " + err.Error())
		panic(err)
	}

	sql = "insert into cellphone_test(id_test, number) values ($1, $2)"
	_, err = tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		println("Error4", err.Error())
		panic(err)
	}

	_, err = tx.Exec(sql, tempPatientId, number)
	if err != nil {
		tx.Rollback()
		println(err.Error())
		panic(err)
	}

	tx.Commit()
	return nil
}
