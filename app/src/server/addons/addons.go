// Package addons 业务插件导入入口
// 导入所有业务插件，触发各插件的 init() 函数完成自动注册
// 新增业务插件只需在此添加 import 即可
package addons

import (
	// 示例业务插件
	_ "whitestone.top/prism-example-site/addons/site-info"
	// 数据总览插件
	_ "whitestone.top/prism-example-site/addons/dashboard"
	// 示例插件
	_ "whitestone.top/prism-example-site/addons/example"
	// 消息记录插件
	_ "whitestone.top/prism-example-site/addons/messages"
)
