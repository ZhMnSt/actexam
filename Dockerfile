FROM golang as test

WORKDIR /tmp/go_test

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

COPY go/ /tmp/go_test/

EXPOSE 8888

RUN go mod init gin
RUN go mod tidy
RUN go test ./
RUN go build gin.go

FROM alpine

COPY --from=test /tmp/go_test/gin /usr/

EXPOSE 8888

CMD ["/usr/gin"]
