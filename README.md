# 简介
> 本项目实现了平台操作k8s内业务镜像预加载的功能

# 架构
![image-20201126101318792](https://picgo-img.oss-cn-beijing.aliyuncs.com/md-img/2020-11-26/1606356798.png)

# 
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o docker-agent .  && docker build . -t mirror-registry.glodon.com/ptc/docker-agent:v3 && docker push mirror-registry.glodon.com/ptc/docker-agent:v3

```


# agent 打包
```
 docker build . -f Dockerfile-agent  -t mirror-registry.glodon.com/ptc/docker-agent:v3  && docker push mirror-registry.glodon.com/ptc/docker-agent:v3 
```
# server 打包
```
 docker build . -f Dockerfile-server -t mirror-registry.glodon.com/ptc/docker-server:v1 && docker push  mirror-registry.glodon.com/ptc/docker-server:v1   
```
