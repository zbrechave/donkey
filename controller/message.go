package controller

import (
	"github.com/davveo/donkey/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"net/http"
	"path/filepath"
)

func MessageUnRead(context *gin.Context) {
	result := common.ReadJson(filepath.Join(BaseDir, "data/message.user.unread.json"))
	dataMap, _ := gjson.Parse(result).Value().(map[string]interface{})
	context.JSON(http.StatusOK, dataMap)
}
