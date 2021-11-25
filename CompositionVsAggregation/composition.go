package main

type Lamp struct {}

type Room struct {
	Lamps *[]Lamp
}


func NewRoom(l *[]Lamp) *Room {
  return &Room{l}
}

type House1 struct {
	Room *Room
}

func NewHouse1(l *[]Lamp) *House1 {
  r := NewRoom(l)
  return &House1{r}
}

type House2 struct {
	Room *Room
}

func NewHouse2(l *[]Lamp) *House2 {
  r := NewRoom(l)
  return &House2{r}
}

func main(){
  lamps := []Lamp{}
  house1 := NewHouse1(&lamps)
  house2 := NewHouse2(&lamps)
}
