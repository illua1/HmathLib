package hmath

import(
	"Src/StringsFactoryCollection"
//	"fmt"
)

var(
	SinStr = StrFactory.FuncBody(
		"Sin",
		StrFactory.MathZone,
	)
	CosStr = StrFactory.FuncBody(
		"Sin",
		StrFactory.MathZone,
	)
	TanStr = StrFactory.FuncBody(
		"Tan",
		StrFactory.MathZone,
	)
	CtgStr = StrFactory.FuncBody(
		"Ctg",
		StrFactory.MathZone,
	)
	AsinStr = StrFactory.FuncBody(
		"Asin",
		StrFactory.MathZone,
	)
	AcosStr = StrFactory.FuncBody(
		"Acos",
		StrFactory.MathZone,
	)
	AtanStr = StrFactory.FuncBody(
		"Atan",
		StrFactory.MathZone,
	)
	ActgStr = StrFactory.FuncBody(
		"Actg",
		StrFactory.MathZone,
	)
	//...
)


var (
	Prog = StrFactory.NewProgrammZone("hmath")
	Version = Prog.Part("this version")
	Context = Prog.Part("context")
)


var (
	HmathError = StrFactory.ErrorMsgMake(
		Prog,
		"error",
	)
	VersionError = StrFactory.ErrorMsgMake(
		Version,
		"can't make",
	)//	return 0, VersionError("Find system solution")
	ContextError = StrFactory.ErrorMsgMake(
		Prog,
		"context",
	)
)
/*
func init(){
	fmt.Println(
		VersionError("find system solution"),
	)
	fmt.Println(
		ContextError(" no have this value"),
	)
}
*/

var (
	ArrNilBodyError = StrFactory.ErrorString("nil body")
	ArrEmptyBodyError = StrFactory.ErrorString("empty body")
	ArrNilEllemetError = StrFactory.ErrorString("nil element")
)
var(
	MathFuncNilTest = func(a ...MathFunc)(error){
		if a == nil{
			return ArrNilBodyError // if call NilTest() <- empty
		}
		if len(a) == 0 {
			return ArrEmptyBodyError
		}
		for i:=range a{
			if a[i] == nil {
				return ArrNilEllemetError
			}
		}
		return nil
	}
	/*
	MathFuncValidalChack = func(lt StrFactory.LenTester, body ...MathFunc)(StrFactory.ErrorPrinter){
		if MathFuncNilTest(body...) {
			return StrFactory.ErrorPrinter(
				func(str string)(error){
					return StrFactory.ErrorString(
						str+StrFactory.ErrorZone("nil element"),
					)
				},
			)
		}
		if !lt(len(body)){
			return StrFactory.ErrorPrinter(
				func(str string)(error){
					return StrFactory.ErrorString(
						str+StrFactory.ErrorZone("Insufficient number of arguments"),
					)
				},
			)
		}
		return nil
	}
	*/
)


