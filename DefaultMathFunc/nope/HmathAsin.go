package hmath

import(
	"math"
)

type AsinT struct{
	mathBodySingl
}

func(asin AsinT)String()(ret string){
	_, body := asin.GetBody()
	ret += "Asin("
	ret += body[0].String()
	ret += ")"
	return 
}

func(asin AsinT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := asin.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Asin(f), nil
}

func(asin AsinT)NeededValues()(ret MathComponents){
	_, body := asin.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Asin(a MathFunc)MathFunc{
	return AsinT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Asin",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Asin dy 0 object count for asin")
					}
					return Asin(input[0]), nil
				},
			),
		),
	)
}