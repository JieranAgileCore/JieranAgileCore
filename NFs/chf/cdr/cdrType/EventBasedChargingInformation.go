package cdrType

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type EventBasedChargingInformation struct {	/* Sequence Type */
	NumberOfEvents	int64 `ber:"tagNum:1"`
	/* Sequence of = 35, FULL Name = struct EventBasedChargingInformation__eventTimeStamps */
	/* TimeStamp */
	EventTimeStamps []TimeStamp `ber:"tagNum:2,optional"`
}

