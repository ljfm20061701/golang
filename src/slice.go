package main 
import "fmt"
func main(){
	/*sli:=[]string{1,2,3,5,0,4,6,7,9,8}
	prstringslice(sli)

	fmt.Prstringln("sli==",sli)
	fmt.Prstringln("sli[1:4]==",sli[1:4])
	fmt.Prstringln("sli[:3]==",sli[:3])
	fmt.Prstringln("sli[6:]==",sli[6:])
	
	p :=make([]string,0,5)
	fmt.Prstringln("p==",p)
	prstringslice(p)*/
	var a [] string
	prstringslice(a)
	
	a = append(a,"a")
	prstringslice(a)
	a = append(a,"cd2","3","4","iooo","hello")
	prstringslice(a)
	b:=make([]string,len(a),cap(a)*2)
	prstringslice(b)
	copy(b,a)
	prstringslice(b)
}
func prstringslice(s[]string){
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(s),cap(s),s)
}
