package main
import "fmt"

type Sbook struct{
	book_name string
	price	float64
	product	string
	author	string
	book_id int
}

func main(){
	var b1,b2 Sbook
	b1.book_name = "C++核心编程"
	b1.price = 32.6
	b1.product = "电子工业出版社"
	b1.author = "谭浩强"
	b1.book_id = 9006889901

	b2.book_name = "golang从入门到精通"
	b2.price = 63.5
	b2.product = "电子工业出版社"
	b2.author = "刘金福"
	b2.book_id = 9006889904
	b1.prt()
	b2.prt()

}
func (b Sbook) prt(){
	fmt.Printf("/*********结构输出*********\n")
	fmt.Printf("%+v\n",b.book_name)
	fmt.Printf("%+v\n",b.price)
	fmt.Printf("%+v\n",b.product)
	fmt.Printf("%+v\n",b.author)
	fmt.Printf("%#v\n",b.book_id)
	fmt.Printf("%+v\n",b)
	fmt.Printf("%#v\n",b)
	fmt.Printf("*********结构输出*********/\n\n")
}
