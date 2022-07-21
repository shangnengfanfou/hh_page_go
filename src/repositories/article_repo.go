package repositories

import (
	"hh_page_go/src/models"
	"hh_page_go/src/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type ArticleRepo struct{}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{}
}

type ArticlePaginate struct {
	PageIndex  int    `json:"pageIndex"`
	PageSize   int    `json:"pageSize"`
	TagId      int    `json:"tagId"`
	UniqueId   string `json:"uniqueId"`
	ViewsCount bool   `json:"viewsCount"`
	StartTime  int    `json:"startTime"`
	EndTime    int    `json:"endTime"`
}

type ArticlePaginateResp struct {
	PageIndex int              `json:"pageIndex"`
	PageSize  int              `json:"pageSize"`
	Count     int              `json:"count"`
	Data      []models.Article `json:"data"`
}

func (h *ArticleRepo) InsertOne(record models.Article) error {
	result := utils.DB.Create(&record)
	return result.Error
}

func (h *ArticleRepo) QueryAndCount(p ArticlePaginate) (ArticlePaginateResp, error) {
	DB := utils.DB.Table("article")
	if p.TagId != 0 {
		DB = DB.Where("tag_id = ?", p.TagId)
	}
	if p.UniqueId != "" {
		DB = DB.Where("unique_id = ?", p.UniqueId)
	}
	if p.StartTime != 0 {
		DB = DB.Where("time > ?", p.StartTime)
	}
	if p.EndTime != 0 {
		DB = DB.Where("time < ?", p.EndTime)
	}
	if p.ViewsCount {
		DB = DB.Order("views_count desc")
	} else {
		DB = DB.Order("time desc")
	}
	var count int
	DB.Count(&count)
	ArticleList := []models.Article{}
	DB.Offset((p.PageIndex - 1) * p.PageSize).Limit(p.PageSize).Find(&ArticleList)
	return ArticlePaginateResp{
		PageIndex: p.PageIndex,
		PageSize:  p.PageSize,
		Count:     count,
		Data:      ArticleList,
	}, nil
}

func (h *ArticleRepo) IncrViewsCount(uniqueId string) error {
	result := utils.DB.Table("article").Where("unique_id = ?", uniqueId).Update("views_count", gorm.Expr("views_count + ?", 1))
	return result.Error
}

type ArticleInfo struct {
	Count      int               `json:"count"`
	CurrCount  int               `json:"currCount"`
	ViewsCount int               `json:"viewsCount"`
	TagsCount  []ArticleTagCount `json:"tagsCount"`
}

type ArticleTagCount struct {
	TagId int `json:"tagId"`
	Count int `json:"count"`
}

func (h *ArticleRepo) Info() (ArticleInfo, error) {
	DB := utils.DB.Table("article")
	var count int
	DB.Count(&count)
	var currCount int
	DB.Where("time >= ?", time.Now().AddDate(0, 0, -30).Unix()).Count(&currCount)
	type Result struct {
		ViewsCount int `json:"viewsCount" gorm:"column:viewsCount"`
	}
	var viewsCountRet Result
	DB.Select("sum(views_count) as viewsCount").Find(&viewsCountRet)
	var tagsCount []ArticleTagCount
	DB.Select("tag_id as tag_id, count(tag_id) as count").Group("tag_id").Scan(&tagsCount)
	return ArticleInfo{
		Count:      count,
		CurrCount:  currCount,
		ViewsCount: viewsCountRet.ViewsCount,
		TagsCount:  tagsCount,
	}, nil
}
