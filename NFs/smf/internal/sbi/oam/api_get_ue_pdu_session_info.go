package oam

import (
	"github.com/gin-gonic/gin"

	"github.com/JieranAgileCore/smf/internal/sbi/producer"
	"github.com/JieranAgileCore/util/httpwrapper"
)

func HTTPGetUEPDUSessionInfo(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["smContextRef"] = c.Params.ByName("smContextRef")

	smContextRef := req.Params["smContextRef"]
	HTTPResponse := producer.HandleOAMGetUEPDUSessionInfo(smContextRef)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
