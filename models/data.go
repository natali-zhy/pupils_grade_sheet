package models

import "time"

type Classes struct {
	Id             int64     `db:"Id"`
	SchoolId       int64     `db:"SchoolId"`
	Name           string    `db:"Name"`
	NumberOfPupils int64     `db:"NumberOfPupils"`
	CreateAt       time.Time `db:"CreateAt"`
	UpdatedAt      time.Time `db:"UpdatedAt"`
}

type Pupils struct {
	Id         int64     `db:"Id"`
	ClassId    int64     `db:"ClassId"`
	Name       string    `db:"Name"`
	Surname    string    `db:"Surname"`
	Patronymic string    `db:"Patronymic"`
	Gender     string    `db:"Gender"`
	Address    string    `db:"Address"`
	CreateAt   time.Time `db:"CreateAt"`
	UpdatedAt  time.Time `db:"UpdatedAt"`
}
type Schools struct {
	Id              int64     `db:"Id"`
	Name            string    `db:"Name"`
	NumberOfClasses int64     `db:"NumberOfClasses"`
	CallCenter      string    `db:"CallCenter"`
	Address         string    `db:"Address"`
	CreateAt        time.Time `db:"CreateAt"`
	UpdatedAt       time.Time `db:"UpdatedAt"`
}

type Scores struct {
	Id        int64     `db:"Id"`
	SubjectId int64     `db:"SubjectId"`
	PupilId   int64     `db:"PupilId"`
	Score     string    `db:"Score"`
	CreateAt  time.Time `db:"CreateAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}

type Subject struct {
	Id        int64     `db:"Id"`
	Name      string    `db:"Name"`
	CreateAt  time.Time `db:"CreateAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}
