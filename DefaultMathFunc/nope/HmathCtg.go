package hmath

import(
	"math"
)

type CtgT struct{
	mathBodySingl
}

func(ctg CtgT)String()(ret string){
	_, body := ctg.GetBody()
	ret += "Ctg("
	ret += body[0].String()
	ret += ")"
	return 
}

func(ctg CtgT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := ctg.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return 1/math.Tan(f), nil
}

func(ctg CtgT)NeededValues()(ret MathComponents){
	_, body := ctg.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Ctg(a MathFunc)MathFunc{
	return CtgT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Ctg",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Ctg dy 0 object count for ctg")
					}
					return Ctg(input[0]), nil
				},
			),
		),
	)
}