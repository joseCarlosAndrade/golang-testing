package foo

import "strconv"


func Fooer(input int) string {
	if div := input % 3; div == 0 {
		return "Foo"
	}

	return strconv.Itoa(input)
}