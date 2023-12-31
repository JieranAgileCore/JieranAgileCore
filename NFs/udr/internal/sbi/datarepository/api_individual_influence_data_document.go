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

// HTTPApplicationDataInfluenceDataInfluenceIdDelete - Delete an individual Influence Data resource
func HTTPApplicationDataInfluenceDataInfluenceIdDelete(c *gin.Context) {
	// New HTTP request
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["influenceId"] = c.Params.ByName("influenceId")

	// Handle request
	rsp := producer.HandleApplicationDataInfluenceDataInfluenceIdDelete(req)

	if rsp.Body != nil {
		// Serialize response body
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
	} else {
		c.Data(rsp.Status, "", nil)
	}
}

// HTTPApplicationDataInfluenceDataInfluenceIdPatch -
// Modify part of the properties of an individual Influence Data resource
func HTTPApplicationDataInfluenceDataInfluenceIdPatch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// HTTPApplicationDataInfluenceDataInfluenceIdPut - Create or update an individual Influence Data resource
func HTTPApplicationDataInfluenceDataInfluenceIdPut(c *gin.Context) {
	// Get HTTP request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.DataRepoLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	// Deserialize request body
	var trafficInfluData models.TrafficInfluData
	err = openapi.Deserialize(&trafficInfluData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.DataRepoLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	// New HTTP request
	req := httpwrapper.NewRequest(c.Request, trafficInfluData)
	req.Params["influenceId"] = c.Params.ByName("influenceId")

	// Handle request
	rsp := producer.HandleApplicationDataInfluenceDataInfluenceIdPut(req)

	// Serialize response body
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
		// Get the first value associated with `Location` key in the header
		c.Header("Location", rsp.Header.Get("Location"))
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
