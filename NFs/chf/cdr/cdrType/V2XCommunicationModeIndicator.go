package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	V2XCommunicationModeIndicatorPresentV2XComSupported    asn.Enumerated = 0
	V2XCommunicationModeIndicatorPresentV2XComNotSupported asn.Enumerated = 1
)

type V2XCommunicationModeIndicator struct {
	Value asn.Enumerated
}
