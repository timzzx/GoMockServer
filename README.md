# GoMockServer

**开发接口之前需要写wiki文档，然后写接口提供客户端联调，但是客户端不可能等接口开发完再调用，而是需要提前提供直接返回假数据的接口。**


**据此写了一个快速提供mock数据的http服务**

> + 下载 git clone https://github.com/timzzx/GoMockServer.git
> + 运行 go run main.go
> + 查看 http://localhost:12345

**使用方法**

> + 比如我们需要一个接口请求地址为 user/info
> + 首先直接访问 http://localhost:12345/user/info

数据返回

```json
{
errno: 200,
msg: "succ",
data: ""
}
```
**修改访问结果**

在data目录下会生成一个api_user_info.json的文本，直接修改即可。


**配合caddy可以快速搭建mock服务，方便个人或者小团队使用**


**缺点：**

> 不能在线修改，需要直接修改json文件，才能生效。


我的想法其实简单，需要有个快速提供mock api的服务，可以本地搭建，让客户端快速访问mock api，也没想过支持外网。如果大家有兴趣可自行修改。