package hmath

import(
	"math"
)

type PowT struct{
	mathBodyCouple
}

func(pow PowT)String()(ret string){
	_, body := pow.GetBody()
	ret += body[0].String()
	ret += "^"
	ret += body[1].String()
	return 
}

func(pow PowT)Get(mc MathContext)(float64, error){
	_, body := pow.GetBody()
	
	f1, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	f2, err := body[1].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Pow(f1, f2), nil
}

func(pow PowT)NeededValues()(ret MathComponents){
	_, body := pow.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Pow(a, b MathFunc)MathFunc{
	return PowT{mathBodyCouple{a, b}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Pow",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) < 2 {
						return nil, MathError("Can't make Pow dy <2 object count for pow")
					}
					return Pow(input[0],input[1]), nil
				},
			),
		),
	)
}