package github.com/illua1/HmathLib

type ValueT struct{
	
	str string
}
func(v ValueT)String()(string){
	return string(v.str)
}
func(v ValueT)Get(mc MathContext)(float64, error){
	if mc.values != nil {
		r, err := mc.GetValue(string(v.str))
		if err == nil {
			return r, nil
		}
		return 0, err
	}
	return 0, MathContextError("Values map is nil")
}

func(v ValueT)NeededValues()(MathComponents, error){
	return mvNew(string(v.str)), nil
}
func(v ValueT)ErrorTest()error{
	return nil
}
func Value(x string)MathFunc{
	return ValueT{str:x}
}
