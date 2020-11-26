package dockerCore

import (
	"bufio"
	"context"
	client2 "image-pull/client"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerInfo struct {
	ImageName, UserName, Password string
}

func (d DockerInfo) PullImage(n int, nodeInfo string) (err error) {
	var (
		encodedJSON []byte
		cli         *client.Client
	)
	ctx := context.Background()
	if cli, err = client2.GetClient(); err != nil {
		return fmt.Errorf("th %d, get docker client-01 err,err:%s \n", n, err.Error())
	}
	if encodedJSON, err = json.Marshal(types.AuthConfig{
		Username: d.UserName,
		Password: d.Password,
	}); err != nil {
		return fmt.Errorf("th %d, image auth err,err:%s \n", n, err.Error())
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("开始下载 %s \n", d.ImageName)
		out, err := cli.ImagePull(
			ctx,
			d.ImageName,
			types.ImagePullOptions{
				RegistryAuth: base64.URLEncoding.EncodeToString(encodedJSON),
			})
		if err != nil {
			err = fmt.Errorf("th %d, retry %d,image pull err,err:%s \n", n, i, err.Error())
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}
		defer out.Close()

		reader := bufio.NewReader(out)
		for {
			line, err := reader.ReadString('\n') //按照字符’\n‘来分割每次读取长度
			if err == io.EOF {
				fmt.Println("读完了")
				break
			}
			if err != nil {
				fmt.Printf("read from dockerpull Out failed, err:%v\n", err)
				continue
			}

			if err = client2.TcpSend(client2.TcpClient, fmt.Sprintf("%s-%s", nodeInfo, line)); err != nil {
				break
			}
		}
		break
	}
	return
}
