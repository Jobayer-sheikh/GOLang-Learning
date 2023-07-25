package dto

type Admin struct {
	Id           int32  `json:"id"`
	AdminNameEng string `json:"adminNameEng"`
	AdminNameBng string `json:"adminNameBng"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
}
