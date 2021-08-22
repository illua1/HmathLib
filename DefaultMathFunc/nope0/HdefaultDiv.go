package defaultF

import(
	"Src/StringsFactoryCollection"
	"Src/MathLibTest"
)

type DivT struct{
	hmath.CoupleMathElements
}

func(a DivT)String()(string){
	{
		_, body := a.GetBody()
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return errTest("").Error()
		}
	}
	return StrDiv(
		a.GetBodyStr()...,
	)
}

func(a DivT)Get(mc hmath.MathContext)(float64, error){
	_, body := a.GetBody()
	{
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return 0, errTest("")
		}
	}
	
	x, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	y, err := body[1].Get(mc)
	if err != nil {
		return 0, err
	}
	return x/y, nil
}

func(a DivT)NeededValues()(ret hmath.MathComponents){
	_, body := a.GetBody()
	{
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return
		}
	}
	ret = ret.Add(body[0].NeededValues())
	return
}

func Div(a, b hmath.MathFunc)hmath.MathFunc{
	return DivT{
		hmath.NewCoupleMathElements(a,b),
	}
}

func init(){
	hmath.DefFunction.Add(
		hmath.NewFunc(
			1,
			"Div",
			hmath.MathBuilder(
				func(input ...hmath.MathFunc)(hmath.MathFunc, error){
					if len(input) != 2 {
						return nil, StrFactory.ErrorString("Need 2 cound of object")
					}
					return Div(input[0],input[1]), nil
				},
			),
		),
	)
}