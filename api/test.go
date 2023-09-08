package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileInfo struct {
	Slack_name string `form:"slack_name"`
	Track      string `form:"track"`
}

func main() {
	r := gin.Default()
	r.GET("/api", getUserInfo)
	r.Run()
}


func getUserInfo(c *gin.Context) {
	var user ProfileInfo

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Inavlid credentials",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"slack_name":      user.Slack_name,
		"current_day":     time.Now().Weekday().String(),
		"Track":           user.Track,
		"github_file_url": "https://github.com/username/repo/blob/main/file_name.ext",
		"github_repo_url": "https://github.com/username/repo",
		"status_code":     http.StatusOK,
		"utc_time":        time.Now().UTC(),
	})
}
