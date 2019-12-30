// 用go语言实现socket服务
//TCp 服务端

import (
	"fmt"
	"net"
)

//处理函数
func process(con net.Conn)  {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err != nil{
			fmt.Println("read from client failed,err:",err)
			break
		}

		recv := string(buf[:n])
		fmt.Println("接受到的数据为：",recv)
	}
}


func main(){

	//1. 监听端口
	Listen,err := net.Listen("tcp","127.0.0.1:20000")
	if err != nil{
		fmt.Println("listen failed,err:",err)
		return 
	}
	for {
		// 等待客户端来链接
		conn, err := Listen.Accept()
		if err!=nil {
			fmt.Println("accpet failed,err:",err)
		}
		go process(conn)	
	}

}