package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type GNbId struct { /* Sequence Type */
	BitLength int64         `ber:"tagNum:0"`
	GNbValue  asn.IA5String `ber:"tagNum:1"`
}
