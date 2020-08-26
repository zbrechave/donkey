package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/davveo/donkey/fake"
	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
)

func LoginAdmin(context *gin.Context) {
	var loginAdminUser fake.LoginAdminUser
	result := common.ReadJson(filepath.Join(BaseDir, "data/login.admin.user.json"))
	err := json.Unmarshal([]byte(result), &loginAdminUser)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  loginAdminUser.Status,
		"message": loginAdminUser.Message,
		"data":    loginAdminUser.Data,
	})

}

func AdminItem(context *gin.Context) {

}

func CheckAdmin(context *gin.Context) {

}


func AdminList(context *gin.Context)  {
	var adminList fake.AdminList
	result := common.ReadJson(filepath.Join(BaseDir, "data/admin.list.json"))
	err := json.Unmarshal([]byte(result), &adminList)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  adminList.Status,
		"message": adminList.Message,
		"data":    adminList.Data,
	})

}