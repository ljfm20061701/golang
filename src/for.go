package main
import(
 "fmt"
 "math"
)

func main(){
	c,k:=4,6;
	cd,mj:=zf(c,k);
	fmt.Printf("长为%d,宽为%d;\n面积为%d,周长为%d\n",c,k,mj,cd);
	fmt.Printf("交换前c = %d,k=%d\n",c,k);
	swp(&c,&k)
		
	fmt.Printf("交换后c = %d,k=%d\n",c,k);
	getsr := func(x float64) float64 {
		return math.Sqrt(x);
	}
	fmt.Println(getsr(900.09));
	//xs()
	//ss()
	//cf()
}
func swp(a ,b *int){
	var t int;
	t = *a;
	*a = *b;
	*b = t;
}
func zf(c,k int) (int,int){
	return (c+k)*2,c*k
}
func cf(){
	var i,j int
	fmt.Println("下面打印乘法口诀:")
	for i=1; i<10;i++{
		for j=1;j<=i;j++{
			fmt.Printf("%d*%d=%2d\t",j,i,j*i)
		}
		fmt.Println("");
	} 
	fmt.Println("上面打印乘法口诀end")
}
func ss(){
	var i,j int
	for i=2; i<200;i++{
		for j=2;j<=i/j;j++{
			if i%j==0 {
				break;
			}
		}
		if(j > i/j){
			fmt.Printf("%4d是素数\n",i);
		}
	}
}
func xs(){
	var b int = 15;
	var a int;
	for a:=0;a<10;a++{
		fmt.Printf("a = %d\n",a);
	}

	for a<b {
		a++;
		fmt.Printf("#####a = %d\n",a)
	}

	var numbers[9] int =[9] int {1,2,5,3,4,6,100}

	for i,j:= range numbers{
		fmt.Printf("第%d位为%d\n",i,j)
	}
	for b<25{
		b++;
		if b ==18{
			continue;
		}
		fmt.Printf("b = %d\n",b)
	}	
}
