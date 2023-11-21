package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	APIDirectionPresentInvocation   asn.Enumerated = 0
	APIDirectionPresentNotification asn.Enumerated = 1
)

type APIDirection struct {
	Value asn.Enumerated
}
