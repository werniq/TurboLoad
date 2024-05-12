package server

import "github.com/gin-gonic/gin"

func applyRoutes(r *gin.Engine) {
	// multiple downloads
	//r.GET("/download-100-gb", download100Gb)
	//r.GET("/download-50-gb", download50Gb)
	r.GET("/download-10-gb", download10Gb)
	r.GET("/download-1-gb", download1Gb)

	r.GET("/get-total-downloads", getTotalDownloads)
	r.GET("/get-all-file-data", getFileData)
}
