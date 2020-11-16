package db

import (
	"database/sql"
	"fmt"
	"github.com/natalizhy/pupils_grade_sheet/models"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func AddSchools(schools models.Schools) (newUserId int64, err error) {
	fmt.Println(schools.Name)

	_, err = DB.Exec("INSERT INTO Schools (Name, NumberOfClasses, CallCenter, Address, CreatedAt, UpdatedAt) "+
		"VALUES ($1, $2, $3, $4, $5, $6)", schools.Name, schools.NumberOfClasses, schools.CallCenter, schools.Address, schools.CreatedAt.Format("2006-01-02 15:04:05"), schools.UpdatedAt.Format("2006-01-02 15:04:05"))
	if err != nil {
		return
	}

	fmt.Println(schools.Name)

	fmt.Println("Created user with id:", newUserId)
	return
}

func AddClasses(classes models.Classes) (newUserId int64, err error) {
	fmt.Println(classes.Name)
	fmt.Println(time.Now().Date())

	_, err = DB.Exec("INSERT INTO Classes (SchoolId, Name, NumberOfPupils, CreatedAt, UpdatedAt) "+
		"VALUES ($1, $2, $3, $4, $5)", 1, classes.Name, classes.NumberOfPupils, classes.CreatedAt.Format("2006-01-02 15:04:05"), classes.UpdatedAt.Format("2006-01-02 15:04:05"))
	if err != nil {
		return
	}

	fmt.Println(classes.Name)

	fmt.Println("Created user with id:", newUserId)
	return
}

func AddPupils(pupils models.Pupils) (newUserId int64, err error) {
	fmt.Println(pupils.Name)

	_, err = DB.Exec("INSERT INTO Pupils (ClassId, Name, Surname, Patronymic, Gender, Address, CreatedAt, UpdatedAt) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", 1, pupils.Surname, pupils.Patronymic, pupils.Gender, pupils.Address, pupils.CreatedAt.Format("2006-01-02 15:04:05"), pupils.UpdatedAt.Format("2006-01-02 15:04:05"), time.Now())
	if err != nil {
		return
	}

	fmt.Println(pupils.Name)

	fmt.Println("Created user with id:", newUserId)
	return
}

func AddScores(scores models.Scores) (newUserId int64, err error) {
	fmt.Println(scores.Score)

	_, err = DB.Exec("INSERT INTO Schools (SubjectId, PupilId, Score, CreatedAt, UpdatedAt) "+
		"VALUES ($1, $2, $3, $4, $5)", 1, 1, scores.Score, scores.CreatedAt.Format("2006-01-02 15:04:05"), scores.UpdatedAt.Format("2006-01-02 15:04:05"), time.Now())
	if err != nil {
		return
	}

	fmt.Println(scores.Score)

	fmt.Println("Created user with id:", newUserId)
	return
}
func AddSubject(subject models.Subject) (newUserId int64, err error) {
	fmt.Println(subject.Name)

	_, err = DB.Exec("INSERT INTO Subject (Name, CreatedAt, UpdatedAt) "+
		"VALUES ($1, $2, $3)", subject.Name, subject.CreatedAt.Format("2006-01-02 15:04:05"), subject.UpdatedAt.Format("2006-01-02 15:04:05"), time.Now())
	if err != nil {
		return
	}

	fmt.Println(subject.Name)

	fmt.Println("Created user with id:", newUserId)
	return
}

func GetSchools() (schools []models.Schools, err error) {
	rows, err := DB.Query("SELECT Id, Name, NumberOfClasses, CallCenter, Address, CreatedAt, UpdatedAt " +
		"FROM schools")
	if err != nil {
		return
	}

	for rows.Next() {
		school := models.Schools{}


		err = rows.Scan(&school.Id, &school.Name, &school.NumberOfClasses, &school.CallCenter, &school.Address, &school.CreatedAt, &school.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		school.CreatedAt.Format("2006-01-02 15:04:05")
		school.UpdatedAt.Format("2006-01-02 15:04:05")
		schools = append(schools, school)
	}
	return
}


func GetClasses() (schools []models.Classes, err error) {
	rows, err := DB.Query("SELECT Id, SchoolId, Name, NumberOfPupils, CreatedAt, UpdatedAt " +
		"FROM classes")
	if err != nil {
		return
	}

	for rows.Next() {
		school := models.Classes{}
		err = rows.Scan(&school.Id, &school.SchoolId, &school.Name, &school.NumberOfPupils, &school.CreatedAt, &school.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		schools = append(schools, school)
	}
	return
}



func GetSchoolsById(id int64) (schools models.Schools, err error) {

	_, err = DB.Query("SELECT Name, NumberOfClasses, CallCenter, Address, CreatedAt, UpdatedAt " +
		"FROM schools WHERE Id=$1", id)
	if err != nil {
		return
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return schools, nil
	case nil:
		return schools, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return schools, err
}

func GetClassesById() (classes []models.Classes, err error) {
	rows, err := DB.Query("SELECT SchoolId, Name, NumberOfPupils, CreatedAt, UpdatedAt " +
		"FROM classes WHERE Id=?")
	if err != nil {
		return
	}

	for rows.Next() {
		class := models.Classes{}
		err = rows.Scan(&class.Id, &class.SchoolId, &class.Name, &class.NumberOfPupils, &class.CreatedAt, &class.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		classes = append(classes, class)
	}
	return
}

func GetPupilsById() (pupils []models.Pupils, err error) {
	rows, err := DB.Query("SELECT Id, ClassId, Name, Surname, Patronymic, Gender, Address, CreatedAt, UpdatedAt " +
		"FROM pupils")
	if err != nil {
		return
	}

	for rows.Next() {
		pupil := models.Pupils{}
		err = rows.Scan(&pupil.Id, &pupil.ClassId, &pupil.Name, &pupil.Surname, &pupil.Patronymic, &pupil.Gender, &pupil.Address, &pupil.CreatedAt, &pupil.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		pupils = append(pupils, pupil)
	}
	return
}
func GetScoresById() (scores []models.Scores, err error) {
	rows, err := DB.Query("SELECT Id, SubjectId, PupilId, Score, CreatedAt, UpdatedAt " +
		"FROM scores")
	if err != nil {
		return
	}

	for rows.Next() {
		score := models.Scores{}
		err = rows.Scan(&score.Id, &score.SubjectId, &score.PupilId, &score.Score, &score.CreatedAt, &score.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		scores = append(scores, score)
	}
	return
}

func GetSubjectById() (subject []models.Subject, err error) {
	rows, err := DB.Query("SELECT Id, Name, CreatedAt, UpdatedAt " +
		"FROM subject")
	if err != nil {
		return
	}

	for rows.Next() {
		sub := models.Subject{}
		err = rows.Scan(&sub.Id, &sub.Name, &sub.CreatedAt, &sub.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		subject = append(subject, sub)
	}
	return
}


func UpSchools(ID int64, schools models.Schools) (err error) {
	fmt.Println(schools.Name)

	_, err = DB.Exec("UPDATE Schools SET (Name, NumberOfClasses, CallCenter, Address, UpdatedAt) "+
		"VALUES ($1, $2, $3, $4, $5) WHERE Id=$6", schools.Name, schools.NumberOfClasses, schools.CallCenter, schools.Address, schools.UpdatedAt.Format("2006-01-02 15:04:05"), ID)
	if err != nil {
		return
	}

	fmt.Println(schools.Name)

	fmt.Println("Created user with id:", ID)
	return
}


func DeleteSchools(schools models.Schools, ID int64) (err error) {
	result := DB.QueryRow("DELETE FROM schools " +
		"WHERE id=$1", ID)

	err = result.Scan(&schools)

	fmt.Println(&schools)
	fmt.Println("Delete schools")

	return
}
func DeleteClasses(classes models.Classes, ID int64) (err error) {
	result := DB.QueryRow("DELETE FROM classes " +
		"WHERE id=$1", ID)

	err = result.Scan(&classes)

	fmt.Println(&classes)
	fmt.Println("Delete classes")

	return
}