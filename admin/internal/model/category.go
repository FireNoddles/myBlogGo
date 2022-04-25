package model

type UpdateCategoryReq struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type DeleteCategoryReq struct {
	Id uint `json:"id"`
}

type GetCategoryReq struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type AddCategoryReq struct {
	Name string `json:"name"`
}

type GetCategoryListResp struct {
	List []*CategoryList `json:"list"`
}

type CategoryList struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
