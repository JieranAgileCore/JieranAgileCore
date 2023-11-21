package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	NetworkFunctionalityPresentCHF         asn.Enumerated = 0
	NetworkFunctionalityPresentSMF         asn.Enumerated = 1
	NetworkFunctionalityPresentAMF         asn.Enumerated = 2
	NetworkFunctionalityPresentSMSF        asn.Enumerated = 3
	NetworkFunctionalityPresentSGW         asn.Enumerated = 4
	NetworkFunctionalityPresentISMF        asn.Enumerated = 5
	NetworkFunctionalityPresentEPDG        asn.Enumerated = 6
	NetworkFunctionalityPresentCEF         asn.Enumerated = 7
	NetworkFunctionalityPresentNEF         asn.Enumerated = 8
	NetworkFunctionalityPresentPGWCSMF     asn.Enumerated = 9
	NetworkFunctionalityPresentMnSProducer asn.Enumerated = 10
)

type NetworkFunctionality struct {
	Value asn.Enumerated
}
