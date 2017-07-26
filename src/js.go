package main
import "fmt"
func main(){
	for i:=0;i<30;i++{
		fmt.Printf("%v\t",fib(i))
		if (i+1)%10 ==0{
		fmt.Println()
		}
	}
	//fmt.Printf("%d的阶乘为%d\n",15,js(15))
}
func js(a int) int{
	if (a==1){
	return 1
	}else{
	return a*js(a-1)
	}
}
func fib(a int) int{
	if(a<2){
	return a
	}else{
	return fib(a-2)+fib(a-1)
	}
}
