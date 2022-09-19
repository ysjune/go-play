package main

import (
	"fmt"
	"go-play/ch20/fedex"
	"go-play/ch20/korea"
)

type DuckInterface interface {
	Fly()
	Walk(distance int) int
}

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("hi %d %s", s.Age, s.Name)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("age: %d\n", s.Age)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

type Sender interface {
	Send(parcel string)
}

type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type Attacker interface {
	Attack()
}

type Closer interface {
	Close()
}

func ReadFile(reader Reader) {
	//c, ok := reader.(Closer)
	//if ok {
	//	c.Close()
	//}

	if c, ok := reader.(Closer); ok {
		c.Close()
	}
}

func main() {

	student := Student{Name: "Su", Age: 12}
	var stringer Stringer

	stringer = student

	fmt.Printf("%s\n", stringer.String())

	sender := &fedex.FedexSender{}
	SendBook("prince", sender)
	SendBook("그리스인 조르바", sender)

	sender2 := &korea.KoreaSender{}
	SendBook("prince2", sender2)
	SendBook("그리스인 조르바2", sender2)

	s := &Student{"su", 15}
	PrintAge(s)
}
