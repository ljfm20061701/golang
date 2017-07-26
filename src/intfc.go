package main
import "fmt"

type lan interface{
	name()
}

type cpp struct{
	named string
}
type java struct{
	named string
}
type golang struct{
	named string
}
type lua struct{
	named string
}
type shell struct{
	named string
}
type php struct{
	named string
}
type python struct{
	named string
}

func (c cpp) name(){
	fmt.Printf("I am %+v cpp\n",c)
}
func (c java) name(){
	fmt.Printf("I am %+v java\n",c)
}
func (c golang) name(){
	fmt.Printf("I am %+v golang\n",c)
}
func (c lua) name(){
	fmt.Printf("I am %+v lua\n",c)
}
func (c shell) name(){
	fmt.Printf("I am %+v shell\n",c)
}
func (c php) name(){
	fmt.Printf("I am %+v php\n",c)
}
func main(){

	var l lan
	l = new(cpp)
	l.name()

	
	l = new(java)
	l.name()
	l = new(golang)
	l.name()
	l = new(shell)
	l.name()
	l = new(php)
	l.name()
}
