package main
import "fmt"
func main(){
	c:=make(chan int,23)
	fmt.Printf("c = %d\n",c);
	c <- 10 ;

	fmt.Printf("c = %d\n",c);
	v := <-c;
	
	fmt.Printf("v = %d\n",v);
}
