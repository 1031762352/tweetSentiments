package model

// Influencer 大V订阅表
// 存储已订阅的Twitter大V基础信息
type Influencer struct {
	BaseModel                      // 嵌入公共基础字段
	TwitterUserID           string `gorm:"column:twitter_user_id;type:varchar(64);not null;uniqueIndex;comment:'Twitter官方用户ID（id_str，全局唯一）'" json:"twitterUserID"`
	Username                string `gorm:"column:username;type:varchar(64);not null;uniqueIndex;comment:'Twitter用户名（如elonmusk，唯一标识）'" json:"username"`
	Name                    string `gorm:"column:name;type:varchar(128);not null;comment:'大V昵称（可包含特殊字符）'" json:"name"`
	AvatarURL               string `gorm:"column:avatar_url;type:varchar(255);comment:'头像URL（高清图）'" json:"avatarURL"`
	Description             string `gorm:"column:description;type:text;comment:'个人简介（支持长文本）'" json:"description"`
	Location                string `gorm:"column:location;type:varchar(128);comment:'所在地（用户填写，非精确位置）'" json:"location"`
	URL                     string `gorm:"column:url;type:varchar(255);comment:'个人网站/博客链接'" json:"url"`
	PublicMetricsFollowers  int64  `gorm:"column:public_metrics_followers;type:bigint;default:0;comment:'粉丝数（避免溢出用bigint）'" json:"publicMetricsFollowers"`
	PublicMetricsFollowing  int    `gorm:"column:public_metrics_following;type:int;default:0;comment:'关注数（普通int足够）'" json:"publicMetricsFollowing"`
	PublicMetricsTweetCount int    `gorm:"column:public_metrics_tweet_count;type:int;default:0;comment:'历史推文总数'" json:"publicMetricsTweetCount"`
	IsVerified              int8   `gorm:"column:is_verified;type:tinyint;not null;default:0;comment:'是否官方认证（1=是，0=否）'" json:"isVerified"`
	IsActive                int8   `gorm:"column:is_active;type:tinyint;not null;default:1;comment:'是否活跃拉取推文（1=是，0=暂停）'" json:"isActive"`
	PullFrequency           int    `gorm:"column:pull_frequency;type:int;not null;default:30;comment:'推文拉取频率（单位：分钟，默认30分钟）'" json:"pullFrequency"`
	LastPullTweetID         string `gorm:"column:last_pull_tweet_id;type:varchar(64);comment:'上次拉取的最后一条推文ID（避免重复拉取）'" json:"lastPullTweetID"`
}

// TableName 指定数据库表名（规范：t_+领域名，小写+下划线）
func (i *Influencer) TableName() string {
	return "t_influencer"
}
