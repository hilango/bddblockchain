package main

/*
本示例作为作者对BDD及区块链学习示例

本示例原代码来自：https://blog.csdn.net/b1303110335/article/details/79243510https://blog.csdn.net/b1303110335/article/details/79243510

主要演示如何使用BDD来实现一个小型区块链

用例：


1. 初始化Web服务
	a). 配置路由
	b). 获取本地址，并打印log
	c). 启动服务

2. 配置路由
	a). 添加Get方法，获取区块链
	b). 添加Post方法，写入区块链

3. 获取区块链

4. 写入区块链

5. 区块链结构体

{{Block

Index|索引位置
TimStamp|时间戳 #string
BPM|心率
Hash|SHA256
PreHash|前一区块的SHA256

calculateHash() string  //计算散列值

generateBlock(oldBlock,BPM) (Block,error)   //生成块

isBlockValid(newBlock,oldBlock) bool  //校验块

replaceChain(newBlocks) //将过期链更新为最新的链

}}

6. 结构体

Blockchain []Block
Message

7. 主方法
获取环境变量
初始化创世纪区块
初始化Web服务

------------

整个链

{{ 块

Index|位置
时间戳
数据
心跳数
Hash|哈希
PreHash|前一个哈希

计算哈希()

生成块()

检验

更新链

}}

run()
	创建路由
	获取端口
	创建http.Server
	监听服务

创建路由()
	Get处理器()
	Post处理器()

Get处理器()
	返回链

Post处理器()
	获取消息
	生成新区块
	添加新区块
	更新链

main()
	加载环境变量
	初始化区块链
		创建创世纪区块
		添加到链
	运行服务

 */
