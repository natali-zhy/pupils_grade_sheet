package main

import (
	"bytes"
	"fmt"
	_ "github.com/gorilla/mux"
	"io"
	"log"
	"strconv"
	_ "strconv"
	"time"

	_ "github.com/natalizhy/pupils_grade_sheet/utils"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/natalizhy/pupils_grade_sheet/db"
	"github.com/natalizhy/pupils_grade_sheet/models"
)

type UserTemp struct {
	AddClasses models.Classes
	AddPupils  models.Pupils
	AddSchools models.Schools
	AddScores  models.Scores
	AddSubject models.Subject
	IsEdit     bool
	Error      map[string]map[string]string
	Classes    []models.Classes
	Pupils     []models.Pupils
	Schools    []models.Schools
	Scores     []models.Scores
	Subject    []models.Subject
}

type Data struct {
	Classes []models.Classes
	Pupils  []models.Pupils
	Schools []models.Schools
	Scores  []models.Scores
	Subject []models.Subject
	IsEdit  bool
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
	//r.Get("/schools/{bookId}", AddSchools)
	//r.Put("/schools/{bookId}", AddSchools)
	//r.Delete("/schools/{bookId}", AddSchools)

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

	fmt.Println(time.Now().Date())

	err := http.ListenAndServe(":3005", r)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSchools(w http.ResponseWriter, r *http.Request) {
	sk, err := db.GetSchools()
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}
	userTemp := UserTemp{IsEdit: true, Schools: sk}

	RenderTempl(w, "templates/schools.html", userTemp)
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	sk, err := db.GetClasses()
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}
	cl := models.Classes{}
	userTemp := UserTemp{IsEdit: true, Classes: sk, AddClasses: cl}

	RenderTempl(w, "templates/classes.html", userTemp)
}

func GetPupils(w http.ResponseWriter, r *http.Request) {
	sk, err := db.GetPupilsById()
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}

	pl := models.Pupils{}
	userTemp := UserTemp{IsEdit: true, Pupils: sk, AddPupils: pl}

	RenderTempl(w, "templates/pupils.html", userTemp)
}

func GetScores(w http.ResponseWriter, r *http.Request) {
	scores, err := db.GetScoresById()
	fmt.Println(scores, err)
	if err != nil {
		fmt.Println(err)
	}

	score := models.Scores{}
	userTemp := UserTemp{IsEdit: true, Scores: scores, AddScores: score}

	RenderTempl(w, "templates/scores.html", userTemp)
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	subjects, err := db.GetSubjectById()
	fmt.Println(subjects, err)
	if err != nil {
		fmt.Println(err)
	}
	userTemp := UserTemp{IsEdit: true, Subject: subjects}

	RenderTempl(w, "templates/subject.html", userTemp)
}

func AddSchools(w http.ResponseWriter, r *http.Request) {

	schools := models.Schools{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	schools.Name = r.FormValue("Name")
	schools.NumberOfClasses, _ = strconv.ParseInt(r.FormValue("NumberOfClasses"), 10, 64)
	schools.CallCenter = r.FormValue("CallCenter")
	schools.Address = r.FormValue("Address")
	schools.CreatedAt.Format("2006-01-02 15:04:05")
	schools.UpdatedAt.Format("2006-01-02 15:04:05")


	sk, err := db.AddSchools(schools)
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}

	sk1, err := db.GetSchools()

	fmt.Println(sk1, err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(schools.NumberOfClasses)
	userTemp := UserTemp{IsEdit: true, AddSchools: schools}

	userTemp.Schools = sk1

	RenderTempl(w, "templates/schools.html", userTemp)
}

func AddClasses(w http.ResponseWriter, r *http.Request) {
	classes := models.Classes{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	classes.Name = r.FormValue("Name")
	classes.NumberOfPupils, _ = strconv.ParseInt(r.FormValue("NumberOfPupils"), 10, 64)

	fmt.Println(classes)

	userTemp := UserTemp{IsEdit: true, AddClasses: classes}

	sk, err := db.AddClasses(userTemp.AddClasses)
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}

	sk1, err := db.GetClasses()
	fmt.Println(sk1, err)
	if err != nil {
		fmt.Println(err)
	}

	userTemp.Classes = sk1

	RenderTempl(w, "templates/classes.html", userTemp)
}

func AddScores(w http.ResponseWriter, r *http.Request) {
	scores := models.Scores{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	scores.Score = r.FormValue("SubjectId")
	scores.Score = r.FormValue("PupilId")
	scores.Score = r.FormValue("Score")

	fmt.Println(scores)

	userTemp := UserTemp{IsEdit: true, AddScores: scores}

	sk, err := db.AddScores(userTemp.AddScores)
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}

	RenderTempl(w, "templates/scores.html", userTemp)
}

func AddSubject(w http.ResponseWriter, r *http.Request) {
	subject := models.Subject{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	subject.Name = r.FormValue("Name")

	fmt.Println(subject)
	userTemp := UserTemp{IsEdit: true, AddSubject: subject}

	sk, err := db.AddSubject(userTemp.AddSubject)
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}

	RenderTempl(w, "templates/subject.html", userTemp)
}

func AddPupils(w http.ResponseWriter, r *http.Request) {
	pupils := models.Pupils{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pupils.Name = r.FormValue("ClassId")
	pupils.Name = r.FormValue("Name")
	pupils.Surname = r.FormValue("Surname")
	pupils.Patronymic = r.FormValue("Patronymic")
	pupils.Gender = r.FormValue("Gender")
	pupils.Address = r.FormValue("Address")

	fmt.Println(pupils)

	userTemp := UserTemp{IsEdit: true, AddPupils: pupils}

	sk, err := db.AddPupils(userTemp.AddPupils)
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}

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
