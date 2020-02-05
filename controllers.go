package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* GET */

func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

/* POST */

func CreateExample(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

/* PUT */

func UpdateExample(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
