package main

import (
	"fmt"
	"os"
)

func main(){
	filename := "files.txt"
	currFile,err := os.Create(filename)
	if err != nil{
		fmt.Println(err,currFile)
	}
	currFile.WriteString("hello world\n")
	currFile,err = os.Open(filename)
	defer currFile.Close()
	buf := make([]byte,1024)
	for {
		n,_:=currFile.Read(buf)
		if n == 0{
			break
		}
		os.Stdout.Write(buf[:n]) 
	}
}