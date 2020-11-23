package models

import "time"

type Classes struct {
	Id             int64     `db:"Id"`
	SchoolId       int64     `db:"SchoolId"`
	Name           string    `db:"Name"`
	NumberOfPupils int64     `db:"NumberOfPupils"`
	CreatedAt      time.Time `db:"CreatedAt"`
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
	CreatedAt  time.Time `db:"CreatedAt"`
	UpdatedAt  time.Time `db:"UpdatedAt"`
}
type Schools struct {
	Id              int64     `db:"id"`
	Name            string    `db:"name"`
	NumberOfClasses int64     `db:"numberofclasses"`
	CallCenter      string    `db:"callcenter"`
	Address         string    `db:"address"`
	CreatedAt       time.Time `db:"createdat"`
	UpdatedAt       time.Time `db:"updatedClassIdt"`
}

type Scores struct {
	Id        int64     `db:"Id"`
	SubjectId int64     `db:"SubjectId"`
	PupilId   int64     `db:"PupilId"`
	Score     string    `db:"Score"`
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}

type Subject struct {
	Id        int64     `db:"Id"`
	Name      string    `db:"Name"`
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}

type User struct {
	Login    string `db:"username" validate:"required,cyr,max=50,min=2"`
	Password string `db:"password" validate:"required,alpha,cyr,max=50,min=4"`
}