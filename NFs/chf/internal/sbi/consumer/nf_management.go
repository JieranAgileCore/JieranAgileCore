package consumer

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	chf_context "github.com/JieranAgileCore/chf/internal/context"
	"github.com/JieranAgileCore/chf/internal/logger"
	"github.com/JieranAgileCore/openapi"
	"github.com/JieranAgileCore/openapi/Nnrf_NFManagement"
	"github.com/JieranAgileCore/openapi/models"
)

func BuildNFInstance(chfContext *chf_context.CHFContext) (profile models.NfProfile, err error) {
	profile.NfInstanceId = chfContext.NfId
	profile.NfType = models.NfType_CHF
	profile.NfStatus = models.NfStatus_REGISTERED
	profile.Ipv4Addresses = append(profile.Ipv4Addresses, chfContext.RegisterIPv4)
	services := []models.NfService{}
	for _, nfService := range chfContext.NfService {
		services = append(services, nfService)
	}
	if len(services) > 0 {
		profile.NfServices = &services
	}
	profile.ChfInfo = &models.ChfInfo{
		// Todo
		// SupiRanges: &[]models.SupiRange{
		// 	{
		// 		//from TS 29.510 6.1.6.2.9 example2
		//		//no need to set supirange in this moment 2019/10/4
		// 		Start:   "123456789040000",
		// 		End:     "123456789059999",
		// 		Pattern: "^imsi-12345678904[0-9]{4}$",
		// 	},
		// },
	}
	return
}

func SendRegisterNFInstance(nrfUri, nfInstanceId string, profile models.NfProfile) (
	resouceNrfUri string, retrieveNfInstanceID string, err error) {
	// Set client and set url
	configuration := Nnrf_NFManagement.NewConfiguration()
	configuration.SetBasePath(nrfUri)
	client := Nnrf_NFManagement.NewAPIClient(configuration)

	var res *http.Response
	for {
		_, res, err = client.NFInstanceIDDocumentApi.RegisterNFInstance(context.TODO(), nfInstanceId, profile)
		if err != nil || res == nil {
			logger.ConsumerLog.Errorf("CHF register to NRF Error[%v]", err)
			time.Sleep(2 * time.Second)
			continue
		}
		defer func() {
			if resCloseErr := res.Body.Close(); resCloseErr != nil {
				logger.ConsumerLog.Errorf("RegisterNFInstance response body cannot close: %+v", resCloseErr)
			}
		}()
		status := res.StatusCode
		if status == http.StatusOK {
			// NFUpdate
			break
		} else if status == http.StatusCreated {
			// NFRegister
			resourceUri := res.Header.Get("Location")
			resouceNrfUri = resourceUri[:strings.Index(resourceUri, "/nnrf-nfm/")]
			retrieveNfInstanceID = resourceUri[strings.LastIndex(resourceUri, "/")+1:]
			break
		} else {
			fmt.Println(fmt.Errorf("handler returned wrong status code %d", status))
			fmt.Println("NRF return wrong status code", status)
		}
	}
	return resouceNrfUri, retrieveNfInstanceID, err
}

func SendDeregisterNFInstance() (problemDetails *models.ProblemDetails, err error) {
	logger.ConsumerLog.Infof("Send Deregister NFInstance")

	chfSelf := chf_context.GetSelf()
	// Set client and set url
	configuration := Nnrf_NFManagement.NewConfiguration()
	configuration.SetBasePath(chfSelf.NrfUri)
	client := Nnrf_NFManagement.NewAPIClient(configuration)

	var res *http.Response

	res, err = client.NFInstanceIDDocumentApi.DeregisterNFInstance(context.Background(), chfSelf.NfId)
	if err == nil {
		return
	} else if res != nil {
		defer func() {
			if resCloseErr := res.Body.Close(); resCloseErr != nil {
				logger.ConsumerLog.Errorf("DeregisterNFInstance response cannot close: %+v", resCloseErr)
			}
		}()
		if res.Status != err.Error() {
			return
		}
		problem := err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails)
		problemDetails = &problem
	} else {
		err = openapi.ReportError("server no response")
	}
	return
}
