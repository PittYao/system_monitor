# 服务器资源监控程序
- [x] Cpu
- [x] Memory
- [x] NetIO
- [x] Ip
- [x] Disk
## https http
- 开关https 
    - @/config/config.json中sever.ssl = [true 开启] [false 关闭] 
## 测试接口
http_test.http
## ws连接实时获取数据测试
ws/index.html

浏览器访问地址:
https://127.0.0.1:8083/static
or
http://127.0.0.1:8083/static
查看实时数据

连接地址 :
ws://localhost:8083/ws?uid=1&to_uid=2
参数1和2任意填，不填则是广播到所有用户

返回结果示例：
```json
{
  "bootTime":"5天21时3分21秒",
  "cpuInfo":{
    "cpuCount":8,
    "usePercent":[
      98.4375,
      96.875,
      100,
      100,
      95.38461538461539,
      93.75,
      95.3125,
      93.75
    ],
    "totalPercent":"96.69%",
    "mhz":"1.801GHz"
  },
  "memoryInfo":{
    "freeMemory":"241.21MB",
    "totalMemory":"7.89GB",
    "usedPercent":"97%"
  },
  "ip":"192.168.99.177",
  "netIoInfo":{
    "sentSpc":"481.18KB/S",
    "recvSpc":"22.62MB/S"
  }
}
```

## 磁盘信息
http://127.0.0.1:8083/disk
```json
{
  "code": 200,
  "message": "获取磁盘信息",
  "data": [
    {
      "device": "C:",
      "usedPercent": "88.13%",
      "total": "237.29GB",
      "used": "209.12GB",
      "free": "28.18GB"
    },
    {
      "device": "D:",
      "usedPercent": "60.34%",
      "total": "102.87GB",
      "used": "62.07GB",
      "free": "40.80GB"
    }
  ]
}
```