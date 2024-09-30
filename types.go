package main

import (
	"math/rand"
	"time"
)

type CreatePatientRequest struct {
	FirstaName string `json:"firstName"`
	LastName   string `json:"lastName"`
}

type CreateDoctorRequest struct {
	FirstaName string `json:"firstName"`
	LastName   string `json:"lastName"`
}

type CreateAppointmentRequest struct {
	Doctor   int       `json:"doctor"`
	Patient  int       `json:"patient"`
	AppType  string    `json:"appType"`
	AppStart time.Time `json:"appStart"`
}

type CheckDoctorAvailabilityRequest struct {
	Doctor int `json:"doctor"`
	TimeStart time.Time `json:"timeStart"`
	TimeEnd time.Time `json:"timeEnd"`
}

type Patient struct {
	ID              int       `json:"id"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	InsuranceNumber int64     `json:"insuranceNumber"`
	CreatedAt       time.Time `json:"createdAt"`
}

type Doctor struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}

type Appointment struct {
	ID        int       `json:"id"`
	Doctor    int       `json:"doctor"`
	Patient   int       `json:"patient"`
	AppType   string    `json:"appType"`
	AppStart  time.Time `json:"appStart"`
	AppEnd    time.Time `json:"appEnd"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewPatient(firstName, lastName string) *Patient {
	return &Patient{
		FirstName:       firstName,
		LastName:        lastName,
		InsuranceNumber: int64(rand.Intn(1000000)),
		CreatedAt:       time.Now(),
	}
}

func NewDoctor(firstName, lastName string) *Doctor {
	return &Doctor{
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}
}

func NewAppointment(doctor int, patient int, appType string, appStart time.Time, appEnd time.Time) *Appointment {
	return &Appointment{
		Doctor:    doctor,
		Patient:   patient,
		AppType:   appType,
		AppStart:  appStart,
		AppEnd:    appEnd,
		CreatedAt: time.Now(),
	}
}
