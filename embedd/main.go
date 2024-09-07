package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type deyplay struct {
	container
	base
	con bool
}

func main() {

	try := deyplay{
		container: container{
			base: base{
				num: 1,
			},
			str: "number one",
		},

		base: base{
			num: 2,
		},

		con: false,
	}

	fmt.Printf("try={num: %v, str: %v, con:%v}\n", try.base.num, try.str, try.con)

	fmt.Println("also num:", try.base.num)

	fmt.Println(try.container.base.num)

	fmt.Println("describe", try.describe())

	type describer interface {
		describe() string
	}

	var d describer = try

	fmt.Println("describer interface:", d.describe())
}
