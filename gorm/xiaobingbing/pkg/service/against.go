package service

import (
	"github.com/gin-gonic/gin"
)

func GetSideBySide(c *gin.Context)  {
	side := c.Param("side")

	if side == ""{
		c.JSON(400,map[string]string{
			"error":"the side can not be null",
		})
	}

	//names := strings.Split(side,",")

}
