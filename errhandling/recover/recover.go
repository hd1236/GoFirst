package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover() //捕获err，使程序不崩溃
		//if err,ok := r.(error); ok{
		//    fmt.Println("Error occurred:",err)
		//}else{
		//    panic(fmt.Sprintf("I don't know what to do:%v",r))
		//}
		switch r.(type) {
		case error:
			fmt.Println("Error occurred:", r.(error))
		case int:
			fmt.Println("Int occurred:", r.(int))
		case string:
			fmt.Println("string occurred:", r.(string))
		default:
			panic(fmt.Sprintf("I don't know what to do:%v", r))
		}
	}()
	//panic(errors.New("this is a error"))

	//b := 0;
	//a := 5/b
	//fmt.Println(a)

	panic(123) //终止程序继续运行，直接返回

	panic("hand")
}

func main() {
	tryRecover()
}
