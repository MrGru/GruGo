package pitfall

//func main() {
//	var remember bool = false
//	if true {
//		remember := true
//		fmt.Println(remember)
//	}
//	fmt.Println(remember)
//}
//
//func shadow() (err error) {
//	x, err := check1() // x is created; err is assigned to
//	if err != nil {
//		return // err correctly returned
//	}
//	if y, err := check2(x); err != nil {
//		return // inner err shadows outer err so nil is wrongly returned!
//	} else {
//		fmt.Println(y)
//	}
//	return
//}
