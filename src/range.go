package main
import "fmt"
func main(){

	sli:=[]int{45,1123,521,512,1243,9}
	sum:=0
	for _,num := range sli{
		sum += num
	}
	fmt.Println("sli==",sli)
	fmt.Printf("sum(sli)=%v,avg(sli)=%v\n",sum,float32(sum)/float32(len(sli)))

	mp:=map[int]string{1:"math",2:"chinese",3:"English",4:"Physics",0:"C++"}
	fmt.Printf("删除前%v\n",mp)
	delete(mp,2)
	fmt.Printf("删除后%v\n",mp)
	for k,v:=range mp{
		fmt.Println(k,v)
	}

	for i,k:=range "1234567890"{
		fmt.Println(i,k)
	}
}
