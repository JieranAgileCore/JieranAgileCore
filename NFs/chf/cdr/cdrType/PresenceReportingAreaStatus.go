package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	PresenceReportingAreaStatusPresentInsideArea  asn.Enumerated = 0
	PresenceReportingAreaStatusPresentOutsideArea asn.Enumerated = 1
	PresenceReportingAreaStatusPresentInactive    asn.Enumerated = 2
	PresenceReportingAreaStatusPresentUnknown     asn.Enumerated = 3
)

type PresenceReportingAreaStatus struct {
	Value asn.Enumerated
}
