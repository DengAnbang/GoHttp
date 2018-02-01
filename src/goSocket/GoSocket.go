package goSocket

import (

	"net"
	"fmt"
	"os"
	"utils"
)

func GoSocket(host, port string) {
	utils.Loge("GoSocket")
	netListen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
	defer netListen.Close()
	for {
		utils.Loge("等待连接1")
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		utils.Loge(conn.RemoteAddr().String(), "连接成功")
		sendMessage(conn, "服务器")
		go handleConnection(conn)
	}
}

//处理连接
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			utils.Loge(conn.RemoteAddr().String(), "connection error: ", err)
			return
		}
		utils.Loge(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		sendMessage(conn, "服务器收到")
	}

}
