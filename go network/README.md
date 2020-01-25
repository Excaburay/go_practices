# Golang 网络编程

go语言提供的**net**包中提供了丰富简洁的网络API

***

### 基础网络方法

#### net.Listen

```go
listener, err := net.Listen("tcp","localhost:8000")
```

* 网络类型必须是 *tcp、tcp4、tcp6、unix、unixpacket*
* listener接口方法
  * Accept()	接收等待并返回下一个连接(conn)给listener
  * Close()     关闭listener, 任何阻塞的Accept操作将结束阻塞并返回错误
  * Addr()     返回listener的网络地址

#### net.Dial

```go
conn, err := net.Dial("tcp", "localhost:8000")
```

* 连接网络类型为 *tcp、tcp4、tcp6、udp、udp4、udp6、ip4、ip6、unix*
* 返回连接conn

#### net.Conn

```go
type Conn interface {
   Read(b []byte) (n int, err error)
   Write(b []byte) (n int, err error)
   Close() error
   LocalAddr() Addr
   RemoteAddr() Addr
   SetDeadline(t time.Time) error
   SetReadDeadline(t time.Time) error
   SetWriteDeadline(t time.Time) error
}
```

*Conn*接口类型，实现了*Read、Write、Close、LocalAddr、RemoteAddr*等方法

可以通过*io.WriteString*方法向*Conn*中写入数据。

可以通过*io.Copy*方法将*Conn*中数据写入指定文件或者终端（*os.Stdout*）



### HTTP网络连接

***

#### http.ListenAndServe

```go
http.ListenAndServe("0.0.0.0:8001", nil)
```

* 监听TCP网络地址并调用
* 第二个参数为 *nil* 时，使用默认的 *DefaultServeMux*

#### http.HandleFunc

```go
http.HandleFunc("/httpTest1", handler1)
```

* 与ListenAndServe配套使用

* 第一个参数是url, 第二个参数是handlerfunction，绑定函数需要两个参数满足ResponseWriter接口和Request结构体，如下图所示

  ```go
  func handler1(w http.ResponseWriter, r *http.Request) {
     io.WriteString(w, "ok")
  }
  ```

#### http.Post

```go
req, err := http.Post(url, "application/json", nil)
```

* 调用者需要在读取数据后关闭resp.Body

* 是对于DefaultClient.Post的wrapper

  ```go
  func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
     return DefaultClient.Post(url, contentType, body)
  }
  ```

* 如果要自定义请求头，使用NewRequest 和 DefaultClent.Do

* 还有 http.Get 等API

#### http.NewRequest

```go
reqest, err := http.NewRequest("GET", "http://localhost:8000/Test1",nil)
```

* NewRequest wraps NewRequestWithContext using the background context

#### http.Client{}.Do()

* Do 发送HTTP请求并返回一个HTTP回应

