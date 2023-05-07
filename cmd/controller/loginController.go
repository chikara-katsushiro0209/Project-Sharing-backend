package controller

import (
	"backend/cmd/domain"
	"backend/cmd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db, _ := model.SqlStart()

	var data struct {
		ID       int    `json:"id"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := model.GetSingleRow(db, data.ID)

	res := domain.ConfirmPassword([]byte(data.Password), []byte(user.Password))

	data.FullName = user.LastName + user.FirstName

	if res {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusBadRequest, data)
	}

	defer db.Close()
}
