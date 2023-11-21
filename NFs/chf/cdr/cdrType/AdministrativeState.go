package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	AdministrativeStatePresentLOCKED       asn.Enumerated = 0
	AdministrativeStatePresentUNLOCKED     asn.Enumerated = 1
	AdministrativeStatePresentSHUTTINGDOWN asn.Enumerated = 2
)

type AdministrativeState struct {
	Value asn.Enumerated
}
