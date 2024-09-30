package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePatient(*Patient) error
	GetPatientByID(int) (*Patient, error)
	GetPatients() ([]*Patient, error)
	DeletePatient(int) error
	CreateDoctor(*Doctor) error
	GetDoctors() ([]*Doctor, error)
	GetDoctorByID(int) (*Doctor, error)
	CreateAppointment(*Appointment) error
	DeleteAppointment(int) error
	GetPatientAppointments(int) ([]*Appointment, error)
	GetDoctorAppointments(int) ([]*Appointment, error)
	GetAppointments() ([]*Appointment, error)
	DeletePatientAppointments(int) error 
	CheckDoctorAvailability(int, time.Time, time.Time) ([]*Appointment, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=go sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	if err := s.CreatePatientTable(); err != nil {
		return err
	}
	if err := s.CreateDoctorTable(); err != nil {
		return err
	}
	if err := s.CreateAppointmentTable(); err != nil {
		return err
	}
	return nil
	
}

func (s *PostgresStore) CreatePatientTable() error {
	query := `create table if not exists patient (
		id serial primary key,
		first_name varchar,
		last_name varchar,
		insurance_number serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateDoctorTable() error {
	query := `create table if not exists doctor (
		id serial primary key,
		first_name varchar,
		last_name varchar,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAppointmentTable() error {
	query := `create table if not exists appointment (
		id serial primary key,
		doctor int,
		patient int,
		appType text,
		appStart timestamp,
		appEnd timestamp,
		created_at timeStamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) DeletePatient(id int) error {
	_, err := s.db.Query("delete from patient where id = $1", id)
	return err
}

func (s *PostgresStore) DeleteAppointment(id int) error {
	_, err := s.db.Query("delete from appintment where id = $1", id)
	return err
}

func (s *PostgresStore) DeletePatientAppointments(id int) error {
	_, err := s.db.Query("delete from appointment where patient = $1", id)
	return err
}




func (s *PostgresStore) CreatePatient(patient *Patient) error {
	
	query := `insert into patient 
	(first_name, last_name, insurance_number, created_at)
	values ($1, $2, $3, $4)`
	resp, err := s.db.Query(query, patient.FirstName, patient.LastName, patient.InsuranceNumber, patient.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) CreateDoctor(doctor *Doctor) error {
	
	query := `insert into doctor 
	(first_name, last_name, created_at)
	values ($1, $2, $3)`
	resp, err := s.db.Query(query, doctor.FirstName, doctor.LastName, doctor.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) CreateAppointment(appointment *Appointment) error {
	
	query := `insert into appointment
	(doctor, patient, appType, appStart, appEnd, created_at)
	values ($1, $2, $3, $4, $5, $6)`
	resp, err := s.db.Query(query, appointment.Doctor, appointment.Patient, appointment.AppType, appointment.AppStart, appointment.AppEnd, appointment.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) GetPatientByID(id int) (*Patient, error) {
	rows, err := s.db.Query("select * from patient where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoPatient(rows)
	}
	return nil, fmt.Errorf("patient %d not found", id)
}

func (s *PostgresStore) GetDoctorByID(id int) (*Doctor, error) {
	rows, err := s.db.Query("select * from doctor where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoDoctor(rows)
	}
	return nil, fmt.Errorf("doctor %d not found", id)
}

func (s *PostgresStore) GetPatients() ([]*Patient, error) {
	rows, err := s.db.Query("select * from patient")
	if err != nil {
		return nil, err
	}

	patients := []*Patient{}
	for rows.Next() {
		patient, err := scanIntoPatient(rows)
		if err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

func (s *PostgresStore) GetDoctors() ([]*Doctor, error) {
	rows, err := s.db.Query("select * from doctor")
	if err != nil {
		return nil, err
	}

	doctors := []*Doctor{}
	for rows.Next() {
		doctor, err := scanIntoDoctor(rows)
		if err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (s *PostgresStore) GetAppointments() ([]*Appointment, error) {
	rows, err := s.db.Query("select * from appointment")
	if err != nil {
		return nil, err
	}

	appointments := []*Appointment{}
	for rows.Next() {
		appointment, err := scanIntoAppintment(rows)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (s *PostgresStore) GetDoctorAppointments(id int) ([]*Appointment, error) {
	rows, err := s.db.Query("select * from appointment where doctor = $1", id)
	if err != nil {
		return nil, err
	}

	appointments := []*Appointment{}
	for rows.Next() {
		appointment, err := scanIntoAppintment(rows)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (s *PostgresStore) CheckDoctorAvailability(id int, timeStart time.Time, timeEnd time.Time) ([]*Appointment, error) {
	
	fmt.Println(id, timeStart, timeEnd)
	rows, err := s.db.Query("select * from appointment where doctor = $1 and (appStart < $3 and appEnd > $2)", id, timeStart, timeEnd)

	if err != nil {
		return nil, err
	}
	appointments := []*Appointment{}
	for rows.Next() {
		appointment, err := scanIntoAppintment(rows)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (s *PostgresStore) GetPatientAppointments(id int) ([]*Appointment, error) {
	rows, err := s.db.Query("select * from appointment where patient = $1", id)
	if err != nil {
		return nil, err
	}

	appointments := []*Appointment{}
	for rows.Next() {
		appointment, err := scanIntoAppintment(rows)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func scanIntoPatient(rows *sql.Rows) (*Patient, error) {
	patient := new(Patient)
	err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.InsuranceNumber, &patient.CreatedAt)

		if err != nil {
			return nil, err
		}

	return patient, err
}

func scanIntoDoctor(rows *sql.Rows) (*Doctor, error) {
	doctor := new(Doctor)
	err := rows.Scan(&doctor.ID, &doctor.FirstName, &doctor.LastName, &doctor.CreatedAt)

		if err != nil {
			return nil, err
		}

	return doctor, err
}

func scanIntoAppintment(rows *sql.Rows) (*Appointment, error) {
	appointment := new(Appointment)
	err := rows.Scan(&appointment.ID, &appointment.Doctor, &appointment.Patient, &appointment.AppType, &appointment.AppStart, &appointment.AppEnd, &appointment.CreatedAt)

		if err != nil {
			return nil, err
		}

	return appointment, err
}