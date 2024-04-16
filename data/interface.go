package data

import "time"

type Id string

type Dose struct {
	Id        Id
	Time      time.Time
	Quantity  string
	Substance string
	Route     string
}

type Storage interface {
	FetchAll() ([]Dose, error)
	Append(Dose) error
	DeleteDose(Id) error
}
