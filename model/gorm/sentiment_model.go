package model

// SentimentResult 情感分析结果表
type SentimentResult struct {
	BaseModel                 // 嵌入公共基础字段
	TweetID           string  `gorm:"column:tweet_id;type:varchar(64);not null;uniqueIndex;comment:'关联的推文ID（关联t_tweet.tweet_id，一对一）'" json:"tweetID"`
	SentimentScore    float64 `gorm:"column:sentiment_score;type:decimal(10,6);not null;comment:'情感得分（范围：-1~1，-1=极度负面，1=极度正面）'" json:"sentimentScore"`
	SentimentLabel    string  `gorm:"column:sentiment_label;type:varchar(16);not null;index;comment:'情感标签（positive=正面，negative=负面，neutral=中性）'" json:"sentimentLabel"`
	Confidence        float64 `gorm:"column:confidence;type:decimal(10,6);not null;comment:'分析置信度（范围：0~1，越接近1越可靠）'" json:"confidence"`
	KeywordScore      float64 `gorm:"column:keyword_score;type:decimal(10,6);default:0;comment:'领域关键词加权分（基于t_sentiment_keyword计算）'" json:"keywordScore"`
	SimilarTweetCount int     `gorm:"column:similar_tweet_count;type:int;default:0;comment:'相似历史推文数量（用于辅助判断）'" json:"similarTweetCount"`
	AnalyzedAt        string  `gorm:"column:analyzed_at;type:varchar(32);not null;comment:'情感分析完成时间（ISO8601格式）'" json:"analyzedAt"`
}

// TableName 指定数据库表名
func (s *SentimentResult) TableName() string {
	return "t_sentiment_result"
}

// NewSentimentResult 无参数构造函数（返回指针）
func NewSentimentResult() *SentimentResult {
	return &SentimentResult{}
}

// InfluencerSentimentSummary 大V情感汇总表
type InfluencerSentimentSummary struct {
	BaseModel                     // 嵌入公共基础字段
	Username              string  `gorm:"column:username;type:varchar(64);not null;index:idx_username_time_range,unique;comment:'大V用户名（关联t_influencer.username）'" json:"username"`
	TimeRange             string  `gorm:"column:time_range;type:varchar(8);not null;index:idx_username_time_range,unique;comment:'统计时间范围（7d=7天，30d=30天，90d=90天）'" json:"timeRange"`
	PositiveCount         int64   `gorm:"column:positive_count;type:bigint;default:0;comment:'正面推文数量'" json:"positiveCount"`
	NegativeCount         int64   `gorm:"column:negative_count;type:bigint;default:0;comment:'负面推文数量'" json:"negativeCount"`
	NeutralCount          int64   `gorm:"column:neutral_count;type:bigint;default:0;comment:'中性推文数量'" json:"neutralCount"`
	AverageSentimentScore float64 `gorm:"column:average_sentiment_score;type:decimal(10,6);comment:'平均情感得分（范围：-1~1）'" json:"averageSentimentScore"`
}

// TableName 指定数据库表名
func (i *InfluencerSentimentSummary) TableName() string {
	return "t_influencer_sentiment_summary"
}

// NewInfluencerSentimentSummary 无参数构造函数（返回指针）
func NewInfluencerSentimentSummary() *InfluencerSentimentSummary {
	return &InfluencerSentimentSummary{}
}
