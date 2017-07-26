package main
import "fmt"
func main(){
	var c1,c2,c3 chan int
	var i1,i2 int
	c1 <- 10;
	//i2 = 20;
	//c3 <-90;
	select {
		case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
     // default:
       //  fmt.Printf("no communication\n")
   }    
	fmt.Printf("c1=%d,c2=%d,c3=%d\ni1 = %d,i2=%d\n",c1,c2,c3,i1,i2)
}
