FROM golang:1.14-alpine

COPY . /opt/app

#front-end
WORKDIR /opt/app/front-end
RUN npm run build:prod


# back-end
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /opt/app/back-end
RUN go build .

EXPOSE 8080
ENTRYPOINT ["gin-vue-element-ui"]
