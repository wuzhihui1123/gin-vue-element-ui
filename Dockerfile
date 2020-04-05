FROM reg.firstshare.cn/base/fs-golang:14

#front-end


# back-end
ENV GOPROXY https://goproxy.cn,direct
COPY . /opt/app
WORKDIR /opt/app/back-end/go-gin-example
RUN go build .

EXPOSE 8080 80
ENTRYPOINT ["./go-gin-example"]
