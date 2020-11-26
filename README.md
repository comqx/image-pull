# 简介
> 本项目主要实现了k8s镜像预加载功能，提供api进行触发，或者可以将功能集成到平台上面实现镜像的预加载。缩短发版时间

# 架构
![image-20201126101318792](https://picgo-img.oss-cn-beijing.aliyuncs.com/md-img/2020-11-26/1606356798.png)


# 后续功能
- 提供web页面
- 通过websocket方式获取每个机器拉取镜像的详细信息

# api简介
```
# 发送要拉取的镜像信息
http://127.0.0.1:8080/api/sendImage POST

# curl方式
curl --location --request POST 'http://127.0.0.1:8080/api/sendImage' \
--header 'Content-Type: application/json' \
--data-raw '{"k8sName":"local",
"dockerPullImage":[
    {"userName":"","passWord":"","imageName":"docker.io/library/busybox"}
    ]
}'

# 获取已经注册的机器信息
http://127.0.0.1:8080/api/getRegisteredList  GET

curl方式：
curl http://127.0.0.1:8080/api/getRegisteredList&k8s_name=local
```

# agent 打包
```
 docker build . -f Dockerfile-agent  -t mirror-registry.xxx.com/ptc/docker-agent:v3  && docker push mirror-registry.xxx.com/ptc/docker-agent:v3 
```
# server 打包
```
 docker build . -f Dockerfile-server -t mirror-registry.xxx.com/ptc/docker-server:v1 && docker push  mirror-registry.xxx.com/ptc/docker-server:v1   
```
