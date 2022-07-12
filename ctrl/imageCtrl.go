/* COPYRIGHT NOTICE
 * 作者     ：ymk
 * 创建时间 ：2022/07/12 11:19
 * 描述     ：管理镜像相关
 */
package ctrl

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func PostImage(c *gin.Context) {
	// 解析表单
	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err.Error())
		Fail(c, "上传出错", nil)
		return
	}
	files := form.File["image"]
	name := c.PostForm("name")
	log.Println(name)
	dst := "upload/image/" + name + "/"
	if err = os.MkdirAll(dst, 0775); err != nil {
		log.Println(err.Error())
		Error(c, "保存失败", nil)
		return
	}
	//遍历所有文件
	for _, file := range files {
		err := c.SaveUploadedFile(file, dst+file.Filename)
		if err != nil {
			log.Println(err.Error())
			err := os.RemoveAll(dst) //上传失败，把失败的东西删掉
			if err != nil {
				Error(c, "保存失败，请管理员查看日志", nil)
				log.Println(err.Error())
				return
			}
			Error(c, "保存失败", nil)
			return
		}
	}
	Success(c, "上传成功", nil)
}
