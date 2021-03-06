package model

type AddArticleReq struct {
	Name    string `gorm:"type:varchar(20);not null" json:"name"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(200)" json:"img"`
}

type DelArticleReq struct {
	Id uint `json:"id"`
}

type GetArticleInfoReq struct {
	Id uint `json:"id"`
}

type GetArticleListReq struct {
	Id       uint   `json:"id"`
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Cid      int    `gorm:"type:int;not null" json:"cid"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

type GetArticleInfoResp struct {
	Id           uint   `json:"id"`
	Name         string `gorm:"type:varchar(20);not null" json:"name"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(200)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

type GetArticleListData struct {
	Id           uint   `json:"id"`
	Name         string `gorm:"type:varchar(20);not null" json:"name"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
	Total        int    `json:"total"`
}

type GetArticleListResp struct {
	List  []*GetArticleListData `json:"list"`
	Total int                   `json:"total"`
}

type UpdateArticleReq struct {
	Id      uint   `json:"id"`
	Name    string `gorm:"type:varchar(20);not null" json:"name"`
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(200)" json:"img"`
}
