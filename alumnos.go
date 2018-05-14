package main

import (
	"fmt"
	"strconv"
)

type user interface {
	getName() [3]string
	getId() [3]string
}

type teachers struct {
	name string
	id   string
	age  int
}

type subjects struct {
	note int
	name string
	id   string
}

type alumnos struct {
	name string
	age  int
	id   string
}

type person struct {
	studnt alumnos
	teach  teachers
	subj   subjects
}

func (first person) getName() [3]string {
	var sl [3]string
	sl[0] = first.studnt.name
	sl[1] = first.subj.name
	sl[2] = first.teach.name
	return sl
}

func (first person) getId() [3]string {
	var sl [3]string
	sl[0] = first.studnt.id
	sl[1] = first.subj.id
	sl[2] = first.teach.id
	return sl
}

func General(s user, weneed string) [3]string {
	var sj [3]string
	switch weneed {
	default:
		{
			fmt.Println("the kind of date isn't valid")
			return sj
		}
	case "name":
		{
			sj = s.getName()
			return sj
		}
	case "id":
		{
			sj = s.getId()
			return sj
		}
	}
	return sj
}
func main() {
	var alumno person
	var inpt string

	fmt.Println("ingrese nombre del alumno id, edad")
	fmt.Scanln(&inpt)
	alumno.studnt.name = inpt
	fmt.Scanln(&inpt)
	alumno.studnt.id = inpt
	fmt.Scanln(&inpt)
	alumno.studnt.age, _ = strconv.Atoi(inpt)

	fmt.Println("ingrese nombre del maestro, id, edad")
	fmt.Scanln(&inpt)
	alumno.teach.name = inpt
	fmt.Scanln(&inpt)
	alumno.teach.id = inpt
	fmt.Scanln(&inpt)
	alumno.teach.age, _ = strconv.Atoi(inpt)

	fmt.Println("ingrese nombre del curso, id, nota")
	fmt.Scanln(&inpt)
	alumno.subj.name = inpt
	fmt.Scanln(&inpt)
	alumno.subj.id = inpt
	fmt.Scanln(&inpt)
	alumno.subj.note, _ = strconv.Atoi(inpt)

	sj := General(alumno, "id")

	fmt.Println("los identificadores son", sj)
}
