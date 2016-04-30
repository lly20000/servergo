package connector

import (
	//"bufio"
	"fmt"
	//"github.com/golang/protobuf/proto"
	//"myproto"
	"net"
	"os"
)

func StartServer(ip string, port string) {
	portstr := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", portstr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConection(conn)
	}

}

//long connect
func handleConection(conn net.Conn) {
	tmpBuffer := make([]byte, 0)

	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), "connection error: ", err)
			return
		}
		tmpBuffer = Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			Log(string(data))
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

//one connect one break connect
func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf []byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		//pares msg
		PaserMsg(buf)

		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
