package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	ThreeGPPPSDataOffStatusPresentActive   asn.Enumerated = 0
	ThreeGPPPSDataOffStatusPresentInactive asn.Enumerated = 1
)

type ThreeGPPPSDataOffStatus struct {
	Value asn.Enumerated
}
