package v1

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"register/models"
	"register/pkg/e"
	"register/pkg/setting"
	"register/pkg/util"
)

var jwtValidityPeriod = setting.JwtValidityPeriod
var LoginList map[string]interface{}

func GetUser(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	// data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

}

func Register(c *gin.Context) {
	// 检查表单
	name := c.Query("name")
	password := c.Query("password")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("用户名不能为空")
	valid.MaxSize(name, 30, "name").Message("用户名最长为30字符")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MaxSize(password, 30, "password").Message("密码最长为30字符")
	valid.MinSize(password, 6, "password").Message("密码最少为6字符")

	crypted, salt, err := util.Encrypt(password)
	if err != nil {
		log.Fatal(2, "Fail to get encrypted password: %v", err)
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistUserByName(name) {
			code = e.SUCCESS
			models.AddUser(name, string(crypted), string(salt))
		} else {
			code = e.ERROR_EXIST_USER
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func Login(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")

	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	session := sessions.Default(c)
	if !valid.HasErrors() {
		if !models.ExistUserByName(name) {
			code = e.ERROR_NOT_EXIST_USER
		} else {
			if models.ExistUserByNameAndPassword(name, password) {
				// 登录成功
				code = e.SUCCESS
				token, err := util.GenerateToken(name, password, jwtValidityPeriod)
				session.Set("uid", name)
				session.Set("status", "online")
				session.Save()
				if err != nil {
					code = e.ERROR_AUTH_TOKEN
				} else {
					data["token"] = token
					data["name"] = name
					data["uid"] = session.Get("uid")
				}
				// Open session
			} else {
				// 密码错误
				code = e.ERROR_PASSWORD_ERROR
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
