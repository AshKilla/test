package main

import "fmt"

//all followed is global var that is not must used
var x, y int

var (
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello my freind"

//define a const var
const name = "chenjia"

//iota
const (
    i = iota
    l
    m
)
func main(){
    g, h := 456, "yeah that is ture" //those are local var that must be otherwise error
    fmt.Println(x,y,a,b,c,d,e,g,"\n",f,h)

    for s := 1; s < 10; s++{
            fmt.Println(h)
    }
}
