# seelog [![Build Status](https://travis-ci.org/xmge/seelog.svg?branch=master)](https://travis-ci.org/xmge/seelog) [![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/xmge/seelog/blob/master/LICENSE)


> 有了seelog,妈妈再也不用担心我登录服务器查看日志le...   
项目地址：https://github.com/xmge/seelog    
演示地址：http://seelog.xmge.top/password
欢迎各位gopher使用指正:smiley:

### 项目背景
> A：我去:confused:,,程序又出问题了，你去看看日志到底是哪里错啦<br>
  B：好好好，我马上登陆服务器看看<br>
  B：对啦，服务器密码是啥来着....<br>
  A：:cold_sweat: 密码是 &#￥&*@*~<br>
  B：额，日志文件在哪呢<br>
  A：...................<br>
  C：要不项目集成seelog吧，集成特别简单，集成后直接打开浏览器就能看日志了

### 项目介绍
* 与golang项目集成、提供浏览器实时查看日志的功能，类似 [tail -f xxx.log](https://www.cnblogs.com/fps2tao/p/7698224.html)
* 支持监控多个日志文件
* 支持多浏览器同时访问
* 支持密码认证
* 支持浏览器 websocket 断线重连
* 支持暂停、清屏、截图、过滤功能
* 查找功能可直接使用浏览器 `Ctrl+F` 来完成

### 集成方式
* 在项目中引入seelog, **go get github.com/xmge/seelog**
* 在代码中引用

```go
	seelog.See("错误日志","err.log")   // 检测 err.log
	seelog.See("调试日志","debug.log") // 检测 debug.log
	seelog.Serve(9000,"password")
```

* 在浏览器中访问 *http://host:port/{password}*

### 项目展示
![image](https://github.com/xmge/seelog/blob/master/demo.gif)
