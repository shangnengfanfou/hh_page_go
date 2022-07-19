package models

type Article struct {
	Id        int    `json:"id" gorm:"column:id;primary_key"`
	TagId     int    `json:"tagId" gorm:"column:tag_id"`
	UniqueId  string `json:"uniqueId" gorm:"column:unique_id"`
	Title     string `json:"title" gorm:"title"`
	Summary   string `json:"summary" gorm:"summary"`
	BannerUrl string `json:"bannerUrl" gorm:"column:banner_url"`
	Time      int64  `json:"time" gorm:"column:time"`
}

func (Article) TableName() string {
	return "article"
}
