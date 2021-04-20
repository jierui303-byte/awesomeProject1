package main

func main() {
	/**
		在gin框架中，Engine被定义为一个结构体，其中包含了路由组，中间件，页面渲染接口，框架配置设置等相关内容
		默认的Engine可以通过gin.Default()进行创建，或者使用gin.New()同样可以创建
		gin.Default()和gin.New()的区别在于default是使用new创建的engine实例，
		但是会默认使用Logger和Recovery中间件
		Logger是负责进行打印并输出的日志的中间件
		Recovery中间件的作用是如果程序执行过程中panic中断了服务，则recovery会恢复程序执行，并返回服务器500内部错误。
		通常使用gin.Default()来创建实例

		engine支持很多种http请求方式：options,head,get,post,put,delete,trace,connect
		常用请求方式：get,post,delete

	*/


}
