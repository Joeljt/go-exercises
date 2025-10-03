package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

// 模拟 new 方法，返回 Tag 类型的实例
func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
}
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
