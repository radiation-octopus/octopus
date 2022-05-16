package tcp

//存放默认常量

//Server 端口
var TcpServerPort int

//Server 回调线程数量
var TcpServerBindingPoolNum int

//Server tcp接受回调方法
var TcpServerTcpAcceptCallBindingMethod string

//Server tcp接受回调注入体
var TcpServerTcpAcceptCallBindingStruct string

//Clinet 端口
var TcpClinetPort int

//Clinet 发送消息队列
var TcpClinetMsgNum int

//Clinet tcp接受回调方法
var TcpClinetTcpAcceptCallBindingMethod string

//Clinet tcp接受回调注入体
var TcpClinetTcpAcceptCallBindingStruct string
