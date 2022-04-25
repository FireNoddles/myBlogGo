package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my_blog/admin/internal/model"
	. "my_blog/utils"
)

func (s *Service) UploadFile(c *gin.Context, req *model.UploadReq) (state int, message string, data *model.UploadFileResp) {
	key, err := s.dmDao.UpLoadFile(c, req.File, req.Size, s.conf.QiNiu.QiNiu.Bucket, s.conf.QiNiu.QiNiu.AccessKey, s.conf.QiNiu.QiNiu.SecretKey)
	if err != nil {
		log.Error("s.dmDao.UpLoadFile err -> ", err)
		state = ERROR
		message = GetErrMsg(state)
		return
	}
	data = new(model.UploadFileResp)
	data.Url = s.conf.QiNiu.QiNiu.Server + key
	state = SUCCSE
	message = GetErrMsg(state)
	return
}
