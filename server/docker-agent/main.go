package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"image-pull/client"
	"image-pull/dockerCore"
	"image-pull/proto"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

func dockerPull(ImageListEnv, nodeInfo string) (err error) {
	var (
		ImageList []dockerCore.DockerInfo
		pullChan  = make(chan dockerCore.DockerInfo, 20)
		wg        sync.WaitGroup
	)
	if err = json.Unmarshal([]byte(ImageListEnv), &ImageList); err != nil {
		fmt.Println("变量：ImageListEnv 为空,", err)
		return
	}
	fmt.Println("开启线程")
	// 开启3个线程，去下载镜像
	wg.Add(3)
	for n := 1; n <= 3; n++ {
		go func(n int) {
			for p := range pullChan {
				if err = p.PullImage(n, nodeInfo); err != nil {
					fmt.Printf("下载出错：线程：%d 镜像名称：%s \n 报错：%s", n, p.ImageName, err.Error())
					return
				}
			}
			wg.Done()
		}(n)
	}
	fmt.Println("开始生成实例")
	for _, i := range ImageList {
		pullChan <- i
	}
	close(pullChan)
	wg.Wait()
	return
}

func main() {
	var (
		err error
		serverAddr,clusterName,nodeIp string
	)
	clusterName= os.Getenv("clusterName")
	if clusterName= os.Getenv("clusterName");clusterName ==""{
		clusterName="local"
	}
	if nodeIp=os.Getenv("nodeIp");nodeIp ==""{
		nodeIp="127.0.0.1"
	}
	if serverAddr= os.Getenv("serverAddr");serverAddr== ""{
		serverAddr="127.0.0.1:30000"
	}
	nodeInfo := fmt.Sprintf("%s/%s", os.Getenv("clusterName"), os.Getenv("nodeIp"))
	for {
		client.TcpClient, err = net.Dial("tcp",serverAddr)
		if err != nil {
			fmt.Println("连接报错，等待server启动：", err)
			time.Sleep(2 * time.Second)
			continue
		}
		defer client.TcpClient.Close()
		// 发送注册信息
		client.TcpSend(client.TcpClient, nodeInfo)

		fmt.Println("开始接受来自server的指令")
		reader := bufio.NewReader(client.TcpClient)
		for {
			recvStr, err := proto.Decode(reader) //直接调用这个包来解决
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			if err != nil {
				fmt.Println("decode failed,err:", err)
				break
			}
			fmt.Println("收到server发来的数据：", recvStr)
			if err = dockerPull(recvStr, nodeInfo); err != nil {
				break
			}
		}
	}

}
