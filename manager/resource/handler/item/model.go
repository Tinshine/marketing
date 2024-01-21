package item

type QueryReq struct {
	AppId int  `form:"app_id"`
	Id    *int `form:"id"`
	Name  *int `form:"id"`
}

type QueryResp struct {
	Items []ItemResp `json:"items"`
	Total int        `json:"total"`
}

type ItemResp struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
