package main

import "fmt"


func main()  {
    ConCap := map[string]string{"bro":"chenbin","father":"chenjunchu","mom":"tangyanfei"}

    for i,y := range ConCap{
        fmt.Println(i,"is",y)
    }

    delete(ConCap,"father")

    for i,y := range ConCap{
        fmt.Println(i,"is",y)
    }
}
