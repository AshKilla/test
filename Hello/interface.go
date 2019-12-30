package main

import "fmt"

type Man interface {
    name() string
    age() int
}

type Woman struct {

}

func (woman Woman) name() string{
    return "chenjiao"
}

func (woman Woman) age() int{
    return 16
}

type Men struct {

}

func (men Men) name() string {
    return "chenjia"
}

func (men Men) age() int{
    return 20
}

func main(){
    var man Man

    man = new(Woman)
    fmt.Println(man.name())
    fmt.Println(man.age())

    man = new(Men)
    fmt.Println(man.name())
    fmt.Println(man.age())
}
