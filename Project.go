package main

import (
	/* 	"fmt" */
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "rootroot"
	dbName := "university"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Student struct {
	Ssn      int
	S_name   string
	Gender   string
	Adress   string
	Religion string
	Faculty  string
	Gpa      float32
	S_Level  int
	Phone    string
	G_ssn    int
	G_name   string
	Job      string
	G_Phone  string
}

var tmpl = template.Must(template.ParseGlob("projectweb/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("select * from guardian inner join student where student.Ssn=guardian.S_ssn ")
	if err != nil {
		panic(err.Error())
	}
	emp := Student{}
	res := []Student{}
	for selDB.Next() {
		var ssn, s_Level, G_ssn, S_ssn int
		var s_name, gender, adres, religion, faculty, Phone, G_name, G_job, G_phone string
		var gpa float32

		err = selDB.Scan(&G_ssn, &G_name, &G_job, &S_ssn, &G_phone, &ssn, &s_name, &gender, &adres, &religion, &faculty, &gpa, &s_Level, &Phone)
		if err != nil {
			panic(err.Error())
		}

		emp.Ssn = ssn
		emp.S_name = s_name
		emp.Gender = gender
		emp.Adress = adres
		emp.Religion = religion
		emp.Faculty = faculty
		emp.Gpa = gpa
		emp.S_Level = s_Level
		emp.Phone = Phone
		emp.G_ssn = G_ssn
		emp.G_name = G_name
		emp.Job = G_job
		emp.G_Phone = G_phone

		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "register", nil)
}
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Home", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {

		Ssn := r.FormValue("nation-id")
		S_name := r.FormValue("name")
		Gender := r.FormValue("gen")
		adress := r.FormValue("address")
		Religion := r.FormValue("rel")
		Faculty := r.FormValue("faculties")
		Gpa := r.FormValue("gpa")
		S_Level := r.FormValue("level")
		phone := r.FormValue("telephone")
		insForm, err := db.Prepare("INSERT INTO Student VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(Ssn, S_name, Gender, adress, Religion, Faculty, Gpa, S_Level, phone)
		/* log.Println("INSERT: Name: " + Ssn + " | City: " + city) */

		G_ssn := r.FormValue("NID")
		G_name := r.FormValue("F_name")
		Job := r.FormValue("job")
		S_ssn := r.FormValue("nation-id")
		Gphone := r.FormValue("F_phone")

		insForm1, err := db.Prepare("INSERT INTO Guardian VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm1.Exec(G_ssn, G_name, Job, S_ssn, Gphone)
	}
	defer db.Close()
	http.Redirect(w, r, "/new", 301)
}
func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs)) 
	


	/* Create Tables */
	db := dbConn()
	CreateStudent, err := db.Query(`create table if not exists  Student ( Ssn varchar(14) primary key,
		S_name varchar(60),
		Gender varchar(10),
		adress varchar(40),
		Religion varchar(10),
		Faculty varchar(60),
		Gpa float,
		S_Level int,
		phone varchar(11) )`)

	if err != nil {
		panic(err.Error())
	}
	defer CreateStudent.Close()
	CreateGuardian, err := db.Query(`create table if not exists Guardian ( G_ssn varchar(14) primary key,
		G_name varchar(60),
		Job varchar(60),
		S_ssn varchar(14)  ,
		phone varchar(11),
		FOREIGN KEY (S_ssn) REFERENCES Student(Ssn) ON DELETE CASCADE ON UPDATE CASCADE
		)`)

	if err != nil {
		panic(err.Error())
	}
	defer CreateGuardian.Close()

	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Home)
	http.HandleFunc("/new", New)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/show", Index)
	http.ListenAndServe(":8080", nil)

}
