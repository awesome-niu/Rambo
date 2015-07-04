package models

/**
 *	爬虫规则实体类
 */
type Rule struct {
	Id      int64
	Header  string //http请求的header参数
	Cookie  string //http请求的cookie参数
	Query   string //http请求的参数
	Login   string //与登陆相关的参数，都放这里
	Content string //采用正则表达式
	Method  string //请求方法，比如：get,post
	SiteId  int64
}
