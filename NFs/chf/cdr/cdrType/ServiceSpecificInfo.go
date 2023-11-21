package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type ServiceSpecificInfo struct { /* Sequence Type */
	ServiceSpecificData *asn.GraphicString `ber:"tagNum:0,optional"`
	ServiceSpecificType *int64             `ber:"tagNum:1,optional"`
}
