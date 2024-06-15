package models

type ServiceResponse struct {
	Id          int64   `json:"id" gorm:"column:id"`
	Image       string  `json:"image" gorm:"column:image"`
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	TypeName    string  `json:"type_name" gorm:"column:type_name"`
	Commission  float64 `json:"commission" gorm:"column:commission"`
	Term        int64   `json:"term" gorm:"column:term"`
	MinSum      float64 `json:"min_sum" gorm:"column:min"`
	MaxSum      float64 `json:"max_sum" gorm:"column:max"`
}

type AccountCreateReq struct {
	ServiceId int64
	Client    *Client
}

type SrvReq struct {
	ClientId  int64   `json:"client_id" gorm:"client_id"`
	ServiceId int64   `json:"service_id" gorm:"column:service_id"`
	Amount    float64 `json:"amount" gorm:"column:amount"`
	Term      int     `json:"term"`
}
