package services

import (
	"hh_page_go/src/models"
	"hh_page_go/src/repositories"
	"time"
)

type ArticleService struct {
	ArticleRepo *repositories.ArticleRepo
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		ArticleRepo: repositories.NewArticleRepo(),
	}
}

type AddArticleBody struct {
	TagId     int    `json:"tagId"`
	UniqueId  string `json:"uniqueId"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	BannerUrl string `json:"bannerUrl"`
}

type ArticlePaginate struct {
	repositories.ArticlePaginate
}

func (h *ArticleService) AddArticle(article AddArticleBody) error {
	a := models.Article{
		TagId:     article.TagId,
		UniqueId:  article.UniqueId,
		Title:     article.Title,
		Summary:   article.Summary,
		BannerUrl: article.BannerUrl,
		Time:      time.Now().Unix(),
	}
	err := h.ArticleRepo.InsertOne(a)
	if err != nil {
		return err
	}
	return nil
}

func (h *ArticleService) GetArticles(p ArticlePaginate) (repositories.ArticlePaginateResp, error) {
	data, err := h.ArticleRepo.QueryAndCount(p.ArticlePaginate)
	if err != nil {
		return repositories.ArticlePaginateResp{}, err
	}
	return data, nil
}

func (h *ArticleService) IncrViewsCount(uniqueId string) error {
	err := h.ArticleRepo.IncrViewsCount(uniqueId)
	if err != nil {
		return err
	}
	return nil
}

type ArticleInfo struct {
	repositories.ArticleInfo
}

func (h *ArticleService) Info() (ArticleInfo, error) {
	data, err := h.ArticleRepo.Info()
	if err != nil {
		return ArticleInfo{}, err
	}
	return ArticleInfo{
		ArticleInfo: data,
	}, nil
}
