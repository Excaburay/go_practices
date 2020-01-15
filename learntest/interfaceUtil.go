package learntest

import "fmt"

//io.write是广泛使用的接口，它负责所有可以写入字节的类型的抽象，包括文件、
//内存缓冲区、网络连接、HTTP客户端、打包器、散列器等。除此之外还有io.Reader io.Closer

type Celsius struct {
	Name string
	Age  string
}

type Ce int

func (c Ce) String() string {
	return fmt.Sprintf("%d 先生", c)
}

func (c *Celsius) String() string {
	return fmt.Sprintf("%s 先生", c.Name)
}
