package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/natalizhy/pupils_grade_sheet/db"
	"github.com/natalizhy/pupils_grade_sheet/models"
)

type UserTemp struct {
	Classes models.Classes
	Pupils  models.Pupils
	Schools models.Schools
	Scores  models.Scores
	Subject models.Subject
	IsEdit  bool
	Error   map[string]map[string]string
}
type Application struct {
	Addr      string
	AddrHttp  string
	Templates string
	Db        Db
}

type Db struct {
	AddrPostgres     string
	UserPostgres     string
	PasswordPostgres string
	Database         string
}

var config = &Application{}

func main() {

	db.InitDB()

	r := chi.NewRouter()
	r.Get("/schools", GetSchools)
	r.Post("/schools", AddSchools)

	r.Get("/classes", GetClasses)
	r.Post("/classes", AddClasses)

	r.Get("/pupils", GetPupils)
	r.Post("/pupils", AddPupils)

	r.Get("/scores", GetScores)
	r.Post("/scores", AddScores)

	r.Get("/subject", GetSubject)
	r.Post("/subject", AddSubject)

	fileHandle := http.FileServer(http.Dir(".")).ServeHTTP

	r.Get("/assets/*", fileHandle)

	fmt.Println("connect :3005")

	err := http.ListenAndServe(":3005", r)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSchools(w http.ResponseWriter, r *http.Request) {
	schools := models.Schools{Id: 0}
	userTemp := UserTemp{IsEdit: true, Schools: schools}

	RenderTempl(w, "templates/schools.html", userTemp)
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	classes := models.Classes{Id: 0}
	userTemp := UserTemp{IsEdit: true, Classes: classes}

	RenderTempl(w, "templates/classes.html", userTemp)
}

func GetPupils(w http.ResponseWriter, r *http.Request) {
	pupils := models.Pupils{Id: 0}
	userTemp := UserTemp{IsEdit: true, Pupils: pupils}

	RenderTempl(w, "templates/pupils.html", userTemp)
}

func GetScores(w http.ResponseWriter, r *http.Request) {
	scores := models.Scores{Id: 0}
	userTemp := UserTemp{IsEdit: true, Scores: scores}

	RenderTempl(w, "templates/scores.html", userTemp)
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	subject := models.Subject{Id: 0}
	userTemp := UserTemp{IsEdit: true, Subject: subject}

	RenderTempl(w, "templates/subject.html", userTemp)
}

func AddSchools(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")
	userIDedit, _ := strconv.ParseInt(userIDstr, 10, 64)
	schools := models.Schools{Id: userIDedit, NumberOfClasses: userIDedit}

	fmt.Println(userIDedit, userIDedit)
	userTemp := UserTemp{IsEdit: true, Schools: schools}

	schools.Name = r.FormValue("Name")
	schools.CallCenter = r.FormValue("CallCenter")
	schools.Address = r.FormValue("Address")

	fmt.Println(schools)

	userTemp.Schools = schools
	fmt.Println(userTemp.Schools)

	RenderTempl(w, "templates/schools.html", userTemp)
}

func AddClasses(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")
	userIDedit, _ := strconv.ParseInt(userIDstr, 10, 64)
	classes := models.Classes{Id: userIDedit, SchoolId: userIDedit}

	fmt.Println(userIDedit, userIDedit)
	userTemp := UserTemp{IsEdit: true, Classes: classes}

	classes.Name = r.FormValue("Name")

	fmt.Println(classes)

	userTemp.Classes = classes
	fmt.Println(userTemp.Classes)

	RenderTempl(w, "templates/classes.html", userTemp)
}

func AddScores(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")
	userIDedit, _ := strconv.ParseInt(userIDstr, 10, 64)
	scores := models.Scores{Id: userIDedit, SubjectId: userIDedit, PupilId: userIDedit}

	fmt.Println(userIDedit, userIDedit)
	userTemp := UserTemp{IsEdit: true, Scores: scores}

	scores.Score = r.FormValue("Score")

	fmt.Println(scores)

	userTemp.Scores = scores
	fmt.Println(userTemp.Scores)

	RenderTempl(w, "templates/scores.html", userTemp)
}

func AddSubject(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")
	userIDedit, _ := strconv.ParseInt(userIDstr, 10, 64)
	subject := models.Subject{Id: userIDedit}

	fmt.Println(userIDedit, userIDedit)
	userTemp := UserTemp{IsEdit: true, Subject: subject}

	subject.Name = r.FormValue("Name")

	fmt.Println(subject)

	userTemp.Subject = subject
	fmt.Println(userTemp.Subject)

	RenderTempl(w, "templates/subject.html", userTemp)
}

func AddPupils(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")
	userIDedit, _ := strconv.ParseInt(userIDstr, 10, 64)
	pupils := models.Pupils{Id: userIDedit, ClassId: userIDedit}

	fmt.Println(userIDedit, userIDedit)
	userTemp := UserTemp{IsEdit: true, Pupils: pupils}

	pupils.Name = r.FormValue("Name")
	pupils.Surname = r.FormValue("Surname")
	pupils.Patronymic = r.FormValue("Patronymic")
	pupils.Gender = r.FormValue("Gender")
	pupils.Address = r.FormValue("Address")

	fmt.Println(pupils)

	userTemp.Pupils = pupils
	fmt.Println(userTemp.Pupils)

	RenderTempl(w, "templates/pupils.html", userTemp)
}

func RenderTempl(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplName)
	fmt.Println(err)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	body := &bytes.Buffer{}

	err = tmpl.Execute(body, data)
	fmt.Println(err)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	w.Write(body.Bytes())
}

