package main

import "time"

type MocksStore struct {
}

// CheckDoctorAvailability implements Storage.
func (m MocksStore) CheckDoctorAvailability(int, time.Time, time.Time) ([]*Appointment, error) {
	return nil, nil
}

// CreateAppointment implements Storage.
func (m MocksStore) CreateAppointment(*Appointment) error {
	return nil
}

// CreateDoctor implements Storage.
func (m MocksStore) CreateDoctor(*Doctor) error {
	return nil
}

// CreatePatient implements Storage.
func (m MocksStore) CreatePatient(*Patient) error {
	return nil
}

// DeleteAppointment implements Storage.
func (m MocksStore) DeleteAppointment(int) error {
	return nil
}

// DeletePatient implements Storage.
func (m MocksStore) DeletePatient(int) error {
	return nil
}

// DeletePatientAppointments implements Storage.
func (m MocksStore) DeletePatientAppointments(int) error {
	return nil
}

// GetAppointments implements Storage.
func (m MocksStore) GetAppointments() ([]*Appointment, error) {
	return nil, nil
}

// GetDoctorAppointments implements Storage.
func (m MocksStore) GetDoctorAppointments(int) ([]*Appointment, error) {
	return nil, nil
}

// GetDoctorByID implements Storage.
func (m MocksStore) GetDoctorByID(int) (*Doctor, error) {
	return nil, nil
}

// GetDoctors implements Storage.
func (m MocksStore) GetDoctors() ([]*Doctor, error) {
	return nil, nil
}

// GetPatientAppointments implements Storage.
func (m MocksStore) GetPatientAppointments(int) ([]*Appointment, error) {
	return nil, nil
}

// GetPatientByID implements Storage.
func (m MocksStore) GetPatientByID(int) (*Patient, error) {
	return nil, nil
}

// GetPatients implements Storage.
func (m MocksStore) GetPatients() ([]*Patient, error) {
	return nil, nil
}
