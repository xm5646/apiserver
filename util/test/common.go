/**
 * 功能描述:  为单元测试提供配置文件解析和数据库初始化
 * @Date: 2019-10-26
 * @author: lixiaoming
 */
package test

import (
	"apiserver/model"
	"apiserver/pkg/config"
	"github.com/spf13/pflag"
)

var (
	// 指定配置文件路径
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
	//cfg = pflag.StringP("config", "c", "/home/cnopens/dashuo/api-server/conf/config.yaml", "apiserver config file path.")
)

// 加载配置文件,初始化数据库, 测试完成请手动关闭数据库: model.DB.Close()
func LoadConfigAndInitDBConnection() {
	// 解析命令行参数
	//pflag.Parse()
	*cfg = "../../conf/config.yaml"
	//dir, err := os.Getwd()
	//if err != nil {
	//
	//} else {
	//	dir = strings.Replace(dir, "\\", "/", -1)
	//}
	//fmt.Println(dir)
	//os.Setenv("CHASSIS_HOME", dir)
	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// 初始化数据库
	model.DB.Init()
}

func GetTestFuncWithDB(f func()) {
	LoadConfigAndInitDBConnection()
	f()
	defer model.DB.Close()
}
