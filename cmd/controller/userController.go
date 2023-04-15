package controller

import (
	"backend/cmd/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllUsers(c *gin.Context) {
	db, _ := model.SqlStart()

	allUsers, err := model.GetRows(db)
	if err != nil {
		log.Fatalf("getRows error err*%v", err)
	}

	out, err := json.Marshal(allUsers)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, out)
	defer db.Close()
}

func GetUser(c *gin.Context) {
	db, _ := model.SqlStart()

	paramUserId := c.Param("id")
	fmt.Println("userId:", paramUserId)

	userId, err := strconv.Atoi(paramUserId)
	if err != nil {
		log.Fatalf("GetuUser strconv.Atoi error err:%v", err)
	}

	var responseData struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Birthday  int    `json:"birthday"`
		Email     string `json:"email"`
	}

	user, _ := model.GetSingleRow(db, userId)

	responseData.ID = user.ID
	responseData.FirstName = user.FirstName
	responseData.LastName = user.LastName
	responseData.Birthday = user.Birthday
	responseData.Email = user.Email

	c.JSON(http.StatusOK, responseData)
	defer db.Close()
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	fmt.Println("userId:", userId)

	var data struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Birthday  int    `json:"birthday"`
		Email     string `json:"email"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("lastName:", data.LastName)

	c.JSON(http.StatusOK, data)
}
