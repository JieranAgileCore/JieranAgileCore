package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	QuotaManagementIndicatorPresentOnlineCharging           asn.Enumerated = 0
	QuotaManagementIndicatorPresentOfflineCharging          asn.Enumerated = 1
	QuotaManagementIndicatorPresentQuotaManagementSuspended asn.Enumerated = 2
)

type QuotaManagementIndicator struct {
	Value asn.Enumerated
}
