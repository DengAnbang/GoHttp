package goSocket

import (
	"net"
	"encoding/binary"
	"main/utils"
)

func sendMessage(conn net.Conn,string string) {
	bytes := []byte(string)
	i := len(bytes)
	utils.Loge(i)
	toInt32 := Int32ToBytes(i)
	conn.Write(toInt32)
	conn.Write(bytes)
}
func Int32ToBytes(i int) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt32(buf []byte) int32 {
	return int32(binary.BigEndian.Uint64(buf))
}