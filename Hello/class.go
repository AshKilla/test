package main
import (
	"fmt"
)

type user struct {
	name string
	email string
}

//值接收者
func (u user) notify(){
	fmt.Println("Email is",u.email)
}

//指针接收者
func (u * user) changeEmail(email string){
	u.email = email
}

//构造一个假的构造函数
func (u * user) New(name,email string){
	u.name = name
	u.email = email
}

func main(){
	//var u user =user{"chenjia","2298981275@qq.com"}
	var u user
	u.New("chenjia","2298981275@qq.com")
	fmt.Println("this is original:")
	u.notify()
	u.changeEmail("1986609238@qq.com")
	fmt.Println("this is another:")
	u.notify()
}