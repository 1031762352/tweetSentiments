package model

// Tweet 推文数据表
// 存储从Twitter拉取的大V推文原始数据
type Tweet struct {
	BaseModel                           // 嵌入公共基础字段
	TweetID                      string `gorm:"column:tweet_id;type:varchar(64);not null;uniqueIndex;comment:'Twitter推文ID（id_str，全局唯一）'" json:"tweetID"`
	TwitterUserID                string `gorm:"column:twitter_user_id;type:varchar(64);not null;index;comment:'推文作者的Twitter用户ID（关联t_influencer.twitter_user_id）'" json:"twitterUserID"`
	Content                      string `gorm:"column:content;type:text;not null;comment:'推文正文内容（支持长文本）'" json:"content"`
	TweetCreatedAt               string `gorm:"column:tweet_created_at;type:varchar(32);not null;index;comment:'推文发布时间（ISO8601格式，如2025-11-16T10:00:00Z）'" json:"createdAt"`
	Lang                         string `gorm:"column:lang;type:varchar(16);comment:'推文语言（如en=英文，zh=中文）'" json:"lang"`
	Source                       string `gorm:"column:source;type:varchar(128);comment:'发布来源（如Twitter for iPhone）'" json:"source"`
	IsOriginal                   int8   `gorm:"column:is_original;type:tinyint;not null;default:1;comment:'是否原创推文（1=原创，0=转发/引用）'" json:"isOriginal"`
	ReferencedTweetID            string `gorm:"column:referenced_tweet_id;type:varchar(64);comment:'引用/转发的原推文ID（非原创时填写）'" json:"referencedTweetID"`
	ReferencedTweetType          string `gorm:"column:referenced_tweet_type;type:varchar(32);comment:'引用类型（retweeted=转发，quoted=引用，replied_to=回复）'" json:"referencedTweetType"`
	PublicMetricsLikeCount       int    `gorm:"column:public_metrics_like_count;type:int;default:0;comment:'点赞数'" json:"publicMetricsLikeCount"`
	PublicMetricsRetweetCount    int    `gorm:"column:public_metrics_retweet_count;type:int;default:0;comment:'转发数'" json:"publicMetricsRetweetCount"`
	PublicMetricsReplyCount      int    `gorm:"column:public_metrics_reply_count;type:int;default:0;comment:'回复数'" json:"publicMetricsReplyCount"`
	PublicMetricsQuoteCount      int    `gorm:"column:public_metrics_quote_count;type:int;default:0;comment:'引用数'" json:"publicMetricsQuoteCount"`
	PublicMetricsImpressionCount int64  `gorm:"column:public_metrics_impression_count;type:bigint;default:0;comment:'曝光量（展示次数）'" json:"publicMetricsImpressionCount"`
	PulledAt                     string `gorm:"column:pulled_at;type:varchar(32);not null;comment:'推文拉取时间（ISO8601格式）'" json:"pulledAt"`
}

// TableName 指定数据库表名
func (t *Tweet) TableName() string {
	return "t_tweet"
}
