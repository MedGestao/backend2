package dao

import (
	"MedGestao/src/connection"
	"MedGestao/src/model"
)

func InsertSpecialty(specialty model.Specialty) (bool, error) {
	db, err := connection.NewConnection()
	defer db.Close()
	if err != nil {
		return false, err
	}

	if err != nil {
		println(err)
		return false, err
	}

	sql := "insert into specialty(description) values($1)"
	_, err = db.Prepare(sql)
	if err != nil {
		println("Error1: ", err.Error())
		return false, err
	}

	_, err = db.Exec(sql,
		specialty.GetDescription())

	if err != nil {
		println("Error2: ", err.Error())
		return false, err
	}

	return true, err

}
