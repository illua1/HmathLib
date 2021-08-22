package hmath

type MathContext struct {
	values map[string]MathFunc
}
func(mc MathContext)String()(ret string){
	ret += "["
	for i := range mc.values{
		ret += i+" = "+mc.values[i].String()+"; "
	}
	if len(ret)>0{
		ret = ret[:len(ret)-len("; ")]
	}
	ret += "]"
	return
}
func MakeMathContext(a... names)(mc MathContext){
	mc.values = map[string]MathFunc{}
	for i:=range a {
		mc.values[a[i].name] = a[i].value
	}
	return
}
func(mc MathContext)Get(name string)(MathFunc, error){
	ret, ok := mc.values[name]
	if !ok {
		return nil, MathContextError("Values map no have value <"+name+">")
	}
	return ret, nil
}
func(mc MathContext)GetValue(name string)(float64, error){
	
	ret, ok := mc.values[name]
	if !ok {
		return 0, MathContextError("Values map no have value <"+name+">")
	}
	
	needed, err := ret.NeededValues()
	{
		if err != nil {
			return 0, err
		}
	}
	if len(needed) == 0 {
		return ret.Get(mc)
	}
	
	test, err := chack([]string{name}, mc, needed)
	{
		if err != nil {
			return 0, err
		}
	}
	if test {
		// MAKE MATH SYSTEM SOLVER!
		return 0, MathError("This version can't use recurent math context system")
	}
	return ret.Get(mc)
}
func chack(history []string, mc MathContext, test MathComponents)(bool, error){
	for name := range test{
		need, err := mc.values[name].NeededValues()
		{
			if err != nil {
				return false, err
			}
		}
		if len(need) != 0 {
			if strTest(
				history,
				need,
			){
				return true, nil
			}
			chack(append(history, name), mc, need)
		}
	}
	return false, nil
}
func strTest(str []string, maps MathComponents)(bool){
	for i := range str {
		if _,ok := maps[str[i]]; ok == true {
			return true
		}
	}
	return false
}