package main
import "fmt"
func main(){

	var sz =[5][2]int{{2,3},{4,55},{9,700},{899,89},{900,400}}
	showSz(sz,5,2)
	a,b:= avrg(sz,5,2)
	fmt.Println(a,b)
}
func showSz(arr [5][2] int, size1 int,size2 int) {

	var i,j int
	for i=0;i<size1;i++{
		for j=0;j<size2;j++{
			fmt.Printf("array[%d][%d] = %4d\t",i,j,arr[i][j]);
			}
		fmt.Println();
	}
}
func avrg(arr [5][2] int, size1 int,size2 int) (int,float64){
	var i,j,sum int
	var avg float64
	for i=0;i<size1;i++{
		for j=0;j<size2;j++{
			sum += arr[i][j]
			}
	}
	avg = float64(sum) / float64(size1*size2);
	fmt.Printf("%+v\n",arr)
	fmt.Printf("数组总和为%+v,平均数为%+v\n",sum,avg)
	fmt.Printf("sum的地址为%P,%X\n",&sum,&sum)
	return sum,avg
}
