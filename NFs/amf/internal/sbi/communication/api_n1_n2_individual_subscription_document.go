/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package communication

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JieranAgileCore/amf/internal/logger"
	"github.com/JieranAgileCore/amf/internal/sbi/producer"
	"github.com/JieranAgileCore/openapi"
	"github.com/JieranAgileCore/openapi/models"
	"github.com/JieranAgileCore/util/httpwrapper"
)

// N1N2MessageUnSubscribe - Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation
func HTTPN1N2MessageUnSubscribe(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueContextId"] = c.Params.ByName("ueContextId")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	rsp := producer.HandleN1N2MessageUnSubscribeRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.CommLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
