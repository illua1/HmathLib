package hmath

import(
	"math"
)

type SinT struct{
	mathBodySingl
}

func(sin SinT)String()(ret string){
	_, body := sin.GetBody()
	ret += "Sin("
	ret += body[0].String()
	ret += ")"
	return 
}

func(sin SinT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := sin.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Sin(f), nil
}

func(sin SinT)NeededValues()(ret MathComponents){
	_, body := sin.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Sin(a MathFunc)MathFunc{
	return SinT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Sin",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Sin dy 0 object count for sin")
					}
					return Sin(input[0]), nil
				},
			),
		),
	)
}