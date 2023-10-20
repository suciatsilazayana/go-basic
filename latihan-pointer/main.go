package main

import "fmt"

type Student struct {
	Name string
}

func (s *Student) setMyName(newname string) {
	s.Name = newname
}

func (s *Student) callMyName() {
	fmt.Println("Hello, my name is", s.Name)
}

func main() {
	student := Student{
		Name: "Suci",
	}

	//memanggil method setmyname
	student.setMyName("Noobee")

	//memanggil method callmyname
	student.callMyName()
}
