# audit-rectification

# Linux环境部署
## 1、编译项目
```shell
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/audit cmd/server/server.go
```
## 2、后台运行
```shell
nohup bin/audit >> out.file 2>&1 &
```

## 修改文件重编译后运行(重编译的路径和文件与旧版本保持一致)
```shell
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o bin/audit cmd/server/server.go
kill -1 oldPID
nohup bin/audit >> out.file 2>&1 &
```