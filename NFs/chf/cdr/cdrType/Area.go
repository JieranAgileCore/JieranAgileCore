package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type Area struct { /* Sequence Type */
	/* Sequence of = 35, FULL Name = struct Area__tacs */
	/* TAC */
	Tacs     []TAC            `ber:"tagNum:0,optional"`
	AreaCode *asn.OctetString `ber:"tagNum:1,optional"`
}
