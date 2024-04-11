package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	CreateUserReq struct {
		Username  string `json:"username"`
		AvatarURL string `json:"avatar"`
	}

	CreateUserReply struct {
		UID int `json:"uid"`
	}
)

func MakeCreateUserHandle(svc UserService) func(c *gin.Context) {

	return func(c *gin.Context) {
		var req CreateUserReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		reply, err := svc.CreateUser(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, reply)
	}
}

type (
	GetUserReq struct {
		UID int `json:"uid"`
	}

	GetUserReply struct {
		UID    int    `json:"uid"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}
)

func MakeGetUserHandle(svc UserService) func(c *gin.Context) {

	return func(c *gin.Context) {
		var req GetUserReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		reply, err := svc.GetUser(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if reply == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, reply)
	}
}
