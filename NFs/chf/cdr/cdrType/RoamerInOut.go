package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	RoamerInOutPresentRoamerInBound  asn.Enumerated = 0
	RoamerInOutPresentRoamerOutBound asn.Enumerated = 1
)

type RoamerInOut struct {
	Value asn.Enumerated
}
