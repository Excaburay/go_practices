# Golang learning practices   
_This is just for the beginner and someone who wanna review basic grammar in golang🤣. To get more parctices,you can go to the official trainning ground [**the tour of go**](https://tour.golang.org/welcome/1)_



## environment   
use __go 1.3__ and **go mod**,so it has the go.mod file in project. To run this example, you can execute  
```
go build
```

-----

## slice map JSON
There are some basic operations about slice and map in **learntest** folder.    

|  part | desc |
|  :--- | :--- |
|  map  |  Map store data by key-value which can search faster than linear search. Key should support == or != operation   |
| slice | Slice is not array but point to array. Once the slice change can impact the array. |
|  JSON |JSON(JavaScript Object Notation) is easy transform to struct in golang|

***

## go test

测试文件名以 _test.go为后缀

常用的测试方法有：testing.T 和 testing.B

```go
const ts := []string{"abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz", "abcd", "efg", "hij", "klm", "nop", "lmn", "ovw", "xyz"}

const res:= "abcd efg hij klm nop lmn ovw xyz abcd efg hij klm nop lmn ovw xyz"

func TestStringAddFunc(t *testing.T) {
	
	if s := StringAddFunc(ts); s != res  {
		t.Errorf("StringAddFunc is not %s/n", s)
	}

}

func BenchmarkStringJoin(b *testing.B) {
	b.N = 10000000 //可以不指定b.N大小，由程序自动指定
	for index := 0; index < b.N; index++ {
		StringJoin(ts)
	}
}
```

* 使用testing.T的函数名须以Test为前缀
* 使用testing.B的函数名须以Benchmark为前缀
* go test -v 执行当前文件中的所有测试方法
* go test -bench="." 执行Benchmark测试方法
* 执行结果如下图所示，同时表明了strings.Join方法相比拼接字符串效率更高

![image-20200125171137122](C:\Users\shuim\AppData\Roaming\Typora\typora-user-images\image-20200125171137122.png)



### go 网络编程

go network 文件夹中实现了基础的客户端与服务端示例，文件夹中有说明文档。

* http1——http3 使用了net/http中的方法构造http网络连接
* listener 使用了net中基础方法构造了socket网络连接
* websocket 使用第三方包github.com/gorilla/websocket构造