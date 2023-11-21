package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	PartialRecordMethodPresentDefault    asn.Enumerated = 0
	PartialRecordMethodPresentIndividual asn.Enumerated = 1
)

type PartialRecordMethod struct {
	Value asn.Enumerated
}
