# 日志收集项目架构设计及Kafka介绍

## 项目背景

日志信息解决问题、数据挖掘、浏览推荐定位

每个业务系统都有日志，当系统出现问题时，需要通过日志信息来定位和解决问题。当系统机器比较少时，登录到服务器上查看即可满足。当系统机器规模巨大，登录到机器上查看几乎不现实（分布式系统，一个系统部署在十几台机器上）

## 解决方案

日志收集-->**存储**-->**建立索引**去获取-->**找到**日志记录-->通过web页面**展示检索**

把机器上的日志实时收集，统一存储到中心系统。再对这些日志建立索引，通过搜索即可快速找到对应的日志记录。通过提供一个界面友好的web页面实现日志展示与检索。

## 面临的问题

实时日志量非常大，每天处理几十亿条。日志准实时收集，延迟控制在分钟级别。系统的架构设计能够支持水平扩展。

## 业界方案

ELK

![image-20230403091122826](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403091122826.png)

filebit从AppServer收集信息给LogstashAgent

## ELK方案的问题

- 运维成本高，每增加一个日志收集项，都需要手动修改配置。
- **监控缺失**，无法准确获取logstash的状态。
- 无法做到定制化开发与维护

# 日志收集系统架构设计

![image-20230403091601178](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403091601178.png)

Log Agent具体部署到每一台业务的服务器上，根据配置项收集日志，配置项是存储到etcd中的。

etcd中写一个web管理页面，管理所有的配置项。

Log Agent将每一个项目的日志收集后存到Kafka中，Kafka既能作为消息队列，也能作为日志存储。

Elastic Search可以检索，Kibana作为web工具

Sys Agent 做系统监控，例如cpu负载，IO，网络速率

## 组件介绍

LoaAgent：日志收集客户端，用来收集服务器上的日志。Kafka：高吞吐量的分布式队列（Linkin开发，apache顶级开源项目）ElasticSearch：开源的搜素引擎，提供基于HTTP RESTful的web接口。

Kibaa：开源的ES数据分析和可视化工具。Hadoop：分布式计算框架，能够对大量数据进行分布式处理的平台。

Strom：一个免费并且开源的分布式实时计算系统。

## *将学到的技能

- 服务端agent开发

- 后端服务组件开发

- etcd的使用

- Kafka和zookeeper的使用

- ES和Kibana的使用


## 消息队列的通信模型

NSQ：将同步消息队列转化为异步的，可以作为进程间的通信(例如：go语言程序和php之间的程序做通信)。

点对点模式：

消息生产者生产消息发送到queue中，然后消息消费者从queue中取出并且消费消息。一条消息被消费以后，queue中就没有了，不存在重复消费。

## 发布/订阅（topic）

消息生产者（发布）将消息发布到topic中，同时有多个消息消费者（订阅）消费该消息。和点对点方式不同，发布到topic的消息会被所有订阅者消费（类似于关注了微信公众号的人都能收到推送的文章）。

补充：发布订阅模式下，当发布者消息量很大时，显然单个订阅者的处理能力是不足的。实际上现实场景中是多个订阅者节点组成一个订阅组负载均衡消费topic消息即分组订阅，这样订阅者很容易实现消费能力线性扩展。可以看成是一个topic下有多个Queue,每个Queue是点对点的方式，Queue之间是发布订阅方式。

# Kafka

Apache Kafka由著名职业社交公司Linkedln开发，最初是被设计用来解决LinkedIn公司内部海量日志传输等问题。Kafka使用Scala语言编写，2011年开源并进入Apache孵化器，2012年10月正式毕业，现在为Apache顶级项目

## 介绍

==Kafka是一个**分布式**（集群）数据流（可以消费，也可以陈列）平台==，可以运行在单台服务器上，也可以在多台服务器上部署形成集群。它提供了发布和订阅功能，使用者可以**发送数据到Kafka中**，也可以**从Kafka中读取数据**（以便进行后续的处理）。Kafka具有高吞吐、低延迟、高容错等特点。

