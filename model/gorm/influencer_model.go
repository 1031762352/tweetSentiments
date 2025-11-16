package model

import "gorm.io/gorm"

// Influencer 大V订阅表（首字母大写，导出）
type Influencer struct {
	BaseModel                      // 嵌入公共基础字段
	TwitterUserID           string `gorm:"column:twitter_user_id;type:varchar(64);not null;uniqueIndex;comment:'Twitter官方用户ID（id_str，全局唯一）'" json:"twitterUserID"`
	Username                string `gorm:"column:username;type:varchar(64);not null;uniqueIndex;comment:'Twitter用户名（如elonmusk，唯一标识）'" json:"username"`
	Name                    string `gorm:"column:name;type:varchar(128);not null;comment:'大V昵称'" json:"name"`
	AvatarURL               string `gorm:"column:avatar_url;type:varchar(255);comment:'头像URL'" json:"avatarURL"`
	Description             string `gorm:"column:description;type:text;comment:'个人简介'" json:"description"`
	Location                string `gorm:"column:location;type:varchar(128);comment:'所在地'" json:"location"`
	URL                     string `gorm:"column:url;type:varchar(255);comment:'个人网站'" json:"url"`
	PublicMetricsFollowers  int64  `gorm:"column:public_metrics_followers;type:bigint;default:0;comment:'粉丝数'" json:"publicMetricsFollowers"`
	PublicMetricsFollowing  int    `gorm:"column:public_metrics_following;type:int;default:0;comment:'关注数'" json:"publicMetricsFollowing"`
	PublicMetricsTweetCount int    `gorm:"column:public_metrics_tweet_count;type:int;default:0;comment:'推文总数'" json:"publicMetricsTweetCount"`
	IsVerified              int8   `gorm:"column:is_verified;type:tinyint;not null;default:0;comment:'是否认证（1=是）'" json:"isVerified"`
	IsActive                int8   `gorm:"column:is_active;type:tinyint;not null;default:1;comment:'是否活跃拉取（1=是）'" json:"isActive"`
	PullFrequency           int    `gorm:"column:pull_frequency;type:int;not null;default:30;comment:'拉取频率（分钟）'" json:"pullFrequency"`
	LastPullTweetID         string `gorm:"column:last_pull_tweet_id;type:varchar(64);comment:'上次拉取的最后一条推文ID'" json:"lastPullTweetID"`
}

// TableName 指定表名（导出方法，首字母大写）
func (i *Influencer) TableName() string {
	return "t_influencer"
}

// NewInfluencer 构造函数（首字母大写，导出！！！）
// 作用：创建Influencer实例，供其他包调用
func NewInfluencer() *Influencer {
	// 注意：这里不需要传入db，db在调用Model方法时传入
	return &Influencer{}
}

// GetByUsername 示例查询方法（导出，供Handler调用）
func (i *Influencer) GetByUsername(db *gorm.DB, username string) (*Influencer, error) {
	var res Influencer
	err := db.Where("username = ?", username).First(&res).Error
	return &res, err
}
