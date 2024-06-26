package main

type Human struct {
	Name string
	Age  int
}

func (h *Human) GetName() string {
	return h.Name
}

type Student struct {
	Human
	University string
	Course     int
	Group      int
}

func main() {
	student := Student{
		Human: Human{
			Name: "Vadim",
			Age:  22,
		},
		University: "SSU",
		Course:     4,
		Group:      411,
	}

	println(student.GetName())
}
