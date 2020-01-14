package arithmatic

import("errors")

func Add(a ,b int)(int){
	return a+b
}
func Subtract(a,b int)(int){
	return a - b
}

func Division(a,b int)(float64,error){
	if b==0{
		return 0,errors.New("cannot divide a number from 0")
	}
	return float64(a)/float64(b) ,nil
}
func Multiply(a,b int)(int){
	return a*b
}
