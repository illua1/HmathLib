package hmath

import(
	"math"
)

type CosT struct{
	mathBodySingl
}

func(cos CosT)String()(ret string){
	_, body := cos.GetBody()
	ret += "Cos("
	ret += body[0].String()
	ret += ")"
	return 
}

func(cos CosT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := cos.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Cos(f), nil
}

func(cos CosT)NeededValues()(ret MathComponents){
	_, body := cos.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Cos(a MathFunc)MathFunc{
	return CosT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Cos",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Cos dy 0 object count for cos")
					}
					return Cos(input[0]), nil
				},
			),
		),
	)
}