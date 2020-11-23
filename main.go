package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	_ "strconv"
	"time"

	"github.com/natalizhy/pupils_grade_sheet/utils"

	"html/template"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/natalizhy/pupils_grade_sheet/db"
	"github.com/natalizhy/pupils_grade_sheet/models"
	"github.com/natalizhy/pupils_grade_sheet/session"
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

const (
	COOKIE_NAME = "sessionId"
)

var config = &Application{}
var inMemorySession *session.Session

func main() {

	file, err := os.Open("../pupils_grade_sheet/config/apiserver.json")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB OK")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		fmt.Println(err)
	}

	db.New(config.Db.UserPostgres, config.Db.PasswordPostgres, config.Db.Database)

	//db.InitDB()

	fmt.Println(config.AddrHttp)
	r := chi.NewRouter()
	inMemorySession = session.NewSession()

	r.Get("/", indexHandler)

	r.Get("/login", GetLoginHandler)
	r.Post("/login", PostLoginHandler)

	r.Get("/schools", GetSchools)
	r.Post("/schools", AddSchools)

	r.Get("/schools/{Id}/{mode}", GetSchool) // редактирование
	r.Post("/schools/{Id}/{mode}", AddSchools)

	r.Get("/schools/{Id}/DeleteSchools", DeleteSchools)

	r.Get("/classes", GetClasses)
	r.Post("/classes", AddClasses)

	r.Get("/classes/{ID}/DeleteClasses", DeleteClasses)

	r.Get("/pupils", GetPupils)
	r.Post("/pupils", AddPupils)

	r.Get("/pupils/{ID}/DeletePupils", DeletePupils)

	r.Get("/scores", GetScores)
	r.Post("/scores", AddScores)

	r.Get("/scores/{ID}/DeleteScores", DeleteScores)

	r.Get("/subject", GetSubject)
	r.Post("/subject", AddSubject)

	r.Get("/subject/{ID}/DeleteSubject", DeleteSubject)

	fileHandle := http.FileServer(http.Dir(".")).ServeHTTP

	r.Get("/assets/*", fileHandle)

	fmt.Println("connect :3005")

	fmt.Println(time.Now().Date())

	err = http.ListenAndServe(config.AddrHttp, r)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)

	if cookie != nil {
		fmt.Println("+++++++++++++")
		fmt.Println(inMemorySession.Get(cookie.Value))

	}
	userTemp := models.User{}

	RenderTempl(w, "templates/login.html", userTemp)
}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	userTemp := models.User{}
	RenderTempl(w, "templates/login.html", userTemp)

}
func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	user.Login = r.FormValue("username")
	user.Password = r.FormValue("password")
	userTemp := UserTemp{}

	var err error

	fmt.Println("userTemp", userTemp)
	userTemp.Error, err = utils.ValidateUser(user)
	fmt.Println(userTemp.Error, err)
	if err != nil {
		fmt.Println("не правильно введенные данные")
	}

	if user.Login == "root" && user.Password == "root" {
	fmt.Println(user.Login, user.Password)
		sessionId := inMemorySession.Init(user.Login)

		cookie := &http.Cookie{
			Name:    COOKIE_NAME,
			Value:   sessionId,
			Expires: time.Now().Add(5 * time.Minute),
		}

		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	RenderTempl(w, "templates/login.html", userTemp)

}

func GetSchool(w http.ResponseWriter, r *http.Request) {

	userIDstr := chi.URLParam(r, "Id")
	schools := models.Schools{}
	sk := []models.Schools{}

	userTemp := UserTemp{IsEdit: true, Schools: sk, AddSchools: schools}

	IsEdit := chi.URLParam(r, "mode")
	if IsEdit == "edit" {
		userTemp.IsEdit = true
	} else {
		userTemp.IsEdit = false
	}

	Id, err := strconv.ParseInt(userIDstr, 10, 64)
	fmt.Println(Id, err)
	if err != nil {
		w.Write([]byte("Профиль не найден"))
		return
	}
	schools, err = db.GetSchoolsById(Id)
	if err != nil {
		w.Write([]byte("Не могу выбрать профиль из базы"))
		return
	}

	userTemp.AddSchools = schools

	sk, err = db.GetSchools()
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}
	userTemp.Schools = sk

	//:= UserTemp{IsEdit: true, Schools: sk, AddSchools:schools}
	RenderTempl(w, "templates/schools.html", userTemp)

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

func UpSchools(w http.ResponseWriter, r *http.Request) {

	userIDstr := chi.URLParam(r, "ID")
	id, err := strconv.ParseInt(userIDstr, 10, 64)
	if err != nil {
		w.Write([]byte("Профиль не найден"))
		return
	}

	schools := models.Schools{
		UpdatedAt: time.Now(),
	}

	schools.Name = r.FormValue("Name")
	schools.NumberOfClasses, _ = strconv.ParseInt(r.FormValue("NumberOfClasses"), 10, 64)
	schools.CallCenter = r.FormValue("CallCenter")
	schools.Address = r.FormValue("Address")
	schools.UpdatedAt.Format("2006-01-02 15:04:05")

	fmt.Println(schools)

	err = db.UpSchools(id, schools)
	fmt.Println(err)
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

func GetClasses(w http.ResponseWriter, r *http.Request) {
	sk, err := db.GetClasses()
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}
	cl := models.Classes{}
	userTemp := UserTemp{Classes: sk, AddClasses: cl}

	RenderTempl(w, "templates/classes.html", userTemp)
}

