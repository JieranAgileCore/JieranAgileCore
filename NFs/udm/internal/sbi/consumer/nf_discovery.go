package consumer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/JieranAgileCore/openapi/Nnrf_NFDiscovery"
	"github.com/JieranAgileCore/openapi/models"
	udm_context "github.com/JieranAgileCore/udm/internal/context"
	"github.com/JieranAgileCore/udm/internal/logger"
	"github.com/JieranAgileCore/udm/internal/util"
)

const (
	NFDiscoveryToUDRParamNone int = iota
	NFDiscoveryToUDRParamSupi
	NFDiscoveryToUDRParamExtGroupId
	NFDiscoveryToUDRParamGpsi
)

func SendNFIntances(nrfUri string, targetNfType, requestNfType models.NfType,
	param Nnrf_NFDiscovery.SearchNFInstancesParamOpts,
) (result models.SearchResult, err error) {
	configuration := Nnrf_NFDiscovery.NewConfiguration()
	configuration.SetBasePath(nrfUri) // addr
	clientNRF := Nnrf_NFDiscovery.NewAPIClient(configuration)

	result, res, err1 := clientNRF.NFInstancesStoreApi.SearchNFInstances(context.TODO(), targetNfType,
		requestNfType, &param)
	if err1 != nil {
		err = err1
		return
	}
	defer func() {
		if rspCloseErr := res.Body.Close(); rspCloseErr != nil {
			logger.ConsumerLog.Errorf("SearchNFInstances response body cannot close: %+v", rspCloseErr)
		}
	}()

	if res != nil && res.StatusCode == http.StatusTemporaryRedirect {
		err = fmt.Errorf("Temporary Redirect For Non NRF Consumer")
	}
	return
}

func SendNFIntancesUDR(id string, types int) string {
	self := udm_context.Getself()
	targetNfType := models.NfType_UDR
	requestNfType := models.NfType_UDM
	localVarOptionals := Nnrf_NFDiscovery.SearchNFInstancesParamOpts{
		// 	DataSet: optional.NewInterface(models.DataSetId_SUBSCRIPTION),
	}
	// switch types {
	// case NFDiscoveryToUDRParamSupi:
	// 	localVarOptionals.Supi = optional.NewString(id)
	// case NFDiscoveryToUDRParamExtGroupId:
	// 	localVarOptionals.ExternalGroupIdentity = optional.NewString(id)
	// case NFDiscoveryToUDRParamGpsi:
	// 	localVarOptionals.Gpsi = optional.NewString(id)
	// }
	fmt.Println(self.NrfUri)
	result, err := SendNFIntances(self.NrfUri, targetNfType, requestNfType, localVarOptionals)
	if err != nil {
		logger.ConsumerLog.Error(err.Error())
		return ""
	}
	for _, profile := range result.NfInstances {
		return util.SearchNFServiceUri(profile, models.ServiceName_NUDR_DR, models.NfServiceStatus_REGISTERED)
	}
	return ""
}
