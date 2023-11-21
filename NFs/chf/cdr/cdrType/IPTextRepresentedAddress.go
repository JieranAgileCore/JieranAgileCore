package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const (
	IPTextRepresentedAddressPresentNothing int = iota /* No components present */
	IPTextRepresentedAddressPresentIPTextV4Address
	IPTextRepresentedAddressPresentIPTextV6Address
)

type IPTextRepresentedAddress struct {
	Present         int            /* Choice Type */
	IPTextV4Address *asn.IA5String `ber:"tagNum:2"`
	IPTextV6Address *asn.IA5String `ber:"tagNum:3"`
}
