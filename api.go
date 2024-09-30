package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func CreateAPISever(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/patient", makeHTTPHandleFunc(s.handlePatient))
	router.HandleFunc("/patient/{id}", makeHTTPHandleFunc(s.handleGetPatientByID))
	router.HandleFunc("/doctor", makeHTTPHandleFunc(s.handleDoctor))
	router.HandleFunc("/doctor/{id}", makeHTTPHandleFunc(s.handleGetDoctorByID))
	router.HandleFunc("/appointment", makeHTTPHandleFunc((s.handleAppointment)))
	router.HandleFunc("/appointment/{id}", makeHTTPHandleFunc(s.handleGetAppointmentByID))
	router.HandleFunc("/appointmentbypatient/{id}", makeHTTPHandleFunc(s.handleGetAppointmentByPatient))
	router.HandleFunc("/appointmentbydoctor/{id}", makeHTTPHandleFunc(s.handleGetAppointmentByDoctor))
	router.HandleFunc("/checkdoctoravailability", makeHTTPHandleFunc(s.handleCheckDoctorAvailability))

	log.Println("API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// Basic CRUD Functionality
func (s *APIServer) handlePatient(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.HandleGetPatient(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreatePatient(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleDoctor(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetDoctor(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateDoctor(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleAppointment(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAppointment(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAppointment(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) HandleGetPatient(w http.ResponseWriter, r *http.Request) error {
	patients, err := s.store.GetPatients()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, patients)
}

func (s *APIServer) handleGetDoctor(w http.ResponseWriter, r *http.Request) error {
	doctors, err := s.store.GetDoctors()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, doctors)
}

func (s *APIServer) handleGetAppointment(w http.ResponseWriter, r *http.Request) error {
	patients, err := s.store.GetAppointments()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, patients)
}

func (s *APIServer) handleGetPatientByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getID(r)
		if err != nil {
			return err
		}

		patient, err := s.store.GetPatientByID(id)
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, patient)
	}

	if r.Method == "DELETE" {
		return s.handleDeletePatient(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleGetDoctorByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getID(r)
		if err != nil {
			return err
		}

		doctor, err := s.store.GetDoctorByID(id)
		if err != nil {
			return err
		}
		return WriteJSON(w, http.StatusOK, doctor)
	}

	if r.Method == "DELETE" {
		return s.handleDeletePatient(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleGetAppointmentByID(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "DELETE" {
		return s.handleDeleteAppointment(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)

}

func (s *APIServer) handleGetAppointmentByPatient(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "DELETE" {
		return s.handleDeletePatientAppointments(w, r)
	}

	id, err := getID(r)
	if err != nil {
		return err
	}

	appointments, err := s.store.GetPatientAppointments(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, appointments)

}

func (s *APIServer) handleGetAppointmentByDoctor(w http.ResponseWriter, r *http.Request) error {

	id, err := getID(r)
	if err != nil {
		return err
	}

	appointments, err := s.store.GetDoctorAppointments(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, appointments)

}

func (s *APIServer) handleCreatePatient(w http.ResponseWriter, r *http.Request) error {
	createPatientReq := new(CreatePatientRequest)
	if err := json.NewDecoder(r.Body).Decode(createPatientReq); err != nil {
		return err
	}

	patient := NewPatient(createPatientReq.FirstaName, createPatientReq.LastName)
	if err := s.store.CreatePatient(patient); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, patient)
}

func (s *APIServer) handleCreateDoctor(w http.ResponseWriter, r *http.Request) error {
	createDoctorReq := new(CreateDoctorRequest)
	if err := json.NewDecoder(r.Body).Decode(createDoctorReq); err != nil {
		return err
	}

	doctor := NewDoctor(createDoctorReq.FirstaName, createDoctorReq.LastName)
	if err := s.store.CreateDoctor(doctor); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, doctor)
}

func (s *APIServer) handleCreateAppointment(w http.ResponseWriter, r *http.Request) error {

	var appEnd time.Time
	createAppointmentReq := new(CreateAppointmentRequest)
	if err := json.NewDecoder(r.Body).Decode(createAppointmentReq); err != nil {
		return err
	}

	if createAppointmentReq.AppType == "quick" {
		appEnd = createAppointmentReq.AppStart.Add(time.Minute * 30)
	}

	if createAppointmentReq.AppType == "extensive" {
		appEnd = createAppointmentReq.AppStart.Add(time.Hour * 1)
	}

	if createAppointmentReq.AppType == "operation" {
		appEnd = createAppointmentReq.AppStart.Add(time.Hour * 2)
	}

	appointment := NewAppointment(createAppointmentReq.Doctor, createAppointmentReq.Patient, createAppointmentReq.AppType, createAppointmentReq.AppStart, appEnd)
	if err := s.store.CreateAppointment(appointment); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, appointment)
}

func (s *APIServer) handleDeletePatient(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeletePatient(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func (s *APIServer) handleDeleteAppointment(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteAppointment(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}

// Special Clinic Functionality

// Cancel all Appointments of a Patient
func (s *APIServer) handleDeletePatientAppointments(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeletePatientAppointments(id); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// Check if Doctor has appointments in the given timeframe
func (s *APIServer) handleCheckDoctorAvailability(w http.ResponseWriter, r *http.Request) error {
	checkDoctorAvailabilityRequest := new(CheckDoctorAvailabilityRequest)
	if err := json.NewDecoder(r.Body).Decode(checkDoctorAvailabilityRequest); err != nil {
		return err
	}

	appointments, err := s.store.CheckDoctorAvailability(checkDoctorAvailabilityRequest.Doctor, checkDoctorAvailabilityRequest.TimeStart, checkDoctorAvailabilityRequest.TimeEnd)

	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, appointments)
}
