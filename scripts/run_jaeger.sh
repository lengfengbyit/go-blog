docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:1.16


# 端口           协议        功能
# 5775          UDP         以 compact 协议接收 zipkin.thrift 数据
# 6831          UDP         以 compact 协议接收 jaeger.thrift 数据
# 6832          UDP         以 binary 协议接收 jaeger.thrift 数据
# 5778          HTTP        Jaeger 的服务配置端口
# 16686         HTTP        Jaeger 的 Web UI
# 14268         HTTP        通过 Client 直接接收 jaeger.thrift 数据
# 9411          HTTP        兼容 Zipkin 的HTTP端口