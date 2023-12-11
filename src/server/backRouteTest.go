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

	router.HandleFunc("/api/patients/{id}/deactivate", routers.DeactivatePatient).Methods(http.MethodPost)

	//DOCTOR ROUTERS
	router.HandleFunc("/api/doctors", routers.CreateDoctor).Methods(http.MethodPost)

	router.HandleFunc("/api/doctors/{id}", routers.GetDoctorById).Methods(http.MethodGet)

	router.HandleFunc("/api/doctors/{id}", routers.EditDoctor).Methods(http.MethodPut)

	router.HandleFunc("/api/doctors/login", routers.ValidateLoginDoctor).Methods(http.MethodPost)

	router.HandleFunc("/api/doctors/deactivate", routers.DeactivateDoctor).Methods(http.MethodPost)

	router.HandleFunc("/api/upload", routers.UploadFile).Methods(http.MethodPost)

	// Inicialize o servidor na porta desejada
	//http.Handle("/", c.Handler(createPatientRouter))
	//http.Handle("/editRequestPatient", c.Handler(selectByIdPatientRouter))

	// server static files
	fs := http.FileServer(http.Dir("./tmp"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	handler := c.Handler(router)

	println("Servidor ligado porta :3001!")
	log.Fatal(http.ListenAndServe(":3001", handler))
}
