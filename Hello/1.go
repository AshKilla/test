/*
	基础的都在这
*/
package main
import "fmt"

type student struct{
	name string
	age int
	sex string
}

// 结构Student的成员函数
func (chen student) getname() string{
	return chen.name
}

// 全局多维数组
var arr = [5][5]int{
	{1,2,3,4,5},
	{1,2,3,4,5},
	{1,2,3,4,5},
	{1,2,3,4,5},
	{1,2,3,4,5},
}


func main(){
	var num1,num2,num3 int = 1,2,3
	var char rune = 's'
	var String string = "chenjia"
	var number = [5]int{1,2,3,4,5}
	var ptr **int		// 定义一个指向指针的指针
	var ptr2 *int = &num1
	ptr = &ptr2

	fmt.Println("go语言实战")
	fmt.Println(num1,num2,num3,char,String) 	// ASCII值
	fmt.Println(len(number))

	for i:= 0 ; i<len(number) ; i++{
		fmt.Println(number[i])
	}

	for i,s:= range number{
		fmt.Println("level",i+1,":",s)
	}

	var mider = MidSort(number,0,5,3)

	fmt.Println("查找的地方在：",mider)
	swap(&num1,&num2)
	fmt.Println(num1," ",num2)


	var s1 student
	s1.name = "chenjia"
	s1.age = 20
	
	fmt.Println("名字是",s1.getname())


	//打印指针
	fmt.Println("num1的地址是",&num1)
	fmt.Println("num的值是",**ptr)
}


//二分查找
func MidSort(num [5]int, left int, right int, n int) int {
	for left <= right {
		var mid = (left + right) / 2
		if num[mid] == n {
			return mid
		}else if( num[mid] > n){
			right = mid - 1
		}else{
			left = mid + 1
		}
	}

	return -1
}


//引用
func swap(a *int, b *int){
	var temp int

	temp = *a
	*a = *b 
	*b = temp
}


//闭包 匿名函数
func getSequence() func() int {
	i:=0
	return func() int {
	   i+=1
	  return i  
	}
}



