# API

## gin

| api         | description                                       |
| ----------- | ------------------------------------------------- |
| Default,New | 返回路由Engine, Default包含Logger和Recovery中间件 |
| H           | map[string]interface{}                            |
| Recovery    | 捕获任何panic                                     |

## gin.Engine

| api                        | description                                                  |
| -------------------------- | ------------------------------------------------------------ |
| GET,POST                   | 路由callback                                                 |
| Run                        | 等价于http.ListenAndServe(addr, router), 监听端口并attach router到server |
| LoadHTMLGlob,LoadHTMLFiles | 加载html文件                                                 |
| SetHTMLTemplate            | 设置html.template.Template                                   |
| Static, StaticFS           | 静态资源, httpFileSystem                                     |
| BasicAuth                  | 基本的认证                                                   |
| Group                      | 具有相同前缀的路由组                                         |
| RunTLS                     | 用pem, key文件做https                                        |
| Use                        | 全局中间件                                                   |
| Delims                     | 设置模板的左右分割符, eg:{[{, }]}                            |
| SetFuncMap                 | s设置模板用到的函数, eg:Date: {[{.now \| formatAsDate}]}     |



## gin.Context

| api                           | description                                               |
| ----------------------------- | --------------------------------------------------------- |
| Next                          | 执行后面的路由handler, 在中间件中使用                     |
| String,Json,HTML,Data         | 对请求返回String,Json,html文件,[]byte                     |
| Bind,ShouldBind,ShouldBindUri | (是否能)绑定到数据源, eg:结构体                           |
| SaveUploadedFile              | 保存上传的文件                                            |
| Param                         | 路径里的,"/user/:id", Param("id")                         |
| Query                         | 问号后的, /path?id=1234&name=Manu&value=, Query("id")     |
| PostForm                      | 参数是表单里的name, 返回值是表单里                        |
| Stream                        | 发送流, 直到返回false                                     |
| SSEvent                       | send server-sent event                                    |
| ClientIP                      | 客户端ip                                                  |
| Redirect                      | 重定向                                                    |
| Header                        | c.Writer.Header().Set(key, value), 如果value=""等价于删除 |

## gin.Context.Writer

| api         | description  |
| ----------- | ------------ |
| CloseNotify | 连接关闭通知 |



## http.Pusher

| api  | description |
| ---- | ----------- |
| Push | 发送js      |

## http.RoundTriper

| api       | description          |
| --------- | -------------------- |
| RoundTrip | 执行一次请求获得返回 |

## http.Server

| api            | description                          |
| -------------- | ------------------------------------ |
| Shutdown       | shutdown gracefully                  |
| ListenAndServe | 相当于router.Run,http.ListenAndServe |

## http.Response

| api        | description         |
| ---------- | ------------------- |
| Body.Close | shutdown gracefully |


## html.template

| api            | description      |
| -------------- | ---------------- |
| Parse, ParseFS | 解析模板         |
| New            | 新对象           |
| Must           | success or panic |

# 结构体tag

| tag         | description                            |
| ----------- | -------------------------------------- |
| form        | 传输的字段名                           |
| time_format | time.Time的传输格式                    |
| time_utc    | 1:utc                                  |
| uri         | 用于路由模式匹配, eg:`uri:"id"`匹配:id |
| binding     | 绑定属性.eg: `binding "required,uuid"` |

# web基础
## url
```
scheme://host[:port#]/path/.../[?query-string][#anchor]
scheme         指定低层使用的协议(例如：http, https, ftp)
host           HTTP服务器的IP地址或者域名
port#          HTTP服务器的默认端口是80，这种情况下端口号可以省略。如果使用了别的端口，必须指明，例如 http://www.cnblogs.com:8080/
path           访问资源的路径
query-string   发送给http服务器的数据
anchor         锚
```
## http

请求

```
GET /domains/example/ HTTP/1.1      //请求行: 请求方法 请求URI HTTP协议/协议版本
Host：www.iana.org             //服务端的主机名
User-Agent：Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.4 (KHTML, like Gecko) Chrome/22.0.1229.94 Safari/537.4            //浏览器信息
Accept：text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8    //客户端能接收的mine
Accept-Encoding：gzip,deflate,sdch     //是否支持流压缩
Accept-Charset：UTF-8,*;q=0.5      //客户端字符编码集
//空行,用于分割请求头和消息体
//消息体,请求资源参数,例如POST传递的参数
```
响应
```
HTTP/1.1 200 OK                     //状态行
Server: nginx/1.0.8                 //服务器使用的WEB软件名及版本
Date:Date: Tue, 30 Oct 2012 04:14:25 GMT        //发送时间
Content-Type: text/html             //服务器发送信息的类型
Transfer-Encoding: chunked          //表示发送HTTP包是分段发的
Connection: keep-alive              //保持连接状态
Content-Length: 90                  //主体内容长度
//空行 用来分割消息头和主体
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"... //消息体

1XX 提示信息 - 表示请求已被成功接收，继续处理
2XX 成功 - 表示请求已被成功接收，理解，接受
3XX 重定向 - 要完成请求必须进行更进一步的处理
4XX 客户端错误 - 请求有语法错误或请求无法实现
5XX 服务器端错误 - 服务器未能实现合法的请求
```
GET/POST区别

* GET请求消息体为空，POST请求带有消息体
* GET提交的数据会放在URL之后，以`?`分割URL和传输数据，参数之间以`&`相连，如`EditPosts.aspx?name=test1&id=123456`。POST方法是把提交的数据放在HTTP包的body中
* GET提交的数据大小有限制（因为浏览器对URL的长度有限制），而POST方法提交的数据没有限制
* GET方式提交数据，会带来安全问题，比如一个登录页面，通过GET方式提交数据时，用户名和密码将出现在URL上，如果页面可以被缓存或者其他人可以访问这台机器，就可以从历史记录获得该用户的账号和密码


# other

* form里`post`的`action`就是路由

* grpc的应用, 就是在路由回调里调一下rpc, 返回调用结果

* curl -X方法 -v详细 -d data

* openssl 生成安全文件,

	* openssl genrsa -out ./testdata/server.key 2048, Generate RSA private key
	
	* openssl req -new -x509 -key ./testdata/server.key -out ./testdata/server.pem -days 365,Generate digital certificate

