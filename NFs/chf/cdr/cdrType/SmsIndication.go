package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	SmsIndicationPresentSMSSupported    asn.Enumerated = 0
	SmsIndicationPresentSMSNotSupported asn.Enumerated = 1
)

type SmsIndication struct {
	Value asn.Enumerated
}
