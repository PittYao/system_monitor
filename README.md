# 服务器资源监控程序
## 测试接口
http_test.http
## ws连接实时获取数据测试
ws/wsClient.html

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