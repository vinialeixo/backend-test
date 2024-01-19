package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinialeixo/backend-test/src/controller"
)

func InitRouters(r *gin.RouterGroup) {

	r.POST("/createUser", controller.CreateUser)
	r.POST("/upload-file", controller.UploadFiles)
}
