package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	MICOModeIndicationPresentMICOMode   asn.Enumerated = 0
	MICOModeIndicationPresentNoMICOMode asn.Enumerated = 1
)

type MICOModeIndication struct {
	Value asn.Enumerated
}
