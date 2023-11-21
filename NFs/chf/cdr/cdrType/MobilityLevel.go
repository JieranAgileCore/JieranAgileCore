package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	MobilityLevelPresentStationary         asn.Enumerated = 0
	MobilityLevelPresentNomadic            asn.Enumerated = 1
	MobilityLevelPresentRestrictedMobility asn.Enumerated = 2
	MobilityLevelPresentFullyMobility      asn.Enumerated = 3
)

type MobilityLevel struct {
	Value asn.Enumerated
}
