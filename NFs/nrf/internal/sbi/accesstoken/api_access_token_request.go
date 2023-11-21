/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package accesstoken

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/JieranAgileCore/nrf/internal/logger"
	"github.com/JieranAgileCore/nrf/internal/sbi/producer"
	"github.com/JieranAgileCore/openapi/models"
	"github.com/JieranAgileCore/util/httpwrapper"
)

// AccessTokenRequest - Access Token Request
func HTTPAccessTokenRequest(c *gin.Context) {
	logger.AccTokenLog.Infoln("In HTTPAccessTokenRequest")
	var accessTokenReq models.AccessTokenReq
	var r *http.Request = c.Request

	// Request parser
	err := r.ParseForm()
	if err != nil {
		logger.AccTokenLog.Errorf(err.Error())
		return
	}
	rt := reflect.TypeOf(accessTokenReq)
	for key, value := range r.PostForm {
		var name string
		var vt reflect.Type
		for i := 0; i < rt.NumField(); i++ {
			if tag := rt.Field(i).Tag.Get("yaml"); tag == key {
				name = rt.Field(i).Name
				vt = rt.Field(i).Type
				break
			}
		}
		if vt.Name() == "string" || vt.Name() == "NfType" {
			reflect.ValueOf(&accessTokenReq).Elem().FieldByName(name).SetString(value[0])
		} else {
			plmnid := models.PlmnId{}
			err = json.Unmarshal([]byte(value[0]), &plmnid)
			if err != nil {
				problemDetail := "[Request Body] " + err.Error()
				rsp := models.ProblemDetails{
					Title:  "Json Unmarshal Error",
					Status: http.StatusBadRequest,
					Detail: problemDetail,
				}
				logger.AccTokenLog.Errorln(problemDetail)
				c.JSON(http.StatusBadRequest, rsp)
				return
			}
			reflectvalue := reflect.ValueOf(&plmnid)
			reflect.ValueOf(&accessTokenReq).Elem().FieldByName(name).Set(reflectvalue)
		}
	}

	err = c.Bind(&accessTokenReq)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.AccTokenLog.Warnln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(c.Request, accessTokenReq)
	req.Params["paramName"] = c.Params.ByName("paramName")

	httpResponse := producer.HandleAccessTokenRequest(req)

	c.JSON(httpResponse.Status, httpResponse.Body)
}
