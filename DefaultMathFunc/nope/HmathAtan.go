package hmath

import(
	"math"
)

type ArctanT struct{
	mathBodySingl
}

func(arctan ArctanT)String()(ret string){
	_, body := arctan.GetBody()
	ret += "ArcTan("
	ret += body[0].String()
	ret += ")"
	return 
}

func(arctan ArctanT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := arctan.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Atan(f), nil
}

func(arctan ArctanT)NeededValues()(ret MathComponents){
	_, body := arctan.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Atan(a MathFunc)MathFunc{
	return ArctanT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Arctan",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Arctan dy 0 object count for arctan")
					}
					return Atan(input[0]), nil
				},
			),
		),
	)
}