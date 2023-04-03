# 日志收集项目架构设计及Kafka介绍

![image-20230403090737229](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403090737229.png)

## 项目背景

日志信息解决问题、数据挖掘、浏览推荐定位

## 解决方案

把机器上的日志实时收集

## 面临的问题

实时日志量非常大，每天处理几十亿条。

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

Kabaa：开源的ES数据分析和可视化工具。Hadoop：分布式计算框架，能够对大量数据进行分布式处理的平台。

Strom：一个免费并且开源的分布式实时计算系统。



![image-20230403092339728](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403092339728.png)

NSQ：将同步消息队列转化为异步的，可以作为进程间的通信。

![image-20230403093921444](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403093921444.png)

![image-20230403093932561](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403093932561.png)

![image-20230403094158016](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403094158016.png)

![image-20230403094603816](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403094603816.png)

![image-20230403094623766](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403094623766.png)

生产者将数据扔到Cluster中，Consumer然后从中获取

Cluster中的Broker被称为节点，三台机器

Partition 提高容错率，分区

![image-20230403095043201](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403095043201.png)

![image-20230403095252305](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230403095252305.png)

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