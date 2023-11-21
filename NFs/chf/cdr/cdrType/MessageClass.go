package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	MessageClassPresentPersonal           asn.Enumerated = 0
	MessageClassPresentAdvertisement      asn.Enumerated = 1
	MessageClassPresentInformationService asn.Enumerated = 2
	MessageClassPresentAuto               asn.Enumerated = 3
)

type MessageClass struct {
	Value asn.Enumerated
}
