package service

// ArithmeticDS ...
type ArithmeticDS struct {
	A  int64
	B  int64
	OP byte
}

// RandomOP ...
func RandomOP(p int) byte {
	switch p {
	case 0:
		return '*'
	case 1:
		return '/'
	case 2:
		return '+'
	case 3:
		return '-'
	default:
		return '.'
	}
}

// GetResult ...
func GetResult(req *ArithmeticDS) int64 {
	switch req.OP {
	case '*':
		return req.A * req.B
	case '/':
		return req.A / req.B
	case '+':
		return req.A + req.B
	case '-':
		return req.A - req.B
	default:
		return -9999999999999
	}
}
