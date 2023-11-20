package wrappers

import (
	"crypto/elliptic"
	"math/big"
)

type Ecwrapper struct {
	curve elliptic.Curve
}

func NewECWrapper(curve elliptic.Curve) *Ecwrapper {
	return &Ecwrapper{
		curve: curve,
	}
}

func (ec *Ecwrapper) Params() *elliptic.CurveParams {
	return ec.curve.Params()
}

func (ec *Ecwrapper) GetPointG() *Ecpoint {
	params := ec.Params()
	return NewECPoint(params.Gx, params.Gy)
}

func (ec *Ecwrapper) IsOnCurve(point *Ecpoint) bool {
	params := ec.Params()
	return params.IsOnCurve(point.Params())
}

func (ec *Ecwrapper) Add(point1, point2 *Ecpoint) *Ecpoint {
	params := ec.Params()
	x1, y1 := point1.Params()
	x2, y2 := point2.Params()
	return NewECPoint(params.Add(x1, y1, x2, y2))
}

func (ec *Ecwrapper) Double(point *Ecpoint) *Ecpoint {
	params := ec.Params()
	return NewECPoint(params.Double(point.Params()))
}

func (ec *Ecwrapper) ScalarMult(k *big.Int, point *Ecpoint) *Ecpoint {
	params := ec.Params()
	x, y := point.Params()
	return NewECPoint(params.ScalarMult(x, y, k.Bytes()))
}
