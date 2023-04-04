# 日志收集项目架构设计及Kafka介绍

## 项目背景

日志信息解决问题、数据挖掘、浏览推荐定位

每个业务系统都有日志，当系统出现问题时，需要通过日志信息来定位和解决问题。当系统机器比较少时，登录到服务器上查看即可满足。当系统机器规模巨大，登录到机器上查看几乎不现实（分布式系统，一个系统部署在十几台机器上）

## 解决方案

把机器上的日志实时收集，统一存储到中心系统。再对这些日志建立索引，通过搜索即可快速找到对应的日志记录。通过提供一个界面友好的web页面实现日志展示与检索。

## 面临的问题

实时日志量非常大，每天处理几十亿条。日志准实时收集，延迟控制在分钟级别。系统的架构设计能够支持水平扩展。

## 业界方案

ELK

![image-20230403091122826](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403091122826.png)

filebit从AppServer收集信息给LogstashAgent

## ELK方案的问题

- 运维成本高，每增加一个日志收集项，都需要手动修改配置。
- 监控缺失，无法准确获取logstash的状态。
- 无法做到定制化开发与维护

# 日志收集系统架构设计

![image-20230403091601178](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403091601178.png)

Log Agent具体部署到每一台业务的服务器上，配置项收集日志，存储到etcd中。

etcd中写一个web管理页面，管理所有的配置项。

Log Agent将每一个项目的日志收集后存到Kafka中，Kafka既能作为消息队列，也能作为日志存储。

Sys Agent 做系统监控，例如cpu负载，IO，网络速率

## 组件介绍

LoaAgent：日志收集客户端，用来收集服务器上的日志。Kafka：高吞吐量的分布式队列（Linkin开发，apache顶级开源项目）ElasticSearch：开源的搜素引擎，提供基于HTTP RESTful的web接口。

Kibaa：开源的ES数据分析和可视化工具。Hadoop：分布式计算框架，能够对大量数据进行分布式处理的平台。

Strom：一个免费并且开源的分布式实时计算系统。

## 将学到的技能

服务端agent开发

后端服务组件开发

etcd的使用

Kafka和zookeeper的使用

ES和Kibana的使用

## 消息队列的通信模型

NSQ：将同步消息队列转化为异步的，可以作为进程间的通信。

消息生产者生产消息发送到queue中，然后消息消费者从queue中取出并且消费消息。一条消息被消费以后，queue中就没有了，不存在重复消费。

## 发布/订阅（topic）

消息生产者（发布）将消息发布到topic中，同时有多个消息消费者（订阅）消费该消息。和点对点方式不同，发布到topic的消息会被所有订阅者消费（类似于关注了微信公众号的人都能收到推送的文章）。补充：发布订阅模式下，当发布者消息量很大时，显然单个订阅者的处理能力是不足的。实际上现实场景中是多个订阅者节点组成一个订阅组负载均衡消费topic消息即分组订阅，这样订阅者很容易实现消费能力线性扩展。可以看成是一个topic下有多个Queue,每个Queue是点对点的方式，Queue之间是发布订阅方式。

## Kafka

Apache Kafka由著名职业社交公司Linkedln开发，最初是被设计用来解决LinkedIn公司内部海量日志传输等问题。Kafka使用Scala语言编写，2011年开源并进入Apache孵化器，2012年10月正式毕业，现在为Apache顶级项目

## 介绍

Kafka是一个分布式数据流平台，可以运行在单台服务器上，也可以在多台服务器上部署形成集群。他提供了发布和订阅功能，使用者可以发送数据到Kafka中，也可以从Kafka中读取数据（以便进行后续的处理）。Kafka具有高吞吐、低延迟、高容错等特点。

<img src="C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403094603816.png" alt="image-20230403094603816" style="zoom:80%;" />

<img src="C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403094623766.png" alt="image-20230403094623766" style="zoom:80%;" />

- Producer：生产者将，消息的产生这，是消息的入口。数据扔到Cluster中，Consumer然后从中获取

- Kafka Cluster：kafka集群，一台或者多台服务器组成。Broker被称为节点，三台机器（代理）

  - ​		Broker：Broker是指部署了Kafka实例的服务器节点。每个服务器上有一个或多个kafka的实列，我们姑且认为每个broker对应一台服务器。每个Kafka集群内的broker都有一个不重复的编号，如图中broker-0、broker-1等...

  - ​		Topic：消息的主题 ，可以理解为消息的分蘖，kafka的数据就保存在同批次。每个broker上都可以创建多个同批次。实际应用中通常是一个业务线建立一个topic

  - ​		Partition：Topic的分区，每个topic可以有多个分区，分区的作用是做负载，提高kafka的吞吐量。同一个topic在不同的分区数据是不重复的，partition的表现形式就是一个一个的文件夹。

  - ​		Replication：每个分区都有多个副本，副本的作用是做备胎。当主分区（Leader）故障的时候会选择一个（Follower）尚未，称为Leader。在kafka中默认副本的最大数量是10个，且副本的数量不能大于Broker的数量，follower和leader在不同的机器中，同一个机器对同一个分区只能存放一个副本（包括自己）

- Consumer：消费者，即消息的消费方，是消息的出口。
  - ​		Consumer Group：我们可以将多个消费组组成一个消费者组，在kafka的设计中，同一个分区的数据只能被消费者组中的某一个消费者消费。同一个消费者组的消费者可以消费同一个topic的不同分区的数据，这也是为了提高kafka的吞吐量。

## 工作流程

![image-20230403095351391](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403095351391.png)

![image-20230403101417190](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403101417190.png)

![image-20230403101526670](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403101526670.png)

![image-20230403102611750](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403102611750.png)

![image-20230403102620825](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403102620825.png)

![image-20230403102838082](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403102838082.png)

![image-20230403102935219](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403102935219.png)

![image-20230403103048304](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403103048304.png)

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

```ba
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

用管理员方式运行cmd:

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