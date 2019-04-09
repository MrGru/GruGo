package bit_operation

type BitFlag int

const (
	Active  BitFlag = 1 << iota // 1 << 0  == 1
	Send                        // 1 << 1 == 2
	Receive                     // 1 << 2 == 4
)

//flag := Active | Send // == 3
