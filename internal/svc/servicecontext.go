// internal/svc/servicecontext.go
package svc

import (
	"gorm.io/gorm"
	"tweetSentiments/internal/config"
)

type ServiceContext struct {
	Config config.Config // 全局配置
	DB     *gorm.DB      // MySQL连接
	// 其他依赖（如Twitter客户端、Qdrant客户端等）
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化DB、客户端等依赖，使用c中的配置（如c.MySQLConf.DSN）
	db := initMySQL(c.MySQLConf)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
