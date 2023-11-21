package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type SMAddressInfo struct { /* Sequence Type */
	SMAddressType   *SMAddressType     `ber:"tagNum:0,optional"`
	SMAddressData   *asn.GraphicString `ber:"tagNum:1,optional"`
	SMAddressDomain *SMAddressDomain   `ber:"tagNum:2,optional"`
}
