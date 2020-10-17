package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type Endpoint struct {

}

func NewEndpoint() *Endpoint {
	return &Endpoint{

	}
}

//รับ INPUT แปลงค่า
func (ep *Endpoint) PingEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Debugf("Check Heartbeat")

	result, err := checkHeartbeat()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)

}
