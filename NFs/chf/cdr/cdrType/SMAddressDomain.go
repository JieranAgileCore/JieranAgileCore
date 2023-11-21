package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type SMAddressDomain struct { /* Sequence Type */
	SMDomainName       *asn.GraphicString `ber:"tagNum:0,optional"`
	ThreeGPPIMSIMCCMNC *PLMNId            `ber:"tagNum:1,optional"`
}
