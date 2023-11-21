package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const ( /* Enum Type */
	MAPDUSteeringFunctionalityPresentMPTCP   asn.Enumerated = 0
	MAPDUSteeringFunctionalityPresentATSSSLL asn.Enumerated = 1
)

type MAPDUSteeringFunctionality struct {
	Value asn.Enumerated
}
