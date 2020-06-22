package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"network/editor"
	"network/model"
	"strconv"
)

type User struct {
	//
}

func (user *User) GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")

	data := rand.Intn(20)
	data = model.Ctx.UserModel.OptData(data)
	userId += strconv.Itoa(data)

	userEditor := editor.Ctx.GetUserEditor(userId)
	fmt.Println(userEditor.UserId)
	fmt.Println(userEditor.UserName)
	userEditor.Save()

	c.JSON(200, gin.H{"message": userId})
}

func (user *User) GetUsersInfo(c *gin.Context) {
	//
}
