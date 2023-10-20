package main

func main() {

	// Pointer pada method
	type Student struct {
		Name  string
		Class string
	}

	func (s *Student) setMyName(newName string) {
		fmt.Println("Try to change from", s.Name, "name to =>", newName)
		s.Name = newName
	}
}
