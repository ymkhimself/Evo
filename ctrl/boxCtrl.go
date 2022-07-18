/* COPYRIGHT NOTICE
 * 作者     ：ymk
 * 创建时间 ：2022/07/12 11:19
 * 描述     ：管理靶机相关
 */
package ctrl

import (
	"Evo/db"
	"Evo/docker"
	"Evo/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加靶机
type BoxForm struct {
	ChallengeId uint   `binding:"required"`
	TeamId      uint   `binding:"required"`
	Ip          string `binding:"required"`
	Image       string `binding:"required"`
	Port        string `binding:"required"`
	SshUser     string `json:"sshUser" binding:"required"`
	SshPwd      string `json:"sshPwd" binding:"required"`
}

func PostBox(c *gin.Context) {

	// 绑定参数
	var boxForm BoxForm
	if err := c.ShouldBind(&boxForm); err != nil {
		log.Println(err.Error())
		Fail(c, "参数绑定失败", nil)
		return
	}

	var box model.Box
	db := db.GetDB()
	db.Where("challenge_id = ? and team_id = ?").First(&box)
	if box.ID != 0 {
		Fail(c, "靶机已存在", nil)
		return
	}

	var challenge model.Challenge
	db.Where("id = ?", boxForm.ChallengeId).First(&challenge)
	if challenge.ID == 0 {
		Fail(c, "队伍不存在", nil)
		return
	}
	name := challenge.Title + strconv.Itoa(int(boxForm.TeamId))
	// 传入镜像名，ip，
	if err := docker.StartContainer(boxForm.Image, name, boxForm.Ip, boxForm.SshPwd); err != nil {
		log.Println(err.Error())
	}

}

// 停止靶机
func StopBox(c *gin.Context) {

}

// 移除靶机
func RemoveBox(c *gin.Context) {

}

// 重启靶机
func RestartBox(c *gin.Context) {

}
