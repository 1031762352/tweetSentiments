package model

import "gorm.io/gorm"

// SentimentKeyword 情感关键词表
type SentimentKeyword struct {
	BaseModel         // 嵌入公共基础字段
	Keyword   string  `gorm:"column:keyword;type:varchar(64);not null;index:idx_keyword_category,unique;comment:'情感关键词（如bullish=看涨，crash=崩盘）'" json:"keyword"`
	Weight    float64 `gorm:"column:weight;type:decimal(5,2);not null;comment:'情感权重（范围：-5~5，正数=正面，负数=负面，绝对值越大权重越高）'" json:"weight"`
	Category  string  `gorm:"column:category;type:varchar(32);not null;index:idx_keyword_category,unique;comment:'关键词分类（如crypto=加密货币，stock=股票，tech=科技）'" json:"category"`
	Frequency int     `gorm:"column:frequency;type:int;not null;default:0;comment:'使用频率（匹配到推文的次数，用于优化关键词权重）'" json:"frequency"`
}

// TableName 指定数据库表名
func (s *SentimentKeyword) TableName() string {
	return "t_sentiment_keyword"
}

// NewSentimentKeyword 无参数构造函数（返回指针）
func NewSentimentKeyword() *SentimentKeyword {
	return &SentimentKeyword{}
}

// 示例CRUD方法
func (s *SentimentKeyword) GetByCategory(db *gorm.DB, category string) ([]*SentimentKeyword, error) {
	var keywords []*SentimentKeyword
	err := db.Where("category = ?", category).Find(&keywords).Error
	return keywords, err
}
