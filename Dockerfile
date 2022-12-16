FROM golang:latest
WORKDIR $GOPATH/src/iss
COPY . .
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN go build -o goapp
EXPOSE 80
ENTRYPOINT ["./goapp"]

