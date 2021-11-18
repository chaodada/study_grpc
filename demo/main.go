package main

import "demo/pb/person"

func main() {


	var h person.Home
	one := h.TestOneOf.(*person.Home_One)
	one.One = "123"


}
