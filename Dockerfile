FROM node:12-alpine3.11 AS front-end
COPY front-end /opt/app/front-end
WORKDIR /opt/app/front-end
RUN npm install --registry=https://registry.npm.taobao.org && npm run build:prod

FROM golang:1.14-alpine
ENV GOPROXY=https://goproxy.cn,direct \
    APP_NAME=gin-vue-element-ui

COPY . /opt/app
COPY --from=front-end /opt/app/back-end/templates /opt/app/back-end/templates

WORKDIR /opt/app/back-end
RUN go build  -o ${GOPATH}/bin/${FS_APP_NAME}

EXPOSE 8080
ENTRYPOINT ["gin-vue-element-ui"]
