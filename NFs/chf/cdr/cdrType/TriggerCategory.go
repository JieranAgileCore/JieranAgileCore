package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	TriggerCategoryPresentImmediateReport asn.Enumerated = 0
	TriggerCategoryPresentDeferredReport  asn.Enumerated = 1
)

type TriggerCategory struct {
	Value asn.Enumerated
}
