package bit_operation

type ByteSize float64

const (
	_           = iota // ignore first value by assignng to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
