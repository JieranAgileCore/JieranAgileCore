package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	SMdeliveryReportRequestedPresentYes asn.Enumerated = 0
	SMdeliveryReportRequestedPresentNo  asn.Enumerated = 1
)

type SMdeliveryReportRequested struct {
	Value asn.Enumerated
}
