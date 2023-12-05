package server

import (
	"MedGestao/src/routers"
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func OpenServerTest() {
	muxRouter := http.NewServeMux()

	//PATIENT
	createPatientRouter := mux.NewRouter()
	selectByIdPatientRouter := mux.NewRouter()
	editPatientRouter := mux.NewRouter()
	validatePatientLoginRouter := mux.NewRouter()
	deactivatePatientRouter := mux.NewRouter()

	//DOCTOR
	createDoctorRouter := mux.NewRouter()
	selectByIdDoctorRouter := mux.NewRouter()
	editDoctorRouter := mux.NewRouter()
	validadeDoctorLoginRouter := mux.NewRouter()
	deactivateDoctorRouter := mux.NewRouter()

	//MEDICAL SCHEDULE
	createMedicalScheduleRouter := mux.NewRouter()

	// Configuração do CORS
	//c := cors.AllowAll()

	// Crie um middleware de CORS com configurações padrão

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"}, // Origens permitidas (adapte ao seu ambiente)
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	//PATIENT ROUTERS
	// Defina uma rota para lidar com solicitações POST para criar um paciente
	createPatientRouter.HandleFunc("/api/createPatient", routers.CreatePatient).Methods("POST")
	// Lide com solicitações createPatient com o middleware CORS
	muxRouter.Handle("/api/createPatient", c.Handler(createPatientRouter))

	selectByIdPatientRouter.HandleFunc("/api/selectByIdPatient", routers.GetPatientById).Methods("POST")
	//selectByIdPatientRouter.HandleFunc("/api/selectByIdPatient/{id}", routers.GetPatientById).Methods("GET")
	// Lide com solicitações selectByIdPatient com o middleware CORS
	muxRouter.Handle("/api/selectByIdPatient", c.Handler(selectByIdPatientRouter))

	editPatientRouter.HandleFunc("/api/editPatient", routers.EditPatient).Methods("POST")
	muxRouter.Handle("/api/editPatient", c.Handler(editPatientRouter))

	validatePatientLoginRouter.HandleFunc("api/login", routers.ValidateLoginPatient).Methods("POST")
	muxRouter.Handle("/api/login", c.Handler(validatePatientLoginRouter))

	deactivatePatientRouter.HandleFunc("/api/deactivatePatient", routers.DeactivatePatient).Methods("POST")
	//selectByIdPatientRouter.HandleFunc("/api/selectByIdPatient/{id}", routers.GetPatientById).Methods("GET")
	// Lide com solicitações selectByIdPatient com o middleware CORS
	muxRouter.Handle("/api/deactivatePatient", c.Handler(deactivatePatientRouter))

	//DOCTOR ROUTERS
	createDoctorRouter.HandleFunc("/api/createDoctor", routers.CreateDoctor).Methods("POST")
	// Lide com solicitações createPatient com o middleware CORS
	muxRouter.Handle("/api/createDoctor", c.Handler(createDoctorRouter))

	selectByIdDoctorRouter.HandleFunc("/api/selectByIdDoctor", routers.GetDoctorById).Methods("POST")
	//selectByIdPatientRouter.HandleFunc("/api/selectByIdPatient/{id}", routers.GetPatientById).Methods("GET")
	// Lide com solicitações selectByIdPatient com o middleware CORS
	muxRouter.Handle("/api/selectByIdDoctor", c.Handler(selectByIdDoctorRouter))

	editDoctorRouter.HandleFunc("/api/editDoctor", routers.EditDoctor).Methods("POST")
	muxRouter.Handle("/api/editDoctor", c.Handler(editDoctorRouter))

	validadeDoctorLoginRouter.HandleFunc("/api/validateLoginDoctor", routers.ValidateLoginDoctor).Methods("POST")
	muxRouter.Handle("/api/validateLoginDoctor", c.Handler(validadeDoctorLoginRouter))

	deactivateDoctorRouter.HandleFunc("/api/deactivateDoctor", routers.DeactivateDoctor).Methods("POST")
	//selectByIdPatientRouter.HandleFunc("/api/selectByIdPatient/{id}", routers.GetPatientById).Methods("GET")
	// Lide com solicitações selectByIdPatient com o middleware CORS
	muxRouter.Handle("/api/deactivateDoctor", c.Handler(deactivateDoctorRouter))

	//MEDICAL SCHEDULE ROUTE
	createMedicalScheduleRouter.HandleFunc("/api/createMedicalSchedule", routers.CreateMedicalSchedule)
	muxRouter.Handle("/api/createMedicalSchedule", c.Handler(createMedicalScheduleRouter))

	// Inicialize o servidor na porta desejada
	//http.Handle("/", c.Handler(createPatientRouter))
	//http.Handle("/editRequestPatient", c.Handler(selectByIdPatientRouter))
	println("Servidor ligado na porta :3001!")
	log.Fatal(http.ListenAndServe(":3001", muxRouter))
}
