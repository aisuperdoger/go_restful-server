FROM golang  

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

# RUN  mkdir /go_restful-server-main/go_client
ADD ./go_client /go_client
WORKDIR /go_client

# RUN cd /go_restful-server-main/go_client
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build
EXPOSE 8080
# RUN pwd
CMD ["./go_client"]
# 宿主机运行命令
# git clone https://github.com/tjumcw/6.824.git
# docker run -it --name=6.824 -v /Users/miaochangwei1/Desktop/6.824:/6.824 6.824