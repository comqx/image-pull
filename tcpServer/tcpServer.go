package tcpServer

import (
	"bufio"
	"image-pull/proto"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

var (
	RegisteredMap map[string]map[string]net.Conn // 注册的agent连接map
	//RegisteredSyncMap sync.Map
)

// 解析数据
func pars(recvStr string) []string {
	return strings.Split(recvStr, "/")
}

// 注册到map中
func registered(reader *bufio.Reader, conn net.Conn) {
	recvStr, err := proto.Decode(reader) //直接调用这个包来解决
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println("decode failed,err:", err)
		return
	}
	recvStrArr := pars(recvStr)
	if len(recvStrArr) != 2 {
		fmt.Println("注册的信息不正确")
		return
	}
	//// 并发map
	//M := &sync.Map{}
	//M.Store(recvStrArr[1], conn)
	//RegisteredSyncMap.Store(recvStrArr[0], M)

	// 普通map
	if RegisteredMap[recvStrArr[0]] == nil {
		RegisteredMap[recvStrArr[0]] = make(map[string]net.Conn, 10)
	}
	RegisteredMap[recvStrArr[0]][recvStrArr[1]] = conn

	b, _ := json.Marshal(RegisteredMap)
	fmt.Printf("注册成功：%s  数据：%s\n", recvStr, string(b))
}

// 接受数据
func Accept(reader *bufio.Reader) {
	for {
		recvStr, err := proto.Decode(reader) //直接调用这个包来解决
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode failed,err:", err)
			return
		}
		fmt.Println(recvStr)
	}
}

// 发送数据指令
func Send(conn net.Conn, sendData string) {
	fmt.Println(">>>>>", sendData)
	b, err := proto.Encode(sendData)
	if err != nil {
		fmt.Println(err)
	}
	conn.Write(b)
}

func Process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// 注册服务到server端
	registered(reader, conn)
	Accept(reader)
	//go send(conn, server2clientChan)
}

func SerInit(listenAddr string) (listen net.Listener, err error) {
	RegisteredMap = make(map[string]map[string]net.Conn, 10)
	//RegisteredSyncMap = sync.Map{}

	listen, err = net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	// 持续监听 提供服务
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("accept failed, err:", err)
				time.Sleep(time.Second)
				continue
			}
			// 把连接信息注册到map中
			go Process(conn)
		}
	}()
	return
}
