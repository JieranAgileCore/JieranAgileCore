package cdrType

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type EnhancedDiagnostics5G struct {	/* Sequence Type */
	/* Sequence of = 35, FULL Name = struct EnhancedDiagnostics5G__rANNASRelCause */
	/* RANNASRelCause */
	RANNASRelCause []RANNASRelCause `ber:"tagNum:0"`
}

