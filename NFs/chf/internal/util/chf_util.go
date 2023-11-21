package util

import (
	"github.com/JieranAgileCore/openapi/Nchf_ConvergedCharging"
)

func GetNchfChargingNotificationCallbackClient() *Nchf_ConvergedCharging.APIClient {
	configuration := Nchf_ConvergedCharging.NewConfiguration()
	client := Nchf_ConvergedCharging.NewAPIClient(configuration)
	return client
}
