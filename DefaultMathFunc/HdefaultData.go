package defaultF

import(
	"Src/StringsFactoryCollection"
	"Src/MathLibTest"
)


var (
	PackagePart = hmath.Prog.Part("defaultF")
)

var (
	ArrayNilError = StrFactory.ErrorString("nil element")
)


var (
	ValidError = "Input error"
	/*
	MulError = StrFactory.ErrorMsgMake(
		PackagePart,
		errorTypes,
	)
	SubError = StrFactory.ErrorMsgMake(
		PackagePart,
		errorTypes,
	)
	DivError = StrFactory.ErrorMsgMake(
		PackagePart,
		errorTypes,
	)
	PowError = StrFactory.ErrorMsgMake(
		PackagePart,
		errorTypes,
	)
	*/
)

/*
var(
	StrMul = func(a ...string)string{
		return strings.Join(
			a,
			"*",
		)
	}
)
func StrSub(a ...string)string{
	if lt := StrFactory.LenTest(len(a), 0); lt != nil{
		return lt(
			StrFactory.ErrorZone(
				"-",
			),
		).Error()
	}
	return "-"+a[0]
}
func StrDiv(a ...string)string{
	if lt := StrFactory.LenTest(len(a), 2); lt != nil{
		return lt(
			StrFactory.ErrorZone(
				"/",
			),
		).Error()
	}
	return a[0]+"/"+a[1]
}
func StrPow(a ...string)string{
	if lt := StrFactory.LenTest(len(a), 2); lt != nil{
		return lt(
			StrFactory.ErrorZone(
				"*",
			),
		).Error()
	}
	return StrFactory.MathZone(a[0])+"^"+StrFactory.MathZone(a[1])
}
*/

var(
	CountLoy = StrFactory.ErrorString("not enough elements")
	CountHight = StrFactory.ErrorString("too many elements")
)


