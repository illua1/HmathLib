package hmath

import(
	"math"
)

type LogT struct{
	mathBodyCouple
}

func(log LogT)String()(ret string){
	_, body := log.GetBody()
	ret += "Log("
	ret += body[0].String()
	ret += ", "
	ret += body[1].String()
	ret += ")"
	return 
}

func(log LogT)Get(mc MathContext)(float64, error){
	//var f1, f2 float64
	_, body := log.GetBody()
	
	f1, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	f2, err := body[1].Get(mc)
	if err != nil {
		return 0, err
	}
	return math.Logb(f1)/math.Logb(f2), nil
	//return math.Log(f1, f2), nil
}

func(log LogT)NeededValues()(ret MathComponents){
	_, body := log.GetBody()
	
	ret = ret.Add(body[0].NeededValues())
	ret = ret.Add(body[1].NeededValues())
	
	return
}

func Log(a, b MathFunc)MathFunc{
	return LogT{mathBodyCouple{a, b}}
}

func init(){
	DefFunction.Add(
		NewFunc(
			2,
			"Log",
			mathBuilder(
				func(input ...MathFunc)(MathFunc, error){
					if len(input) < 2 {
						return nil, MathError("Can't make Log dy <2 object count for log")
					}
					return Log(input[0],input[1]), nil
				},
			),
		),
	)
}