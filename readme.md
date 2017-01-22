# QRTC
Qrtc 是连麦服务端工具，可以更方便地创建，查看，删除连麦房间以及生成RoomToken
文档详情： [Server 连麦](http://developer.qiniu.com/article/pili/sdk/server-rtc-sdk.html)

## 下载
[qrtc](https://github.com/qiniudemo/qrtc/raw/master/qrtc.zip)

## Usage
### 配置
使用本工具首先需要配置`conf.json`文件，填写相应的七牛ak/sk
当然你也可以不设置，进入工具的页面后，如果检查到没有设置ak/sk，会进入一个登录页面来设置ak/sk。
> 工具不会对ak/sk正确性进行校验

### 使用
首先到```qrtc```所在目录，给`qrtc`可执行权限：`chmod +x qrtc`，然后通过命令行运行本工具。
然后访问 localhost:8080 即可进入本工具
windows用户双击 qrtc-win 即可使用


## API
获取连麦token：
```
POST /room/<RoomName>/user/<UserID>/token
Host: <localhost>
```

> e.g.
> curl -X POST 127.0.0.1:8080/stream/aaa

 
RoomName: 房间名称。
UserID: 请求加入房间的用户ID。
该API生成权限(perm)为`user`,有效时间为永久的RoomToken
相关文档：[RoomToken计算](http://developer.qiniu.com/article/pili/sdk/server-rtc-sdk.html#6)


**连麦房间名(room_name)和流名(streamKey)不需要一致，以下两个API把room_name作为streamKey获取推流和播放地址，只是为了方便测试 :)**

获取推流地址：
```
POST /stream/<RoomName>
Host: <localhost>
```

> e.g.
> curl -X POST 127.0.0.1:8080/stream/haha19/play

RoomName: 房间名称。


获取播放地址：
```
GET /stream/<RoomName>/play
Host: <localhos>
```

> e.g.
> curl 127.0.0.1:8080/stream/haha19/play

RoomName: 房间名称。

