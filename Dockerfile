# Debian 10
FROM golang:1.15-alpine

RUN mkdir -p /data/backup

ENV EASY_ETCD_BACKUP_DIR=/data/backup

RUN mkdir -p /var/www/html

# 安装nodejs
RUN apk add --update nodejs npm

# 淘宝镜像
RUN npm install -g cnpm --registry=https://registry.npm.taobao.org

# 构建go程序
WORKDIR /go/src/app
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go build -o /usr/local/bin


# 构建静态UI文件
WORKDIR /go/src/app/www
RUN cnpm install
RUN npm run build
# Go程序中指定的静态资源目录
RUN mv ./dist/* /var/www/html

ENTRYPOINT [ "/usr/local/bin/easy-etcd" ]
