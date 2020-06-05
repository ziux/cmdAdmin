package conf

import (
	"zixu.com/cmdAdmin/cache"
	"zixu.com/cmdAdmin/conf/vars"
	"zixu.com/cmdAdmin/model"
	"zixu.com/cmdAdmin/util"
)




// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	//godotenv.Load()

	// 设置日志级别
	//util.BuildLogger(os.Getenv("LOG_LEVEL"))
	util.BuildLogger(vars.LogLevel)

	// 读取翻译文件
	//if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	//	util.Log().Panic("翻译文件加载失败", err)
	//}

	// 连接数据库
	model.Database(vars.DB.MakeStr())
	if vars.Redis.Used{
		cache.Redis()

	}
}
