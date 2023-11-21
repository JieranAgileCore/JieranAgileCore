package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type SMInterface struct { /* Sequence Type */
	InterfaceId   *asn.GraphicString `ber:"tagNum:0,optional"`
	InterfaceText *asn.GraphicString `ber:"tagNum:1,optional"`
	InterfacePort *asn.GraphicString `ber:"tagNum:2,optional"`
	InterfaceType *SMInterfaceType   `ber:"tagNum:3,optional"`
}
