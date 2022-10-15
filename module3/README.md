### 1、先编译go程序

`GOOS=linux GOARCH=amd64 go build -o ./bin/amd64/httpserver`

### 2、 创建Dockerfile

 `FROM ubuntu
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
ENTRYPOINT /httpserver`

### 3、执行Dockerfile

` sudo docker build -t htttpserver .`

![./1.png](./1.png)

### 4、登陆docker login

`docker login -u xxxx`

### 5、创建tag标签 docker tag httpserver xxxx/httpserver:v1

### 6、推送镜像docker push xxxx/httpserver:v1

![2](./2.png)

