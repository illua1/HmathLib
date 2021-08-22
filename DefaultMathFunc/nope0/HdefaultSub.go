package defaultF

import(
	//"Src/StringsFactoryCollection"
	"Src/MathLibTest"
)

type SubT struct{
	hmath.SingleMathElements
}

func(a SubT)String()(string){
	{
		_, body := a.GetBody()
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return errTest("").Error()
		}
	}
	return StrSub(
		a.GetBodyStr()...,
	)
}

func(a SubT)Get(mc hmath.MathContext)(float64, error){
	_, body := a.GetBody()
	{
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return 0, errTest("")
		}
	}
	
	
	ret, err := body[0].Get(mc)
	if err != nil {
		return 0, err
	}
	return -ret, nil
}

func(a SubT)NeededValues()(ret hmath.MathComponents){
	_, body := a.GetBody()
	{
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return
		}
	}
	ret = ret.Add(body[0].NeededValues())
	return
}

func Sub(a hmath.MathFunc)hmath.MathFunc{
	return SubT{
		hmath.NewSingleMathElements(a),
	}
}

func init(){
	hmath.DefFunction.Add(
		hmath.NewFunc(
			1,
			"Sub",
			hmath.MathBuilder(
				func(input ...hmath.MathFunc)(hmath.MathFunc, error){
					return Sub(input[0]), nil
				},
			),
		),
	)
}