package model

import "mime/multipart"

type UploadReq struct {
	File multipart.File `json:"file"`
	Size int64          `json:"size"`
}

type UploadFileResp struct {
	Url string `json:"url"`
}
