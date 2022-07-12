/* COPYRIGHT NOTICE
 * 作者     ：ymk
 * 创建时间 ：2022/07/09 20:19
 * 描述     ：测试config能否初始化
 */
package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	//获取一个路径,到当前目录
	wd, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wd)
	err := viper.ReadInConfig()
	if err != nil {
		t.Log(err.Error())
	}
	//读取config中关于数据库的设置
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	pwd := viper.GetString("datasource.pwd")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utfmb4&parseTime=True&loc=Local",
		username, pwd, host, port, database)
	t.Log(dsn)
}
