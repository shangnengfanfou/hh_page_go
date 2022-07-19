package repositories

import (
	"hh_page_go/src/models"
	"hh_page_go/src/utils"
)

type ArticleRepo struct{}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{}
}

type ArticlePaginate struct {
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	TagId     int    `json:"tagId"`
	UniqueId  string `json:"uniqueId"`
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
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
	var count int
	DB.Count(&count)
	ArticleList := []models.Article{}
	DB.Offset((p.PageIndex - 1) * p.PageSize).Limit(p.PageSize).Order("time desc").Find(&ArticleList)
	return ArticlePaginateResp{
		PageIndex: p.PageIndex,
		PageSize:  p.PageSize,
		Count:     count,
		Data:      ArticleList,
	}, nil
}