func GetPupils(w http.ResponseWriter, r *http.Request) {
	sk, err := db.GetPupilsById()
	fmt.Println(sk, err)
	if err != nil {
		fmt.Println(err)
	}

	pl := models.Pupils{}
	userTemp := UserTemp{Pupils: sk, AddPupils: pl}

	RenderTempl(w, "templates/pupils.html", userTemp)
}

func GetScores(w http.ResponseWriter, r *http.Request) {
	scores, err := db.GetScoresById()
	fmt.Println(scores, err)
	if err != nil {
		fmt.Println(err)
	}

	score := models.Scores{}
	userTemp := UserTemp{Scores: scores, AddScores: score}

	RenderTempl(w, "templates/scores.html", userTemp)
}

func GetSubject(w http.ResponseWriter, r *http.Request) {
	subjects, err := db.GetSubjectById()
	fmt.Println(subjects, err)
	if err != nil {
		fmt.Println(err)
	}
	userTemp := UserTemp{Subject: subjects}

	RenderTempl(w, "templates/subject.html", userTemp)
}

func AddSchools(w http.ResponseWriter, r *http.Request) {

	userIDstr := chi.URLParam(r, "Id")
	fmt.Println("userIDstr", userIDstr)
	//if userIDstr != ""{
	Id, _ := strconv.ParseInt(userIDstr, 10, 64)

	fmt.Println("Id", Id)
	//	fmt.Println(Id, err)
	//	if err != nil {
	//		w.Write([]byte("Профиль не найден"))
	//		return
	//	}
	//
	//}

	schools := models.Schools{
		Id:        Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userTemp := UserTemp{IsEdit: true, AddSchools: schools}

	schools.Name = r.FormValue("Name")
	schools.NumberOfClasses, _ = strconv.ParseInt(r.FormValue("NumberOfClasses"), 10, 64)
	schools.CallCenter = r.FormValue("CallCenter")
	schools.Address = r.FormValue("Address")
	schools.CreatedAt.Format("2006-01-02 15:04:05")
	schools.UpdatedAt.Format("2006-01-02 15:04:05")

	userTemp.AddSchools = schools

	fmt.Println("userTemp.Error", userTemp.Error)
	if len(userTemp.Error) == 0 {

		if userIDstr != "" {
			id, err := strconv.ParseInt(userIDstr, 10, 64)
			if err != nil {
				w.Write([]byte("Неправельный ID"))
				return
			}
			fmt.Println(strconv.FormatInt(id, 10))

			err = db.UpSchools(id, schools)
			fmt.Println("+++++++++++++++++++++++++", err, id, Id)
			if err != nil {
				fmt.Println(err)
				w.Write([]byte("Юзер не добавлен"))
				return
			}

		} else {
			sk, err := db.AddSchools(schools)
			fmt.Println(sk, err)
			if err != nil {
				fmt.Println(err)
			}

		}

		sk1, err := db.GetSchools()

		fmt.Println(sk1, err)
		if err != nil {
			fmt.Println(err)
		}
		userTemp.Schools = sk1

		http.Redirect(w, r, "/schools/"+strconv.FormatInt(Id, 10)+"/edit", http.StatusSeeOther)
		return
	}
	//

	//schools := models.Schools{
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//}
	//
	//schools.Name = r.FormValue("Name")
	//schools.NumberOfClasses, _ = strconv.ParseInt(r.FormValue("NumberOfClasses"), 10, 64)
	//schools.CallCenter = r.FormValue("CallCenter")
	//schools.Address = r.FormValue("Address")
	//schools.CreatedAt.Format("2006-01-02 15:04:05")
	//schools.UpdatedAt.Format("2006-01-02 15:04:05")
	//
	//
	//sk, err := db.AddSchools(schools)
	//fmt.Println(sk, err)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//sk1, err := db.GetSchools()
	//
	//fmt.Println(sk1, err)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(schools.NumberOfClasses)
	//userTemp := UserTemp{IsEdit: true, AddSchools: schools}
	//
	//userTemp.Schools = sk1

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

func DeleteSchools(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "ID")

	schools := models.Schools{}

	userID, err := strconv.ParseInt(userIDstr, 10, 64)

	err = db.DeleteSchools(schools, userID)

	if err != nil {
		http.Redirect(w, r, "/schools", http.StatusTemporaryRedirect)
		return
	}
}
func DeleteClasses(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "ID")

	classes := models.Classes{}

	userID, err := strconv.ParseInt(userIDstr, 10, 64)

	err = db.DeleteClasses(classes, userID)

	if err != nil {
		http.Redirect(w, r, "/classes", http.StatusTemporaryRedirect)
		return
	}
}
func DeletePupils(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")

	user := models.Schools{}

	userID, err := strconv.ParseInt(userIDstr, 10, 64)

	err = db.DeleteSchools(user, userID)

	if err != nil {
		http.Redirect(w, r, "/schools", http.StatusTemporaryRedirect)
		return
	}
}
func DeleteScores(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")

	user := models.Schools{}

	userID, err := strconv.ParseInt(userIDstr, 10, 64)

	err = db.DeleteSchools(user, userID)

	if err != nil {
		http.Redirect(w, r, "/schools", http.StatusTemporaryRedirect)
		return
	}
}
func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")

	user := models.Schools{}

	userID, err := strconv.ParseInt(userIDstr, 10, 64)

	err = db.DeleteSchools(user, userID)

	if err != nil {
		http.Redirect(w, r, "/schools", http.StatusTemporaryRedirect)
		return
	}
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
