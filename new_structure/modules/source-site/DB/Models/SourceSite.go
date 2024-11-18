package Models

type SourceSite struct {
	ID            int                    `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	Title         string                 `gorm:"column:title;type:varchar(45)"`
	Config        map[string]interface{} `gorm:"column:config;serializer:json"`
	CrawlInterval int                    `gorm:"column:crawl_interval;type:mediumint"`
	MaxPosts      int                    `gorm:"column:Max_posts;type:mediumint"`
}
