/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datarepository

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JieranAgileCore/openapi"
	"github.com/JieranAgileCore/openapi/models"
	"github.com/JieranAgileCore/udr/internal/logger"
	"github.com/JieranAgileCore/udr/internal/sbi/producer"
	"github.com/JieranAgileCore/util/httpwrapper"
)

// HTTPGetOdbData - Retrieve ODB Data data by SUPI or GPSI
func HTTPGetOdbData(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := producer.HandleGetOdbData(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.DataRepoLog.Errorln(err)
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
