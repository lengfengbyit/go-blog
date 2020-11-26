# go-blog

### 编译
> 在打包的时候，将版本信息设置到生成的二进制文件中，方便日后追查
```bash
go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.gitCommitID=`git rev-parse HEAD` -X main.buildVersion=1.0.0"
```
- -ldflags -X 参数，可在打包的时候给变量赋值

### 安装运行
1. 增加配置文件并修改配置
```bash
cp configs/config_example.yaml configs/config.yaml
``` 

2. 运行
```bash
go run main.go
```
- 第一次运行会自动安装依赖包

3. 使用 go-bindata 生成 config.go 配置文件
- go build 不会吧 config.yaml 等非 .go 文件打包
- 使用 go-bindata 可以把配置文件生成 .go 文件，一起打包

```bash
go-bindata -o configs/config.go -pkg=configs configs/config.yaml
```
- -o 指定生成 .go 配置文件的输出路径
- -pkg 指定生成 .go 配置文件的包名

**读取配置**
```go
config := configs.MustAsset("configs/config.yaml")
```

**注意事项**
- 将第三方文件打包进二进制文件后，会增大二进制文件
- 无法做到文件的热更新和监听，必须要重新打包并重启服务才能更新内容

### 交叉编译
1. 编译成 linux 平台的可执行文件
```bash
CGO_ENABLE=0 GOOS=linux go build -a .
```
- CGO_ENABLE 是否开启cgo, 0关闭 1开启
- GOOS 目标操作系统， 如：linux, darwin, windows
- -a 强制重新编译依赖包