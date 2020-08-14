#介绍
go的几道练习题

## 第一题
建立一个tcp服务器，telnet该服务器，输入一个字符串进去，立马返回该字符串的逆序字符串


## 第二题
建立一个tcp服务器，telnet该服务器，输入一个表达式，可以计算结果(可以使用外部模块来解析表达式，至少支持加减乘除)


## 第三题
建立一个tcp服务器，telnet该服务器，返回一个真随机算子 (利用gorouine中的知识)


## 第四题
使用cgo封装C++ STL中的Set类型，支持所有主要函数（insert, find, size, erase） 
我这里打算封装所有类型 但是没有实现  改成单独int就比较好实现

## 第五题 
使用colly爬取站点 http://tumregels.github.io/Network-Programming-with-Go/ 所有页面(在本地打开页面后展开的效果和服务器上差不多)
[样例](https://ruilisi.github.io/golang-crawler-example/)

## 第六题 
建立一个http服务器，上传一个彩色图片，返回该彩色的黑白图片
