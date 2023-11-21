package cdrType

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type SessionAMBR struct {	/* Sequence Type */
	AmbrUL	Bitrate `ber:"tagNum:1"`
	AmbrDL	Bitrate `ber:"tagNum:2"`
}

