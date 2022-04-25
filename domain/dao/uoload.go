package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

type UploadDao interface {
	UpLoadFile(c *gin.Context, file multipart.File, fileSize int64, bucket string, accessKey string,
		secretKey string) (key string, err error)
}

func (d *dao) UpLoadFile(c *gin.Context, file multipart.File, fileSize int64, bucket string, accessKey string,
	secretKey string) (key string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return
	}
	key = ret.Key
	return

}
