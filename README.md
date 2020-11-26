# go-blog

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

