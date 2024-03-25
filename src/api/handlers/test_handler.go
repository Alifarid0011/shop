package handlers

import (
	"fmt"
	"net/http"

	"github.com/Alifarid0011/shop/src/api/helper"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// Test Sawagger
// @Summary testing swagger
// @Description just test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse{} "Success"
// @Router /v1/test [get]
func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, helper.BaseHttpResponse{"test", true, 200, nil})
}

// GetUser  Sawagger
// @Summary GetUser
// @Description just test
// @Tags kir
// @Accept json
// @Param id path int true "user id"
// @Param test path string false "test unrequited param"
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse{} "Success"
// @Failure 400 {object} helper.BaseHttpResponse{} "Failed"
// @Router /v1/test/user/{id} [post]
// @Security AuthBearer
func (h *TestHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, helper.BaseHttpResponse{gin.H{
		"data": "user by id",
		"id":   c.Param("id"),
	}, true, 200, nil})
}
func (h *TestHandler) GetUserByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "user by id",
		"id":     c.Param("name"),
	})
}

func (h *TestHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "create user",
	})
}

type TestBodyFetch struct {
	Name   string `json:"name"`
	Adders string `json:"address"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func (h *TestHandler) GetQueryString(c *gin.Context) {
	tbf := TestBodyFetch{}
	err := c.ShouldBindQuery(&tbf)
	Ids := c.QueryArray("id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": tbf,
		"ids":    Ids,
	})
}

// GetBody  Sawagger
// @Summary GetBody
// @Description just test
// @Tags test
// @Accept json
// @Produce json
// @Param data body TestBodyFetch true "test body data"
// @Success 200 {object} helper.BaseHttpResponse{} "Success"
// @Failure 400 {object} helper.BaseHttpResponse{} "Failed"
// @Router /v1/test/body [post]
func (h *TestHandler) GetBody(c *gin.Context) {
	tbf := TestBodyFetch{}
	err := c.ShouldBindJSON(&tbf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": tbf,
	})
}

func (h *TestHandler) MultiFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]
	for _, file := range files {
		fmt.Println(file.Filename)
		err := c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

type TestHeaderFetch struct {
	AppName string
	Code    int
}

func (h *TestHandler) GetHeader(c *gin.Context) {
	thf := TestHeaderFetch{}
	err := c.ShouldBindHeader(&thf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": thf,
	})
}
