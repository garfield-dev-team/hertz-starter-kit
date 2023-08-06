# 指定基础镜像的版本，确保每次构建都是幂等的
FROM golang:1.20.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first, because of caching reasons.
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . ./
# Compile project
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .
RUN ./build.sh

# 使用体积更小的基础镜像
# FROM alpine:3.15 AS final
# Golang 项目推荐 scratch 镜像进一步减小体积
# 需要注意 scratch 存在一些问题，比如没有 apk 命令，也不像 alpine 镜像内置 tzdata
# 因此如果用 scratch 镜像，需要在 Golang 应用内部做时区配置
FROM alpine:3.15 as final

# 不要使用 root 权限运行应用
RUN addgroup -g 10000 hertzapp
RUN adduser -D -u 10000 -G hertzapp hertzapp

USER hertzapp

# 如果需要是用 TZ 环境变量 实现时区控制，需要安装 tzdata 这个包
# debian 的基础镜像默认情况下已经安装了 tzdata，而 ubuntu 并没有
# RUN apk add --no-cache tzdata

# 设置时区
# 在使用 Docker 容器时，系统默认的时区就是 UTC 时间（0 时区），和我们实际需要的北京时间相差八个小时
ENV LANG=en_US.UTF-8 LANGUAGE=en_US:en LC_ALL=en_US.UTF-8 TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=builder /app/output .

# 默认暴露 80 端口
EXPOSE 8080

CMD ["./output/bootstrap.sh"]
