package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	PreemptionCapabilityPresentNOTPREEMPT asn.Enumerated = 0
	PreemptionCapabilityPresentMAYPREEMPT asn.Enumerated = 1
)

type PreemptionCapability struct {
	Value asn.Enumerated
}
