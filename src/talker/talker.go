package talker

import "strconv"

type Talker struct {
	Name string
	Age  int
}

func (t Talker) Talk() string {
	return "Name: " + t.Name + " Age: " + strconv.Itoa(t.Age)
}
