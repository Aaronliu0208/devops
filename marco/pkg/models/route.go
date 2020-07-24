package models

import "casicloud.com/ylops/marco/pkg/nginx"

//Route represent nginx location
type Route struct {
	/* 匹配符号
	https://blog.csdn.net/yuan_xw/article/details/51333451
	= 开头表示精确匹配
	"" 空字符串 默认匹配
	^~开头表示uri以某个常规字符串开头，理解为匹配 url路径即可。
	nginx不对url做编码，因此请求为/static/20%/aa，
	可以被规则^~ /static/ /aa匹配到（注意是空格）
	~ 开头表示区分大小写的正则匹配
	~* 开头表示不区分大小写的正则匹配
	!~和!~*分别为区分大小写不匹配及不区分大小写不匹配的正则
	*/
	Pattern string
	// Path 正则表达式匹配的路径
	Path   string
	Root   string
	Extras nginx.Options
}

//Marshal implements directive Marshaler
func (r Route) Marshal() ([]nginx.Directive, error) {
	location := "location " + r.Pattern + " " + r.Path
	locationBlk := nginx.NewBlock(location)
	if len(r.Root) > 0 {
		locationBlk.AddKVOption("root", r.Root)
	}

	locationBlk.AddInterface(r.Extras)
	var ds nginx.Directives
	ds = append(ds, locationBlk)
	return ds, nil
}
