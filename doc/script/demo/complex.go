package demo

import (
	"fmt"
)

type Stest struct {
	Name string
	Id   int
}

type Squote struct {
	id  int
	val string
	p   *Squote
}

func Hander() {

}

func Gquote() {
	s1 := &Squote{
		id:  1,
		val: "爷爷",
		p: &Squote{
			id:  2,
			val: "爸爸",
			p: &Squote{
				id:  2,
				val: "儿子",
			},
		},
	}
	s2 := new(Stest)
	s3 := &Stest{}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(&s1)
	fmt.Println(*s2)
}
