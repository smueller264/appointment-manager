package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Patient Tests
func TestHandleGetPatient(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/patient", nil)

	server.HandleGetPatient(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but god %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleGetPatientByID(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/patient/{id}", nil)

	server.handleGetPatientByID(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleCreatePatient(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/patient", nil)

	server.handleCreatePatient(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleDeletePatient(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/patient/{id}", nil)

	server.handleDeletePatient(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

//Doctor Tests

func TestHandleGetDoctor(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/doctor", nil)

	server.handleGetDoctor(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleGetDoctorByID(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/doctor/{id}", nil)

	server.handleGetDoctorByID(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleCreateDoctor(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/doctor", nil)

	server.handleCreateDoctor(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleCheckDoctorAvailability(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/checkdoctoravailability", nil)

	server.handleCheckDoctorAvailability(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

// Appointment tests
func TestHandleGetAppointment(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/appointment", nil)

	server.handleGetAppointment(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but god %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleGetAppointmentByID(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/appointment/{id}", nil)

	server.handleGetAppointmentByID(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleGetAppointmentByPatient(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/appointmentbypatient/{id}", nil)

	server.handleGetAppointmentByPatient(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleGetAppointmentByDoctor(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/appointmentbydoctor/{id}", nil)

	server.handleGetAppointmentByDoctor(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleCreateAppointment(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/appointment", nil)

	server.handleCreateAppointment(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}

func TestHandleDeleteAppointment(t *testing.T) {
	store := MocksStore{}
	server := CreateAPISever(":5173", store)

	r := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/appointment/{id}", nil)

	server.handleCreateAppointment(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", r.Result().StatusCode)
	}

	defer r.Result().Body.Close()

	_, err := io.ReadAll(r.Result().Body)
	if err != nil {
		t.Error(err)
	}

}
