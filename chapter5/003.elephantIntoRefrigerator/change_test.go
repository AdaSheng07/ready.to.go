package main

import "testing"

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age  int
}

func (s Student) ChangeName(newName string) {
	s.Name = newName
}

func (s *Student) ChangeAge(newAge int) {
	s.Age = newAge
}

func (s Student) ChangeNameObject() {
	s.Name = "Ian"
}

func (s *Student) ChangeNamePointer() {
	s.Name = "Mandy"
}

//func TestName(t *testing.T) {
//	s := Student{
//		Name: "Roy",
//	}
//	s.ChangeNameObject()
//	fmt.Println("change name object:", s.Name)
//	s.ChangeNamePointer()
//	fmt.Println("change name pointer:", s.Name)
//}

func TestName(t *testing.T) {
	//var c1 Change = Student{
	//	Name: "Roy",
	//	Age:  34,
	//}
	// Cannot use 'Student{ Name: "Roy", Age: 34, }' (type Student) as the type Change Type does not implement 'Change' as the 'ChangeAge' method has a pointer receiver

	var c2 Change = &Student{
		Name: "Roy",
		Age:  34,
	}
	c2.ChangeName("Tommy")
	c2.ChangeAge(23)
}
