package defaultF

import(
	//"Src/StringsFactoryCollection"
	"Src/MathLibTest"
)

type MulT struct{
	hmath.MultipleMathElements
}

func(a MulT)String()(string){
	{
		_, body := a.GetBody()
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return errTest("").Error()
		}
	}
	return StrMul(
		a.GetBodyStr()...,
	)
}

func(a MulT)Get(mc hmath.MathContext)(float64, error){
	_, body := a.GetBody()
	{
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return 0, errTest("")
		}
	}
	
	var f float64
	var err error
	if len(body) != 0 {
		f, err = body[0].Get(mc)
		if err != nil {
			return 0, err
		}
	}else{
		return 0, nil
	}
	
	for i := 1; i < len(body); i++{
		ret, err := body[i].Get(mc)
		if err != nil {
			return 0, err
		}
		f *= ret
	}
	return f, nil
}

func(a MulT)NeededValues()(ret hmath.MathComponents){
	_, body := a.GetBody()
	{
		if errTest := hmath.MathFuncNilChack(body...); errTest != nil {
			return ret
		}
	}
	for i := range body{
		ret = ret.Add(body[i].NeededValues())
	}
	return
}

func Mul(a ...hmath.MathFunc)hmath.MathFunc{
	return MulT{
		hmath.NewMultipleMathElements(a...),
	}
}

func init(){
	hmath.DefFunction.Add(
		hmath.NewFunc(
			255,
			"Mul",
			hmath.MathBuilder(
				func(input ...hmath.MathFunc)(hmath.MathFunc, error){
					return Mul(input...), nil
				},
			),
		),
	)
}