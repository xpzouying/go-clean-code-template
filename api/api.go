package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/xpzouying/go-clean-code-template/internal/controller"
	"github.com/xpzouying/go-clean-code-template/internal/middleware"
)

// UserService is a service for user.
// It should be implemented by the business logic layer.
// Inject it into the API layer.
type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserReply, error)
}

func RegisterHTTPServer(r *gin.Engine, svc UserService) {
	setAPIRouter(r)

	setUserRouter(r, svc)
}

func SetRouter(r *gin.Engine, svc UserService) {

	r.Use(middleware.GinZapLogger())

	r.GET("/status", controller.HandleGetStatus)

	setAPIRouter(r)
}

func setAPIRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")

	apiRouter.GET("/status", controller.HandleGetStatus)
}

func setUserRouter(r *gin.Engine, svc UserService) {

	userRouter := r.Group("/user")

	userRouter.POST("/create", MakeCreateUserHandle(svc))
}
