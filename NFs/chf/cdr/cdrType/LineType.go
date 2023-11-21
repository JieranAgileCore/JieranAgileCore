package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	LineTypePresentDSL asn.Enumerated = 0
	LineTypePresentPON asn.Enumerated = 1
)

type LineType struct {
	Value asn.Enumerated
}
