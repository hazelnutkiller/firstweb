FROM golang:alpine
#docker中的工作目錄 Docker 運作的資料夾位置，通常透過此指令指定 Server 的運作資料夾
WORKDIR $GOPATH/src/gin_docker
# 将当前目录同步到docker工作目录下，也可以只配置需要的目录和文件（配置目录、编译后的程序等）
ADD . ./
# 由于所周知的原因，某些包会出现下载超时。这里在docker里也使用go module的代理服务
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
#指定编译完成后的文件名，可以不设置使用默认的，最后一步要执行该文件名
RUN go build -o gin_docker .
EXPOSE 9999
# 这里跟编译完的文件名一致
ENTRYPOINT ["./gin_docker"]