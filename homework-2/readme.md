## 说明

### 内容一
```
现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。

作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 kubernetes 集群，以下是你可以思考的维度

优雅启动
优雅终止
资源需求和 QoS 保证
探活
日常运维需求，日志等级
配置和代码分离
```

###内容二
```
模块八：课后作业（第二部分）

除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）

Service
Ingress
可以考虑的细节

如何确保整个应用的高可用
如何通过证书保证 httpServer 的通讯安全
```

###内容三
```txt
为 HTTPServer 添加 0-2 秒的随机延时
为 HTTPServer 项目添加延时 Metric
将 HTTPServer 部署至测试集群，并完成 Prometheus 配置
从 Promethus 界面中查询延时指标数据
（可选）创建一个 Grafana Dashboard 展现延时分配情况
```

## 实现内容

```
1. 优雅启动:
readiness，liveiness探针检测

2. 优雅终止：
检测信号 signal.SIGTERM signal.SIGINT
    拒绝新请求
    等待所有请求处理完毕
    释放资源
    关闭服务


3. 资源需求和Qos保证：
request，limited

4.探活
/healthz http路径检测状态

5. 日常运维需求，日志等级

6. 配置和代码分离
configmap
secret

使用service， ingress将服务发布给集群外的调用方：
 如何保证高可用
 如果通过证书保证httpserver通讯安全
 
```
 