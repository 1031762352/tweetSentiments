// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"

	"tweetSentiments/internal/config"
	"tweetSentiments/internal/handler"
	"tweetSentiments/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/tweetsentiments-api.yaml", "the config file path")

func main() {
	// 解析命令行参数（指定配置文件路径）
	flag.Parse()

	// 加载yaml配置到Config结构体
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 创建REST服务（使用yaml中的端口、超时等配置）
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop() // 程序退出时关闭服务

	// 初始化服务上下文（注入配置、DB、外部客户端等依赖）
	ctx := svc.NewServiceContext(c)

	// 注册所有API处理器
	handler.RegisterHandlers(server, ctx)

	// 启动提示
	fmt.Printf("✅ Server is starting at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	server.Start()
}
