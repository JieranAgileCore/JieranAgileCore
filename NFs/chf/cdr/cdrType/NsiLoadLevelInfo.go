package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type NsiLoadLevelInfo struct { /* Sequence Type */
	LoadLevelInformation *int64           `ber:"tagNum:0,optional"`
	Snssai               *SingleNSSAI     `ber:"tagNum:1,optional"`
	NsiId                *asn.OctetString `ber:"tagNum:2,optional"`
}
