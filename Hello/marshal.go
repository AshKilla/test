// 序列化

package main

import "encoding/json"
import "fmt"

type person struct{
	Name string
	Age   int
}

func main(){
	p1 := person{
		Name:"chenjia",
		Age:12,
	}

	//序列化
	b,err := json.Marshal(p1)
	if err != nil{
		fmt.Printf("marshal failed,err:%v",err)
		return
	}

	fmt.Printf("%v\n",string(b))

	//反序列化
	str := `{"name":"理性","age":18}`
	var p2 person
	json.Unmarshal([]byte(str),&p2)
	fmt.Println(p2)
}