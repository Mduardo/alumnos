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
	var alumno []person
	var l person
	var inpt string
	for i := 0; i < 2; i++ {
		fmt.Println("give me the student's name,id, age")
		fmt.Scanln(&inpt)
		l.studnt.name = inpt
		fmt.Scanln(&inpt)
		l.studnt.id = inpt
		fmt.Scanln(&inpt)
		l.studnt.age, _ = strconv.Atoi(inpt)

		fmt.Println("give me subject's name, note")
		fmt.Scanln(&inpt)
		l.subj.name = inpt
		fmt.Scanln(&inpt)
		l.subj.note, _ = strconv.Atoi(inpt)

		fmt.Println("give me the teacher's name , id, age")
		fmt.Scanln(&inpt)
		l.teach.name = inpt
		fmt.Scanln(&inpt)
		l.teach.id = inpt
		fmt.Scanln(&inpt)
		l.teach.age, _ = strconv.Atoi(inpt)

		alumno = append(alumno, l)

	}
	aver := average(alumno)
	fmt.Println("the avarege of notes is", aver)
	for i := 0; i < len(alumno); i++ {
		r := General(alumno[i], "id")
		fmt.Println("the ids are...", r)
	}

}

func average(source []person) int {
	aver := 0
	numitmes := 0
	for i, content := range source {
		aver += content.subj.note
		numitmes = i + 1
	}
	return aver / numitmes
}
