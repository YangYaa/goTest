package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context, profile DbModel) {
	err := c.ShouldBindJSON(profile)
	if err != nil {
		fmt.Println("The Unmarshal JSON failed")
		return
	}
	if err := profile.Check(); err != nil {
		return
	}
	if err := profile.Create(); err != nil {
		fmt.Println("Insert into database error")
		return
	}
	fmt.Println("Insert into database success")
}

func AddHandler(c *gin.Context, profile DbModel) {
	err := c.ShouldBindJSON(profile)
	if err != nil {
		fmt.Println("The Unmarshal JSON failed")
		return
	}
	if err := profile.Check(); err != nil {
		return
	}
	if err := profile.Create(); err != nil {
		fmt.Println("Insert into database error")
		return
	}
	fmt.Println("Insert into database success")
}

func UpdateHandler(c *gin.Context, profile DbModel) {

	err := c.ShouldBindJSON(profile)
	if err != nil {
		fmt.Println("The Unmarshal JSON failed")
		return
	}
	if err := profile.Check(); err != nil {
		return
	}
	if err := profile.Update(); err != nil {
		fmt.Println("Update into database error")
		return
	}
	fmt.Println("Update into database success")
}

func DeleteHandler(c *gin.Context, profile DbModel) {

	err := c.ShouldBindJSON(profile)
	if err != nil {
		fmt.Println("The Unmarshal JSON failed")
		return
	}
	if err := profile.Check(); err != nil {
		return
	}
	if r, err := profile.Query(); err != nil {
		fmt.Println("Delete from database error")
		return
	} else {
		count := len(r)
		if count <= 0 {
			return
		}
		for _, value := range r {
			if err := value.Delete(); err != nil {
				return
			} else {
				fmt.Println("Delete from database success")
			}
		}
	}
}
