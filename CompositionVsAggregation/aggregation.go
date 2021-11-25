package main

type LibraryCard struct {
}

type Student struct {
   LibCard *LibraryCard
}

func NewStudent(l *LibraryCard) *Student {
  return &Student{l}
}

type Classroom1 struct {
	Student *Student
}

func NewClassroom1(s *Student) *Classroom1 {
  return &Classroom1{s}
}

type Classroom2 struct {
	Student *Student
}

func NewClassroom2(s *Student) *Classroom2 {
  return &Classroom2{s}
}

func main(){
  lc := new(LibraryCard)
  student := NewStudent(lc)
  classroom1 := NewClassroom1(student)
  classroom2 := NewClassroom2(student)
}