<img src="C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403094603816.png" alt="image-20230403094603816" style="zoom:80%;" />

生产者把消息丢到管道里，做分区后消费者拿出

<img src="C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403094623766.png" alt="image-20230403094623766" style="zoom:80%;" />

- Producer：生产者，消息的产生者，是消息的入口。数据扔到Cluster中，Consumer然后从中获取

- Kafka Cluster：kafka集群，**一台或者多台==服务器==组成**。Broker被称为节点，三台机器（代理）

  + Broker：Broker是指部署了Kafka实例的服务器节点。每个服务器上有一个或多个kafka的实列，我们姑且认为每个broker对应一台服务器。每个Kafka集群内的broker都有一个不重复的编号，如图中broker-0、broker-1等...

  - ​		Topic：**消息的主题** ，什么类型的日志消息？可以理解为消息的分区，kafka的数据就保存在同批次。每个broker上都可以创建多个同批次。实际应用中通常是一个业务线建立一个topic

  - ​		Partition：Topic的分区，每个topic可以有多个分区，分区的作用是做负载，提高kafka的吞吐量。同一个topic**在不同的分区数据是不重复**的，所有的相同的partition0的分区数据是一样的，partition的表现形式就是一个一个的文件夹。（partition0和partition1不一样）

  - ​		Replication：每个分区都有多个副本，副本的作用是做备胎。当主分区（Leader）故障的时候会选择一个（Follower）尚未，称为Leader。在kafka中默认副本的最大数量是10个，且副本的数量不能大于Broker的数量，follower和leader在不同的机器中，同一个机器对同一个分区只能存放一个副本（包括自己）

- Consumer：消费者，即消息的消费方，是消息的出口。
  - ​		Consumer Group：我们可以将多个消费组组成一个消费者组，在kafka的设计中，同一个分区的数据只能被消费者组中的某一个消费者消费。同一个消费者组的消费者可以消费同一个topic的不同分区的数据，这也是为了提高kafka的吞吐量。

## 工作流程

![image-20230403095351391](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403095351391.png)

把消息给leader，查看leader是谁

## kafka选择partition的原则

在kafka中，如果topic有多个partition，producer将数据发送到哪个partition中呢？

1. ==**指定**==：partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition
2. ==**hash**==：如果没有指定的partition，但是设置了数据的key，则会根据key的值hash出一个partition
3. ==**轮询**==：如果既没有指定的partition，又没有设置key，则会用轮询方式，即每次取出一小段时间的数据写入某个partition，下一段的时间写入下一个partition

## ACK应答机制（生产者往kafka发送数据的模式）

producer在向kafka 写入消息的时候，可以设置参数来确定是否确认kafka接收到数据

- 0 代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效率最高。

- 1 代表producer往集群发送数据只要leader应答就可以发送下一条，只确保leader发送成功。

- all ：把数据发送给leader，确保follower从leader拉取数据回复ack给leader，leader再回复ack

  代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，确保leader发送成功和所有的副本都完成备份。安全性最高，但是效率最低。

注意：如果往不存在的topic里写数据，kafka会自动创建topic，partition和replication的数量默认配置都是1。

## Topic和数据日志

```topic```是同一类别的消息记录record的集合。在Kafka中，一个主题通常有多个订阅者。对于每个主题，Kafka集群维护了一个分区数据日志文件结构如下：

![image-20230403102620825](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403102620825.png)

![image-20230403102838082](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403102838082.png)

![image-20230403102935219](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403102935219.png)

![image-20230403103048304](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403103048304.png)

一个组内的两个消费者不能从同一个partition（P0-P3）获取数据。

![image-20230403103455941](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403103455941.png)

为什么Kafka比较快？

将随机读转化为顺序读，记录了每个索引在文件里的位置。

![image-20230403103650083](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403103650083.png)

![image-20230403103705245](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403103705245.png)

![image-20230403103804907](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403103804907.png)

![image-20230403103820348](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403103820348.png)

bin 可执行文件

config配置文件

kafkalogs数据文件

