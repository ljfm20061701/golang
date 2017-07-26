package main
import (
"fmt"
)
func getsequence() func() int{
	i:=0
	return func() int{
		i++
	return i
	}
}
type square struct{
	ck int
}
type Circle struct{
	radius float64
}
const (
	PI=3.1415926
)
func (c Circle) getArea() float64{
	return PI*c.radius*c.radius
}
func main(){
	fmt.Println(PI)
	var c1 Circle
	c1.radius = 10.0
	fmt.Printf("半径为%.2f的圆面积为%.2f\n",c1.radius,c1.getArea())
	gc := getsequence();
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	fmt.Println(gc())
	
	
	
	gc1 := getsequence();
	fmt.Println(gc1())
	fmt.Println(gc1())
	fmt.Println(gc1())
	fmt.Println(gc1())
	fmt.Println(gc1())
	fmt.Println(gc1())
	
}
