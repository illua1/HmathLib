package hmath

import(
	"math"
)

type TanT struct{
	mathBodySingl
}

func(tan TanT)String()(ret string){
	_, body := tan.GetBody()
	ret += "Tan("
	ret += body[0].String()
	ret += ")"
	return 
}

func(tan TanT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := tan.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Tan(f), nil
}

func(tan TanT)NeededValues()(ret MathComponents){
	_, body := tan.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Tan(a MathFunc)MathFunc{
	return TanT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Tan",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Tan dy 0 object count for tan")
					}
					return Tan(input[0]), nil
				},
			),
		),
	)
}