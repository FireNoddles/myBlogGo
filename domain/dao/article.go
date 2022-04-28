package dao

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"my_blog/domain/model"
)

type ArticleDao interface {
	GetArticleList(c *gin.Context, pageSize int, pageNum int, where string, paras ...interface{}) (articles []*model.Article, total int, err error)
	GetArticleInfo(c *gin.Context, where string, paras ...interface{}) (articles []*model.Article, err error)
	UpdateArticle(c *gin.Context, article *model.Article) (err error)
	AddArticle(c *gin.Context, article *model.Article) (err error)
	DelArticle(c *gin.Context, article *model.Article) (err error)
}

func (d *dao) GetArticleList(c *gin.Context, pageSize int, pageNum int, where string, paras ...interface{}) (articles []*model.Article, total int, err error) {
	var wg errgroup.Group

	wg.Go(func() error {
		articles, err = d.GetArticleListInTx(c, pageSize, pageNum, where, paras...)
		return err
	})

	wg.Go(func() error {
		total, err = d.GetArticleListCountInTx(c, where, paras...)
		return err
	})

	if err = wg.Wait(); err != nil {
		log.Error("d.GetArticleList err ->", err)
		return
	}
	return
}

func (d *dao) GetArticleInfo(c *gin.Context, where string, paras ...interface{}) (articles []*model.Article, err error) {
	articles, err = d.GetArticleInfoInTx(c, where, paras...)
	return articles, err
}

func (d *dao) UpdateArticle(c *gin.Context, article *model.Article) (err error) {
	data := map[string]interface{}{
		"name":         article.Name,
		"updated_time": article.UpdatedTime,
		"cid":          article.Cid,
		"desc":         article.Desc,
		"content":      article.Content,
		"img":          article.Img,
	}
	return d.UpdateArticleInTx(c, article, data)
}

func (d *dao) AddArticle(c *gin.Context, article *model.Article) (err error) {
	return d.AddArticleInTx(c, article)
}

func (d *dao) DelArticle(c *gin.Context, article *model.Article) (err error) {
	data := map[string]interface{}{
		"state":        article.State,
		"updated_time": article.UpdatedTime,
	}
	return d.UpdateArticleInTx(c, article, data)
}

func (d *dao) GetArticleListInTx(c *gin.Context, pageSize int, pageNum int, where string, paras ...interface{}) (articles []*model.Article, err error) {
	err = d.DbMyBlog.Where(where, paras...).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil {
		log.Error("d.GetArticleListInTx err ->", err)
		return
	}
	return
}

func (d *dao) GetArticleListCountInTx(c *gin.Context, where string, paras ...interface{}) (count int, err error) {
	err = d.DbMyBlog.Model(&model.Article{}).Where(where, paras...).Count(&count).Error
	if err != nil {
		log.Error("d.GetArticleListCount err ->", err)
		return
	}
	return
}

func (d *dao) GetArticleInfoInTx(c *gin.Context, where string, paras ...interface{}) (article []*model.Article, err error) {
	err = d.DbMyBlog.Where(where, paras...).Find(&article).Error
	if err != nil {
		log.Error("d.GetCategoryListInTx err ->", err)
		return
	}
	return
}

func (d *dao) UpdateArticleInTx(c *gin.Context, article *model.Article, data map[string]interface{}) (err error) {
	err = d.DbMyBlog.Model(&article).Update(data).Error
	if err != nil {
		log.Error("d.UpdateArticleInTx err -> ", err)
		return
	}
	return
}

func (d *dao) AddArticleInTx(c *gin.Context, article *model.Article) (err error) {
	err = d.DbMyBlog.Create(&article).Error
	if err != nil {
		log.Error("d.AddArticleInTx err ->", err)
		return
	}
	return
}
