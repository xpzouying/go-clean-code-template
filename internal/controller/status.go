package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xpzouying/go-clean-code-template/internal/constant"
)

type RespStatus struct {
	Version     string `json:"version"`
	BuildTime   string `json:"build_time"`
	GitBranch   string `json:"git_branch"`
	GitRevision string `json:"git_revision"`
}

// HandleGetStatus handles the GET /status API.
func HandleGetStatus(c *gin.Context) {

	WriteSuccessResponse(c, &RespStatus{
		Version:     constant.Version,
		BuildTime:   constant.BuildTime,
		GitBranch:   constant.GitBranch,
		GitRevision: constant.GitRevision,
	})
}
