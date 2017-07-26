package main
import "fmt"
import "unsafe"
const (
	U = "abcd" 
	M = len(U)
	F = unsafe.Sizeof(U)
)
func main(){
	const name = "三毛"
	const L,W = 4,6
	area :=L*W
	fmt.Println(area) 
	fmt.Printf("面积为:%d\n",area)
	fmt.Println("Hello,World")
	fmt.Println("中华人民共和国")
	var age int
	age = 28;
	age +=2
	var year int = 3
	fmt.Println(age+year);
	var egg int
	fmt.Println(egg)
	str:= "刘金福"
	fmt.Println(str)
	a,b,c,d := 1,",二",3,"南无阿弥陀佛"
	fmt.Println(a,b,c,d)

	f:= "sssssssssssssss"
	fmt.Println("hahahaha",f)
	_,b,c,d=1,"",3,""
	fmt.Println(U,M,F)
	fmt.Println(a,b,c,d)
	var p_a *int = &a;
	fmt.Printf("a的地址为:%X\n",p_a);
	fmt.Printf("a的地址为:%X\n",&a);
	fmt.Printf("p_a的类型为:%T\n",p_a);
	
	if (age < 30) {
	fmt.Printf("age小于30\n")
	} else if(age ==30) {
	fmt.Printf("age等于30\n")
	} else {
	fmt.Printf("age大于30\n")
	}
	var grade string="B"
	var score int = 90
	switch score{
		case 90:
			grade = "A"
		case 80:
			grade = "B"
		case 50,60,70:
			grade = "c"
		default: grade = "D"
	}
	fmt.Printf("grade is %s\n",grade)
	switch grade {
		case "A":
			fmt.Println("优秀");
		default:
			fmt.Println("要加油哦");
	}
}
