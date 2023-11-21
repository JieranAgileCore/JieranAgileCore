package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	PriorityTypePresentLow    asn.Enumerated = 0
	PriorityTypePresentNormal asn.Enumerated = 1
	PriorityTypePresentHigh   asn.Enumerated = 2
)

type PriorityType struct {
	Value asn.Enumerated
}
