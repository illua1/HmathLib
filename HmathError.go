package hmath


type MathContextError string
func(mr MathContextError)Error()(string){
	return "MathContext Error: "+string(mr)
}

type MathError string
func(mr MathError)Error()(string){
	return string(mr)
}
