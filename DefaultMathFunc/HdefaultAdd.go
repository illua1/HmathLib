package defaultF

import(
	"Src/StringsFactoryCollection"
	"Src/MathLibTest"
	
	"strings"
)
type AddT struct{
	body hmath.MathArray
}

func(a AddT)String()(string){
	body, err := a.body.Get()
	{
		if err != nil {
			return AddError(err).Error()
		}
	}
	return StrAdd(
		hmath.MathFuncPrint(body...)...,
	)
}

func(a AddT)Get(mc hmath.MathContext)(float64, error){
	body, err := a.body.Get()
	{
		if err != nil {
			return 0, AddError(err)
		}
	}
	var f float64
	for i := range body{
		ret, err := body[i].Get(mc)
		if err != nil {
			return 0, err
		}
		f += ret
	}
	return f, nil
}

func(a AddT)NeededValues()(hmath.MathComponents, error){
	body, err := a.body.Get()
	{
		if err != nil {
			return nil, AddError(err)
		}
	}
	
	{
		var ret hmath.MathComponents
		
		for i := range body{
			need, err := body[i].NeededValues()
			if err != nil {
				return nil, err
			}
			ret = ret.Add(need)
		}
		return ret, nil
	}
}

func(a AddT)ErrorTest()error{
	_, err := a.body.Get()
	
	return err
}

func Add(a ...hmath.MathFunc)hmath.MathFunc{
	return AddT{
		hmath.MathArrayMake(
			AddCountTest,
			a...,
		),
	}
}

var (
	AddName = "Add"
	
	AddError = StrFactory.ErrorMsgMake(
		PackagePart.Part(
			AddName,
		),
		ValidError,
	)
	
	StrAdd = func(a ...string)string{
		return StrFactory.MathZone(
			strings.Join(
				a,
				"+",
			),
		)
	}
	
	AddCountTest = StrFactory.LenTesterError(
		func(x int)error{
			if x == 0 {
				return CountLoy
			}
			return nil
		},
	)
)

func init(){
	hmath.DefFunction.Add(
		hmath.NewFunc(
			255,
			AddName,
			hmath.MathBuilder(
				func(input ...hmath.MathFunc)(hmath.MathFunc, error){
					ret := Add(input...)
					
					if err := ret.ErrorTest(); err != nil{
						return nil, err
					}
					
					return Add(input...), nil
				},
			),
		),
	)
}