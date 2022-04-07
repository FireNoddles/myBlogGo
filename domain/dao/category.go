package dao

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/domain/model"
)

type CategoryDao interface {
	GetCategoryList(c *gin.Context, where string, paras ...interface{}) (categorys []*model.Category, err error)
	UpdateCategory(c *gin.Context, category *model.Category) (err error)
	AddCategory(c *gin.Context, category *model.Category) (err error)
	DelCategory(c *gin.Context, category *model.Category) (err error)
}

func (d *dao) GetCategoryList(c *gin.Context, where string, paras ...interface{}) (categorys []*model.Category, err error) {
	return d.GetCategoryListInTx(c, where, paras...)
}

func (d *dao) UpdateCategory(c *gin.Context, category *model.Category) (err error) {
	data := map[string]interface{}{
		"name":         category.Name,
		"updated_time": category.UpdatedTime,
	}
	return d.UpdateCategoryInTx(c, category, data)
}

func (d *dao) AddCategory(c *gin.Context, category *model.Category) (err error) {
	return d.AddCategoryInTx(c, category)
}

func (d *dao) DelCategory(c *gin.Context, category *model.Category) (err error) {
	data := map[string]interface{}{
		"state":        category.State,
		"updated_time": category.UpdatedTime,
	}
	return d.UpdateCategoryInTx(c, category, data)
}

func (d *dao) GetCategoryListInTx(c *gin.Context, where string, paras ...interface{}) (categorys []*model.Category, err error) {
	err = d.DbMyBlog.Where(where, paras...).Find(&categorys).Error
	if err != nil {
		log.Error("d.GetCategoryListInTx err ->", err)
		return
	}
	return
}

func (d *dao) UpdateCategoryInTx(c *gin.Context, category *model.Category, data map[string]interface{}) (err error) {
	err = d.DbMyBlog.Model(&category).Update(data).Error
	if err != nil {
		log.Error("d.UpdateCategoryInTx err -> ", err)
		return
	}
	return
}

func (d *dao) AddCategoryInTx(c *gin.Context, category *model.Category) (err error) {
	err = d.DbMyBlog.Create(&category).Error
	if err != nil {
		log.Error("d.AddCategoryInTx err ->", err)
		return
	}
	return
}

func (d *dao) DelCategoryInTx(c *gin.Context, category *model.Category) (err error) {
	//TODO implement me
	panic("implement me")
}
