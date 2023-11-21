package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	DelayToleranceIndicatorPresentDTSupported    asn.Enumerated = 0
	DelayToleranceIndicatorPresentDTNotSupported asn.Enumerated = 1
)

type DelayToleranceIndicator struct {
	Value asn.Enumerated
}
