package controller

import (
	"MedGestao/src/dao"
	"MedGestao/src/model"
)

func RegisterMedicalSchedule(medicalSchedule model.MedicalSchedule) (bool, error) {
	var success bool
	var err error

	if medicalSchedule == (model.MedicalSchedule{}) {
		println("Nenhum dado recebido")
		return success, err
	}

	success, err = dao.MedicalScheduleInsert(medicalSchedule)
	if err != nil {
		return success, err
	}

	return success, err
}

func SearchAllMedicalSchedule() ([]model.MedicalSchedule, error) {
	medicalScheduleList, err := dao.MedicalScheduleSelectAll()
	if err != nil {
		return []model.MedicalSchedule{}, err
	}

	return medicalScheduleList, err

}

func SearchByIdMedicalSchedule(id int) (model.MedicalSchedule, error) {
	var medicalSchedule model.MedicalSchedule
	var err error

	if id == 0 {
		return medicalSchedule, err
	}

	medicalSchedule, err = dao.MedicalScheduleSelectById(id)
	if err != nil {
		return medicalSchedule, err
	}

	return medicalSchedule, err
}

func EditMedicalSchedule(medicalSchedule model.MedicalSchedule) (bool, error) {
	var success bool
	var err error
	if medicalSchedule == (model.MedicalSchedule{}) {
		return success, err
	}

	success, err = dao.MedicalScheduleEdit(medicalSchedule)
	if err != nil {
		return success, err
	}

	return success, err
}

func OffMedicalSchedule(id int) (bool, error) {
	var success bool
	var err error
	if id == 0 {
		return success, err
	}

	success, err = dao.MedicalScheduleDelete(id)
	if err != nil {
		return success, err
	}

	return success, err
}
