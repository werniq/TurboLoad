package server

import "github.com/gin-gonic/gin"

func updateFileData(c *gin.Context) {

}

// getFileData
func getFileData(c *gin.Context) {
	filesData, err := database.GetFileInfo()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": filesData,
	})
}

// updateFileInfoAfterDownload increases file total downloads amount
func updateFileInfoAfterDownload(filename string, duration int64) error {
	return database.AfterResponseUpdate(filename, duration)
}
