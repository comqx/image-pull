package client

import (
	"image-pull/proto"
	"fmt"
	"net"

	"github.com/docker/docker/client"
)

var (
	dockerclient *client.Client
	clienterr    error
	TcpClient    net.Conn
)

func init() {
	dockerclient, clienterr = client.NewEnvClient()
	if clienterr != nil {
		panic(clienterr)
	}
}

func GetClient() (*client.Client, error) {
	if dockerclient == nil {
		return nil, clienterr
	}
	return dockerclient, nil
}

// 发送数据
func TcpSend(conn net.Conn, message string) (err error) {
	// 调用协议编码数据
	var (
		b []byte
	)
	b, err = proto.Encode(message)
	if err != nil {
		fmt.Println("encode failed,err:", err)
		return
	}
	if _, err = conn.Write(b); err != nil {
		fmt.Println("发送数据失败：", err)
		conn.Close()
		return
	}
	return
}

//// 接受数据
//func Accept(reader *bufio.Reader) {
//	for {
//		recvStr, err := proto.Decode(reader) //直接调用这个包来解决
//		if err == io.EOF {
//			return
//		}
//		if err != nil {
//			fmt.Println("decode failed,err:", err)
//			return
//		}
//		b, err := proto.Encode(recvStr)
//		if err != nil {
//			fmt.Println("encode failed,err:", err)
//			return
//		}
//		fmt.Println(b)
//		TcpClient.Write(b)
//
//	}
//}
