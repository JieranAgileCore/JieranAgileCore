package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	OperationalStatePresentENABLED  asn.Enumerated = 0
	OperationalStatePresentDISABLED asn.Enumerated = 1
)

type OperationalState struct {
	Value asn.Enumerated
}
