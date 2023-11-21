package cdrType

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type PSCellInformation struct {	/* Sequence Type */
	NRcgi	*Ncgi `ber:"tagNum:0,optional"`
	Ecgi	*Ecgi `ber:"tagNum:1,optional"`
}

