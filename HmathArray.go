package hmath

import(
	"github.com/illua1/StringsFactoryCollection"
)



type MathArray struct{
	body []MathFunc
	countTest StrFactory.LenTesterError
}

func(mr MathArray)Get()([]MathFunc, error){
	
	{
		test := mr.countTest(
			len(mr.body),
		)
		if test != nil{
			return nil, test
		}
	}
	
	{
		nilTest := MathFuncNilTest(mr.body...)
		if nilTest != nil {
			return nil, nilTest
		}
	}
	
	return mr.body, nil
}

func MathArrayMake(
	countTest StrFactory.LenTesterError, 
	body ...MathFunc,
)(MathArray){
	return MathArray{
		body : body,
		countTest : countTest,
	}
}
