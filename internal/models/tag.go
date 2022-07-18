package models

import "github.com/gin-gonic/gin"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

var _ Resource = (*Tag)(nil)

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Get(c *gin.Context) {
	panic("implement me")
}

func (t *Tag) List(c *gin.Context) {
	panic("implement me")
}

func (t *Tag) Create(c *gin.Context) {
	panic("implement me")
}

func (t *Tag) Update(c *gin.Context) {
	panic("implement me")
}

func (t *Tag) Delete(c *gin.Context) {
	panic("implement me")
}

func (t Tag) TableName() string {
	return "blog_tag"
}
