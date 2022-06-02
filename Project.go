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
	Ssn   int
	S_name string
	Gender string
	adress string 
	Religion string
	Faculty string
	Gpa float32
	S_Level int
	phone string

}
type Guardian struct {
	G_ssn string
	G_name string
	Job string
	S_ssn string
	phone string
}
var tmpl = template.Must(template.ParseGlob("projectweb/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Student ")
    if err != nil {
        panic(err.Error())
    }
	selDB1 ,err:=db.Query("SELECT * FROM Guardian ")
    if err != nil {
        panic(err.Error())
    }
    emp := Student{}
    res := []Student{}
	emp1:=Guardian{}
	res1:=[]Guardian{}
    for selDB.Next() {
	    var	ssn  ,s_Level int
	    var s_name ,gender ,adres,religion ,faculty ,Phone string
	    var gpa float32
		
        err = selDB.Scan(&ssn, &s_name,&gender,&adres,&religion,&faculty,&gpa, &s_Level,&Phone)
        if err != nil {
            panic(err.Error())
        }
		
        emp.Ssn = ssn
        emp.S_Level = s_Level
        emp.S_name = s_name
		emp.Gender=gender
		emp.adress=adres 
		emp.Religion=religion
		emp.Faculty=faculty
		emp.phone=Phone
	
        res = append(res, emp)
    }
	for selDB1.Next() {
		var G_ssn ,G_name ,Job ,S_ssn ,phone string
		err = selDB1.Scan(&G_ssn, &G_name,&Job, &S_ssn,&phone)
				if err != nil {
					panic(err.Error())
				}
			emp1.G_ssn = G_ssn
				emp1.G_name = G_name
				emp1.Job=Job
				emp1.S_ssn=S_ssn
				emp1.phone=phone
				
				res1 = append(res1, emp1)
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

	Ssn :=  r.FormValue("nation-id")
	S_name :=  r.FormValue("name")
	Gender :=  r.FormValue("gen")
	 adress :=  r.FormValue("address") 
	Religion :=  r.FormValue("rel")
	Faculty :=  r.FormValue("faculties")
	Gpa :=  r.FormValue("gpa")
	S_Level :=  r.FormValue("level")
	phone:= r.FormValue("telephone")
        insForm, err := db.Prepare("INSERT INTO Student VALUES(?,?,?,?,?,?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(Ssn, S_name,Gender,adress,Religion,Faculty,Gpa,S_Level,phone)
        /* log.Println("INSERT: Name: " + Ssn + " | City: " + city) */

		G_ssn :=r.FormValue("NID")
		G_name :=r.FormValue("F_name")
		Job :=r.FormValue("job")
		S_ssn :=  r.FormValue("nation-id")
		Gphone :=r.FormValue("F_phone")	

		insForm1, err := db.Prepare("INSERT INTO Guardian VALUES(?,?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm1.Exec(G_ssn, G_name,Job,S_ssn,Gphone)
    }
    defer db.Close()
    http.Redirect(w, r, "/new", 301)
}
func main() {	
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