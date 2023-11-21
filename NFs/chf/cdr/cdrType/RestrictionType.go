package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	RestrictionTypePresentAllowedAreas    asn.Enumerated = 0
	RestrictionTypePresentNotAllowedAreas asn.Enumerated = 1
)

type RestrictionType struct {
	Value asn.Enumerated
}
