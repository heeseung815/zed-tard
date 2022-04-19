package tlog

import (
	"fmt"
	"net/http"
	"tard/mods/storage"

	"dev.azure.com/carrotins/hdm/hdm-go.git/logging"
	"github.com/gin-gonic/gin"
)

func HandleGetTLog() func(c *gin.Context) {
	log := logging.GetLog("tlog-handler")

	return func(c *gin.Context) {
		tripId := c.Param("tripId")
		tripSerial, _ := c.GetQuery("tripSerial")
		if tripSerial == "" {
			tripSerial = "0"
		}

		blobName := fmt.Sprintf("%s-%s", tripId, tripSerial)
		result := storage.DownloadBlockBlob(blobName)

		c.JSON(http.StatusOK, gin.H{"result": result})
		log.Infof("%s %s success", tripId, tripSerial)
	}
}

func HandlePostTLog() func(c *gin.Context) {
	log := logging.GetLog("tlog-handler")

	return func(c *gin.Context) {
		log.Info("post tlog called")
	}
}

type TLogMeta struct {
	tripId     string
	tripSerial string
}
