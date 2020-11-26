package service

import (
	"image-pull/tcpServer"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type dockerInfo struct {
	UserName  string `json:"userName"`
	Password  string `json:"passWord"`
	ImageName string `json:"imageName"`
}

type imageInfoStruct struct {
	K8sName         string       `json:"k8sName"`
	DockerPullImage []dockerInfo `json:"dockerPullImage"`
}

func sendImageInfo(ctx *gin.Context) {
	var (
		imageInfo imageInfoStruct
		err       error
	)
	if err = ctx.BindJSON(&imageInfo); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, err)
		return
	}
	// 查到匹配的客户端列表
	val, ok := tcpServer.RegisteredMap[imageInfo.K8sName]
	b, _ := json.Marshal(val)
	fmt.Println("查询匹配到的客户端列表：", string(b))
	if !ok {
		fmt.Println("没有查询到对应的client连接地址")
		time.Sleep(time.Second)
	}

	// sync map
	//valSyncMap, ok := tcpServer.RegisteredSyncMap.Load(imageInfo.K8sName)
	//if !ok {
	//	fmt.Println("没有查询到对应的client连接地址")
	//	time.Sleep(time.Second)
	//}
	//发送指令给客户端
	// 普通map
	for _, conn := range val {
		b, _ := json.Marshal(imageInfo.DockerPullImage)
		go tcpServer.Send(conn, string(b))
	}

	//// sync map
	//if valSyncMap == nil {
	//	ctx.JSON(http.StatusOK, fmt.Sprintf("%s k8s名字对应的机器没有注册到server.", imageInfo.K8sName))
	//	return
	//}
	//valSyncMap.(*sync.Map).Range(func(k, conn interface{}) bool {
	//	b, _ := json.Marshal(imageInfo.DockerPullImage)
	//	go tcpServer.Send(conn.(net.Conn), string(b))
	//	return true
	//})

	ctx.JSON(http.StatusOK, "请求成功")
}

func getRegisteredList(ctx *gin.Context) {
	var (
		resMap  map[string][]string
		k8sName string
	)
	resMap = make(map[string][]string, 5)
	k8sName = ctx.Query("k8s_name")

	// 并发map
	//valSyncMap, ok := tcpServer.RegisteredSyncMap.Load(k8sName)
	//if !ok {
	//	fmt.Println("没有查询到对应的client连接地址")
	//	time.Sleep(time.Second)
	//}
	//valSyncMap.(*sync.Map).Range(func(k, conn interface{}) bool {
	//	resMap[k8sName] = append(resMap[k8sName], k.(string))
	//	return true
	//})

	// 普通map
	for k, _ := range tcpServer.RegisteredMap[k8sName] {
		resMap[k8sName] = append(resMap[k8sName], k)
	}
	ctx.JSON(http.StatusOK, resMap)
}

//
//// Client 单个 websocket 信息
//type wsClientViewSet struct {
//	views.GenericViewSet
//	FileName string
//	Socket   *websocket.Conn
//	Message  chan []byte
//}
//
//func (view wsClientViewSet) sendFile() {
//	var (
//		err     error
//		line    string
//		fileObj *os.File
//	)
//	filePath := fmt.Sprintf("/data/Software/mydan/CI/logs/build/%s", view.FileName)
//	fileObj, err = os.Open(filePath)
//	if err != nil {
//		fmt.Printf("open file faild, err:%v\n", err)
//		return
//	}
//	//关闭文件
//	defer fileObj.Close()
//	reader := bufio.NewReader(fileObj)
//	//msg := make([]byte, 512)
//	for {
//		line, err = reader.ReadString('\n') //按照字符’\n‘来分割每次读取长度
//		if err == io.EOF {
//			time.Sleep(50 * time.Millisecond)
//			continue
//		}
//		if err != nil {
//			fmt.Printf("read from file failed, err:%v\n", err)
//			return
//		}
//		if err = view.Socket.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
//			fmt.Printf("Send: %s\n", err)
//			view.Socket.Close()
//			break
//		}
//	}
//}
//
//func (view *wsClientViewSet) wsClient() {
//	var (
//		err      error
//		upGrader = websocket.Upgrader{
//			// cross origin domain
//			CheckOrigin: func(r *http.Request) bool {
//				return true
//			},
//			// 处理 Sec-WebSocket-Protocol Header
//			Subprotocols: []string{view.Ctx.GetHeader("Sec-WebSocket-Protocol")},
//		}
//	)
//
//	if view.Socket, err = upGrader.Upgrade(view.Ctx.Writer, view.Ctx.Request, nil); err != nil {
//		return
//	}
//	view.FileName = view.Ctx.Query("job_name")
//	fmt.Println(view.FileName)
//	view.sendFile()
//}
