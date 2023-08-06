
#!/bin/bash
RUN_NAME="hertz_app"
# 获取版本号
VERSION="v1.0"
# 获取Git commit
COMMIT=$(git rev-parse HEAD)
# 获取分支信息
BRANCH=$(git rev-parse --abbrev-ref HEAD)
# 获取构建时间
BUILD_TIME=$(date +'%Y-%m-%d-%H:%M:%S')

mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
# CI 环境推荐传递 `-a` 参数，禁用编译缓存，从零开始编译所有的依赖包
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT -X main.buildBranch=$BRANCH -X main.buildTime=$BUILD_TIME" \
-a -o output/bin/${RUN_NAME} .
