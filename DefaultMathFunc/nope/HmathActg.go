package hmath

import(
	"math"
)

type ArcctgT struct{
	mathBodySingl
}

func(arcctg ArcctgT)String()(ret string){
	_, body := arcctg.GetBody()
	ret += "ArcCtg("
	ret += body[0].String()
	ret += ")"
	return 
}

func(arcctg ArcctgT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := arcctg.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return (math.Pi/2)-math.Atan(f), nil
}

func(arcctg ArcctgT)NeededValues()(ret MathComponents){
	_, body := arcctg.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Actg(a MathFunc)MathFunc{
	return ArcctgT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"ArcCtg",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make ArcCtg dy 0 object count for crcCtg")
					}
					return Actg(input[0]), nil
				},
			),
		),
	)
}