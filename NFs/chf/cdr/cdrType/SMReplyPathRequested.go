package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	SMReplyPathRequestedPresentNoReplyPathSet asn.Enumerated = 0
	SMReplyPathRequestedPresentReplyPathSet   asn.Enumerated = 1
)

type SMReplyPathRequested struct {
	Value asn.Enumerated
}
