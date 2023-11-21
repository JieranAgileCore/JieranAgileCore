/*
 * NRF UriList
 */

package context

import (
	"github.com/JieranAgileCore/openapi/models"
)

type UriList struct {
	NfType models.NfType `json:"nfType,omitempty" bson:"nfType,omitempty"`
	Link   Links         `json:"_link" bson:"_link" mapstructure:"_link"`
}
