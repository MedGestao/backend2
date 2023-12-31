package server

import (
	"MedGestao/src/routers"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func OpenServerTest() {
	router := mux.NewRouter()

	// Configuração do CORS
	//c := cors.AllowAll()

	// Crie um middleware de CORS com configurações padrão

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001", "http://localhost:3000"}, // Origens permitidas (adapte ao seu ambiente)
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	//PATIENT ROUTERS
	// Defina uma rota para lidar com solicitações POST para criar um paciente
	router.HandleFunc("/api/patients", routers.CreatePatient).Methods(http.MethodPost)

	router.HandleFunc("/api/patients/{id}", routers.GetPatientById).Methods(http.MethodGet)

	router.HandleFunc("/api/patients/{id}", routers.EditPatient).Methods(http.MethodPut)

	router.HandleFunc("/api/patients/login", routers.ValidateLoginPatient).Methods(http.MethodPost)

	router.HandleFunc("/api/patients/deactivate", routers.DeactivatePatient).Methods(http.MethodPost)

	//DOCTOR ROUTERS
	router.HandleFunc("/api/doctors", routers.CreateDoctor).Methods(http.MethodPost)

	router.HandleFunc("/api/doctors", routers.GetDoctorsAll).Queries("doctorName", "{doctorName}").Methods(http.MethodGet)

	router.HandleFunc("/api/doctors/{id}", routers.GetDoctorById).Methods(http.MethodGet)

	router.HandleFunc("/api/doctors/{id}", routers.EditDoctor).Methods(http.MethodPut)

	router.HandleFunc("/api/doctors/login", routers.ValidateLoginDoctor).Methods(http.MethodPost)

	router.HandleFunc("/api/validate-email", routers.ValidateEmailDoctor).Queries("email", "{email}").Methods(http.MethodGet)

	router.HandleFunc("/api/validate-cpf", routers.ValidateCPFDoctor).Queries("cpf", "{cpf}").Methods(http.MethodGet)

	router.HandleFunc("/api/doctors/deactivate", routers.DeactivateDoctor).Methods(http.MethodPost)

	router.HandleFunc("/api/specialties", routers.GetSpecialty).Methods(http.MethodGet)

	router.HandleFunc("/api/upload", routers.UploadFile).Methods(http.MethodPost)

	//MEDICAL SCHEDULE ROUTERS
	router.HandleFunc("/api/doctors/schedule", routers.CreateMedicalSchedule).Methods(http.MethodPost)

	router.HandleFunc("/api/doctors/{id}/schedule", routers.GetMedicalScheduleAllByIdDoctor).
		Queries("selectedDate", "{selectedDate}", "selectedDay", "{selectedDay}").
		Methods(http.MethodGet)

	router.HandleFunc("/api/medicalSchedule/{id}", routers.GetMedicalScheduleById).Methods(http.MethodGet)

	router.HandleFunc("/api/medicalSchedule/{id}", routers.EditMedicalSchedule).Methods(http.MethodPut)

	// PATIENT DOCTOR CONSULTATION ROUTERS
	router.HandleFunc("/api/patientDoctorConsultation", routers.CreatePatientDoctorConsutation).Methods(http.MethodPost)

	router.HandleFunc("/api/patientDoctorConsultation/searchByDoctor/{id}", routers.GetPatientDoctorConsultationAllByIdDoctor).Methods(http.MethodPut)

	router.HandleFunc("/api/patientDoctorConsultation/searchByPatient/{id}", routers.GetPatientDoctorConsultationAllByIdPatient).Methods(http.MethodGet)

	router.HandleFunc("/api/patientDoctorConsultation/{id}", routers.GetPatientDoctorConsultationById).Methods(http.MethodGet)

	router.HandleFunc("/api/patientDoctorConsultation/{id}", routers.EditPatientDoctorConsultation).Methods(http.MethodPut)

	router.HandleFunc("/api/patientDoctorConsultation/completePatientDoctorConsultation/{id}", routers.CompletePatientDoctorConsultation).Methods(http.MethodGet)

	router.HandleFunc("/api/patientDoctorConsultation/deactivatePatientDoctorConsultation/{id}", routers.DeactivatePatientDoctorConsultation).Methods(http.MethodGet)

	// server static files
	fs := http.FileServer(http.Dir("./tmp"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	handler := c.Handler(router)

	println("Servidor ligado na porta :3001!")
	log.Fatal(http.ListenAndServe(":3001", handler))
}
