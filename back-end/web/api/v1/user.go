package v1

import (
	"fmt"
	"gin-vue-element-ui/models"
	"gin-vue-element-ui/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	data := make([]models.User, 0)
	for i := 1; i <= 10; i++ {
		data = append(data, models.User{Name: fmt.Sprintf("user-%d", i), Age: i * 10})
	}

	resp := web.Success(data)
	c.JSON(http.StatusOK, resp)
}
