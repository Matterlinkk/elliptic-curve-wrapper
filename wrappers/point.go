package wrappers

import (
	"fmt"
	"log"
	"math/big"
	"strings"
)

type Ecpoint struct {
	x *big.Int
	y *big.Int
}

func NewECPoint(x, y *big.Int) *Ecpoint {
	return &Ecpoint{
		x: x,
		y: y,
	}
}

func (ecp *Ecpoint) Params() (*big.Int, *big.Int) {
	return ecp.x, ecp.y
}

func (ecp *Ecpoint) Print(base int) {
	fmt.Printf("X:%s\nY:%s\n", ecp.x.Text(base), ecp.y.Text(base))
}

func (ecp *Ecpoint) IsEqualTo(otherPoint *Ecpoint) bool {
	x1, y1 := ecp.Params()
	x2, y2 := otherPoint.Params()
	return x1.Cmp(x2) == 0 && y1.Cmp(y2) == 0
}

func ECPointToString(point *Ecpoint, base int) string {
	x, y := point.Params()
	return x.Text(base) + "," + y.Text(base)
}

func StringToECPoint(str string, base int) *Ecpoint {
	params := strings.Split(str, ",")
	if len(params) != 2 {
		log.Panicf("len more or less then 2, len is: %d", len(params))
	}
	var x, y *big.Int
	x, ok := x.SetString(params[0], base)
	if !ok {
		log.Panicf("Invalid x-value")
	}
	y, ok = y.SetString(params[0], base)
	if !ok {
		log.Panicf("Invalid y-value")
	}

	return NewECPoint(x, y)
}
