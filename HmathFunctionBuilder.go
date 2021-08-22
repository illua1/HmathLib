package hmath
import(
	"strconv"
)
var(
	DefFunction = MathBuildContext{}
)

type MathBuilder func(...MathFunc)(MathFunc, error)

type mathBuildFunction struct{
	inputCount uint8
	name string
	body MathBuilder
}

func(mbf mathBuildFunction)String()string{
	return "Function "+mbf.name+", Input count "+strconv.Itoa(int(int8(mbf.inputCount)))
}
func NewFunc(inputCount uint8, name string, body MathBuilder)mathBuildFunction{
	return mathBuildFunction{
		inputCount : inputCount,
		name : name,
		body : body,
	}
}
func(mbf mathBuildFunction)MakeFunc(el ...MathFunc)(MathFunc, error){
	return mbf.body(el...)
}


type MathBuildContext map[string]mathBuildFunction
func(mbc MathBuildContext)String()(ret string){
	ret += "["
	for i := range mbc {
		ret += mbc[i].String() + "; "
	}
	if len(mbc) > 0 {
		ret = ret[:len(ret)-len("; ")]
	}
	ret += "]"
	return 
}

func NewMbc(elements ...mathBuildFunction)(ret MathBuildContext){
	ret = make(MathBuildContext)
	for i := range elements {
		ret[elements[i].name] = elements[i]
	}
	return
}

func(mbc MathBuildContext)Add(elements ...mathBuildFunction)(ret MathBuildContext){
	ret = mbc
	for i := range elements {
		ret[elements[i].name] = elements[i]
	}
	return
}