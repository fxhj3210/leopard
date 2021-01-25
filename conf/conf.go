package conf

import (
	"github.com/joho/godotenv"
	"leopard/cache"
	"leopard/global"
	"leopard/model"
	"os"
)

//var Log = logrus.New()
func Init() {
	//读取环境变量
	_ = godotenv.Load()

	//设置日志
	global.LogrusInit()

	//连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))

	//连接Redis
	cache.Redis()
}