# windows下kafka安装启动以及使用

[(48条消息) windows下kafka安装启动以及使用_topEngineerr的博客-CSDN博客](https://blog.csdn.net/topdeveloperr/article/details/105676361)

命令行报错：弄到一个新的没有额外文件夹的路径下

zookeeper.propertises  日志存在哪个目录下

missing go.sum entry解决方法：[(48条消息) Go构建项目的时候，解决missing go.sum entry for module providing package ＜package_name＞_model/router.go:6:2: missing go.sum entry for modu_junlz0413的博客-CSDN博客](https://blog.csdn.net/tmis_mysql/article/details/116780817)

```bash
go build -mod=mod
```

kafka日志存入F:\goland\go_project\kafka\tmp\kafkalogs

kafka文件F:\kafka_2.13-3.4.0

cmd启用：

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\zookeeper-server-start.bat config\zookeeper.properties
```

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\kafka-server-start.bat config\server.properties
```

![image-20230410160904121](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230410160904121.png)

![image-20230403145322430](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403145322430.png)



![image-20230403091601178](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403091601178.png)



## LogAgent的工作流程：

### 1、读日志(有新增的则读，否则等待) --tail包

```bash

go get github.com/hpcloud/tail

set GO111MODULE=on

SET GOPROXY=http://goproxy.cn
cls
go mod init
go get github.com/hpcloud/tail
```

tail_demo

```go
import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,                                 //轮询文件更改
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}

		fmt.Println("line:", line.Text)
	}
}

```

### 2、往kafka里写日志  --sarama第三方库

初始化go.mod

```bash
SET GO111MODULE=on
SET GOPROXY=http://goproxy.cn
//go mod init
go mod init example.com/m
//PS F:\goland\go_project\21weeks\21weeks_go\project1\kafka_demo> go mod init example.com/m
```

go.mod

```go
require github.com/Shopify/sarama v1.19.0
```

```bash
go get
//或者 go mod download
```

kafka_demo

```go
package main

import (
   "fmt"
   "github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client
func main() {
   config := sarama.NewConfig()
   config.Producer.RequiredAcks = sarama.WaitForAll          //发送完整数据类型leader和follow都确认
   config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
   config.Producer.Return.Successes = true                   //成功交付的消息将在success channel返回
   //构造一个消息
   msg := &sarama.ProducerMessage{}
   msg.Topic = "web_log"
   msg.Value = sarama.StringEncoder("this is a test log")
   //连接kafka
   client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
   if err != nil {
      fmt.Println("prodicer closed,err:", err)
      return
   }
   fmt.Println("连接kafka成功")
   defer client.Close()
   //发送消息
   pid, offset, err := client.SendMessage(msg)
   fmt.Println("xxx")
   if err != nil {
      fmt.Println("send msg failed,err:", err)
      return
   }
   fmt.Printf("pid:%v offset:%v\n", pid, offset)
   fmt.Println("发送成功！")
}
```

运行tail.demo才能再终端中显示

## Logagent实现156

### 建立文件夹156_logagent

main.go

```go
package main

import (
   "example.com/logagent/kafka"
   "example.com/logagent/taillog"
   "fmt"
   "time"
)

func run() {
   //1、读取日志
   for {
      select {
      case line := <-taillog.ReadChan():
         //2、发送到kafka
         kafka.SendToKafka("web_log", line.Text)
      default:
         time.Sleep(time.Second)
      }

   }
}

// logAgent入口程序
func main() {
   //1、初始化kafka连接
   err := kafka.Init([]string{"127.0.0.1:9092"})
   if err != nil {
      fmt.Printf("init Kafka failed,err:%v\n", err)
      return
   }
   fmt.Println("init kafka success")
   //2、打开日志文件准备收集日志
   err = taillog.Init("F:\\goland\\go_project\\21weeks\\my.log")
   if err != nil {
      fmt.Printf("init taillog failed,err:%v\n", err)
      return
   }
   fmt.Println("init taillog success")
   run()
}
```

### 建立两个子文件夹

kafka 下 kafka.go

```go
package kafka

import (
   "fmt"
   "github.com/Shopify/sarama"
)

//打开往kafka写日志的模块

var (
   client sarama.SyncProducer //声明一个全局的连接kafka的生产者client
)

// Init 初始化client
func Init(addrs []string) (err error) {
   config := sarama.NewConfig()
   config.Producer.RequiredAcks = sarama.WaitForAll          //发送完整数据类型leader和follow都确认
   config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
   config.Producer.Return.Successes = true                   //成功交付的消息将在success channel返回
   //连接kafka
   client, err = sarama.NewSyncProducer(addrs, config)
   if err != nil {
      fmt.Println("producer closed,err:", err)
      return
   }
   return
}
func SendToKafka(topic, data string) {
   //构造一个消息
   msg := &sarama.ProducerMessage{}
   msg.Topic = topic
   msg.Value = sarama.StringEncoder(data)
   //发送到kafka
   pid, offset, err := client.SendMessage(msg)
   fmt.Println("xxx")
   if err != nil {
      fmt.Println("send msg failed,err:", err)
      return
   }
   fmt.Printf("pid:%v offset:%v\n", pid, offset)
   fmt.Println("发送成功！")
}
```

taillog 下 taillog.go

```go
package taillog

import (
   "fmt"
   "github.com/hpcloud/tail"
)

var (
   tailObj *tail.Tail
   LogChan chan string
)

func Init(fileName string) (err error) {
   config := tail.Config{
      ReOpen:    true,                                 //重新打开
      Follow:    true,                                 //是否跟随
      Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
      MustExist: false,                                //文件不存在不报错
      Poll:      true,                                 //轮询文件更改
   }
   tailObj, err = tail.TailFile(fileName, config)
   if err != nil {
      fmt.Println("tail file failed,err:", err)
      return
   }
   return
}
func ReadChan() <-chan *tail.Line {
   return tailObj.Lines
}
```

conf 

config.go

```go
package conf

type Sumconfig struct {
   KafkaConf   `ini:"kafka"`
   TaillogConf `ini:"taillog"`
}
type KafkaConf struct {
   Address string `ini:"address"`
   Topic   string `ini:"topic"`
}
type TaillogConf struct {
   Filename string `ini:"filename"`
}
```

config,ini

```go
[kafka]
address=127.0.0.1:9092
topic=web_log

[taillog]
filename=F:\goland\go_project\21weeks\my.log
```

### 运行：

主文件夹在终端中打开，输入以下命令：

```bash
set GO111MODULE=on
SET GOPROXY=http://goproxy.cn
//初始化go mod
go mod init example.com/logagent
//整理一下
go mod tidy
```

将go.mod里的版本改为

require github.com/Shopify/sarama v1.19.0

```bash
go get//或者go mod download
```

用管理员启用cmd：

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\zookeeper-server-start.bat config\zookeeper.properties
```

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\kafka-server-start.bat config\server.properties
```

运行Goland

用管理员方式运行cmd:程序已经发送到kafka中，下面就用kafka读取数据即可（kafka-终端-消费者）

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\kafka-console-consumer.bat --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning
```

打开文件my.log输入内容，在cmd中显示

额外命令：

```bash
go env   //查看环境配置
```

--------------------------------------------------------

### BUG原因及调试

错误：终端打开文件夹156_logagent，输入go get后出现以下错误：

```bash 
example.com/logagent imports
        21weeks/21weeks_go/156_logagent/kafka: package 21weeks/21weeks_go/156_logagent/kafka is not in GOROOT (F:\goland\install\go1.20.1\src\21weeks\21weeks_go\156_logagent\kafka)    
example.com/logagent imports
        21weeks/21weeks_go/156_logagent/taillog: package 21weeks/21weeks_go/156_logagent/taillog is not in GOROOT (F:\goland\install\go1.20.1\src\21weeks\21weeks_go\156_logagent\taillo
g)
```

在为初始化go mod之前是可以调用这两个package的，初始化后则标红。

原因：

将go mod的配置打开后，在默认情况下是会使用go mod的管理方式。而打开的情况下很重要的一点，当前文件不能在GOPATH的src路径之下。这样go会默认使用GOPATH来进行包的管理，go mod 相关的配置就不起作用了，进而就无法进行go mod模块化的管理。

我的go mod init初始化为：

```bash
go mod init example.com/logagent
```

解决办法：

直接用go mod的路径调取这两个包，而不是根据它们原有的目录

```go
import (
	"example.com/logagent/kafka"
	"example.com/logagent/taillog"
)
```

## 157、配置文件版LogAgent

创建配置文件ini：

创建文件夹conf：创建文件config.ini

ini中文网使用教程：[开始使用 - go-ini/ini (unknwon.cn)](https://ini.unknwon.cn/docs/intro/getting_started)

```go
[kafka]
address=127.0.0.1:9092
topic=web_log

[taillog]
filename=F:\goland\go_project\21weeks\my.log
```

### 写法1

main.go 

```go
//0.加载配置文件
cfg, err = ini.Load("./conf/config.ini")
if err != nil {
   fmt.Printf("Fail to read file: %v", err)
   os.Exit(1)
}
//init 不能写注释
//fmt.Println(cfg.Section("kafka").Key("address").String())
//fmt.Println(cfg.Section("kafka").Key("topic").String())
//fmt.Println(cfg.Section("taillog").Key("filename).String())
其它的改一下
```

main.go 完整

```go
package main

import (
   "example.com/logagent/kafka"
   "example.com/logagent/taillog"
   "fmt"
   "gopkg.in/ini.v1"
   "os"
   "time"
)

var (
   err error
   cfg *ini.File
)

func run() {
   //1、读取日志
   for {
      select {
      case line := <-taillog.ReadChan():
         //2、发送到kafka
         kafka.SendToKafka(cfg.Section("kafka").Key("topic").String(), line.Text)
      default:
         time.Sleep(time.Second)
      }
   }
}

// logAgent入口程序
func main() {
   //0.加载配置文件
   cfg, err = ini.Load("./conf/config.ini")
   if err != nil {
      fmt.Printf("Fail to read file: %v", err)
      os.Exit(1)
   }
   //init 不能写注释
   fmt.Println(cfg.Section("kafka").Key("address").String())
   fmt.Println(cfg.Section("kafka").Key("topic").String())
   fmt.Println(cfg.Section("taillog").Key("filename").String())
   //1、初始化kafka连接
   err = kafka.Init([]string{cfg.Section("kafka").Key("address").String()})
   if err != nil {
      fmt.Printf("init Kafka failed,err:%v\n", err)
      return
   }
   fmt.Println("init kafka success")
   //2、打开日志文件准备收集日志
   err = taillog.Init(cfg.Section("taillog").Key("filename").String())
   if err != nil {
      fmt.Printf("init taillog failed,err:%v\n", err)
      return
   }
   fmt.Println("init taillog success")
   run()
}
```

### 写法2

在conf文件夹里创建一个结构体：

```go
package conf

type Sumconfig struct {
   KafkaConf   `ini:"kafka"`
   TaillogConf `ini:"taillog"`
}
type KafkaConf struct {
   Address string `ini:"address"`
   Topic   string `ini:"topic"`
}
type TaillogConf struct {
   Filename string `ini:"filename"`
}
```

main.go

```go
package main

import (
   "example.com/logagent/conf"
   "example.com/logagent/kafka"
   "example.com/logagent/taillog"
   "fmt"
   "gopkg.in/ini.v1"
   "time"
)

var (
   cfg = new(conf.Sumconfig)
)

func run() {
   //1、读取日志
   for {
      select {
      case line := <-taillog.ReadChan():
         //2、发送到kafka
         kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
      default:
         time.Sleep(time.Second)
      }
   }
}

// logAgent入口程序
func main() {
   // 将配置文件加载出来映射到cfg对象里面
   err := ini.MapTo(cfg, "./conf/config.ini")
   if err != nil {
      fmt.Printf("Fail to read file: %v", err)
      return
   }
   //1、初始化kafka连接
   err = kafka.Init([]string{cfg.KafkaConf.Address})
   if err != nil {
      fmt.Printf("init Kafka failed,err:%v\n", err)
      return
   }
   fmt.Println("init kafka success")
   //2、打开日志文件准备收集日志
   err = taillog.Init(cfg.TaillogConf.Filename)
   if err != nil {
      fmt.Printf("init taillog failed,err:%v\n", err)
      return
   }
   fmt.Println("init taillog success")
   run()
}
```

# 内容回顾

## go module

依赖管理工具

## context

goroutine管理

```context.Context```

两个根节点:```context.Backgrpund()```,```context.TODO()```

四个方法:```context.withCancel()```、```context.withTimeout()```、```context.withDeadline()```、```context.withValue()```

## 日志收集项目

### ELK：目前业界主流的日志收集方案

缺点：部署的时候比较麻烦。每一个filebeat都需要配置一个配置文件。

解决：**使用etcd来管理被收集的日志项**。

### 项目架构：

![image-20230403091601178](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403091601178.png)

1、kafka：消息队列

2、tailf：读日志的第三方库

3、go-ini :解析配置文件

# etcd

## 基本

使用etcd==优化==日志收集项目

协议：

1、了解Raft协议

​	1、选举

在Raft协议中，每个节点都可以处于三种状态之一：follower、candidate和leader。所有节点在初始状态下都是follower，leader负责处理客户端的请求，follower和candidate负责将请求转发给leader。如果follower在一段时间内没有收到leader的消息，它会自动转变成candidate状态，向其他节点发送选举请求，并等待其他节点的投票，如果它获得了半数以上节点的投票，就会成为新的leader。如果在选举过程中出现了平票的情况，那么就进行多次投票，直到出现了胜者。

​	2、日志复制机制

Raft协议采用的是日志复制模式，即leader向follower广播日志，follower接收并存储日志，并将日志复制到其他follower节点中。当leader节点宕机时，选举出新的leader节点，并通过复制日志来保证与原来的leader节点具有相同的状态

​	3、异常处理（脑裂）

在Raft协议中，由于网络问题或者节点宕机等原因，可能会导致脑裂的问题，即分布式系统中的不同节点之间彼此无法连接，形成两个不同的集群。在这种情况下，Raft采用“领导者复制”（Leader-based replication）的方式，即将集群中只有leader节点的集群视为合法状态，负责启动一个新的选举过程，重新选举出新的leader节点，保持分布式系统的正常运行。

​	4、zookeeper的zab协议的区别

Raft协议和Zookeeper的Zbab协议类似，都是解决分布式系统中的一致性问题的协议。但是它们之间有一些区别：

- 选举方式不同：Zab协议使用了类Paxos算法的选举机制，而Raft协议使用了Term的概念，使得选举过程更加直观。
- 数据同步机制不同：Zab协议采用的是两阶段提交机制，而Raft协议是leader向follower广播日志的方式。
- 容错机制不同：Zab协议在主节点挂掉的时候，其他节点也无法工作，而Raft协议则可以在leader节点挂掉的情况下，自动进行选举。

watch底层实现的原理

回答为什么不用ELK？

etcd的watch 

etcd的底层如何实现watch给客户发通知websocket



服务注册发现

![img](https://www.liwenzhou.com/images/Go/etcd/etcd_01.png)

![image-20230411151750024](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230411151750024.png)

etcd部署：[etcd部署 (qq.com)](https://docs.qq.com/doc/DTndrQXdXYUxUU09O)

## 配置中心

将一些配置信息放到etcd上进行集中管理

这类场景的使用方式为：应用在启动的时候主动从etcd获取一次配置信息，同时，在etcd节点上注册一个watch并等待，以后每次配置有更新的时候，etcd都会实时通知订阅者，以此达到获取最新配置信息的目的。

## 分布式锁

因为 etcd 使用 Raft 算法保持了数据的强一致性，某次操作存储到集群中的值必然是全局一致的，所以很容易实现分布式锁。锁服务有两种使用方式，一是保持独占，二是控制时序。

全班同学拿一瓶水

- **保持独占即所有获取锁的用户==最终只有一个==可以得到**。etcd 为此提供了一套实现分布式锁原子操作 CAS（`CompareAndSwap`）的 API。通过设置`prevExist`值，可以保证在多个节点同时去创建某个目录时，只有一个成功。而创建成功的用户就可以认为是获得了锁。
- ==控制时序==，即所有想要获得锁的用户都会被安排执行，但是**获得锁的顺序也是全局唯一的，同时决定了执行顺序**。etcd 为此也提供了一套 API（自动创建有序键），对一个目录建值时指定为`POST`动作，这样 etcd 会自动在目录下生成一个当前最大的值为键，存储这个新的值（客户端编号）。同时还可以使用 API 按顺序列出所有当前目录下的键值。此时这些键的值就是客户端的时序，而这些键中存储的值可以是代表客户端的编号。

![img](https://www.liwenzhou.com/images/Go/etcd/etcd_02.png)

## 为什么用 etcd 而不用ZooKeeper？

etcd 实现的这些功能，ZooKeeper都能实现。那么为什么要用 etcd 而非直接使用ZooKeeper呢？

### 为什么不选择ZooKeeper？

1. 部署维护复杂，其使用的`Paxos`强一致性算法复杂难懂。官方只提供了`Java`和`C`两种语言的接口。
2. 使用`Java`编写引入大量的依赖。运维人员维护起来比较麻烦。
3. 最近几年发展缓慢，不如`etcd`和`consul`等后起之秀。

### 为什么选择etcd？

1. 简单。使用 Go 语言编写部署简单；支持HTTP/JSON API,使用简单；使用 Raft 算法保证强一致性让用户易于理解。
2. etcd 默认数据一更新就进行持久化。
3. etcd 支持 SSL 客户端安全认证。

最后，etcd 作为一个年轻的项目，正在高速迭代和开发中，这既是一个优点，也是一个缺点。优点是它的未来具有无限的可能性，缺点是无法得到大项目长时间使用的检验。然而，目前 `CoreOS`、`Kubernetes`和`CloudFoundry`等知名项目均在生产环境中使用了`etcd`，所以总的来说，etcd值得你去尝试。

![image-20230411152934148](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230411152934148.png)

![image-20230411153044842](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230411153044842.png)

## put和get操作

### 终端执行：

以管理员身份打开etcd.exe

终端执行

```cd /d F:\goland\go_project\etcd\etcd-v3.5.7-windows-amd64```

```etcdctl.exe --endpoints=http://127.0.0.1:2379 put qikey "sh1"```

```etcdctl.exe --endpoints=http://127.0.0.1:2379 get qikey```

### go执行 初始化

安装：```go get go.etcd.io/etcd/clientv3```

代码：

```go
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	//导入的包，v3表示版本
	// 创建etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		//节点
		Endpoints: []string{"localhost:2379"},
		//5s钟都连不上就超时了
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put 创建一个键值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "q1mi", "dsb")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get 获取一个键的值
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "q1mi")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	//Kvs多个键值对，一个个遍历出来
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
```

初始化go.mod ```go mod init etcd_demo.com/etcd1```

整理依赖关系```go mod tidy```

```go build```出现以下问题：

```go
C:\Users\Administrator\go\pkg\mod\github.com\coreos\etcd@v3.3.27+incompatible\clientv3\balancer\resolver\endpoint\endpoint.go:182:40: undefined: resolver.ResolveNowOption
# github.com/coreos/etcd/clientv3/balancer/picker
C:\Users\Administrator\go\pkg\mod\github.com\coreos\etcd@v3.3.27+incompatible\clientv3\balancer\picker\err.go:37:53: undefined: balancer.PickOptions
```

### 错误1：

[(48条消息) `golang` 调用 `etcdv3` 报错 `undefined: balancer.PickOptions`_斜杠打卡小程序的博客-CSDN博客](https://blog.csdn.net/qq_32828933/article/details/107179973)

[(48条消息) golang etcd 报错 undefined: resolver.BuildOption 解决方案_wohu1104的博客-CSDN博客](https://blog.csdn.net/wohu1104/article/details/107923944)

错误信息：

```go
# github.com/coreos/etcd/clientv3/balancer/picker
undefined: balancer.PickOptions
undefined: balancer.PickOptions

# github.com/coreos/etcd/clientv3/balancer/resolver/endpoint
undefined: resolver.BuildOption
undefined: resolver.ResolveNowOption
```

解决方法：

```go
将 grpc 版本替换成 v1.26.0
```

修改依赖为 v1.26.0

```go
go mod edit -require=google.golang.org/grpc@v1.26.0
```

获取 v1.26.0 版本的 grpc

```go 
go get -u -x google.golang.org/grpc@v1.26.0
```

命令行：

```cd /d F:\goland\go_project\etcd\etcd-v3.4.24-windows-amd64```

```etcdctl.exe --endpoints=http://127.0.0.1:2379 put qikey "sh1"```

```etcdctl.exe --endpoints=http://127.0.0.1:2379 get qikey```

### 错误2：

```go
connect to etcd success
{"level":"warn","ts":"2023-04-06T15:41:59.416+0800","caller":"clientv3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"endpoint://client-23dc86cb-73cb-4fb0-
ac93-f8bd79d9ddfa/127.0.0.1:2379","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = context deadline exceeded"}
put to etcd failed, err:context deadline exceeded
```

报错，put失败，超出上下文截止时间

url改为：

```go
localhost:2379
```

# 从etcd中获取日志收集项的配置信息



用管理员启用cmd：

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\zookeeper-server-start.bat config\zookeeper.properties
```

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\kafka-server-start.bat config\server.properties
```



双击**etcd.exe**就是启动了etcd。其他平台解压之后在bin目录下找etcd可执行文件。

默认会在2379端口监听客户端通信，在2380端口监听节点间通信。

**etcdctl.exe**可以理解为一个客户端或本机etcd的控制端。

运行Goland

用管理员方式运行cmd:程序已经发送到kafka中，下面就用kafka读取数据即可（kafka-终端-消费者）

```bash
cd /d F:\kafka_2.13-3.4.0
bin\windows\kafka-console-consumer.bat --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning
```

打开文件my.log输入内容，在cmd中显示





Logagent：把文件从日志中读取出来发送到kafka

Logtransfer：从kafka把日志取出来写入ES，使用Kibana做可视化展示

系统监控：gopsutil做系统监控信息的采集，写入influxDB，使用grafana作展示

prometheus监控：采集性能呢指标数据，保存起来，使用grafnan作展示

# LogTransfer

## ES(Elastic search)

 ES搭建指南：docs.qq.com/doc/DTmZxQUdHeFRXU2dP

下载

```bin\elasticsearch.bat```

## Kibana

下载

切换不同版本的JDK：[(50条消息) Win10同时安装使用Java JDK8和11两个版本如何设置环境变量_Yeoh1999的博客-CSDN博客](https://blog.csdn.net/Yeoh1999/article/details/108248254)

jdk1.8适用于 ES7.6.1

jdk20适用于ES8.7.0，但是似乎不稳定

Postman：

执行go程序，然后获得index和type。打开postman

GET      127.0.0.1:9200/user/_doc/_search

# Kafka消费

根据topic找到分区

根据每一个分区去消费数据

```go
package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}
```

# LogTaranfer实现

# 项目总结

项目名称改一下：服务端agent开发

1、项目架构（图）

2、监控可以不说

3、为什么不用ELK

4、logAgent里面如何保证日志不丢/重启之后继续收集日志（记录读取文件的offset）

5、kafka课上整理的点

6、etcd的watch原理

7、es相关知识点（搜索引擎的底层如何实现的）



找工作：

1、找开发的算法和数据结构

2、运维会前端

3、Boss直聘花钱

