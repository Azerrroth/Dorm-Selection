package v1

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"login/models"
	"login/pkg/e"
)

func GetDormList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	num, _ := strconv.Atoi(c.Query("num"))

	valid := validation.Validation{}
	valid.Required(page, "page").Message("页码不能为空")
	// valid.MaxSize(page, 30, "page").Message("用户名最长为30字符")
	valid.Required(num, "num").Message("页个数不能为空 ")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if int(models.GetDormCount("true")) > (page * num) {
			code = e.SUCCESS
		} else {
			code = e.ERROR_NO_MORE_ITEMS
		}
	}
	data := models.GetDormPage((page-1)*num, num)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
