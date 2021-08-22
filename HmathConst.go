package hmath

import (
	"strconv"
)
type ConstT struct{
	value float64
}
func(c ConstT)String()(string){
	return strconv.FormatFloat(float64(c.value), 'g', -1, 64)
}
func(c ConstT)Get(mc MathContext)(float64, error){
	return float64(c.value), nil
}
func(c ConstT)NeededValues()(MathComponents, error){
	return nil, nil
}
func(c ConstT)ErrorTest()error{
	return nil
}
func Constant(x float64)MathFunc{
	return ConstT{value:x}
}
