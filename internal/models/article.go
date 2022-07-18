package models

import "github.com/gin-gonic/gin"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

var _ Resource = (*Article)(nil)

func NewArticle() *Article {
	return &Article{}
}

func (a *Article) Get(c *gin.Context) {
	panic("implement me")
}

func (a *Article) List(c *gin.Context) {
	panic("implement me")
}

func (a *Article) Create(c *gin.Context) {
	panic("implement me")
}

func (a *Article) Update(c *gin.Context) {
	panic("implement me")
}

func (a *Article) Delete(c *gin.Context) {
	panic("implement me")
}

func (a *Article) TableName() string {
	return "blog_article"
}
