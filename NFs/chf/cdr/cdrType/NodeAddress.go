package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

const (
	NodeAddressPresentNothing int = iota /* No components present */
	NodeAddressPresentIPAddress
	NodeAddressPresentDomainName
)

type NodeAddress struct {
	Present    int                /* Choice Type */
	IPAddress  *IPAddress         `ber:"tagNum:0"`
	DomainName *asn.GraphicString `ber:"tagNum:1"`
}
