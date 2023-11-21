package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	ChChSelectionModePresentServingNodeSupplied  asn.Enumerated = 0
	ChChSelectionModePresentSubscriptionSpecific asn.Enumerated = 1
	ChChSelectionModePresentAPNSpecific          asn.Enumerated = 2
	ChChSelectionModePresentHomeDefault          asn.Enumerated = 3
	ChChSelectionModePresentRoamingDefault       asn.Enumerated = 4
	ChChSelectionModePresentVisitingDefault      asn.Enumerated = 5
	ChChSelectionModePresentFixedDefault         asn.Enumerated = 6
)

type ChChSelectionMode struct {
	Value asn.Enumerated
}
