1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
```
redis-benchmark -d 10 -t get,set
====== SET ======
  100000 requests completed in 1.91 seconds
====== GET ======
  100000 requests completed in 1.93 seconds

redis-benchmark -d 20 -t get,set
====== SET ======
  100000 requests completed in 1.88 seconds
====== GET ======
  100000 requests completed in 1.88 seconds

redis-benchmark -d 50 -t get,set
====== SET ======
  100000 requests completed in 1.87 seconds
====== GET ======
  100000 requests completed in 1.90 seconds

redis-benchmark -d 100 -t get,set
====== SET ======
  100000 requests completed in 1.87 seconds
====== GET ======
  100000 requests completed in 1.83 seconds

redis-benchmark -d 100 -t get,set
====== SET ======
  100000 requests completed in 1.85 seconds
====== GET ======
  100000 requests completed in 1.76 seconds

redis-benchmark -d 1000 -t get,set
====== SET ======
  100000 requests completed in 1.85 seconds
====== GET ======
  100000 requests completed in 1.88 seconds

redis-benchmark -d 5000 -t get,set
====== SET ======
  100000 requests completed in 2.10 seconds
====== GET ======
  100000 requests completed in 2.02 seconds
```
2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
5k * 500000	31.948 MB
1k:* 500000	10.490 MB
10bit * 500000	5.245 MB


