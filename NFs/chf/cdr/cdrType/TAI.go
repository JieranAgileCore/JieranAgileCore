package cdrType

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type TAI struct {	/* Sequence Type */
	PLMNId	PLMNId `ber:"tagNum:0"`
	Tac	TAC `ber:"tagNum:1"`
}

