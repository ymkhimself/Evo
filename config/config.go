/* COPYRIGHT NOTICE
 * 作者     ：ymk
 * 创建时间 ：2022/07/09 14:14
 * 描述     ：从配置文件初始化一些配置
 */
package config

import (
	"os"

	"github.com/spf13/viper"
)

// 初始化设置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
