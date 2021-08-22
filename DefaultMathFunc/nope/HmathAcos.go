package hmath

import(
	"math"
)

type AcosT struct{
	mathBodySingl
}

func(acos AcosT)String()(ret string){
	_, body := acos.GetBody()
	ret += "Acos("
	ret += body[0].String()
	ret += ")"
	return 
}

func(acos AcosT)Get(mc MathContext)(float64, error){
	var f float64
	_, body := acos.GetBody()
	
	f, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Acos(f), nil
}

func(acos AcosT)NeededValues()(ret MathComponents){
	_, body := acos.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Acos(a MathFunc)MathFunc{
	return AcosT{mathBodySingl{a}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Acos",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) == 0 {
						return nil, MathError("Can't make Acos dy 0 object count for acos")
					}
					return Acos(input[0]), nil
				},
			),
		),
	)
}