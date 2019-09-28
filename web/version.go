package web

import (
	"github.com/gin-gonic/gin"
)

func (srv *server) Version(c *gin.Context) {
	name := c.Param("name")
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.JSON(200, &vesionOut{
		Code:           "0",
		Message:        "OK",
		ReleaseTime:    srv.config.Release.ReleaseTime,
		ReleaseVersion: srv.config.Release.ReleaseVersion,
		Name:           name,
	},
	)
}

type vesionOut struct {
	Code           string `json:"code"`
	ReleaseTime    string `json:"releaseTime"`
	ReleaseVersion string `json:"releaseVersion"`
	Message        string `json:"message"`
	Name           string `json:"name"`
}
