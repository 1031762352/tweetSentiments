package svc

import (
	"time"
	"tweetSentiments/internal/config"
	"tweetSentiments/model/gorm"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config                     config.Config
	DB                         *gorm.DB
	Influencer                 *model.Influencer // 指针类型
	Tweet                      *model.Tweet
	SentimentResult            *model.SentimentResult
	SentimentKeyword           *model.SentimentKeyword
	InfluencerSentimentSummary *model.InfluencerSentimentSummary
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := initMySQL(c.MySQLConf)
	if err != nil {
		logx.Error("MySQL初始化失败：", err)
		panic(err)
	}

	// 关键修正：调用构造函数时不要传 db 参数
	return &ServiceContext{
		Config:                     c,
		DB:                         db,
		Influencer:                 model.NewInfluencer(),                 // 无参数！！！
		Tweet:                      model.NewTweet(),                      // 无参数！！！
		SentimentResult:            model.NewSentimentResult(),            // 无参数！！！
		SentimentKeyword:           model.NewSentimentKeyword(),           // 无参数！！！
		InfluencerSentimentSummary: model.NewInfluencerSentimentSummary(), // 无参数！！！
	}
}

// initMySQL 初始化MySQL连接（GORM v2）
func initMySQL(mysqlConf config.MySQLConf) (*gorm.DB, error) {
	// GORM配置
	gormConfig := &gorm.Config{
		// 日志配置：开发环境打印SQL，生产环境可调整为warn级别
		Logger: logger.Default.LogMode(logger.Info),
		// 禁用默认事务（提高性能，按需开启）
		SkipDefaultTransaction: true,
	}

	// 连接MySQL
	db, err := gorm.Open(mysql.Open(mysqlConf.DSN), gormConfig)
	if err != nil {
		return nil, err
	}

	// 获取底层sql.DB连接，设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 连接池配置（根据业务调整）
	sqlDB.SetMaxOpenConns(100)                 // 最大打开连接数
	sqlDB.SetMaxIdleConns(20)                  // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 连接最大生命周期
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // 连接最大空闲时间

	logx.Info("MySQL连接初始化成功！")
	return db, nil
}
