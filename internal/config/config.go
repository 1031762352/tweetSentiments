// internal/config/config.go
package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

// Config 全局配置结构体（与yaml文件字段一一对应）
type Config struct {
	RestConf      rest.RestConf   `yaml:"rest"`      // HTTP服务配置
	MySQLConf     MySQLConf       `yaml:"mysql"`     // MySQL数据库配置
	QdrantConf    QdrantConf      `yaml:"qdrant"`    // Qdrant向量数据库配置
	TwitterConf   TwitterConf     `yaml:"twitter"`   // Twitter API配置
	EmbeddingConf EmbeddingConf   `yaml:"embedding"` // 嵌入服务配置
	CacheConf     cache.CacheConf `yaml:"cache"`     // 缓存配置（可选）
	CronConf      CronConf        `yaml:"cron"`      // 定时任务配置（推文拉取、情感分析）
}

// RestConf HTTP服务配置（继承go-zero的rest.RestConf，扩展自定义字段）
// 注：go-zero的rest.RestConf已包含Host、Port、Timeout等基础字段，无需重复定义

// MySQLConf MySQL数据库配置
type MySQLConf struct {
	DSN string `yaml:"dsn"` // 连接地址：user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
}

// QdrantConf Qdrant向量数据库配置（存储推文嵌入向量，用于相似性查询）
type QdrantConf struct {
	Endpoint       string `yaml:"endpoint"`       // 服务地址（如 http://127.0.0.1:6334）
	APIKey         string `yaml:"apiKey"`         // 认证密钥（无则留空）
	CollectionName string `yaml:"collectionName"` // 集合名（如 tweet_sentiments）
}

// TwitterConf Twitter API配置（用于拉取大V推文）
type TwitterConf struct {
	BearerToken string `yaml:"bearerToken"` // 认证Token（从Twitter开发者平台获取）
	Timeout     int    `yaml:"timeout"`     // 请求超时时间（单位：秒，默认30）
	MaxRetries  int    `yaml:"maxRetries"`  // 失败重试次数（默认3）
}

// EmbeddingConf 嵌入服务配置（用于生成推文文本嵌入向量）
type EmbeddingConf struct {
	Endpoint string `yaml:"endpoint"` // 服务地址（如本地模型或第三方API）
	APIKey   string `yaml:"apiKey"`   // 认证密钥（无则留空）
	Model    string `yaml:"model"`    // 模型名称（如 text-embedding-3-small）
}

// CronConf 定时任务配置
type CronConf struct {
	TweetPullCron string `yaml:"tweetPullCron"` // 推文拉取定时表达式（如 "0 */30 * * * *" 每30分钟）
	SentimentCron string `yaml:"sentimentCron"` // 情感分析定时表达式（如 "5 */30 * * * *" 每30分钟第5秒）
	SummaryCron   string `yaml:"summaryCron"`   // 情感汇总定时表达式（如 "0 0 * * * *" 每天凌晨）
}
