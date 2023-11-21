package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	SharingLevelPresentSHARED    asn.Enumerated = 0
	SharingLevelPresentNONSHARED asn.Enumerated = 1
)

type SharingLevel struct {
	Value asn.Enumerated
}
