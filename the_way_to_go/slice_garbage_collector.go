package the_way_to_go

import (
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

//garbage collector not work. Can't release the array
//func FindDigits(filename string) []byte {
//	b, _ := ioutil.ReadFile(filename)
//	return digitRegexp.Find(b)
//}

//func FindDigits(filename string) []byte {
//	b, _ := ioutil.ReadFile(filename)
//	b = digitRegexp.Find(b)
//	c := make([]byte, len(b))
//	copy(c, b)
//	return c
//}
