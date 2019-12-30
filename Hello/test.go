package main

import "fmt"

var (
    i int = 1
    s float32 = 12.0
    d float64 = 12.11
    //b bool = ture
    //p Pointer = & b
)

var name = "my name is chenjia"

var fool = "i think you are totally foolish"

func main()  {
    fmt.Println("what about you")
    fmt.Println(name,"haha",fool)
    fmt.Println(i,s,d)
}
