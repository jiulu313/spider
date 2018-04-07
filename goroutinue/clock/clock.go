package main

import (
	"net"
	"log"
	"io"
	"time"
)

func main()  {
	listener , err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil{
		panic("err listener")
	}

	for{
		conn,err := listener.Accept()
		if err != nil{
			log.Printf("connect break")
			continue
		}

		go handleConn(conn)
	}


}
func handleConn(c net.Conn) {
	defer c.Close()
	for{
		_,err := io.WriteString(c,time.Now().Format("15:04:05\n"))
		if err != nil{
			return
		}

		time.Sleep(1 * time.Second)
	}

}




