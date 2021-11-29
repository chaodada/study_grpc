package main

import "demo/pb/person"

func main() {


	var h person.Homes
	one := h.TestOneOf.(*person.Homes_One)
	one.One = "123"


}
