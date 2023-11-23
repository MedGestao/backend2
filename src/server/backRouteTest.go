package server

import (
	"MedGestao/src/routers"
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//type Paciente struct {
//	Email    string `json:"email"`
//	password int    `json:"password"`
//	// Outros campos do paciente
//}

func OpenServerTest() {
	println("Entrou!")
	muxRouter := http.NewServeMux()
	createPatientRouter := mux.NewRouter()
	selectByIdPatientRouter := mux.NewRouter()
	editPatientRouter := mux.NewRouter()
	validatePatientLoginRouter := mux.NewRouter()
	deactivatePatientRouter := mux.NewRouter()

	// Configuração do CORS
	//c := cors.AllowAll()

	// Crie um middleware de CORS com configurações padrão

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"}, // Origens permitidas (adapte ao seu ambiente)
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

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

	// Inicialize o servidor na porta desejada
	//http.Handle("/", c.Handler(createPatientRouter))
	//http.Handle("/editRequestPatient", c.Handler(selectByIdPatientRouter))
	log.Fatal(http.ListenAndServe(":3001", muxRouter))
	print("Servidor ligado porta :3001!")
}
