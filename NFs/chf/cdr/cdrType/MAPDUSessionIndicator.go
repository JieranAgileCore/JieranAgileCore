package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	MAPDUSessionIndicatorPresentMAPDURequest               asn.Enumerated = 0
	MAPDUSessionIndicatorPresentMAPDUNetworkUpgradeAllowed asn.Enumerated = 1
)

type MAPDUSessionIndicator struct {
	Value asn.Enumerated
}
