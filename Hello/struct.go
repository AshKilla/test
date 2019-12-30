package main

import "fmt"

type books struct {
    title string
    author string
    subject string
    book_id int
}

//define shuju

var you = []int {101010,11,12}

//define slice
var slice = make([]int,len(you))

func main(){
    var book1 books

    book1.title = "create"
    book1.author = "ash kila"
    book1.subject = "art"
    book1.book_id = 123
    fmt.Println(book1)

}
