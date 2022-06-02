create database university;
use university;
create table Student(
Ssn varchar(14) primary key,
S_name varchar(60),
pass varchar(20),
Gender varchar(10),
adress varchar(40),
Religion varchar(10),
Governorate varchar(30),
E_mail varchar(30),
State varchar(3),#new or old
Faculty varchar(60),
Gpa float,
S_Level int,
phone varchar(11)
);

create table Guardian(
G_ssn varchar(14) primary key,
G_name varchar(60),
Job varchar(60),
S_ssn varchar(14)  ,
phone varchar(11),
FOREIGN KEY (S_ssn) REFERENCES Student(Ssn) ON DELETE CASCADE ON UPDATE CASCADE
);




insert into Student values("98745216987415","aya","asd","female",Date'2000-5-5',"street15","muslim","Sohag","email@gamil.com","new","FCI",2.0,1,"01100000000",true);
insert into Guardian values("15151841315123","Mustafa","writer","98745216987415","01234569000");


