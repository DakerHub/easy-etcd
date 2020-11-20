# Debian 10
FROM golang:1.15

RUN mkdir -p /var/www/html

# 安装nodejs
RUN apt-get -yqq update
RUN curl -sL https://deb.nodesource.com/setup_15.x | bash -
RUN apt-get install -y nodejs

# 构建go程序
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o /usr/local/bin


# 构建静态UI文件
WORKDIR /go/src/app/www
RUN npm install
RUN npm run build
# Go程序中指定的静态资源目录
RUN mv ./dist/* /var/www/html

ENTRYPOINT [ "/usr/local/bin/easy-etcd" ]
