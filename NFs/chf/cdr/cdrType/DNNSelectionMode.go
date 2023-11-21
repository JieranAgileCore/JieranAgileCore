package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	DNNSelectionModePresentUEorNetworkProvidedSubscriptionVerified asn.Enumerated = 0
	DNNSelectionModePresentUEProvidedSubscriptionNotVerified       asn.Enumerated = 1
	DNNSelectionModePresentNetworkProvidedSubscriptionNotVerified  asn.Enumerated = 2
)

type DNNSelectionMode struct {
	Value asn.Enumerated
}
