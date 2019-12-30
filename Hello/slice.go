/*
cap():return array's size of  space was distributed      len() 方法获取长度。
len():build-in function,return size of current space   cap() 可以测量切片最长可以达到多少。
append():
*/

package main
import "fmt"


func main()  {
    numbers := []int {0,1,2,3,4,5,6,7}
    var num []int
    PrintSlice(numbers)
    fmt.Println(numbers[1:4],"\n")
    numbers = append(numbers,9,10)
    PrintSlice(numbers)
    num = append(numbers,9,10)
    PrintSlice(num)

}

func PrintSlice(x []int)  {
    fmt.Printf("len:%d,cap:%d,%v",len(x),cap(x),x)
}
