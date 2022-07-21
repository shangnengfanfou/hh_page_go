package apis

import (
	"hh_page_go/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
	articleService *services.ArticleService
}

func NewArticleApi() *ArticleApi {
	return &ArticleApi{articleService: services.NewArticleService()}
}

func (h *ArticleApi) Add(c *gin.Context) {
	var body services.AddArticleBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	err = h.articleService.AddArticle(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK"})
}

func (h *ArticleApi) Page(c *gin.Context) {
	var body services.ArticlePaginate
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	ret, err := h.articleService.GetArticles(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK", "data": ret})
}

func (h *ArticleApi) IncrViewsCount(c *gin.Context) {
	uniqueId := c.Param("uid")
	if uniqueId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "500", "msg": "uniqueId is required"})
		return
	}
	err := h.articleService.IncrViewsCount(uniqueId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "500", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK"})
}

func (h *ArticleApi) Info(c *gin.Context) {
	ret, err := h.articleService.Info()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "500", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "OK", "data": ret})
}
