package server

import "github.com/gin-gonic/gin"

func applyRoutes(r *gin.Engine) {
	r.GET("/download-10-gb", download10Gb)
	r.GET("/get-total-download", getTotalDownloads)
	r.GET("/get-file-data", getFileData)
}
