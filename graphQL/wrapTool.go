package graphQL

import (
	
	M	"github.com/gerard94/util/misc"
	SC	"strconv"

)

func Wrap (i ...interface{}) *OutputObjectValue {
	o := NewOutputObjectValue()
	for j, in := range i {
		o.InsertOutputField("data" + SC.Itoa(j), MakeAnyValue(in))
	}
	return o
} //Wrap

func Unwrap (rootValue *OutputObjectValue, n int) interface{} {
	f := rootValue.First()
	for i := 0; i < n && f != nil; i++ {
		f = rootValue.Next(f)
	}
	M.Assert(f != nil, 20)
	switch v := f.Value.(type) {
	case *AnyValue:
		return v.Any
	case *NullValue:
		return v
	default:
		M.Halt(100)
		return nil
	}
} //Unwrap
