1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。
```text
fix length
长度协议，每次发送固定缓冲区大小的数据
使用固定长度的粘包和拆包场景，可以使用FixedLengthFrameDecoder，该解码一器会每次读取固定长度的消息，如果当前读取到的消息不足指定长度，那么就会等待下一个消息到达后进行补足。

delimiter based
分隔符协议，基于分隔符区分每一个请求发送的数据
分隔符协议需要 Server 不停的检测，耗费性能

length field based frame decoder
TCP 协议头里面写入每次发送请求的长度。 客户端在协议头里面带入数据长度，服务器在接收到请求后，根据协议头里面的数据长度来决定接受多少数据,只有在读取到足够长度的消息之后才算是读到了一个完整的消息。之后会按照参数指定的包长度偏移量数据对接收到的数据进行解码，从而得到目标消息体数据。
```


2. 实现一个从 socket connection 中解码出 goim 协议的解码器。
