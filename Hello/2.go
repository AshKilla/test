/*
	高级的
	slice:
		len()
		cap()
		append()
	range:range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素

	map(集合):无序的键值对的集合
		delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key	delete(countryCapitalMap, "France")


	类型转换：和c++一样
*/

package main

import "fmt"

type student struct {
	name string
	age int
}

//构造函数
func NewStudent(name string,age int) student{
	return student{
		name : name,
		age : age,
	}
}

//成员函数
func (stu student) getname(){
	fmt.Println(stu.name)
}


func main()  {

	sum := 0;
	var chen  = [5]int{1,2,3,4,5}

	//var strdub map[string]string


	for _,num := range chen{
		sum += num
	}

	fmt.Println("sum:",sum)

	stu := NewStudent("chenjia",20)
	stu.getname()

}
