package user

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"network/model"
	"strconv"
)

type User struct {
	//
}

func (user *User) GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")

	data := rand.Intn(20)
	data = model.Models.UserModel.OptData(data)
	userId += strconv.Itoa(data)

	c.JSON(200, gin.H{"message": userId})
}

func (user *User) GetUsersInfo(c *gin.Context) {
	//
}
