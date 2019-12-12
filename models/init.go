package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(
					// 	new(Fl),
					//   new(FlContext),
					//   new(Report)
					)
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("")
	return prefix + name
}

// BackendUserTBName 获取 BackendUser 对应的表名称
// func FlTBName() string {
// 	return TableName("fl")
// }

// //  document
// func FlContextTBName() string {
// 	return TableName("fl_context")
// }

// //  doc
// func ReportTBName() string {
// 	return TableName("report")
// }
