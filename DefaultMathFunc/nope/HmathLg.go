package hmath

import(
	"math"
)

type LgT struct{
	mathBodySingl
}

func(lg LgT)String()(ret string){
	_, body := lg.GetBody()
	ret += "Lg("
	ret += body[0].String()
	ret += ")"
	return 
}

func(lg LgT)Get(mc MathContext)(float64, error){
	_, body := lg.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Log10(f), nil
}

func(lg LgT)NeededValues()(ret MathComponents){
	_, body := lg.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	
	return
}

func Lg(a MathFunc)MathFunc{
	return LgT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Lg",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Lg dy <2 object count for lg")
					}
					return Lg(input[0]), nil
				},
			),
		),
	)
}