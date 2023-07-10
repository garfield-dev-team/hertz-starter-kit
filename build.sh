
#!/bin/bash
RUN_NAME="hertz_app"
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
# CI 环境推荐传递 `-a` 参数，禁用编译缓存，从零开始编译所有的依赖包
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o output/bin/${RUN_NAME} .
