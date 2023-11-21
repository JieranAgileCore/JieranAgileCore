package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	ManagementOperationStatusPresentOPERATIONSUCCEEDED asn.Enumerated = 0
	ManagementOperationStatusPresentOPERATIONFAILED    asn.Enumerated = 1
)

type ManagementOperationStatus struct {
	Value asn.Enumerated
}
