package models

type News struct {
	Id           int64
	TagId        int64
	SiteId       int64
	UniqueCodeId int64
	Url          string
	Title        string
	CreateTime   int64
}
