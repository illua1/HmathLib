package hmath

type MathFunc interface {
	//MathArray
	
	Get(MathContext)(float64, error) // Solving
	//GetDerivative()MathFunc
	//GetAntiderivative()MathFunc
	//Simplifying()
	String()string // Printing
	NeededValues()(MathComponents, error) // Testin by all mc.Get("X") for solving
	ErrorTest()error
}

func MathFuncPrint(a ...MathFunc)(ret []string){
	for i := range a {
		ret = append(ret, a[i].String())
	}
	return
}

// Single math object body is array of mathematical object
/*
type MathArray interface {
	GetBody()(bool, []MathFunc)
	//Types()MathTypeId
}
*/

type names struct{
	name string
	value MathFunc
}
func(n names)String()string{
	return "F("+n.name+") = " + n.value.String()
}
func Name(str string, value MathFunc)names{
	return names{
		name : str,
		value : value,
	}
}

type MathComponents map[string]struct{}

func(mv MathComponents)String()(ret string){
	ret += "["
	for i,_ := range mv {
		ret += i + ", "
	}
	if len(mv) > 0 {
		ret = ret[:len(ret)-len(", ")]
	}else{
		ret += "No needed values"
	}
	ret += "]"
	return
}
func(mv MathComponents)Add(ellements ...MathComponents)(ret MathComponents){
	ret = make(MathComponents)
	for i,_ := range mv{
		ret[i] = struct{}{}
	}
	for i,_ := range ellements{
		for e,_ := range ellements[i]{
			ret[e] = struct{}{}
		}
	}
	return
}
func mvNew(str string)MathComponents{
	return MathComponents{str:struct{}{}}
}